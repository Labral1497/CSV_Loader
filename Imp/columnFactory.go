package imp

// GetColumn operation
type GetColumn struct {
	index int
}

func NewGetColumn(index int) *GetColumn {
	return &GetColumn{index: index}
}

func (g *GetColumn) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		if g.index < len(row) {
			res = append(res, []string{row[g.index]})
		}
	}
	return res
}

// ForEveryColumn operation
type ForEveryColumn struct {
	columnFunc func(string) string
}

func NewForEveryColumn(f func(string) string) *ForEveryColumn {
	return &ForEveryColumn{columnFunc: f}
}

func (f *ForEveryColumn) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		var newRow []string
		for _, cell := range row {
			newRow = append(newRow, f.columnFunc(cell))
		}
		res = append(res, newRow)
	}
	return res
}
