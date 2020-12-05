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
	Calls DB function to save quiz to database, after validating its questions.
	Can be used as a POST or PUT function.
*/
func SaveQuiz(DB database.DB, q quiz.Quiz) error {
	if q.IsPublished {
		return errors.New("Cannot save a published quiz")
	}
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
	return DB.SaveScore(quiz.ID, calculateScore(quiz))
}

func GetScore(email string, quizID int) (score, maxScore int, err error){
	return database.DBImplementation{}.GetScore(email, quizID)
}

func PublishQuiz(DB database.DB, q quiz.Quiz) error {
	q.IsPublished = true
	return SaveQuiz(DB, q)
}

func GetCompletedByUser(email string) ([]quiz.Completed, error) {
	return 	database.DBImplementation{}.GetCompletedByUser(email)
} 

func DeleteQuiz(quizID int64) error {
	return nil
}