package repo

import (
	u "github.com/PentaGol/admin_service/genproto/admin"
)

type AdminStoreI interface {
	CreateAdmin(*u.AdminRequest) (*u.AdminResponse, error)
	GetAdminById(*u.IdRequest) (*u.AdminResponse, error)
	CheckFiedld(*u.CheckFieldReq) (*u.CheckFieldRes, error)
	GetByEmail(*u.EmailReq) (*u.AdminResponse, error)
	UpdateToken(*u.RequestForTokens) (*u.AdminResponse, error)
}
