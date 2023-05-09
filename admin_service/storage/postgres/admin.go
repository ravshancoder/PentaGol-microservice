package postgres

import (
	"fmt"
	"log"

	u "github.com/PentaGol/admin_service/genproto/admin"
)

func (r *AdminRepo) CreateAdmin(admin *u.AdminRequest) (*u.AdminResponse, error) {
	var res u.AdminResponse
	err := r.db.QueryRow(`
		insert into 
			admins(name, email, password, refresh_token) 
		values
			($1, $2, $3, $4) 
		returning 
			id, name, email, refresh_token, created_at, updated_at`, admin.Name, admin.Email, admin.Password, admin.RefreshToken).
		Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.RefreshToken,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create admin")
		return &u.AdminResponse{}, err
	}

	return &res, nil
}

func (r *AdminRepo) GetAdminById(admin *u.IdRequest) (*u.AdminResponse, error) {
	var res u.AdminResponse
	err := r.db.QueryRow(`
		select 
			id, name, email, refresh_token, created_at, updated_at
		from 
			admins 
		where id = $1 and deleted_at is null`, admin.Id).
		Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.RefreshToken,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get Admin")
		return &u.AdminResponse{}, err
	}

	return &res, nil
}

func (r *AdminRepo) CheckFiedld(req *u.CheckFieldReq) (*u.CheckFieldRes, error) {
	query := fmt.Sprintf("SELECT 1 FROM admins WHERE %s=$1", req.Field)
	res := &u.CheckFieldRes{}
	temp := -1
	err := r.db.QueryRow(query, req.Value).Scan(&temp)
	if err != nil {
		res.Exists = false
		return res, nil
	}
	if temp == 0 {
		res.Exists = true
	} else {
		res.Exists = false
	}
	return res, nil
}

func (r *AdminRepo) GetByEmail(req *u.EmailReq) (*u.AdminResponse, error) {
	res := u.AdminResponse{}
	err := r.db.QueryRow(`
	SELECT 
		id, 
		name,
		email, 
		password,
		refresh_token,
		created_at, 
		updated_at 
	FROM 
		admins 
	WHERE 
		email=$1 AND deleted_at IS NULL`, req.Email).
	Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.Password,
		&res.RefreshToken,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error while getting Admin login")
		return &u.AdminResponse{}, err
	}

	return &res, nil
}

func (r *AdminRepo) UpdateToken(admin *u.RequestForTokens) (*u.AdminResponse, error) {
	res := u.AdminResponse{}
	err := r.db.QueryRow(`
		update
			admins
		set
			refresh_token = $1
		where 
			id = $2
		returning 
			id, name, email, refresh_token, created_at, updated_at`, admin.RefreshToken, admin.Id).
		Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.RefreshToken,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to update admin")
		return &u.AdminResponse{}, err
	}

	return &res, err
}
