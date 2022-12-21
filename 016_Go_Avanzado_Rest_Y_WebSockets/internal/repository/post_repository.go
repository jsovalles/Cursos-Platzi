package repository

import (
	"database/sql"
	"fmt"
	"github.com/jsovalles/rest-ws/internal/models"
	"github.com/jsovalles/rest-ws/internal/utils"
)

type PostRepository interface {
	CreatePost(post models.Post) (err error)
	GetPostById(id string) (user models.Post, err error)
	UpdatePost(post models.Post) (err error)
	DeletePostByIdAndUserId(id string, userId string) (err error)
	ListPosts(page uint64) (posts []models.Post, err error)
}

type postRepository struct {
	env utils.Env
	db  utils.Database
}

func NewPostRepository(env utils.Env, db utils.Database) PostRepository {
	return &postRepository{env: env, db: db}
}

const (
	postsTable  = "posts"
	createPost  = "INSERT INTO " + postsTable + " (id, post_content, user_id) VALUES (:id, :post_content, :user_id)"
	getPostById = "SELECT * FROM " + postsTable + " WHERE id = $1"
	updatePost  = "UPDATE " + postsTable + " SET post_content = :post_content WHERE id = :id and user_id = :user_id"
	deletePost  = "DELETE FROM " + postsTable + " WHERE id = $1 and user_id = $2"
	listPosts   = "SELECT * FROM " + postsTable + " limit $1 OFFSET $2"
)

func (p *postRepository) CreatePost(post models.Post) (err error) {
	_, err = p.db.Client.NamedExec(createPost, post)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (p *postRepository) GetPostById(id string) (post models.Post, err error) {
	err = p.db.Client.Get(&post, getPostById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this post, please validate")
			return
		}
		return
	}
	return
}

func (p *postRepository) UpdatePost(post models.Post) (err error) {
	_, err = p.db.Client.NamedExec(updatePost, post)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (p *postRepository) DeletePostByIdAndUserId(id string, userId string) (err error) {
	_, err = p.db.Client.Exec(deletePost, id, userId)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	return
}

func (p *postRepository) ListPosts(page uint64) (posts []models.Post, err error) {
	err = p.db.Client.Select(&posts, listPosts, 2, page*2)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this post, please validate")
			return
		}
		return
	}
	return
}
