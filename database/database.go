package database

import "../quiz"
/*
	Contains functions and types relating to database operations.

*/

/* 
	Provices a framework to mock database access.
	This is useful for testing API functions before the DB has been implemented,
	and allows testing without touching the database before we have set up a system
	to use a different database in tests, for example.
*/
type DB interface{
	SaveQuiz(q quiz.Quiz) error
	SubmitAnswers() error
	GetScore() error
	PublishQuiz() error
	EditQuiz() error
	GetCompletedByUser() error
}

type MockDB struct{}
type DBImplementation struct{}

func (m MockDB) SaveQuiz(q quiz.Quiz) error {
	return nil
}

func (m MockDB) SubmitAnswers(q quiz.Quiz) error {
	return nil
}

func (m MockDB) GetScore(email string, quizID int) (score, maxScore int) {
	return 0, 0
}

func (m MockDB) PublishQuiz(QuizID int) error {
	return nil
}

func (m MockDB) EditQuiz(q quiz.Quiz) error {
	return nil
}

func (m MockDB) GetCompletedByUser(email string) []quiz.Completed {
	return nil
} 

func (db DBImplementation) SaveQuiz(q quiz.Quiz) error {
	return nil
}

func (db DBImplementation) SubmitAnswers(q quiz.Quiz) error {
	return nil
}

func (db DBImplementation) GetScore(email string, quizID int) (score, maxScore int) {
	return 0, 0
}

func (db DBImplementation) PublishQuiz(QuizID int) error {
	return nil
}

func (db DBImplementation) EditQuiz(q quiz.Quiz) error {
	return nil
}

func (db DBImplementation) GetCompletedByUser(email string) []quiz.Completed {
	return nil
} 

