package main

import (
	"fmt"

	simpleQuiz "github.com/panvalkar1994/simpleQuiz/quiz"
)

func main() {
	var user simpleQuiz.User
	user.Reset()
	fmt.Println(simpleQuiz.Welcome)
	fmt.Println("Enter Your Username")
	fmt.Scanln(&user.Name)
	fmt.Print(simpleQuiz.GetIntroduction(user.Name))
	var option int
	quiz, err := simpleQuiz.LoadQuiz("questions.csv")
	if err != nil {
		panic(err)
	}
	fmt.Print(simpleQuiz.GetMenu("Start"))
	fmt.Scanln(&option)
	// TODO: Refactoring this
	for {
		if option == 2 {
			fmt.Println("Exiting...")
			return
		} else if option == 0 {
			fmt.Print(simpleQuiz.GetIntroduction(user.Name))
			fmt.Print(simpleQuiz.GetMenu("Start"))
			fmt.Scanln(&option)
		} else {
			for {
				if quiz.End() {
					user.Score()
					fmt.Print(simpleQuiz.GetMenu("Re-start"))
					fmt.Scanln(&option)
					if option == 1 {
						quiz.Reset()
						user.Reset()
						continue
					}
					break
				} else {
					quiz.AskQuestion(&user)
				}
			}
		}
	}

}
