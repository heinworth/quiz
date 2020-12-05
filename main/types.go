package main

type Quiz struct {
	ID int
	Questions []Question
}

type Question struct {
	QuizID int
	Options [4]string
	Answer string
}

type CompletedQuiz struct {
	Quiz
	CompletedByEmail string
	Score int
}