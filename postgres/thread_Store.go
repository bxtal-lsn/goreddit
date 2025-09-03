package postgres

import (
	"fmt"

	"github.com/bxtal-lsn/goreddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}

	return t, nil
}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	panic("not implemented")
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	panic("not implemented")
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	panic("not implemented")
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	panic("not implemented")
}
