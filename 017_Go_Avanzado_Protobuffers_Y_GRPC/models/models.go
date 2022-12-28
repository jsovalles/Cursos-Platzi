package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id       string `json:"id" db:"id"`
	Question string `json:"question" db:"question"`
	Answer   string `json:"answer" db:"answer"`
	TestId   string `json:"test_id" db:"test_id"`
}

type Enrollment struct {
	StudentId string `json:"student_id" db:"student_id"`
	TestId    string `json:"test_id" db:"test_id"`
}
