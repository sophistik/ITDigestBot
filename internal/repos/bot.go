package repos

import "github.com/sophistik/ITDigestBot/internal/entities"

type UserRepo interface {
	Create(user entities.User) error
	Update(user entities.User) error
	Find(id int64) (*entities.User, error)
	Delete(id int64) error
}

type LastUserInput interface {
	Create(id int64, input string) error
	Delete(id int64) error
	Find(id int64) (string, error)
}
