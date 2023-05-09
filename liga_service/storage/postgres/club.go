package postgres

import (
	"log"

	p "github.com/PentaGol/liga_service/genproto/liga"
)

func (r *Repo) CreateClub(club *p.ClubRequest) (*p.ClubResponse, error) {
	var res p.ClubResponse
	err := r.db.QueryRow(`
		insert into 
			clubs(name, points) 
		values
			($1, $2) 
		returning 
			id, name, points, created_at, updated_at`, club.Name, club.Points).
		Scan(
			&res.Id,
			&res.Name,
			&club.Points,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create club")
		return &p.ClubResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetClubById(liga *p.IdRequest) (*p.ClubResponse, error) {
	res := p.ClubResponse{}
	err := r.db.QueryRow(`
		select 
			id, name, points, created_at, updated_at 
		from 
			clubs 
		where 
			id = $1 and deleted_at is null`, liga.Id).
		Scan(
			&res.Id,
			&res.Name,
			&res.Points,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get club")
		return &p.ClubResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetAllClubs(req *p.AllClubRequest) (*p.Clubs, error) {
	res := p.Clubs{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, name, points, created_at, updated_at 
		from 
			clubs 
		where 
			deleted_at is null 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get club")
		return &p.Clubs{}, err
	}

	for rows.Next() {
		temp := p.ClubResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Points,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning club")
			return &p.Clubs{}, err
		}

		res.Clubs = append(res.Clubs, &temp)
	}

	return &res, nil
}
