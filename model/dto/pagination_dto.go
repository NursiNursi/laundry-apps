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
	Page int
	RowsPerPage int
	TotalRows int
	TotalPages int
}

// example
// product 50

// Paging {Page: 1. RowsPerPage: 10, TotalRows: 50, TotalPages: 5}