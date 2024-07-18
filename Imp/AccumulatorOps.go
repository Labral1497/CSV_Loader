package imp

import (
	"fmt"
	"strconv"
)

// Avg Operation
type Avg struct{
	sum float64
	count int
}

func NewAvg() *Avg {
	return &Avg{}
}

func (a *Avg) Apply(rows [][]string) [][]string {
	for _, row := range rows {
		if row != nil {
			result, _ := strconv.ParseFloat(row[0], 64)
			a.count++
			a.sum += result
		}
	}
	return nil // Return nil to skip writing the row
}

func (a *Avg) Final() []string {
	if a.count > 0 {
		avg := a.sum / float64(a.count)
		return []string{fmt.Sprintf("%.2f", avg)}
	}
	return nil
}

// SumCol Operation
type SumCol struct {
	sum float64
}

func NewSumCol() *SumCol {
	return &SumCol{}
}

func (s *SumCol) Apply(rows [][]string) [][]string {
	for _, row := range rows {
		if row != nil && len(row) > 0 {
			result, _ := strconv.ParseFloat(row[0], 64)
			s.sum += result
		}
	}
	return nil
}

func (s *SumCol) Final() []string {
	return []string{fmt.Sprintf("%.2f", s.sum)}
}
