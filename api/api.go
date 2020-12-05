package api

import (
	"../quiz"
	"../database"
	"fmt"
	errors "github.com/pkg/errors"
)

/**
	Represents the quiz's external API. 
	These functions are intended to be called by HTTP handler functions after request data has been 
	parsed and validated. When the HTTP API is served from this package, these functions should 
	be unexported.
*/

/*
	Created a new quiz and saves it to database.
*/
func CreateQuiz(DB database.DB, q quiz.Quiz) error {
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

/*
	Submit answers to a quiz. Errors if user has already completed this quiz before.
*/
func SubmitAnswers(DB database.DB, quiz quiz.Completed) error {
	alreadyCompleted, err := DB.GetCompletedByUser(quiz.CompletedByEmail)
	if err != nil {
		return errors.Wrap(err, "DB error")
	} 		
	for _, v := range alreadyCompleted {
		if v.ID == quiz.ID {
			return errors.New("User has already completed this quiz")
		}
	}
	return nil
}

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
