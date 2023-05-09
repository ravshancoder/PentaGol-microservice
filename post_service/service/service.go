package service

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	p "github.com/PentaGol/post_service/genproto/post"
	"github.com/PentaGol/post_service/pkg/logger"
	grpcclient "github.com/PentaGol/post_service/service/grpc_client"
	"github.com/PentaGol/post_service/storage"
)

type PostService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpcclient.Clients
}

func NewPostService(db *sqlx.DB, log logger.Logger, client grpcclient.Clients) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
		Client:  client,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *p.PostRequest) (*p.PostResponse, error) {
	res, err := s.storage.Post().CreatePost(req)
	if err != nil {
		log.Println("failed to create post: ", err)
		return &p.PostResponse{}, err
	}

	return res, nil
}

func (s *PostService) GetPostById(ctx context.Context, req *p.IdRequest) (*p.PostResponse, error) {
	res, err := s.storage.Post().GetPostById(req)
	if err != nil {
		log.Println("failed to get post by id: ", err)
		return &p.PostResponse{}, err
	}

	return res, nil
}

func (s *PostService) GetAllPosts(ctx context.Context, req *p.AllPostRequest) (*p.Posts, error) {
	res, err := s.storage.Post().GetAllPosts(req)
	if err != nil {
		log.Println("failed to get all post: ", err)
		return &p.Posts{}, err
	}

	return res, nil
}

func (s *PostService) SearchByTitle(ctx context.Context, req *p.Search) (*p.Posts, error) {
	res, err := s.storage.Post().SearchByTitle(req)
	if err != nil {
		log.Println("failed to get post by search title: ", err)
		return &p.Posts{}, err
	}

	return res, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *p.UpdatePostRequest) (*p.PostResponse, error) {
	err := s.storage.Post().UpdatePost(req)
	if err != nil {
		log.Println("failed to update post: ", err)
		return &p.PostResponse{}, err
	}

	return &p.PostResponse{}, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *p.IdRequest) (*p.PostResponse, error) {
	res, err := s.storage.Post().DeletePost(req)
	if err != nil {
		log.Println("failed to delete post: ", err)
		return &p.PostResponse{}, err
	}

	return res, err
}

func (s *PostService) GetNews(ctx context.Context, req *p.AllPostRequest) (*p.Posts, error) {
	res, err := s.storage.Post().GetNews(req)
	if err != nil {
		log.Println("failed to get post by id: ", err)
		return &p.Posts{}, err
	}

	return res, nil
}
