package main

import(
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"time"
	
)

func main()  {
	csvFilename := flag.String("csv", "quiz.csv", "a csv file of question and answer")
	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s\n", *csvFilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file")
	}
	now := time.Now()
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems{
		fmt.Printf("problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a{
			fmt.Println("correct!")
			correct++
		}
	}
	t := time.Now()
	elapsed := t.Sub(now)
	fmt.Printf("you used %v seconds \n", int(elapsed))
	fmt.Printf("you scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
	
}
type problem struct{
	q string
	a string
}
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}