package storage

import (
	"github.com/sophistik/ITDigestBot/internal/entities"
	"github.com/sophistik/ITDigestBot/internal/repos"
)

type UserStorage struct {
	Tags map[int64]entities.User
}

func NewUserStorage() (repos.UserRepo, error) {
	return &UserStorage{
		Tags: make(map[int64]entities.User, 0),
	}, nil
}

func (s UserStorage) Create(user entities.User) error {
	s.Tags[user.ID] = user

	return nil
}

func (s *UserStorage) Update(user entities.User) error {
	s.Tags[user.ID] = user

	return nil
}

func (s *UserStorage) Find(id int64) (*entities.User, error) {
	result, _ := s.Tags[id]
	return &result, nil
}

func (s *UserStorage) Delete(id int64) error {
	delete(s.Tags, id)

	return nil
}
