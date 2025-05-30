// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package authors

import (
	"context"
	"database/sql"
	"fmt"
)

const createAuthor = `-- name: CreateAuthor :execlastid
INSERT INTO authors (
          name, bio
) VALUES (
  $1, $2
)
`

func (q *Queries) CreateAuthor(ctx context.Context) (int64, error) {
	result, err := q.db.ExecContext(ctx, createAuthor)
	if err != nil {
		return 0, fmt.Errorf("query CreateAuthor: %w", err)
	}
	return result.LastInsertId()
}

const deleteAuthorExec = `-- name: DeleteAuthorExec :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthorExec(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAuthorExec)
	if err != nil {
		err = fmt.Errorf("query DeleteAuthorExec: %w", err)
	}
	return err
}

const deleteAuthorExecLastID = `-- name: DeleteAuthorExecLastID :execlastid
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthorExecLastID(ctx context.Context) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteAuthorExecLastID)
	if err != nil {
		return 0, fmt.Errorf("query DeleteAuthorExecLastID: %w", err)
	}
	return result.LastInsertId()
}

const deleteAuthorExecResult = `-- name: DeleteAuthorExecResult :execresult
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthorExecResult(ctx context.Context) (sql.Result, error) {
	result, err := q.db.ExecContext(ctx, deleteAuthorExecResult)
	if err != nil {
		err = fmt.Errorf("query DeleteAuthorExecResult: %w", err)
	}
	return result, err
}

const deleteAuthorExecRows = `-- name: DeleteAuthorExecRows :execrows
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthorExecRows(ctx context.Context) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteAuthorExecRows)
	if err != nil {
		return 0, fmt.Errorf("query DeleteAuthorExecRows: %w", err)
	}
	return result.RowsAffected()
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	if err != nil {
		err = fmt.Errorf("query GetAuthor: %w", err)
	}
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, fmt.Errorf("query ListAuthors: %w", err)
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, fmt.Errorf("query ListAuthors: %w", err)
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("query ListAuthors: %w", err)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("query ListAuthors: %w", err)
	}
	return items, nil
}
