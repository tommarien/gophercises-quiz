package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "input csv file")

	flag.Parse()

	problems, err := readProblems(*csvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	quiz := NewQuiz(problems)

	quiz.Run()

	quiz.PrintScore()
}

func readProblems(csvFile string) ([]Problem, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("could not read from file: %w", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not parse file as csv: %w", err)
	}

	problems := make([]Problem, len(rows))

	for i, row := range rows {
		if len(row) != 2 {
			return nil, fmt.Errorf("row %d does not have 2 colums", i)
		}

		problems[i] = Problem{
			Question: row[0],
			Answer:   row[1],
		}
	}

	return problems, nil
}
