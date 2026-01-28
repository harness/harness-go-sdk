#!/bin/bash
set -euo pipefail

# ===========================================
# Configuration
# ===========================================

# Script information
SCRIPT_NAME=$(basename "${BASH_SOURCE[0]}")
SCRIPT_VERSION="2.0.0"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Directory configuration
GENERATED_DIR="${SCRIPT_DIR}/generated"
TEMP_DIR="${SCRIPT_DIR}/temp_$(date +%s)"
BACKUP_DIR="${SCRIPT_DIR}/backup_$(date +%Y%m%d_%H%M%S)"

# Default values
VERBOSE=${VERBOSE:-false}
DRY_RUN=${DRY_RUN:-false}
AUTO_FIX=${AUTO_FIX:-false}
BUILD_JOBS=${BUILD_JOBS:-4}

# ===========================================
# Utility Functions
# ===========================================

# Get current timestamp in ISO 8601 format
get_timestamp() {
    date +"%Y-%m-%d %H:%M:%S"
}

# Find the module root by looking for go.mod in parent directories
find_module_root() {
    local dir="${1:-$PWD}"
    local max_depth=10
    local depth=0
    
    while [[ "$dir" != "/" && $depth -lt $max_depth ]]; do
        if [ -f "${dir}/go.mod" ]; then
            echo "$dir"
            return 0
        fi
        dir="$(dirname "$dir")"
        ((depth++))
    done
    
    log_error "Could not find go.mod in any parent directory (searched $max_depth levels)"
    return 1
}

# Set TARGET_DIR to the module root
export TARGET_DIR="$(find_module_root "$(pwd)")" || {
    echo "Error: Could not find go.mod in any parent directory" >&2
    exit 1
}

echo "Using module root: $TARGET_DIR"
cd "$TARGET_DIR"

# Now set SWAGGER_FILE after TARGET_DIR is defined
SWAGGER_FILE="${TARGET_DIR}/harness/chaos/api/swagger.json"

# Only process files in the harness/chaos directory
CHAOS_DIR="$TARGET_DIR/harness/chaos"
if [ ! -d "$CHAOS_DIR" ]; then
    echo "Error: harness/chaos directory not found in $TARGET_DIR" >&2
    exit 1
fi

# Set GO_FILES to only include files under harness/chaos
GO_FILES=$(find "$CHAOS_DIR" -name "*.go" -not -path "*/vendor/*" -not -path "*/testdata/*")

# Ensure we're in the correct directory
if [ ! -f "go.mod" ]; then
    echo "Error: This script must be run from the module root directory" >&2
    exit 1
fi

# Ensure we're in the correct directory
cd "$TARGET_DIR"

# Cleanup function
cleanup() {
    # Clean up temporary directories
    if [ -d "$TEMP_DIR" ]; then
        log_debug "Cleaning up temporary directory: $TEMP_DIR"
        rm -rf "$TEMP_DIR"
    fi
}

# Function to prompt user for action
prompt_user() {
    local message="$1"
    local default_action="${2:-abort}"  # Default to 'abort' if not specified
    
    while true; do
        read -p "$message [F]ix manually, [A]bort (default: ${default_action:0:1}): " -n 1 -r
        echo
        case $REPLY in
            [Ff]* ) return 1;;  # Return non-zero for fix
            [Aa]* ) return 0;;  # Return zero for abort
            '' ) if [ "$default_action" = "abort" ]; then return 0; else return 1; fi;;  # Handle empty input
            * ) echo "Please enter 'F' to fix manually or 'A' to abort.";;
        esac
    done
}

# Register cleanup on script exit
trap cleanup EXIT

# Colors for output
if [ -t 1 ]; then
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m' # No Color
else
    RED='' GREEN='' YELLOW='' BLUE='' NC=''
fi

