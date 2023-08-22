package dto

// untuk disimpan di parameter
type PaginationParam struct {
	Page int
	Offset int
	Limit int
}

// untuk disimpan di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// untuk disimpan di response
type Paging struct {
	Page        int `json:"paging"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}

// example
// product 50

// Paging {Page: 1. RowsPerPage: 10, TotalRows: 50, TotalPages: 5}