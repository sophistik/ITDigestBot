package services

import (
	"fmt"

	"github.com/sophistik/ITDigestBot/internal/entities"
	"github.com/sophistik/ITDigestBot/internal/repos"
	"github.com/sophistik/ITDigestBot/pkg/postgres"
	"github.com/sophistik/ITDigestBot/pkg/slices"
)

type BotAPIService struct {
	userRepository repos.UserRepo
}

func NewBotAPIService(
	utr repos.UserRepo,
) *BotAPIService {
	return &BotAPIService{
		userRepository: utr,
	}
}

func (a *BotAPIService) AddTags(id int64, tags []string) error {
	user, err := a.userRepository.Find(id)
	if err != nil {
		if err != postgres.ErrNotFound {
			return fmt.Errorf("can't get user %d: %w", id, err)
		}
	}

	if user == nil {
		user = &entities.User{
			ID:   id,
			Tags: tags,
		}

		if err := a.userRepository.Create(*user); err != nil {
			return fmt.Errorf("can't create user %d tags, %w", id, err)
		}

		return nil
	}

	user.Tags = append(user.Tags, tags...)
	user.Tags = slices.Unique(user.Tags)

	if err := a.userRepository.Update(*user); err != nil {
		return fmt.Errorf("can't update user's %d tags: %w", id, err)
	}

	return nil
}

func (a *BotAPIService) RemoveTags(id int64, tags []string) error {
	user, err := a.userRepository.Find(id)
	if err != nil {
		return fmt.Errorf("can't get user %d: %w", id, err)
	}

	user.Tags = slices.Delete(user.Tags, tags)

	if err := a.userRepository.Update(*user); err != nil {
		return fmt.Errorf("can't update user's %d tags: %w", id, err)
	}

	return nil
}

func (a *BotAPIService) GetTags(id int64) ([]string, error) {
	user, err := a.userRepository.Find(id)
	if err != nil {
		return nil, fmt.Errorf("can't find user %d: %w", id, err)
	}

	return user.Tags, nil
}

func (a *BotAPIService) SetLastInput(id int64, input string) error {
	user, err := a.userRepository.Find(id)
	if err != nil {
		if err != postgres.ErrNotFound {
			return fmt.Errorf("can't get user %d: %w", id, err)
		}
	}

	if user == nil {
		user = &entities.User{
			ID:        id,
			LastInput: &input,
		}

		if err := a.userRepository.Create(*user); err != nil {
			return fmt.Errorf("can't create user %d tags, %w", id, err)
		}

		return nil
	}

	user.LastInput = &input

	if err := a.userRepository.Update(*user); err != nil {
		return fmt.Errorf("can't update user's %d tags: %w", id, err)
	}

	return nil
}

func (a *BotAPIService) RemoveLastUpdate(id int64) error {
	user, err := a.userRepository.Find(id)
	if err != nil {
		return fmt.Errorf("can't get user %d: %w", id, err)
	}

	user.LastInput = nil

	if err := a.userRepository.Update(*user); err != nil {
		return fmt.Errorf("can't update user's %d tags: %w", id, err)
	}

	return nil
}

func (a *BotAPIService) GetLastInput(id int64) (string, error) {
	user, err := a.userRepository.Find(id)
	if err != nil {
		return "", fmt.Errorf("can't find user %d: %w", id, err)
	}

	return *user.LastInput, nil
}