# Show help message
show_help() {
    cat << EOF
Usage: $0 [OPTIONS]

Options:
  -v, --verbose       Enable verbose output
  -h, --help          Show this help message and exit
  --dry-run           Run in dry-run mode (no changes will be made)
  --auto-fix          Automatically fix formatting and import issues
  -j, --jobs N        Number of parallel build jobs (default: 4)

Environment variables:
  VERBOSE       Set to 'true' to enable verbose output
  DRY_RUN       Set to 'true' to enable dry-run mode
  AUTO_FIX      Set to 'true' to automatically fix formatting issues
EOF
}

# Parse command line arguments
parse_arguments() {
    # Set default values
    VERBOSE=${VERBOSE:-false}
    DRY_RUN=${DRY_RUN:-false}
    AUTO_FIX=${AUTO_FIX:-false}
    
    while [[ $# -gt 0 ]]; do
        case $1 in
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -h|--help)
                show_help
                exit 0
                ;;
            --dry-run)
                DRY_RUN=true
                shift
                ;;
            --auto-fix)
                AUTO_FIX=true
                shift
                ;;
            -j|--jobs)
                BUILD_JOBS="$2"
                if ! [[ "$BUILD_JOBS" =~ ^[0-9]+$ ]]; then
                    log_error "Number of jobs must be a positive integer"
                    exit 1
                fi
                shift 2
                ;;
            *)
                log_error "Unknown option: $1"
                exit 1
                ;;
        esac
    done
}

# ===========================================
# Logging Functions
# ===========================================

log_debug() {
    if [ "$VERBOSE" = true ]; then
        echo -e "$(get_timestamp) ${BLUE}[DEBUG]${NC} $*"
    fi
}

log_info() {
    echo -e "$(get_timestamp) ${BLUE}[INFO]${NC} $*"
}

log_success() {
    echo -e "$(get_timestamp) ${GREEN}[SUCCESS]${NC} $*"
}

log_warning() {
    echo -e "$(get_timestamp) ${YELLOW}[WARNING]${NC} $*" >&2
}

log_error() {
    echo -e "$(get_timestamp) ${RED}[ERROR]${NC} $*" >&2
}

