// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: heights.sql

package repository

import (
	"context"
)

const getHeight = `-- name: GetHeight :one
SELECT
    height
FROM
    chains
WHERE
    id = $1
`

func (q *Queries) GetHeight(ctx context.Context, id int32) (int64, error) {
	row := q.queryRow(ctx, q.getHeightStmt, getHeight, id)
	var height int64
	err := row.Scan(&height)
	return height, err
}

const getHeightByName = `-- name: GetHeightByName :one
SELECT
    height
FROM
    chains
WHERE
    chain = $1
`

func (q *Queries) GetHeightByName(ctx context.Context, chain string) (int64, error) {
	row := q.queryRow(ctx, q.getHeightByNameStmt, getHeightByName, chain)
	var height int64
	err := row.Scan(&height)
	return height, err
}

const setHeight = `-- name: SetHeight :exec
UPDATE
    chains
SET
    height = $1
WHERE
    id = $2
`

type SetHeightParams struct {
	Height int64
	ID     int32
}

func (q *Queries) SetHeight(ctx context.Context, arg *SetHeightParams) error {
	_, err := q.exec(ctx, q.setHeightStmt, setHeight, arg.Height, arg.ID)
	return err
}