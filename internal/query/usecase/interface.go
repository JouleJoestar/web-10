// internal/query/usecase/interface.go
package usecase

type Usecase interface {
	FetchUser(name string) (string, error)
	CreateUser(name string) error
}
