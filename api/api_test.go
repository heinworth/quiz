package api

import (
	"../quiz"
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
	} {
		err := CreateQuiz(test.quiz)
		assert.Equal(t, test.shouldError, err != nil)
	}
}


