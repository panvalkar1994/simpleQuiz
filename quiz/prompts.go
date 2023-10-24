package quiz

import "fmt"

const Welcome = "Welcome to the Quiz Game"

const Introduction = `Hi %s, I am a simple Quiz Game written in Go.
I am a part of the Gophercises series of exercises.
I am a command line application that asks you a series of questions from a CSV file and lets you answer them.
I will tell you how many questions you got right and how many you got wrong at the end of the quiz.
`
const Menu = `Select an option:
0. Print the introduction
1. %s the quiz
2. Exit the quiz
3. Show Menu
`

func GetMenu(option string) string {
	return fmt.Sprintf(Menu, option)
}

func GetIntroduction(user string) string {
	return fmt.Sprintf(Introduction, user)
}