# Check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to verify Go code formatting and imports
verify_go_code() {
    local context=${1:-"verification"}
    local has_errors=0
    local needs_manual_fix=0
    local output
    local max_retries=3
    local retry_count=0
    
    log_info "Verifying Go code formatting and imports ($context)..."
    
    # Check gofmt
    while true; do
        log_debug "Running gofmt..."
        local gofmt_issues
        gofmt_issues=$(gofmt -l $GO_FILES 2>/dev/null)
        
        if [ -n "$gofmt_issues" ]; then
            if [ $retry_count -eq 0 ]; then
                log_warning "gofmt found issues in files:"
                echo "$gofmt_issues" | sed 's/^/  /'
                if [ "$DRY_RUN" = true ]; then
                    log_info "[DRY RUN] Would fix formatting issues..."
                    break
                else
                    log_info "Attempting to fix formatting issues..."
                fi
            fi
            
            # Skip actual file modifications in dry-run mode
            if [ "$DRY_RUN" = false ]; then
                local files_to_fix=($(echo "$gofmt_issues" | tr '\n' ' '))
                for file in "${files_to_fix[@]}"; do
                    if [ -f "$file" ]; then
                        gofmt -w "$file"
                    fi
                done
            else
                break
            fi
            
            # Check if we've fixed the issues
            output=$(gofmt -l $GO_FILES)
            if [ -z "$output" ]; then
                log_success "✓ All formatting issues fixed"
                break
            elif [ $retry_count -lt $max_retries ]; then
                retry_count=$((retry_count + 1))
                log_info "Retrying formatting fixes (attempt $retry_count/$max_retries)..."
                continue
            else
                echo -e "\n${YELLOW}Could not automatically fix all formatting issues. Please fix them manually:${NC}"
                echo "$output"
                needs_manual_fix=1
                has_errors=1
                break
            fi
        else
            log_success "✓ No formatting issues found"
            break
        fi
    done
    
    # Reset retry counter for goimports
    retry_count=0
    
    # Check goimports if available
    if command -v goimports >/dev/null 2>&1; then
        while true; do
            log_debug "Running goimports..."
            local goimports_issues
            goimports_issues=$(goimports -l $GO_FILES 2>/dev/null)
            
            if [ -n "$goimports_issues" ]; then
                if [ $retry_count -eq 0 ]; then
                    log_warning "goimports found issues in files:"
                    echo "$goimports_issues" | sed 's/^/  /'
                    if [ "$DRY_RUN" = true ]; then
                        log_info "[DRY RUN] Would fix import issues..."
                        break
                    else
                        log_info "Attempting to fix import issues..."
                    fi
                fi
                
                # Skip actual file modifications in dry-run mode
                if [ "$DRY_RUN" = false ]; then
                    local files_to_fix=($(echo "$goimports_issues" | tr '\n' ' '))
                    for file in "${files_to_fix[@]}"; do
                        if [ -f "$file" ]; then
                            goimports -w "$file"
                        fi
                    done
                else
                    break
                fi
                
                # Check if we've fixed the issues
                output=$(goimports -l $GO_FILES)
                if [ -z "$output" ]; then
                    log_success "✓ All import issues fixed"
                    break
                elif [ $retry_count -lt $max_retries ]; then
                    retry_count=$((retry_count + 1))
                    log_info "Retrying import fixes (attempt $retry_count/$max_retries)..."
                    continue
                else
                    echo -e "\n${YELLOW}Could not automatically fix all import issues. Please fix them manually:${NC}"
                    echo "$output"
                    needs_manual_fix=1
                    has_errors=1
                    break
                fi
            else
                log_success "✓ No import issues found"
                break
            fi
        done
    else
        log_warning "goimports not found, skipping import verification"
    fi
    
    # Handle manual fixes if needed
    if [ $needs_manual_fix -ne 0 ]; then
        echo -e "\n${YELLOW}Some issues require manual intervention.${NC}"
        echo "Please fix the issues shown above in your editor."
        echo -e "${YELLOW}After fixing the issues, press Enter to continue or Ctrl+C to abort...${NC}"
        read -p "> "
        
        # Verify fixes after manual intervention
        log_info "Verifying fixes after manual intervention..."
        
        # Check gofmt again
        output=$(gofmt -l $GO_FILES)
        if [ -n "$output" ]; then
            echo -e "\n${RED}The following files still have formatting issues:${NC}"
            echo "$output"
            has_errors=1
        fi
        
        # Check goimports again if available
        if command -v goimports >/dev/null 2>&1; then
            output=$(goimports -l $GO_FILES)
            if [ -n "$output" ]; then
                echo -e "\n${RED}The following files still have import issues:${NC}"
                echo "$output"
                has_errors=1
            fi
        fi
        
        if [ $has_errors -ne 0 ]; then
            log_error "There are still issues that need to be fixed."
            return 1
        fi
    fi
    
    if [ $has_errors -eq 0 ]; then
        log_success "✓ All code verification passed!"
        return 0
    else
        log_error "Code verification failed. Please fix the issues and try again."
        return 1
    fi
}

# Function to verify required tools are installed
check_dependencies() {
    local missing_deps=0
    
    if ! command -v swagger-codegen &> /dev/null; then
        log_error "swagger-codegen is not installed. Please install it first."
        missing_deps=$((missing_deps + 1))
    fi
    
    if ! command -v go &> /dev/null; then
        log_error "Go is not installed or not in PATH"
        missing_deps=$((missing_deps + 1))
    fi
    
    if [ $missing_deps -gt 0 ]; then
        log_error "Missing $missing_deps required dependencies. Please install them and try again."
        return 1
    fi
    
    return 0
}

