package quiz

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	questions []Question
	position  int
}

func (q *Quiz) Reset() {
	q.position = 0
	q.shuffle()
}

func LoadQuiz(filePath string) (*Quiz, error) {
	fmt.Println("Loading Quiz...")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error in Opening file", err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error with reading csv")
		return nil, err
	}
	quiz := new(Quiz)
	quiz.questions = []Question{}
	for _, record := range records {
		question := Question{
			Question: record[0],
			Answer:   record[1],
		}
		quiz.questions = append(quiz.questions, question)
	}

	return quiz, nil
}

func (q *Quiz) Next() error {
	if q.End() {
		return fmt.Errorf("all questions exhuasted")
	}
	q.position++
	return nil
}

func (q *Quiz) End() bool {
	return len(q.questions) <= q.position
}

func (q *Quiz) GetQuestion() (*Question, error) {
	if q.End() {
		return nil, fmt.Errorf("all questions exhuasted %d/%d", q.position, q.position)
	}
	question := q.questions[q.position]
	return &question, nil
}

func (q *Quiz) AskQuestion(u *User) {
	var answer string
	ques, err := q.GetQuestion()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Question> %s\n", ques.Question)
	fmt.Scanln(&answer)
	correct := ques.Answer == answer
	u.UpdateScore(correct)
	q.Next()
}

func (q *Quiz) shuffle() {
	// shuffle the questions
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range q.questions {
		j := random.Intn(len(q.questions))
		q.questions[i], q.questions[j] = q.questions[j], q.questions[i]
	}

}
