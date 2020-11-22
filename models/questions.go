package models

type Question struct {
	ID       uint
	Question string
}

func (Question) TableName() string {
	return "project_random_questions.questions"
}