# Function to create backup of files
create_backup() {
    local dir="${1:-$CHAOS_DIR}"
    
    # Ensure we're only backing up the chaos directory or its subdirectories
    if [[ ! "$dir" =~ ^${CHAOS_DIR}(/|$) ]]; then
        log_error "Backup target must be within $CHAOS_DIR"
        return 1
    fi
    
    local backup_dir="${dir}_backup_$(date +%Y%m%d_%H%M%S)"
    
    if [ "$DRY_RUN" = true ]; then
        log_warning "[DRY RUN] Would create backup of $dir at $backup_dir"
        echo "$backup_dir"
        return 0
    fi
    
    if [ ! -d "$dir" ]; then
        log_error "Source directory does not exist: $dir"
        return 1
    fi
    
    log_info "Creating backup of $dir at $backup_dir"
    
    # Create backup directory structure
    mkdir -p "$(dirname "$backup_dir")" || {
        log_error "Failed to create backup directory: $(dirname "$backup_dir")"
        return 1
    }
    
    # Copy files to backup
    if command -v rsync >/dev/null 2>&1; then
        if ! rsync -a "$dir/" "$backup_dir"; then
            log_error "Failed to create backup of $dir"
            return 1
        fi
    else
        if ! cp -r "$dir" "$backup_dir" 2>/dev/null; then
            log_error "Failed to create backup of $dir"
            return 1
        fi
    fi
    
    log_success "Created backup at: $backup_dir"
    echo "$backup_dir"
    return 0
}

# Function to restore from backup
restore_from_backup() {
    local backup_dir="$1"
    local target_dir="${2:-$CHAOS_DIR}"
    
    # Ensure we're only restoring to the chaos directory or its subdirectories
    if [[ ! "$target_dir" =~ ^${CHAOS_DIR}(/|$) ]]; then
        log_error "Restore target must be within $CHAOS_DIR"
        return 1
    fi
    
    if [ "$DRY_RUN" = true ]; then
        log_warning "[DRY RUN] Would restore from backup: $backup_dir to $target_dir"
        return 0
    fi
    
    if [ ! -d "$backup_dir" ]; then
        log_error "Backup directory does not exist: $backup_dir"
        return 1
    fi
    
    log_info "Restoring from backup: $backup_dir to $target_dir"
    
    # Verify target directory is inside CHAOS_DIR
    if [[ ! "$target_dir" =~ ^${CHAOS_DIR}(/|$) ]]; then
        log_error "Cannot restore outside of $CHAOS_DIR"
        return 1
    fi
    
    # Create parent directory if it doesn't exist
    local parent_dir=$(dirname "$target_dir")
    if [ ! -d "$parent_dir" ]; then
        log_debug "Creating parent directory: $parent_dir"
        mkdir -p "$parent_dir" || {
            log_error "Failed to create directory: $parent_dir"
            return 1
        }
    fi
    
    # Copy files from backup
    if command -v rsync >/dev/null 2>&1; then
        if ! rsync -a --delete "$backup_dir/" "$target_dir"; then
            log_error "Failed to restore from backup: $backup_dir to $target_dir"
            return 1
        fi
    else
        # Remove existing target directory if it exists
        if [ -d "$target_dir" ]; then
            log_debug "Removing existing directory: $target_dir"
            rm -rf "$target_dir" || {
                log_error "Failed to remove directory: $target_dir"
                return 1
            }
        fi
        
        if ! cp -r "$backup_dir" "$(dirname "$target_dir")/"; then
            log_error "Failed to restore from backup: $backup_dir to $target_dir"
            return 1
        fi
    fi
    
    log_success "Successfully restored from backup: $backup_dir"
    return 0
}

# ===========================================
# Build Verification
# ===========================================

