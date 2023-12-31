// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Filters struct {
	// Specify the end date till when you want the settlement reconciliation details.
	EndDate string `json:"end_date"`
	// Specify the start date from when you want the settlement reconciliation details.
	StartDate string `json:"start_date"`
}

func (o *Filters) GetEndDate() string {
	if o == nil {
		return ""
	}
	return o.EndDate
}

func (o *Filters) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

// Pagination - To fetch the next set of settlements, pass the cursor received in the response to the next API call.
//
//	To receive the data for the first time, pass the cursor as null.
//	Limit would be number of settlements that you want to receive.
type Pagination struct {
	// Specifies from where the next set of settlement details should be fetched.
	Cursor *string `json:"cursor,omitempty"`
	// Number of settlements you want to fetch in the next iteration. Maximum limit is 1000, default value is 10.
	Limit int64 `json:"limit"`
}

func (o *Pagination) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *Pagination) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

type FetchPGReconRequest struct {
	Filters Filters `json:"filters"`
	// To fetch the next set of settlements, pass the cursor received in the response to the next API call.
	//  To receive the data for the first time, pass the cursor as null.
	//  Limit would be number of settlements that you want to receive.
	Pagination Pagination `json:"pagination"`
}

func (o *FetchPGReconRequest) GetFilters() Filters {
	if o == nil {
		return Filters{}
	}
	return o.Filters
}

func (o *FetchPGReconRequest) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}
