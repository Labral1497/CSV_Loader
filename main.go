package main

import (
	"coralogx_EX/imp"
	"coralogx_EX/processing"
	"fmt"
	"strconv"
)

const (
    InputFile   = "input.csv"
    OutputFile1 = "output1.csv"
    OutputFile2 = "output2.csv"
    OutputFile3 = "output3.csv"
)

func main() {
    runTest1()
    runTest2()
    runTest3()
}

func runTest1() {
    err := processing.NewCsvLoader(InputFile).
        With(imp.NewDuplicateRows()).
        With(imp.NewGetRows(3, 5)).
        With(imp.NewGetColumn(3)).
        With(imp.NewAvg()).
        With(imp.NewCeil()).
        Write(OutputFile1).
        Run()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Processing %s is complete\n", InputFile)
    }
}

func runTest2() {
    err := processing.NewCsvLoader(InputFile).
        With(imp.NewGetRows(3, 5)).
        With(imp.NewGetColumn(3)).
        With(imp.NewForEveryColumn(func(cell string) string {
            n, _ := strconv.Atoi(cell)
            return strconv.Itoa(n*2)
        })).
        Write(OutputFile2).
        Run()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Processing %s is complete\n", InputFile)
    }
}

func runTest3() {
    err := processing.NewCsvLoader(InputFile).
    With(imp.NewFilterRows(func(record []string) bool {
        return record[2] == "Carlos Soltero"
    })).
    With(imp.NewGetColumn(3)).
    With(imp.NewSumCol()).
    Write(OutputFile3).
    Run()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Processing %s is complete\n", InputFile)
    }
}
