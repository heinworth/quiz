package database

import "../quiz"

/*
	Contains functions and types relating to database operations.

*/

/* 
	Provides a framework to mock database access.
	This is useful for testing API functions before the DB has been implemented,
	and allows testing without touching the database before we have set up a system
	to use a different database in tests, for example.
*/
type DB interface{
	SaveQuiz(q quiz.Quiz) error
	SaveScore(quizID int, score int) error
	GetScore(email string, quizID int) (score, maxScore int, err error)
	PublishQuiz(quizID int) error
	EditQuiz(q quiz.Quiz) error
	GetCompletedByUser(email string) ([]quiz.Completed, error)
}

type MockDB struct{
	alreadyCompleted []quiz.Completed
}
type DBImplementation struct{}

func (m MockDB) SaveQuiz(q quiz.Quiz) error {
	return nil
}

func (m MockDB) SaveScore(quizID int, score int) error {
	return nil
}

func (m MockDB) GetScore(email string, quizID int) (score, maxScore int, err error) {
	return
}

func (m MockDB) PublishQuiz(QuizID int) error {
	return nil
}

func (m MockDB) EditQuiz(q quiz.Quiz) error {
	return nil
}

func (m MockDB) GetCompletedByUser(email string) ([]quiz.Completed, error) {
	return m.alreadyCompleted, nil
} 

func (m *MockDB) SetAlreadyCompleted(completed []quiz.Completed) {
	m.alreadyCompleted = completed
}

func (db DBImplementation) SaveQuiz(q quiz.Quiz) error {
	//TODO: implement database
	return nil
}

func (db DBImplementation) SaveScore(quizID int, score int) error {
	//TODO: implement database
	return nil
}

func (db DBImplementation) GetScore(email string, quizID int) (score, maxScore int, err error) {
	//TODO: implement database
	return
}

func (db DBImplementation) PublishQuiz(QuizID int) error {
	//TODO: implement database
	return nil
}

func (db DBImplementation) EditQuiz(q quiz.Quiz) error {
	//TODO: implement database
	return nil
}

func (db DBImplementation) GetCompletedByUser(email string) ([]quiz.Completed, error) {
	//TODO: implement database
	return nil, nil
} 
