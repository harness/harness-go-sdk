package nextgen

import (
	"context"

	"github.com/antihax/optional"
)

// Iterator provides a generic way to iterate over paginated results.
// Call Next() in a loop; it returns false when exhausted or on error.
// After iteration, check Err() for any errors encountered.
type Iterator[T any] struct {
	items   []T
	pos     int
	done    bool
	err     error
	fetcher func(page int32) ([]T, bool, error)
	page    int32
}

// Next advances the iterator. Returns false when there are no more items.
func (it *Iterator[T]) Next() bool {
	if it.done {
		return false
	}
	it.pos++
	if it.pos < len(it.items) {
		return true
	}
	if it.pos > 0 && it.pos >= len(it.items) {
		items, isLast, err := it.fetcher(it.page)
		if err != nil {
			it.err = err
			it.done = true
			return false
		}
		if len(items) == 0 {
			it.done = true
			return false
		}
		it.items = items
		it.pos = 0
		it.page++
		if isLast {
			it.done = true
		}
		return true
	}
	return false
}

// Value returns the current item. Only valid after a successful Next() call.
func (it *Iterator[T]) Value() T {
	return it.items[it.pos]
}

// Err returns the error that stopped iteration, if any.
func (it *Iterator[T]) Err() error {
	return it.err
}

// Collect drains the iterator into a slice. Useful when you want all results at once.
func (it *Iterator[T]) Collect() ([]T, error) {
	var result []T
	for it.Next() {
		result = append(result, it.Value())
	}
	return result, it.Err()
}

// ListAllPipelines returns an iterator over all pipelines in a project.
func ListAllPipelines(ctx context.Context, client *APIClient, accountID, orgID, projectID string) *Iterator[PmsPipelineSummaryResponse] {
	it := &Iterator[PmsPipelineSummaryResponse]{
		pos:  -1,
		page: 0,
	}
	it.fetcher = func(page int32) ([]PmsPipelineSummaryResponse, bool, error) {
		resp, httpResp, err := client.PipelinesApi.GetPipelineList(ctx, accountID, orgID, projectID, &PipelinesApiGetPipelineListOpts{
			Page: optional.NewInt32(page),
			Size: optional.NewInt32(25),
		})
		if err != nil {
			return nil, false, err
		}
		if httpResp != nil {
			httpResp.Body.Close()
		}
		if resp.Data == nil {
			return nil, true, nil
		}
		return resp.Data.Content, resp.Data.Last, nil
	}
	// Trigger initial fetch
	items, isLast, err := it.fetcher(0)
	if err != nil {
		it.err = err
		it.done = true
	} else {
		it.items = items
		it.page = 1
		if isLast || len(items) == 0 {
			it.done = true
		}
	}
	return it
}

// ListAllExecutions returns an iterator over all pipeline executions in a project.
func ListAllExecutions(ctx context.Context, client *APIClient, accountID, orgID, projectID string, opts *ExecutionDetailsApiGetListOfExecutionsOpts) *Iterator[PipelineExecutionSummary] {
	it := &Iterator[PipelineExecutionSummary]{
		pos:  -1,
		page: 0,
	}
	it.fetcher = func(page int32) ([]PipelineExecutionSummary, bool, error) {
		listOpts := &ExecutionDetailsApiGetListOfExecutionsOpts{}
		if opts != nil {
			*listOpts = *opts
		}
		listOpts.Page = optional.NewInt32(page)
		if !listOpts.Size.IsSet() {
			listOpts.Size = optional.NewInt32(25)
		}

		resp, httpResp, err := client.ExecutionDetailsApi.GetListOfExecutions(ctx, accountID, orgID, projectID, listOpts)
		if err != nil {
			return nil, false, err
		}
		if httpResp != nil {
			httpResp.Body.Close()
		}
		if resp.Data == nil {
			return nil, true, nil
		}
		return resp.Data.Content, resp.Data.Last, nil
	}
	// Trigger initial fetch
	items, isLast, err := it.fetcher(0)
	if err != nil {
		it.err = err
		it.done = true
	} else {
		it.items = items
		it.page = 1
		if isLast || len(items) == 0 {
			it.done = true
		}
	}
	return it
}
