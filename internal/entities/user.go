package entities

type User struct {
	ID        int64    `db:"id"`
	Tags      []string `db:"-"`
	LastInput *string  `db:"last_input"`
}
