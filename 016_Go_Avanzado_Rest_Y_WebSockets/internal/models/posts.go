package models

import "time"

type Post struct {
	Id          string    `json:"id" db:"id"`
	PostContent string    `json:"postContent" db:"post_content"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UserId      string    `json:"userId" db:"user_id"`
}

type CreatePostRequest struct {
	PostContent string `json:"postContent"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}
