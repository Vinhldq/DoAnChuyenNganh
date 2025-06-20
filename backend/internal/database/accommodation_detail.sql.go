// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: accommodation_detail.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/shopspring/decimal"
)

const checkAccommodationDetailExists = `-- name: CheckAccommodationDetailExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
        WHERE
            ` + "`" + `id` + "`" + ` = ?
    )
`

func (q *Queries) CheckAccommodationDetailExists(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkAccommodationDetailExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const countAccommodationDetail = `-- name: CountAccommodationDetail :one
SELECT
    COUNT(*)
FROM
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
`

func (q *Queries) CountAccommodationDetail(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAccommodationDetail)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccommodationDetail = `-- name: CreateAccommodationDetail :exec
INSERT INTO
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + ` (
        ` + "`" + `id` + "`" + `,
        ` + "`" + `accommodation_id` + "`" + `,
        ` + "`" + `name` + "`" + `,
        ` + "`" + `guests` + "`" + `,
        ` + "`" + `beds` + "`" + `,
        ` + "`" + `facilities` + "`" + `,
        ` + "`" + `available_rooms` + "`" + `,
        ` + "`" + `price` + "`" + `,
        ` + "`" + `created_at` + "`" + `,
        ` + "`" + `updated_at` + "`" + `
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateAccommodationDetailParams struct {
	ID              string
	AccommodationID string
	Name            string
	Guests          uint8
	Beds            json.RawMessage
	Facilities      json.RawMessage
	AvailableRooms  uint8
	Price           decimal.Decimal
	CreatedAt       uint64
	UpdatedAt       uint64
}

func (q *Queries) CreateAccommodationDetail(ctx context.Context, arg CreateAccommodationDetailParams) error {
	_, err := q.db.ExecContext(ctx, createAccommodationDetail,
		arg.ID,
		arg.AccommodationID,
		arg.Name,
		arg.Guests,
		arg.Beds,
		arg.Facilities,
		arg.AvailableRooms,
		arg.Price,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteAccommodationDetail = `-- name: DeleteAccommodationDetail :exec
UPDATE ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
SET
    ` + "`" + `is_deleted` + "`" + ` = 1
WHERE
    ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteAccommodationDetail(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAccommodationDetail, id)
	return err
}

const getAccommodationDetail = `-- name: GetAccommodationDetail :one
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `accommodation_id` + "`" + `,
    ` + "`" + `name` + "`" + `,
    ` + "`" + `guests` + "`" + `,
    ` + "`" + `beds` + "`" + `,
    ` + "`" + `facilities` + "`" + `,
    ` + "`" + `available_rooms` + "`" + `,
    ` + "`" + `price` + "`" + `
FROM
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `accommodation_id` + "`" + ` = ?
    and ` + "`" + `is_deleted` + "`" + ` = 0
`

type GetAccommodationDetailParams struct {
	ID              string
	AccommodationID string
}

type GetAccommodationDetailRow struct {
	ID              string
	AccommodationID string
	Name            string
	Guests          uint8
	Beds            json.RawMessage
	Facilities      json.RawMessage
	AvailableRooms  uint8
	Price           decimal.Decimal
}

func (q *Queries) GetAccommodationDetail(ctx context.Context, arg GetAccommodationDetailParams) (GetAccommodationDetailRow, error) {
	row := q.db.QueryRowContext(ctx, getAccommodationDetail, arg.ID, arg.AccommodationID)
	var i GetAccommodationDetailRow
	err := row.Scan(
		&i.ID,
		&i.AccommodationID,
		&i.Name,
		&i.Guests,
		&i.Beds,
		&i.Facilities,
		&i.AvailableRooms,
		&i.Price,
	)
	return i, err
}

const getAccommodationDetails = `-- name: GetAccommodationDetails :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `accommodation_id` + "`" + `,
    ` + "`" + `name` + "`" + `,
    ` + "`" + `guests` + "`" + `,
    ` + "`" + `beds` + "`" + `,
    ` + "`" + `discount_id` + "`" + `,
    ` + "`" + `facilities` + "`" + `,
    ` + "`" + `available_rooms` + "`" + `,
    ` + "`" + `price` + "`" + `
FROM
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
WHERE
    ` + "`" + `accommodation_id` + "`" + ` = ?
    and ` + "`" + `is_deleted` + "`" + ` = 0
`

type GetAccommodationDetailsRow struct {
	ID              string
	AccommodationID string
	Name            string
	Guests          uint8
	Beds            json.RawMessage
	DiscountID      sql.NullString
	Facilities      json.RawMessage
	AvailableRooms  uint8
	Price           decimal.Decimal
}

func (q *Queries) GetAccommodationDetails(ctx context.Context, accommodationID string) ([]GetAccommodationDetailsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccommodationDetails, accommodationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAccommodationDetailsRow
	for rows.Next() {
		var i GetAccommodationDetailsRow
		if err := rows.Scan(
			&i.ID,
			&i.AccommodationID,
			&i.Name,
			&i.Guests,
			&i.Beds,
			&i.DiscountID,
			&i.Facilities,
			&i.AvailableRooms,
			&i.Price,
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

const getAccommodationDetailsByIDs = `-- name: GetAccommodationDetailsByIDs :many
SELECT
    id, accommodation_id, name, guests, beds, facilities, available_rooms, price, discount_id, is_verified, is_deleted, created_at, updated_at
FROM
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
WHERE
    ` + "`" + `id` + "`" + ` IN (/*SLICE:ids*/?)
    AND ` + "`" + `accommodation_id` + "`" + ` = ?
`

type GetAccommodationDetailsByIDsParams struct {
	Ids             []string
	AccommodationID string
}

func (q *Queries) GetAccommodationDetailsByIDs(ctx context.Context, arg GetAccommodationDetailsByIDsParams) ([]EcommerceGoAccommodationDetail, error) {
	query := getAccommodationDetailsByIDs
	var queryParams []interface{}
	if len(arg.Ids) > 0 {
		for _, v := range arg.Ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(arg.Ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	queryParams = append(queryParams, arg.AccommodationID)
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EcommerceGoAccommodationDetail
	for rows.Next() {
		var i EcommerceGoAccommodationDetail
		if err := rows.Scan(
			&i.ID,
			&i.AccommodationID,
			&i.Name,
			&i.Guests,
			&i.Beds,
			&i.Facilities,
			&i.AvailableRooms,
			&i.Price,
			&i.DiscountID,
			&i.IsVerified,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getAccommodationDetailsWithPagination = `-- name: GetAccommodationDetailsWithPagination :many
SELECT
    ` + "`" + `id` + "`" + `,
    ` + "`" + `accommodation_id` + "`" + `,
    ` + "`" + `name` + "`" + `,
    ` + "`" + `guests` + "`" + `,
    ` + "`" + `beds` + "`" + `,
    ` + "`" + `discount_id` + "`" + `,
    ` + "`" + `facilities` + "`" + `,
    ` + "`" + `available_rooms` + "`" + `,
    ` + "`" + `price` + "`" + `
FROM
    ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
LIMIT
    ?
OFFSET
    ?
`

type GetAccommodationDetailsWithPaginationParams struct {
	Limit  int32
	Offset int32
}

type GetAccommodationDetailsWithPaginationRow struct {
	ID              string
	AccommodationID string
	Name            string
	Guests          uint8
	Beds            json.RawMessage
	DiscountID      sql.NullString
	Facilities      json.RawMessage
	AvailableRooms  uint8
	Price           decimal.Decimal
}

func (q *Queries) GetAccommodationDetailsWithPagination(ctx context.Context, arg GetAccommodationDetailsWithPaginationParams) ([]GetAccommodationDetailsWithPaginationRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccommodationDetailsWithPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAccommodationDetailsWithPaginationRow
	for rows.Next() {
		var i GetAccommodationDetailsWithPaginationRow
		if err := rows.Scan(
			&i.ID,
			&i.AccommodationID,
			&i.Name,
			&i.Guests,
			&i.Beds,
			&i.DiscountID,
			&i.Facilities,
			&i.AvailableRooms,
			&i.Price,
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

const getInfoAvailableRoomOfAccommodationDetailByOrderID = `-- name: GetInfoAvailableRoomOfAccommodationDetailByOrderID :many
SELECT
    egad.id,
    egad.available_rooms,
    egod.quantity
FROM
    ` + "`" + `ecommerce_go_order_detail` + "`" + ` egod
    JOIN ` + "`" + `ecommerce_go_accommodation_detail` + "`" + ` egad ON egod.accommodation_detail_id = egad.id
WHERE
    egod.order_id = ?
`

type GetInfoAvailableRoomOfAccommodationDetailByOrderIDRow struct {
	ID             string
	AvailableRooms uint8
	Quantity       uint8
}

func (q *Queries) GetInfoAvailableRoomOfAccommodationDetailByOrderID(ctx context.Context, orderid string) ([]GetInfoAvailableRoomOfAccommodationDetailByOrderIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getInfoAvailableRoomOfAccommodationDetailByOrderID, orderid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInfoAvailableRoomOfAccommodationDetailByOrderIDRow
	for rows.Next() {
		var i GetInfoAvailableRoomOfAccommodationDetailByOrderIDRow
		if err := rows.Scan(&i.ID, &i.AvailableRooms, &i.Quantity); err != nil {
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

const updateAccommodationDetail = `-- name: UpdateAccommodationDetail :exec
UPDATE ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
SET
    ` + "`" + `name` + "`" + ` = ?,
    ` + "`" + `guests` + "`" + ` = ?,
    ` + "`" + `beds` + "`" + ` = ?,
    ` + "`" + `facilities` + "`" + ` = ?,
    ` + "`" + `available_rooms` + "`" + ` = ?,
    ` + "`" + `price` + "`" + ` = ?,
    ` + "`" + `updated_at` + "`" + ` = ?
WHERE
    ` + "`" + `id` + "`" + ` = ?
    and ` + "`" + `accommodation_id` + "`" + ` = ?
    and ` + "`" + `is_deleted` + "`" + ` = 0
`

type UpdateAccommodationDetailParams struct {
	Name            string
	Guests          uint8
	Beds            json.RawMessage
	Facilities      json.RawMessage
	AvailableRooms  uint8
	Price           decimal.Decimal
	UpdatedAt       uint64
	ID              string
	AccommodationID string
}

func (q *Queries) UpdateAccommodationDetail(ctx context.Context, arg UpdateAccommodationDetailParams) error {
	_, err := q.db.ExecContext(ctx, updateAccommodationDetail,
		arg.Name,
		arg.Guests,
		arg.Beds,
		arg.Facilities,
		arg.AvailableRooms,
		arg.Price,
		arg.UpdatedAt,
		arg.ID,
		arg.AccommodationID,
	)
	return err
}

const updateAvailableRoom = `-- name: UpdateAvailableRoom :exec
UPDATE ` + "`" + `ecommerce_go_accommodation_detail` + "`" + `
SET
    ` + "`" + `available_rooms` + "`" + ` = ?
WHERE
    ` + "`" + `id` + "`" + ` = ?
`

type UpdateAvailableRoomParams struct {
	AvailableRoom uint8
	ID            string
}

func (q *Queries) UpdateAvailableRoom(ctx context.Context, arg UpdateAvailableRoomParams) error {
	_, err := q.db.ExecContext(ctx, updateAvailableRoom, arg.AvailableRoom, arg.ID)
	return err
}
