// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type QueryResponseResults struct {
}

type QueryResponse struct {
	Results      []map[string]QueryResponseResults
	Aggregations []AggregationResponse
	TotalCount   *int64
}

func (o *QueryResponse) GetResults() []map[string]QueryResponseResults {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *QueryResponse) GetAggregations() []AggregationResponse {
	if o == nil {
		return nil
	}
	return o.Aggregations
}

func (o *QueryResponse) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}
