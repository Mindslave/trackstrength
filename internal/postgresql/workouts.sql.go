// Code generated by sqlc. DO NOT EDIT.
// source: workouts.sql

package postgresql

import (
	"context"
	"database/sql"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO workouts (
  time_estimate 
) VALUES (
  $1
)
RETURNING id, time_estimate
`

func (q *Queries) CreateWorkout(ctx context.Context, timeEstimate sql.NullInt32) (Workout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout, timeEstimate)
	var i Workout
	err := row.Scan(&i.ID, &i.TimeEstimate)
	return i, err
}

const getWorkout = `-- name: GetWorkout :one
SELECT id, time_estimate FROM workouts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWorkout(ctx context.Context, id int64) (Workout, error) {
	row := q.db.QueryRowContext(ctx, getWorkout, id)
	var i Workout
	err := row.Scan(&i.ID, &i.TimeEstimate)
	return i, err
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT id, time_estimate FROM workouts
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListWorkoutsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWorkouts(ctx context.Context, arg ListWorkoutsParams) ([]Workout, error) {
	rows, err := q.db.QueryContext(ctx, listWorkouts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workout
	for rows.Next() {
		var i Workout
		if err := rows.Scan(&i.ID, &i.TimeEstimate); err != nil {
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