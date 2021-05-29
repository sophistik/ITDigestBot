package repos

type UserTagsRepo interface {
	Create(id int64, tags []string) error
	Update(id int64, tags []string) error
	Upsert(id int64, tags []string) error
	Get(id int64) ([]string, error)
}

type LastUserInput interface {
	Add(id int64, input string) error
	Delete(id int64) error
	Get(id int64) (string, error)
}
