// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: email.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEmail = `-- name: CreateEmail :one
INSERT INTO emails (
    id,
    user_id,
    subject,
    content,
    title,
    email_address
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, subject, content, title, email_address, create_at
`

type CreateEmailParams struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	Title        string    `json:"title"`
	EmailAddress string    `json:"email_address"`
}

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (Email, error) {
	row := q.db.QueryRow(ctx, createEmail,
		arg.ID,
		arg.UserID,
		arg.Subject,
		arg.Content,
		arg.Title,
		arg.EmailAddress,
	)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Subject,
		&i.Content,
		&i.Title,
		&i.EmailAddress,
		&i.CreateAt,
	)
	return i, err
}

const getEmailById = `-- name: GetEmailById :one
SELECT id, user_id, subject, content, title, email_address, create_at FROM emails
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEmailById(ctx context.Context, id uuid.UUID) (Email, error) {
	row := q.db.QueryRow(ctx, getEmailById, id)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Subject,
		&i.Content,
		&i.Title,
		&i.EmailAddress,
		&i.CreateAt,
	)
	return i, err
}

const getEmailsByUserId = `-- name: GetEmailsByUserId :many
SELECT id, user_id, subject, content, title, email_address, create_at FROM emails
WHERE user_id = $1
`

func (q *Queries) GetEmailsByUserId(ctx context.Context, userID uuid.UUID) ([]Email, error) {
	rows, err := q.db.Query(ctx, getEmailsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Email{}
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Subject,
			&i.Content,
			&i.Title,
			&i.EmailAddress,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmail = `-- name: UpdateEmail :one
UPDATE emails 
SET 
    subject = COALESCE($1, subject),
    content = COALESCE($2, content),
    title = COALESCE($3, title),
    email_address = COALESCE($4, email_address)
WHERE id = $5
RETURNING id, user_id, subject, content, title, email_address, create_at
`

type UpdateEmailParams struct {
	Subject      pgtype.Text `json:"subject"`
	Content      pgtype.Text `json:"content"`
	Title        pgtype.Text `json:"title"`
	EmailAddress pgtype.Text `json:"email_address"`
	ID           uuid.UUID   `json:"id"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) (Email, error) {
	row := q.db.QueryRow(ctx, updateEmail,
		arg.Subject,
		arg.Content,
		arg.Title,
		arg.EmailAddress,
		arg.ID,
	)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Subject,
		&i.Content,
		&i.Title,
		&i.EmailAddress,
		&i.CreateAt,
	)
	return i, err
}
