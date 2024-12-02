package manager

import (
	"github.com/project-sistem-voucher/api/service"
	"github.com/project-sistem-voucher/config"
)

type ServiceManager interface {
	VoucherService() service.VoucherService
	RedeemService() service.RedeemService
	UserService() service.UserService
	AuthService() service.AuthService
}

type serviceManager struct {
	repoManager RepoManager
	redisClient config.Cacher
}

func NewServiceManager(repo RepoManager, redis config.Cacher) ServiceManager {
	return &serviceManager{
		repoManager: repo,
		redisClient: redis,
	}
}

func (m *serviceManager) VoucherService() service.VoucherService {
	return service.NewVoucherService(m.repoManager.VoucherRepo())
}

func (m *serviceManager) RedeemService() service.RedeemService {
	return service.NewRedeemService(m.repoManager.RedeemRepo(), m.repoManager.VoucherRepo())
}

func (m *serviceManager) UserService() service.UserService {
	return service.NewUserService(m.repoManager.UserRepo())
}

func (m *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(m.repoManager.UserRepo(), m.redisClient)
}
