// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	DeleteSong(ctx context.Context, id int) error
	GetVerses(ctx context.Context, arg GetVersesParams) ([]GetVersesRow, error)
	InsertSong(ctx context.Context, arg InsertSongParams) (int, error)
	UpdateSong(ctx context.Context, arg UpdateSongParams) error
}

var _ Querier = (*Queries)(nil)
