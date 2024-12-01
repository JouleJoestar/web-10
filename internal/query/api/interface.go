package api

type Usecase interface {
	FetchUser(name string) (string, error)
	CreateUser(name string) error
}
