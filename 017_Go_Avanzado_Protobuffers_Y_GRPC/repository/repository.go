package repository

import (
	"database/sql"
	"fmt"
	"github.com/jsovalles/grpc/models"
	"github.com/jsovalles/grpc/utils"
)

type Repository interface {
	GetStudent(id string) (student models.Student, err error)
	SetStudent(student models.Student) (err error)
	GetTest(id string) (test models.Test, err error)
	SetTest(test models.Test) (err error)
	SetQuestion(question models.Question) (err error)
	SetEnrollment(enrollment models.Enrollment) (err error)
	GetStudentsPerTest(testId string) (students []models.Student, err error)
	GetQuestionsPerTest(testId string) (questions []models.Question, err error)
}

type repository struct {
	env utils.Env
	db  utils.Database
}

func (u *repository) GetStudent(id string) (student models.Student, err error) {
	err = u.db.Client.Get(&student, getStudentById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func (u *repository) SetStudent(student models.Student) (err error) {
	_, err = u.db.Client.NamedExec(createStudent, student)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (u *repository) GetTest(id string) (test models.Test, err error) {
	err = u.db.Client.Get(&test, getTestById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func (u *repository) SetTest(test models.Test) (err error) {
	_, err = u.db.Client.NamedExec(createTest, test)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (u *repository) SetQuestion(question models.Question) (err error) {
	_, err = u.db.Client.NamedExec(createQuestions, question)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (u *repository) SetEnrollment(enrollment models.Enrollment) (err error) {
	_, err = u.db.Client.NamedExec(createEnrollment, enrollment)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (u *repository) GetStudentsPerTest(testId string) (students []models.Student, err error) {
	err = u.db.Client.Select(&students, getStudentPerTest, testId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func (u *repository) GetQuestionsPerTest(testId string) (questions []models.Question, err error) {
	err = u.db.Client.Select(&questions, getQuestionsPerTest, testId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func NewRepository(env utils.Env, db utils.Database) Repository {
	return &repository{env: env, db: db}
}

const (
	studentsTable       = "students"
	testTable           = "tests"
	questionTable       = "questions"
	enrollmentTable     = "enrollments"
	createStudent       = "INSERT INTO " + studentsTable + " (id, name, age) VALUES(:id, :name, :age)"
	getStudentById      = "SELECT * FROM " + studentsTable + " WHERE id = $1"
	createTest          = "INSERT INTO " + testTable + " (id, name) VALUES(:id, :name)"
	getTestById         = "SELECT * FROM " + testTable + " WHERE id = $1"
	createQuestions     = "INSERT INTO " + questionTable + " (id, answer, question, test_id) VALUES(:id, :answer, :question, :test_id)"
	getStudentPerTest   = "SELECT * FROM " + studentsTable + " WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)"
	createEnrollment    = "INSERT INTO " + enrollmentTable + " (student_id, test_id) VALUES(:student_id, :test_id)"
	getQuestionsPerTest = "SELECT * FROM " + questionTable + " WHERE test_id = $1"
)
