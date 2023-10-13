package repository

import (
	"context"
	"database/sql"
	"github.com/thyagofr/voting-app/worker/internal/model"
)

const insertQuery = "INSERT INTO votes (voting_date, category) VALUES ($1, $2)"

type VoteRepository interface {
	Save(ctx context.Context, vote *model.Vote) error
}

type VotePostgreSQL struct {
	db *sql.DB
}

var _ VoteRepository = &VotePostgreSQL{}

func NewVotePostgreSQL(db *sql.DB) *VotePostgreSQL {
	return &VotePostgreSQL{db: db}
}

func (v *VotePostgreSQL) Save(ctx context.Context, vote *model.Vote) error {
	if _, err := v.db.ExecContext(ctx, insertQuery, vote.CreationDate, vote.Category); err != nil {
		return err
	}
	return nil
}
