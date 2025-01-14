package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// type Result struct {
// 	RollNumber string
// 	Score1     int
// 	Score2     int
// 	Score3     int
// }

type Result struct {
	RollNumber string `json:"rollnumber"`
	Score1     int    `json:"score1"`
	Score2     int    `json:"score2"`
	Score3     int    `json:"score3"`
}

// Just Marshal
// func writeJSON(results []Result) {
// 	jd, err := json.Marshal(results)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	f, err := os.Create("results.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// 	_, err = f.Write(jd)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// Marshal and Indent separately
// func writeJSON(results []Result) {
// 	jd, err := json.Marshal(results)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var ijd bytes.Buffer
// 	json.Indent(&ijd, jd, "", "\t")
// 	// fmt.Println(ijd.String())

// 	f, err := os.Create("results.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// 	_, err = f.Write(ijd.Bytes())
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// Marshal and Indent together
// func writeJSON(results []Result) {
// 	jd, err := json.MarshalIndent(results, "", "\t")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// fmt.Println(string(jd))

// 	f, err := os.Create("results.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// 	_, err = f.Write(jd)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// With encoder
func writeJSON(results []Result) {
	f, err := os.Create("results.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	enc.Encode(results)
}

// With Decoder
// func readJSON() []Result {
// 	f, err := os.Open("results.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer f.Close()
// 	var results []Result
// 	err = json.NewDecoder(f).Decode(&results)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return results
// }

// With Unmarshal
func readJSON() []Result {
	content, err := os.ReadFile("results.json")
	if err != nil {
		log.Fatalln(err)
	}
	var results []Result
	json.Unmarshal(content, &results)
	return results
}

func main() {
	results := []Result{
		{RollNumber: "A001", Score1: 47, Score2: 92, Score3: 93},
		{RollNumber: "A002", Score1: 86, Score2: 65, Score3: 94},
		{RollNumber: "A003", Score1: 53, Score2: 96, Score3: 98},
		{RollNumber: "A004", Score1: 91, Score2: 92, Score3: 93},
	}

	writeJSON(results)

	results1 := readJSON()

	fmt.Println(results1)
}
