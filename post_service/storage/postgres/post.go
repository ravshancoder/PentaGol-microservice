package postgres

import (
	"fmt"
	"log"
	"time"

	p "github.com/PentaGol/post_service/genproto/post"
)

func (r *PostRepo) CreatePost(post *p.PostRequest) (*p.PostResponse, error) {
	var res p.PostResponse
	err := r.db.QueryRow(`
		insert into 
			posts(title, description, img_url) 
		values
			($1, $2, $3) 
		returning 
			id, title, description, img_url, created_at, updated_at`, post.Title, post.Description, post.ImgUrl).
		Scan(
			&res.Id, 
			&res.Title, 
			&res.Description,
			&post.ImgUrl,
			&res.CreatedAt, 
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create post")
		return &p.PostResponse{}, err
	}

	return &res, nil
}

func (r *PostRepo) GetPostById(post *p.IdRequest) (*p.PostResponse, error) {
	res := p.PostResponse{}
	err := r.db.QueryRow(`
		select 
			id, title, description, img_url, created_at, updated_at 
		from 
			posts 
		where 
			id = $1 and deleted_at is null`, post.Id).
		Scan(
			&res.Id, 
			&res.Title, 
			&res.Description, 
			&res.ImgUrl,
			&res.CreatedAt, 
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get post")
		return &p.PostResponse{}, err
	}

	return &res, nil
}

func (r *PostRepo) GetAllPosts(req *p.AllPostRequest)(*p.Posts, error){
	res := p.Posts{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, title, description, img_url, created_at, updated_at 
		from 
			posts 
		where 
			deleted_at is null 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get all post")
		return &p.Posts{}, err
	}

	for rows.Next(){
		temp := p.PostResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Title,
			&temp.Description,
			&temp.ImgUrl,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning post")
			return &p.Posts{}, err
		}

		res.Posts = append(res.Posts, &temp)
	}

	return &res, nil
}

func (r *PostRepo) SearchByTitle(title *p.Search) (*p.Posts, error) {
	res := p.Posts{}
	query := fmt.Sprintf("select id, title, description, img_url, created_at, updated_at from posts where title ilike '%" + title.Search + "%' and deleted_at is null")

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("failed to search post")
		return &p.Posts{}, nil
	}

	for rows.Next() {
		post := p.PostResponse{}

		err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.ImgUrl,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning post")
			return &p.Posts{}, nil
		}

		res.Posts = append(res.Posts, &post)
	}

	return &res, nil
}

func (r *PostRepo) UpdatePost(post *p.UpdatePostRequest) error {
	res, err := r.db.Exec(`
		update
			posts 
		set 
			title = $1, description = $2, img_url = $3, updated_at = $4
		where 
			id = $5`, post.Title, post.Description, post.ImgUrl, time.Now(), post.Id)
	if err != nil {
		log.Println("failed to update post")
		return err
	}

	fmt.Println(res.RowsAffected())

	return nil
}

func (r *PostRepo) DeletePost(id *p.IdRequest) (*p.PostResponse, error) {
	post := p.PostResponse{}
	err := r.db.QueryRow(`
		update 
			posts 
		set 
			deleted_at = $1 
		where 
			id = $2 
		returning 
			id, title, description, img_url, created_at, updated_at`, time.Now(), id.Id).
		Scan(
			&post.Id, 
			&post.Title, 
			&post.Description, 
			&post.ImgUrl,
			&post.CreatedAt, 
			&post.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to delete post")
		return &p.PostResponse{}, err
	}

	return &post, nil
}

func (r *PostRepo) GetNews(req *p.AllPostRequest)(*p.Posts, error){
	res := p.Posts{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, title, description, img_url, created_at, updated_at 
		from 
			posts 
		where 
			deleted_at is null 
		limit $1 offset $2
		ORDER BY created_at DESC`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get post")
		return &p.Posts{}, err
	}

	for rows.Next(){
		temp := p.PostResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Title,
			&temp.Description,
			&temp.ImgUrl,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning post")
			return &p.Posts{}, err
		}

		res.Posts = append(res.Posts, &temp)
	}

	return &res, nil
}