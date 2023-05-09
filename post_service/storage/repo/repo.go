package repo

import (
	p "github.com/PentaGol/post_service/genproto/post"
)

type PostStorageI interface {
	CreatePost(*p.PostRequest) (*p.PostResponse, error)
	GetPostById(*p.IdRequest) (*p.PostResponse, error)
	GetAllPosts(*p.AllPostRequest)(*p.Posts, error)
	SearchByTitle(*p.Search) (*p.Posts, error)
	UpdatePost(*p.UpdatePostRequest) error
	DeletePost(*p.IdRequest) (*p.PostResponse, error)
	GetNews(*p.AllPostRequest)(*p.Posts, error)
}