# Function to verify Go build
verify_go_build() {
    local phase="${1:-unknown}"
    local build_cmd="go build -v -p $BUILD_JOBS ./..."
    
    log_info "Verifying Go build ($phase phase)..."
    log_debug "Build command: $build_cmd"
    
    local start_time=$(date +%s)
    
    # Run the build command and capture output
    local build_output
    if ! build_output=$($build_cmd 2>&1); then
        log_error "Build verification failed during $phase phase"
        
        # Provide more context for common build errors
        if echo "$build_output" | grep -q "cannot find package"; then
            log_error "Dependency error: Missing Go package"
            log_info "Try running 'go mod tidy' to update dependencies"
        elif echo "$build_output" | grep -q "undefined:"; then
            log_error "Compilation error: Undefined symbol"
        fi
        
        # Show first few lines of build output
        echo -e "\n${YELLOW}Build output:${NC}"
        echo "$build_output" | head -n 20
        
        # If there's more output, show a summary
        local line_count=$(echo "$build_output" | wc -l)
        if [ "$line_count" -gt 20 ]; then
            echo "... and $((line_count - 20)) more lines"
        fi
        
        return 1
    fi
    
    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    
    log_success "Build verification passed in ${duration}s ($phase phase)"
    return 0
}

# ===========================================
# Code Generation
# ===========================================

# Function to generate code using swagger-codegen
generate_code() {
    local swagger_file="${1:?Swagger file path is required}"
    local version="${2:?Version is required}"
    local temp_dir="${TEMP_DIR:?TEMP_DIR is not set}"
    
    log_info "Starting code generation..."
    log_debug "Swagger file: $swagger_file"
    log_debug "Version: $version"
    log_debug "Temporary directory: $temp_dir"
    
    # Validate swagger file
    if [ ! -f "$swagger_file" ]; then
        log_error "Swagger file not found: $swagger_file"
        return 1
    fi
    
    # Create temporary directory
    mkdir -p "$temp_dir"
    
    # Generate code using swagger-codegen
    log_info "Running swagger-codegen..."
    local start_time=$(date +%s)
    
    if ! swagger-codegen generate \
        -i "$swagger_file" \
        -l go \
        -o "$temp_dir" \
        --additional-properties=packageName=chaos,packageVersion="$version" \
        --http-user-agent="Harness-Chaos-SDK/$version" \
        --git-user-id="harness" \
        --git-repo-id="harness-go-sdk" \
        --release-note="Generated by $SCRIPT_NAME v$SCRIPT_VERSION"; then
        log_error "Failed to generate code from $swagger_file"
        return 1
    fi
    
    local end_time=$(date +%s)
    log_debug "Code generation completed in $((end_time - start_time)) seconds"
    
    # Move generated files to the target directory
    log_info "Organizing generated files..."
    
    # Create target directories if they don't exist
    mkdir -p "$CHAOS_DIR/api"
    mkdir -p "$CHAOS_DIR/model"
    
    # Move files with error handling
    local moved_files=0
    local skipped_files=0
    
    # Move API files
    if [ -d "$temp_dir/api" ]; then
        for file in "$temp_dir/api/"*.go; do
            if [ -f "$file" ]; then
                local base_name=$(basename "$file")
                local target_file="$CHAOS_DIR/api/${base_name#*_}"  # Remove prefix if any
                
                log_debug "Moving $file to $target_file"
                mv "$file" "$target_file"
                ((moved_files++))
            fi
        done
    fi
    
    # Move model files
    if [ -d "$temp_dir/model" ]; then
        for file in "$temp_dir/model/"*.go; do
            if [ -f "$file" ]; then
                log_debug "Moving $file to $CHAOS_DIR/model/"
                mv "$file" "$CHAOS_DIR/model/"
                ((moved_files++))
            fi
        done
    fi
    
    # Move root files
    for file in "$temp_dir/"*.go; do
        if [ -f "$file" ]; then
            log_debug "Moving $file to $CHAOS_DIR/"
            mv "$file" "$CHAOS_DIR/"
            ((moved_files++))
        fi
    done
    
    # Clean up
    log_debug "Cleaning up temporary files..."
    rm -rf "$temp_dir"
    
    log_success "Code generation completed: $moved_files files generated"
    
    if [ $skipped_files -gt 0 ]; then
        log_warning "Skipped $skipped_files files (see debug logs for details)"
    fi
    
    return 0
}

# ===========================================
# Error Handling
# ===========================================

