package api

/**
	Represents the quiz's external API. 
	These functions are intended to be called by HTTP handler functions after request data has been 
	parsed and validated. When the HTTP API is served from this package, these functions should 
	be unexported.
*/

func CreateQuiz(q Quiz) {}
func SubmitAnswers() {}
func GetScore(email string, quizID int) (score int, outOf int){
	return 0, 0
}
func PublishQuiz(quizID int) {}
func EditQuiz(q Quiz) {}
func GetCompletedByUser(email string) []Quiz {
	return nil
} 
