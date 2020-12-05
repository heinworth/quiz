package api

import (
	"../quiz"
	"../database"
	"fmt"
	"errors"
)

/**
	Represents the quiz's external API. 
	These functions are intended to be called by HTTP handler functions after request data has been 
	parsed and validated. When the HTTP API is served from this package, these functions should 
	be unexported.
*/

var DB = database.MockDB{}

func CreateQuiz(q quiz.Quiz) error {
	for _, question := range q.Questions {
		var foundOptions []string
		for _, option := range question.Options {
			if containsString(option, foundOptions){
				return errors.New(fmt.Sprintf("Duplicate option detected: %s", option))
			}
			foundOptions = append(foundOptions, option)
		}
	}
	return DB.SaveQuiz(q)
}
func SubmitAnswers() {}
func GetScore(email string, quizID int) (score, maxScore int){
	return 0, 0
}
func PublishQuiz(quizID int) {}
func EditQuiz(q quiz.Quiz) {}
func GetCompletedByUser(email string) []quiz.Completed {
	return nil
} 

func containsString(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
