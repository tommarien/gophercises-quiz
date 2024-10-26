package main

import "fmt"

type Problem struct {
	Question, Answer string
}

type Quiz struct {
	score    uint
	problems []Problem
}

func NewQuiz(problems []Problem) Quiz {
	return Quiz{
		problems: problems,
	}
}

func (q *Quiz) Run() {
	for i, p := range q.problems {
		fmt.Printf("Problem %2d: %s: ", i+1, p.Question)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.Answer {
			q.score++
		}
	}
}

func (q *Quiz) PrintScore() {
	fmt.Printf("Your score is %d/%d\n", q.score, len(q.problems))
}
