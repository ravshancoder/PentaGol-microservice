package postgres

import (
	"log"
	"time"

	p "github.com/PentaGol/liga_service/genproto/liga"
)

func (r *Repo) CreateLiga(liga *p.LigaRequest) (*p.LigaResponse, error) {
	var res p.LigaResponse
	err := r.db.QueryRow(`
		insert into 
			ligas(name) 
		values
			($1) 
		returning 
			id, name, created_at, updated_at`, liga.Name).
		Scan(
			&res.Id,
			&res.Name,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create liga")
		return &p.LigaResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetLigaById(liga *p.IdRequest) (*p.LigaResponse, error) {
	res := p.LigaResponse{}
	err := r.db.QueryRow(`
		select 
			id, name, created_at, updated_at 
		from 
			ligas 
		where 
			id = $1 and deleted_at is null`, liga.Id).
		Scan(
			&res.Id,
			&res.Name,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get liga")
		return &p.LigaResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetAllLigas(req *p.AllLigaRequest) (*p.Ligas, error) {
	res := p.Ligas{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, name, created_at, updated_at 
		from 
			ligas 
		where 
			deleted_at is null 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get liga")
		return &p.Ligas{}, err
	}

	for rows.Next() {
		temp := p.LigaResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning liga")
			return &p.Ligas{}, err
		}

		res.Ligas = append(res.Ligas, &temp)
	}

	return &res, nil
}

func (r *Repo) DeleteLiga(id *p.IdRequest) (*p.LigaResponse, error) {
	liga := p.LigaResponse{}
	err := r.db.QueryRow(`
		update 
			ligas 
		set 
			deleted_at = $1 
		where 
			id = $2 and deleted_at is null
		returning 
			id, name, created_at, updated_at`, time.Now(), id.Id).
		Scan(
			&liga.Id,
			&liga.Name,
			&liga.CreatedAt,
			&liga.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to delete liga")
		return &p.LigaResponse{}, err
	}

	return &liga, nil
}
