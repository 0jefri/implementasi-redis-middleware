package manager

import "github.com/project-sistem-voucher/api/repository"

type RepoManager interface {
	VoucherRepo() repository.VoucherRepository
	RedeemRepo() repository.RedeemRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) VoucherRepo() repository.VoucherRepository {
	return repository.NewVoucherRepository(m.infraManager.Conn())
}

func (m *repoManager) RedeemRepo() repository.RedeemRepository {
	return repository.NewRedeemRepository(m.infraManager.Conn())
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}
