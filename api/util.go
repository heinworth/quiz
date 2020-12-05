package api

import "../quiz"

/*
	Helper functions that are not part of the external API.
*/

func calculateScore(quiz quiz.Completed) (score int) {
	for _, question := range quiz.Questions {
		if question.CorrectAnswer == question.ChosenAnswer {
			score++
		}
	}
	return score
}

func containsString(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
