package postgres

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/lib/pq"
	"github.com/sophistik/ITDigestBot/internal/entities"
	"github.com/sophistik/ITDigestBot/pkg/postgres"
)

const userTable string = "users"

type UserTagsStorage struct {
	db   *postgres.DB
	goqu *goqu.Database
}

func NewUserTagsStorage(db *postgres.DB, goqu *goqu.Database) (*UserTagsStorage, error) {
	storage := &UserTagsStorage{
		db:   db,
		goqu: goqu,
	}
	queries := []string{
		userCreateQuery,
		userUpdateQuery,
	}

	if err := db.ValidateQueries(queries); err != nil {
		return nil, fmt.Errorf("invalid query exists: %w", err)
	}

	return storage, nil
}

func (s *UserTagsStorage) Find(id int64) (*entities.User, error) {
	query := s.findBaseQuery(s.goqu).Where(goqu.I("id").Eq(id))

	sqlQuery, params, err := query.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query: %w", err)
	}

	var u entities.User

	if err := s.db.Session.QueryRow(sqlQuery, params...).Scan(
		&u.ID,
		pq.Array(&u.Tags),
		&u.LastInput,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, postgres.ErrNotFound
		}

		return nil, fmt.Errorf("can't exec query: %w", err)
	}

	return &u, nil
}

func (s *UserTagsStorage) findBaseQuery(db postgres.GoquDB) *goqu.SelectDataset {
	selectDataset := db.From(goqu.T(userTable))

	return selectDataset.
		Select(
			"id",
			"tags",
			"last_input",
		)
}

const userCreateQuery = `
	INSERT INTO users (
		id,
		tags,
		last_input
	)
	VALUES (
		$1, $2, $3
	)
`

func (s *UserTagsStorage) Create(u entities.User) error {
	_, err := s.db.Session.Exec(
		userCreateQuery,
		u.ID,
		pq.Array(u.Tags),
		u.LastInput,
	)

	if err != nil {
		return fmt.Errorf("can't exec query: %w", err)
	}

	return nil
}

const userUpdateQuery = `
	UPDATE users 
	SET 
		tags = $2,
		last_input = $3
	WHERE id = $1
`

func (s *UserTagsStorage) Update(u entities.User) error {
	_, err := s.db.Session.Exec(
		userUpdateQuery,
		u.ID,
		pq.Array(u.Tags),
		u.LastInput,
	)

	if err != nil {
		return fmt.Errorf("can't exec query: %w", err)
	}

	return nil
}

func (s *UserTagsStorage) Delete(id int64) error {
	delete := s.goqu.
		Delete(userTable).
		Where(goqu.Ex{
			"id": id,
		}).Prepared(true).
		Executor()

	if _, err := delete.Exec(); err != nil {
		return fmt.Errorf("couldn't exec delete: %w", err)
	}

	return nil
}