# Function to handle errors and rollback if needed
handle_error() {
    local exit_code=$?
    local error_line=$1
    local error_cmd=$2
    
    log_error "Script failed with exit code $exit_code"
    log_error "Line: $error_line, command: $error_cmd"
    
    if [ -d "$BACKUP_DIR" ] && [ "$DRY_RUN" = false ]; then
        log_warning "Attempting to rollback changes..."
        if restore_from_backup "$BACKUP_DIR" "$TARGET_DIR/harness/chaos"; then
            log_success "Rollback completed successfully"
        else
            log_error "Rollback failed. Manual intervention required."
            log_error "Backup directory: $BACKUP_DIR"
        fi
    fi
    
    # Clean up temporary files
    cleanup
    
    exit $exit_code
}

# Set up error handling
trap 'handle_error ${LINENO} "$BASH_COMMAND"' ERR

# Main execution
main() {
    # Parse command line arguments
    parse_arguments "$@"
    
    local swagger_file="${SWAGGER_FILE}"
    local version="$(date +%Y%m%d.0)"
    
    log_info "=== Starting Chaos Go SDK Code Generation ==="
    log_info "Target directory: ${TARGET_DIR}"
    log_info "Swagger file: ${swagger_file}"
    log_info "Version: ${version}"
    log_info "Verbose mode: ${VERBOSE}"
    log_info "Dry run: ${DRY_RUN}"
    
    # Verify dependencies
    log_info "Verifying dependencies..."
    if ! check_dependencies; then
        log_error "Dependency check failed"
        exit 1
    fi
    log_success "All dependencies verified"
    
    # Check if swagger file exists
    if [ ! -f "$swagger_file" ]; then
        log_error "Swagger file not found: $swagger_file"
        exit 1
    fi
    
    # Create necessary directories
    mkdir -p "$GENERATED_DIR"
    
    # Verify code before making changes
    log_info "Running pre-generation verification..."
    if ! verify_go_build "pre-generation"; then
        log_error "Pre-generation build verification failed. Please fix the issues before proceeding."
        exit 1
    fi
    log_success "Pre-generation verification passed"
    
    # Create backup of current files
    if [ "$DRY_RUN" = true ]; then
        log_warning "[DRY RUN] Would create backup of current files"
    else
        log_info "Creating backup of current files..."
        if ! create_backup; then
            log_error "Failed to create backup. Aborting..."
            exit 1
        fi
        log_success "Backup created successfully"
        
        # Set up error handler
        trap 'handle_error' ERR
    fi
    
    # Generate new client code
    log_info "Generating client code..."
    if [ "$DRY_RUN" = true ]; then
        log_warning "[DRY RUN] Would generate new client code"
    else
        if ! generate_code "$swagger_file" "$version"; then
            log_error "Code generation failed"
            exit 1
        fi
        log_success "Code generation completed"
        
        # Verify generated code
        log_info "Verifying generated code..."
        if ! verify_go_build "post-generation"; then
            log_error "Generated code verification failed"
            exit 1
        fi
        log_success "Generated code verified successfully"
        
        # Verify code formatting
        log_info "Verifying code formatting and imports..."
        if ! verify_go_code "post-generation"; then
            log_warning "Code formatting/imports verification failed. Please fix the reported issues."
            if [ "$DRY_RUN" = false ] && [ "$AUTO_FIX" = true ]; then
                log_info "Attempting to fix formatting issues automatically..."
                if ! verify_go_code "fix"; then
                    log_error "Failed to fix formatting issues automatically"
                    exit 1
                fi
                log_success "Formatting issues fixed automatically"
            else
                log_info "You can run 'make format' to fix formatting issues"
            fi
        else
            log_success "Code formatting and imports verified"
        fi
        
        # Final build verification
        log_info "Performing final build verification..."
        if ! verify_go_build "final"; then
            log_error "Final build verification failed"
            exit 1
        fi
        log_success "Final build verification passed"
    fi
    
    log_success "\n=== Code Generation Completed Successfully ==="
    log_info "You can now review the changes and commit them to version control."
}

main "$@"
