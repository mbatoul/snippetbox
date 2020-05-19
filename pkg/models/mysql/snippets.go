package mysql

import (
	"database/sql"
	"errors"

	"mathisbatoul.com/snippetbox/pkg/models"
)

// SnippetModel wraps a SQL connection pool
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires string) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `
		SELECT * FROM snippets
		WHERE expires > UTC_TIMESTAMP()
		AND id = ?
	`

	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Created)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
