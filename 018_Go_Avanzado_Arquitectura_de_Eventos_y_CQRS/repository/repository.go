package repository

import (
	"database/sql"
	"fmt"
	"github.com/jsovalles/cqrs/models"
	"github.com/jsovalles/cqrs/utils"
)

type Repository interface {
	InsertFeed(feed models.Feed) (err error)
	ListFeeds(feeds []models.Feed, err error)
}

type repository struct {
	env utils.Env
	db  utils.Database
}

func (r *repository) InsertFeed(feed models.Feed) (err error) {
	_, err = r.db.Client.NamedExec(createFeed, feed)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (r *repository) ListFeeds(feeds []models.Feed, err error) {
	err = r.db.Client.Select(&feeds, listFeeds)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

//func (u *repository) SetEnrollment(enrollment models.Enrollment) (err error) {
//	_, err = u.db.Client.NamedExec(createEnrollment, enrollment)
//	if err != nil {
//		fmt.Errorf(err.Error())
//		return
//	}
//	return
//}
//
//func (u *repository) GetStudentsPerTest(testId string) (students []models.Student, err error) {
//	err = u.db.Client.Select(&students, getStudentPerTest, testId)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			err = fmt.Errorf("there are no results for this user, please validate")
//			return
//		}
//		return
//	}
//	return
//}
//
//func (u *repository) GetQuestionsPerTest(testId string) (questions []models.Question, err error) {
//	err = u.db.Client.Select(&questions, getQuestionsPerTest, testId)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			err = fmt.Errorf("there are no results for this user, please validate")
//			return
//		}
//		return
//	}
//	return
//}

func NewRepository(env utils.Env, db utils.Database) Repository {
	return &repository{env: env, db: db}
}

const (
	feedTable  = "feeds"
	createFeed = "INSERT INTO " + feedTable + " (id, title, description) VALUES(:id, :title, :description)"
	listFeeds  = "SELECT * FROM " + feedTable
)
