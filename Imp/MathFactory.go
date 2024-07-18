package imp

import (
	"math"
	"strconv"
)

// Ceil Operation
type Ceil struct{}

func NewCeil() *Ceil {
	return &Ceil{}
}

func (c *Ceil) Apply(rows [][]string) [][]string {
	var res [][]string
	for _, row := range rows {
		result, _ := strconv.ParseFloat(row[0], 64)
		res = append(res, []string{strconv.Itoa(int(math.Ceil(result)))})
	}
	return res
}
