package quiz

import "fmt"

type User struct {
	TotalQuestions int
	CorrectAnswers int
	WrongAnswers   int
	Name           string
}

func (u *User) Score() {
	fmt.Printf("%s, you scored %d/%d and %d missed\n", u.Name, u.CorrectAnswers, u.TotalQuestions, u.WrongAnswers)
}

func (u *User) UpdateScore(correct bool) {
	if correct {
		u.CorrectAnswers++
	} else {
		u.WrongAnswers++
	}
	u.TotalQuestions++
}

func (u *User) Reset() {
	u.TotalQuestions = 0
	u.CorrectAnswers = 0
	u.WrongAnswers = 0
}
