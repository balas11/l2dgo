package main

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Result struct {
	RollNumber string
	Score1     int
	Score2     int
	Score3     int
	Average    float32
}

func readCSV(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func process(records [][]string) []Result {
	length := len(records) - 1
	results := make([]Result, 0, length)

	for _, record := range records[1:] {
		var r Result
		r.RollNumber = record[0]
		r.Score1, _ = strconv.Atoi(strings.TrimSpace(record[1]))
		r.Score2, _ = strconv.Atoi(strings.TrimSpace(record[2]))
		r.Score3, _ = strconv.Atoi(strings.TrimSpace(record[3]))
		r.Average = (float32(r.Score1) + float32(r.Score2) + float32(r.Score3)) / 3
		results = append(results, r)
	}
	return results
}

func writeCSV(results []Result, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"RollNumber", "Score1", "Score2", "Score3", "Average"})
	for _, result := range results {
		record := []string{
			result.RollNumber,
			strconv.Itoa(int(result.Score1)),
			strconv.Itoa(int(result.Score2)),
			strconv.Itoa(int(result.Score3)),
			strconv.FormatFloat(float64(result.Average), 'f', 2, 32),
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	records, err := readCSV("scores.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	results := process(records)

	slices.SortFunc(results, func(a, b Result) int {
		if n := cmp.Compare(a.Average, b.Average); n != 0 {
			return n * -1
		}
		return strings.Compare(a.RollNumber, b.RollNumber)
	})

	err = writeCSV(results, "results.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(results)
}
