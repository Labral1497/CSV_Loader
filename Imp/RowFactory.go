package imp

import (
	"coralogx_EX/processing"
	"sort"
	"strconv"
)

type GetRows struct {
	indices []int
}

func NewGetRows(indices ...int) *GetRows {
	sort.Ints(indices)
	return &GetRows{indices: indices}
}

func (f *GetRows) Apply(rows [][]string) [][]string {
	if processing.RowCounter <= f.indices[len(f.indices)-1] && processing.RowCounter >= f.indices[0] {
		return rows
	}
	return nil
}

// FilterRowsByFunc operation
type FilterRows struct {
	filterFunc func([]string) bool
}

func NewFilterRows(f func([]string) bool) *FilterRows {
	return &FilterRows{filterFunc: f}
}

func (f *FilterRows) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		if f.filterFunc(row) {
			res = append(res, row)
		}
	}
	return res
}

// SumRow operation
type SumRow struct{}

func NewSumRow() *SumRow {
	return &SumRow{}
}

func (s *SumRow) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		sum := 0
		for _, cell := range row {
			num, _ := strconv.Atoi(cell)
			sum += num
		}
		res = append(res, []string{strconv.Itoa(sum)})
	}
	return res
}

// DuplicateRow operation
type DuplicateRow struct{}

func NewDuplicateRows() *DuplicateRow {
	return &DuplicateRow{}
}

func (d *DuplicateRow) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		res = append(res, row, row)
	}
	return res
}
