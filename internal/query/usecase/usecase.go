package usecase

import "github.com/ValeryBMSTU/web-10/internal/query/provider"

type Usecase struct {
	dbProvider *provider.DatabaseProvider
}

func NewUsecase(dp *provider.DatabaseProvider) *Usecase {
	return &Usecase{dbProvider: dp}
}

func (uc *Usecase) FetchUser(name string) (string, error) {
	return uc.dbProvider.SelectUser(name)
}

func (uc *Usecase) CreateUser(name string) error {
	return uc.dbProvider.InsertUser(name)
}
