# CSV_Loader
This repository contains a Go implementation of a CSV processing pipeline. The pipeline allows for various operations to be applied to each row of a CSV file, including filtering rows, extracting columns, performing mathematical operations, and accumulating results. We assume that each test or process will include only one accumulator operation, and that string to integer conversions will always succeed without errors.

Files
main.go
The entry point of the application. This file initializes the CSVLoader, adds various operations, and runs the processing pipeline. Example usage is provided to demonstrate how to set up and execute the pipeline.

MathFactory.go
Contains the implementation of mathematical operations that can be applied to CSV rows. This includes operations such as summing row values and calculating the average. These operations are designed to work seamlessly within the processing pipeline.

columnFactory.go
Provides functions for operations related to column manipulation. This includes extracting specific columns from rows and performing operations on column values. The functions in this file are used to customize the processing of CSV data by focusing on specific columns.

AccumulatorFactory.go
Implements accumulator operations which are special types of operations that need to run after all rows have been processed. Examples include summing column values across all rows. Only one accumulator operation is assumed per test or process.

RowFactory.go
Contains operations for filtering and modifying rows based on certain criteria. This includes functions to filter rows based on row indices or custom conditions. These operations are used to selectively process rows in the CSV file.

Process.go
Defines the core CSVLoader struct and its methods, including With, Write, and Run. The CSVLoader manages the sequence of operations to be applied to the CSV data and ensures proper execution of these operations, including handling of the accumulator operation.

Assumptions
Single Accumulator Operation: Each test or process will include only one accumulator operation.
Error-Free Conversions: Conversions from strings to integers will always succeed without errors.


Running Instructions:
Enter on your terminal "go run main.go" and watch as 3 output files are produced by 3 tests in your directory.
