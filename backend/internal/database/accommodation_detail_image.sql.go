// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: accommodation_detail_image.sql

package database

import (
	"context"
)

const deleteAccommodationDetailImage = `-- name: DeleteAccommodationDetailImage :exec
DELETE FROM ` + "`" + `ecommerce_go_accommodation_detail_image` + "`" + `
WHERE
    ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteAccommodationDetailImage(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAccommodationDetailImage, id)
	return err
}

const getAccommodationDetailImages = `-- name: GetAccommodationDetailImages :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `image` + "`" + `
FROM
    ` + "`" + `ecommerce_go_accommodation_detail_image` + "`" + `
WHERE
    ` + "`" + `accommodation_detail_id` + "`" + ` = ?
`

type GetAccommodationDetailImagesRow struct {
	ID    string
	Image string
}

func (q *Queries) GetAccommodationDetailImages(ctx context.Context, accommodationDetailID string) ([]GetAccommodationDetailImagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccommodationDetailImages, accommodationDetailID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAccommodationDetailImagesRow
	for rows.Next() {
		var i GetAccommodationDetailImagesRow
		if err := rows.Scan(&i.ID, &i.Image); err != nil {
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

const updateAccommodationDetailImages = `-- name: UpdateAccommodationDetailImages :exec
INSERT INTO
    ` + "`" + `ecommerce_go_accommodation_detail_image` + "`" + ` (
        ` + "`" + `id` + "`" + `,
        ` + "`" + `accommodation_detail_id` + "`" + `,
        ` + "`" + `image` + "`" + `,
        ` + "`" + `created_at` + "`" + `,
        ` + "`" + `updated_at` + "`" + `
    )
VALUES
    (?, ?, ?, ?, ?)
`

type UpdateAccommodationDetailImagesParams struct {
	ID                    string
	AccommodationDetailID string
	Image                 string
	CreatedAt             uint64
	UpdatedAt             uint64
}

func (q *Queries) UpdateAccommodationDetailImages(ctx context.Context, arg UpdateAccommodationDetailImagesParams) error {
	_, err := q.db.ExecContext(ctx, updateAccommodationDetailImages,
		arg.ID,
		arg.AccommodationDetailID,
		arg.Image,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}
