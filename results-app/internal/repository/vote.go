package repository

import (
	"context"
	"database/sql"
	"github.com/thyagofr/results-app/internal/model"
)

const reportQuery = "SELECT category, total_per_category FROM public.votes_report"

type VoteRepository interface {
	GetReport(ctx context.Context) ([]*model.VotesPerCategory, error)
}

type VotePostgreSQL struct {
	db *sql.DB
}

var _ VoteRepository = &VotePostgreSQL{}

func NewVotePostgreSQL(db *sql.DB) *VotePostgreSQL {
	return &VotePostgreSQL{db: db}
}

func (v *VotePostgreSQL) GetReport(ctx context.Context) ([]*model.VotesPerCategory, error) {
	report := make([]*model.VotesPerCategory, 0)

	rows, err := v.db.QueryContext(ctx, reportQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		resume := new(model.VotesPerCategory)
		if err = rows.Scan(&resume.Category, &resume.Votes); err != nil {
			return nil, err
		}
		report = append(report, resume)
	}

	return report, nil
}
