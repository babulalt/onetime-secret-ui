// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: secret.sql

package db

import (
	"context"
)

const createSecret = `-- name: CreateSecret :one
INSERT INTO secrets (
  content,
  creator,
  token,
  hashpassword,
  isview
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING id, content, creator, token, hashpassword, isview, created_at
`

type CreateSecretParams struct {
	Content      string `json:"content"`
	Creator      string `json:"creator"`
	Token        string `json:"token"`
	Hashpassword string `json:"hashpassword"`
	Isview       bool   `json:"isview"`
}

func (q *Queries) CreateSecret(ctx context.Context, arg CreateSecretParams) (Secret, error) {
	row := q.db.QueryRowContext(ctx, createSecret,
		arg.Content,
		arg.Creator,
		arg.Token,
		arg.Hashpassword,
		arg.Isview,
	)
	var i Secret
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Creator,
		&i.Token,
		&i.Hashpassword,
		&i.Isview,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSecret = `-- name: DeleteSecret :exec
DELETE FROM secrets
WHERE id = $1
`

func (q *Queries) DeleteSecret(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSecret, id)
	return err
}

const getSecret = `-- name: GetSecret :one
SELECT id, content, creator, token, hashpassword, isview, created_at FROM secrets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSecret(ctx context.Context, id int64) (Secret, error) {
	row := q.db.QueryRowContext(ctx, getSecret, id)
	var i Secret
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Creator,
		&i.Token,
		&i.Hashpassword,
		&i.Isview,
		&i.CreatedAt,
	)
	return i, err
}

const listSecrets = `-- name: ListSecrets :many
SELECT id, content, creator, token, hashpassword, isview, created_at FROM secrets
WHERE creator = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListSecretsParams struct {
	Creator string `json:"creator"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) ListSecrets(ctx context.Context, arg ListSecretsParams) ([]Secret, error) {
	rows, err := q.db.QueryContext(ctx, listSecrets, arg.Creator, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Secret{}
	for rows.Next() {
		var i Secret
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Creator,
			&i.Token,
			&i.Hashpassword,
			&i.Isview,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSecret = `-- name: UpdateSecret :one
UPDATE secrets
SET isview = $2
WHERE id = $1
RETURNING id, content, creator, token, hashpassword, isview, created_at
`

type UpdateSecretParams struct {
	ID     int64 `json:"id"`
	Isview bool  `json:"isview"`
}

func (q *Queries) UpdateSecret(ctx context.Context, arg UpdateSecretParams) (Secret, error) {
	row := q.db.QueryRowContext(ctx, updateSecret, arg.ID, arg.Isview)
	var i Secret
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Creator,
		&i.Token,
		&i.Hashpassword,
		&i.Isview,
		&i.CreatedAt,
	)
	return i, err
}