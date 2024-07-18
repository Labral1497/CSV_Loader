package imp

import (
	"fmt"
	"strconv"
)

var Count = 0
var Sum = 0.0

// Avg Operation
type Avg struct{}

func NewAvg() *Avg {
	return &Avg{}
}

func (a *Avg) Apply(rows [][]string) [][]string {
	for _, row := range rows {
		if row != nil {
			result, _ := strconv.ParseFloat(row[0], 64)
			Count++
			Sum += result
		}
	}
	return nil // Return nil to skip writing the row
}

func (a *Avg) Final() []string {
	if Count > 0 {
		avg := Sum / float64(Count)
		return []string{fmt.Sprintf("%.2f", avg)}
	}
	return nil
}

// SumCol Operation
type SumCol struct {
	sum int
}

func NewSumCol() *SumCol {
	return &SumCol{}
}

func (s *SumCol) Apply(rows [][]string) [][]string {
	for _, row := range rows {
		if row != nil && len(row) > 0 {
			result, _ := strconv.Atoi(row[0])
			s.sum += result
		}
	}
	return nil
}

func (s *SumCol) Final() []string {
	return []string{strconv.Itoa(s.sum)}
}
