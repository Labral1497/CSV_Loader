package processing

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

var RowCounter = 0

type CSVLoader struct {
	filePath           string
	operations         []Operation
	accumulatorOp      AccumulatorOp
	postAccumulatorOps []Operation
	outputPath         string
}

func NewCsvLoader(filePath string) *CSVLoader {
	return &CSVLoader{filePath: filePath, accumulatorOp: nil}
}

type Operation interface {
	Apply(rows [][]string) [][]string
}

// AccumulatorOp interface for operations that need to run after all rows are processed
type AccumulatorOp interface {
	Operation
	Final() []string
}

func (c *CSVLoader) With(operation Operation) *CSVLoader {
    // Once we meet an accumulator op we guaranteed than no accumulator op will come afterwards
	if accumulatorOp, ok := operation.(AccumulatorOp); ok {
		c.accumulatorOp = accumulatorOp
		c.operations = append(c.operations, operation)
	} else if c.accumulatorOp != nil {
		c.postAccumulatorOps = append(c.postAccumulatorOps, operation)
	} else {
		c.operations = append(c.operations, operation)
	}

	return c
}

func (c *CSVLoader) Write(outputPath string) *CSVLoader {
	c.outputPath = outputPath
	return c
}


func (c *CSVLoader) Run() error {
	if c.filePath == "" || c.outputPath == "" {
		return errors.New("file paths not specified")
	}

	inputFile, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	outputFile, err := os.Create(c.outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        RowCounter++
        c.processRows(row, writer, false /* postAccumulatorFlag */)
    }

    // Process the Accumulator op and post accumulator operations
    if c.accumulatorOp != nil {
        row := c.accumulatorOp.Final()
        c.processRows(row, writer, true /* postAccumulatorFlag */)
    }

    RowCounter = 0
    return nil
}

func (c *CSVLoader) processRows(row []string, writer *csv.Writer, postAccumulatorFlag bool) {
    rows := [][]string{row}
    operations := c.operations
    if postAccumulatorFlag {
        operations = c.postAccumulatorOps
    }
    for _, op := range operations {
        rows = op.Apply(rows)
    }
    for _, row := range rows {
        if row == nil {
            continue
        }
        if err := writer.Write(row); err != nil {
            panic(err)
        }
    }
}