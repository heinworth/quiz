package api

import (
	"../quiz"
	"../database"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuiz(t *testing.T) {
	for _, test := range []struct{
		quiz quiz.Quiz
		shouldError bool
		name string
	} {
		{
			name: "Should not allow duplicate options",
			quiz: quiz.Quiz{
				Questions: []quiz.Question{
					quiz.Question{
						Options: [4]string{
							"option 1",
							"option 1",
							"option 1",
							"option 1",
						},
					},
				},
			},
			shouldError: true,
		},
		{
			name: "Should allow 4 different options",
			quiz: quiz.Quiz{
				Questions: []quiz.Question{
					quiz.Question{
						Options: [4]string{
							"option 1",
							"option 2",
							"option 3",
							"option 4",
						},
					},
				},
			},
			shouldError: false,
		},
		{
			name: "Should not be able to save a published quiz",
			quiz: quiz.Quiz{
				Questions: []quiz.Question{
					quiz.Question{
						Options: [4]string{
							"option 1",
							"option 2",
							"option 3",
							"option 4",
						},
					},
				},
				IsPublished: true,
			},
			shouldError: true,
		},
	} {
		err := SaveQuiz(database.MockDB{}, test.quiz)
		assert.Equal(t, test.shouldError, err != nil)
	}
}

func TestSubmitAnswers(t *testing.T) {
	DB := database.MockDB{}
	DB.SetAlreadyCompleted([]quiz.Completed {
		quiz.Completed{
			Quiz: quiz.Quiz{ID: 1},
		},
	})
	request := quiz.Completed {
		Quiz: quiz.Quiz{ID: 1},
	}
	assert.NotNil(t, SubmitAnswers(DB, request))

	DB.SetAlreadyCompleted(nil)
	assert.Nil(t, SubmitAnswers(DB, request))
}

func testCalculateScore(t *testing.T) {
	var completed quiz.Completed

	completed.Questions = []quiz.Question{
		quiz.Question{
			CorrectAnswer: "foobar",
			ChosenAnswer: "foobar",
		},
	}

	assert.Equal(t, 1, calculateScore(completed))

	completed.Questions = []quiz.Question{
		quiz.Question{
			CorrectAnswer: "foobar",
			ChosenAnswer: "",
		},
	}

	assert.Equal(t, 0, calculateScore(completed))
}


