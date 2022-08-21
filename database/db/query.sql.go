// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email FROM users WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return &i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, email FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]*User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertUser = `-- name: InsertUser :execlastid
INSERT INTO users (name, email) VALUES (?, ?)
`

type InsertUserParams struct {
	Name  sql.NullString
	Email string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, insertUser, arg.Name, arg.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
