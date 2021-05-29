package storage

import (
	"github.com/sophistik/ITDigestBot/internal/repos"
)

type UserTagsStorage struct {
	Tags map[int64][]string
}

func NewUserTagsStorage() (repos.UserTagsRepo, error) {
	return &UserTagsStorage{
		Tags: make(map[int64][]string, 0),
	}, nil
}

func (s *UserTagsStorage) Upsert(id int64, tags []string) error {
	if _, ok := s.Tags[id]; ok {
		return s.Update(id, tags)
	}

	return s.Create(id, tags)
}

func (s *UserTagsStorage) Create(id int64, tags []string) error {
	s.Tags[id] = tags

	return nil
}

func (s *UserTagsStorage) Update(id int64, tags []string) error {
	s.Tags[id] = tags

	return nil
}

func (s *UserTagsStorage) Get(id int64) ([]string, error) {
	return s.Tags[id], nil
}
