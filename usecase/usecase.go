package usecase

import (
	"go-cassandra-crud/entity"
	"go-cassandra-crud/repo"
)

type Usecase struct {
	repo repo.Repo
}

func New(repo repo.Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) FetchAll() ([]entity.CartCount, error) {
	return u.repo.FetchAll()
}

func (u *Usecase) FetchOne(userID string) (entity.CartCount, error) {
	return u.repo.FetchOne(userID)
}

func (u *Usecase) Insert(req entity.CartCount) error {
	return u.repo.Insert(req)
}

func (u *Usecase) Delete(userID string) error {
	return u.repo.Delete(userID)
}
