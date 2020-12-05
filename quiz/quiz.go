package quiz

/*
	Contains types core to the app's functionality
*/

//Represents a Quiz and its questions.
type Quiz struct {
	ID int
	Questions []Question
	IsPublished bool
}

// Each question is linked to a quiz, and must have 4 options. Answer must 
// equal one of these options.
type Question struct {
	QuizID int
	Options [4]string
	CorrectAnswer string
	ChosenAnswer string // answer chosen by the user
}

// Represents a quiz completed by a user.
type Completed struct {
	Quiz
	CompletedByEmail string
	Score int
}