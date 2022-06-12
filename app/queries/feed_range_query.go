package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FeedRangeQueries struct {
    *sqlx.DB
}

func (q *FeedRangeQueries)GetFeedRanges() ([]models.FeedRange, error) {
    feedRanges := []models.FeedRange{}

    query := `SELECT * FROM feed_range`

    err := q.Get(&feedRanges, query)
    if err != nil {
        return feedRanges, err
    }

    return feedRanges, nil
}

func (q *FeedRangeQueries) GetFeedRange(id uuid.UUID) (models.FeedRange, error) {
    feedRange := models.FeedRange{}

    query := `SELECT * FROM feed_range WHERE ID = $1`

    err := q.Get(&feedRange, query, id)
    if err != nil {
        return feedRange, err
    }

    return feedRange, nil
}

func (q *FeedRangeQueries) CreateFeedRange(f *models.FeedRange) error {
    query := `INSERT INTO feed_range VALUES ($1, $2, $3, $4, )`

    _, err := q.Exec(query, f.ID, f.CreatedAt, f.UpdatedAt, f.Name, f.Description, f.IsEnabled)
    if err != nil {
        return err
    }

    return nil
}
