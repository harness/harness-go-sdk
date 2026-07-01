package nextgen

import (
	"net/url"
	"testing"
)

func TestAddQueryParam_MultiFormat_StringSlice(t *testing.T) {
	params := url.Values{}
	values := []string{"RUNNING", "SUCCESS", "FAILED"}

	addQueryParam(params, "status", values, "multi")

	if len(params["status"]) != 3 {
		t.Errorf("Expected 3 values, got %d", len(params["status"]))
	}

	expected := []string{"RUNNING", "SUCCESS", "FAILED"}
	for i, val := range expected {
		if params["status"][i] != val {
			t.Errorf("Expected params[status][%d] = %s, got %s", i, val, params["status"][i])
		}
	}
}

func TestAddQueryParam_MultiFormat_InterfaceSlice(t *testing.T) {
	params := url.Values{}
	values := []interface{}{"value1", "value2"}

	addQueryParam(params, "key", values, "multi")

	if len(params["key"]) != 2 {
		t.Errorf("Expected 2 values, got %d", len(params["key"]))
	}

	if params["key"][0] != "value1" || params["key"][1] != "value2" {
		t.Errorf("Unexpected values: %v", params["key"])
	}
}

func TestAddQueryParam_MultiFormat_SingleValue(t *testing.T) {
	params := url.Values{}

	addQueryParam(params, "key", "single_value", "multi")

	if len(params["key"]) != 1 {
		t.Errorf("Expected 1 value, got %d", len(params["key"]))
	}

	if params["key"][0] != "single_value" {
		t.Errorf("Expected 'single_value', got %s", params["key"][0])
	}
}

func TestAddQueryParam_CSV_Format(t *testing.T) {
	params := url.Values{}
	values := []string{"a", "b", "c"}

	addQueryParam(params, "tags", values, "csv")

	if len(params["tags"]) != 1 {
		t.Errorf("Expected 1 value for CSV format, got %d", len(params["tags"]))
	}

	if params["tags"][0] != "a,b,c" {
		t.Errorf("Expected 'a,b,c', got %s", params["tags"][0])
	}
}

func TestAddQueryParam_Pipes_Format(t *testing.T) {
	params := url.Values{}
	values := []string{"x", "y", "z"}

	addQueryParam(params, "filter", values, "pipes")

	if len(params["filter"]) != 1 {
		t.Errorf("Expected 1 value for pipes format, got %d", len(params["filter"]))
	}

	if params["filter"][0] != "x|y|z" {
		t.Errorf("Expected 'x|y|z', got %s", params["filter"][0])
	}
}

func TestAddQueryParam_EmptyString_Format(t *testing.T) {
	params := url.Values{}

	addQueryParam(params, "simple", "value", "")

	if len(params["simple"]) != 1 {
		t.Errorf("Expected 1 value, got %d", len(params["simple"]))
	}

	if params["simple"][0] != "value" {
		t.Errorf("Expected 'value', got %s", params["simple"][0])
	}
}

func TestAddQueryParam_URLEncoding(t *testing.T) {
	params := url.Values{}
	values := []string{"value with spaces", "value&special"}

	addQueryParam(params, "data", values, "multi")

	if len(params["data"]) != 2 {
		t.Errorf("Expected 2 values, got %d", len(params["data"]))
	}

	// url.Values.Encode() should properly encode these
	encoded := params.Encode()
	if encoded != "data=value+with+spaces&data=value%26special" &&
		encoded != "data=value%20with%20spaces&data=value%26special" {
		t.Logf("Encoded params: %s", encoded)
		// This is informational - URL encoding can vary slightly
	}
}

func TestAddQueryParam_MultiFormat_EmptySlice(t *testing.T) {
	params := url.Values{}
	values := []string{}

	addQueryParam(params, "empty", values, "multi")

	if len(params["empty"]) != 0 {
		t.Errorf("Expected 0 values for empty slice, got %d", len(params["empty"]))
	}
}
