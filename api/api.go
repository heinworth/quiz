package api

import (
	"../quiz",
	"../database"
)

/**
	Represents the quiz's external API. 
	These functions are intended to be called by HTTP handler functions after request data has been 
	parsed and validated. When the HTTP API is served from this package, these functions should 
	be unexported.
*/


var DB = database.MockDB{}

func CreateQuiz(q quiz.Quiz) {

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
