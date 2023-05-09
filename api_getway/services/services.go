package services

import (
	"fmt"

	"github.com/PentaGol/api_getway/config"
	pu "github.com/PentaGol/api_getway/genproto/admin"
	pl "github.com/PentaGol/api_getway/genproto/liga"
	pp "github.com/PentaGol/api_getway/genproto/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	AdminService() pu.AdminServiceClient
	PostService() pp.PostServiceClient
	LigaService() pl.LigaServiceClient
}

type serviceManager struct {
	adminService pu.AdminServiceClient
	postService pp.PostServiceClient
	ligaService pl.LigaServiceClient
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connAdmin, err := grpc.Dial(
		fmt.Sprintf("%s:%s", conf.AdminServiceHost, conf.AdminServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%s", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connLiga, err := grpc.Dial(
		fmt.Sprintf("%s:%s", conf.LigaServiceHost, conf.LigaServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		adminService: pu.NewAdminServiceClient(connAdmin),
		postService: pp.NewPostServiceClient(connPost),
		ligaService: pl.NewLigaServiceClient(connLiga),
	}

	return serviceManager, nil
}

func (s *serviceManager) AdminService() pu.AdminServiceClient {
	return s.adminService
}

func (s *serviceManager) PostService() pp.PostServiceClient {
	return s.postService
}

func (s *serviceManager) LigaService() pl.LigaServiceClient {
	return s.ligaService
}
