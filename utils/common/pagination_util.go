package common

import (
	"math"
	"os"
	"strconv"

	"github.com/NursiNursi/laundry-apps/model/dto"
	"github.com/NursiNursi/laundry-apps/utils/exceptions"
)

func GetPaginationParams(params dto.PaginationParam) dto.PaginationQuery {
	err := LoadEnv()
	exceptions.CheckErr(err)

	var (page, take, skip int)

	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0{
		n, _ := strconv.Atoi(os.Getenv("DEFAULT_ROWS_PER_PAGE"))
		take = n
	} else {
		take = params.Limit
	}

	// rumus offset / pagination
	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return dto.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

func Paginate(page, limit, totalRows int) dto.Paging {
	return dto.Paging{
		Page: page,
		RowsPerPage: limit,
		TotalRows: totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}