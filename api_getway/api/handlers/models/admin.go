package models

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AdminRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IdAdminRequest struct {
	Id int64 `json:"id"`
}

type GetAllAdminsRequest struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

type Admins struct {
	Admins []Admin `json:"admins"`
}

type LoginAdmin struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	AccesToken   string `json:"acces_token"`
	Refreshtoken string `json:"refresh_token"`
}
