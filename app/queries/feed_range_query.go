package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TimeFrameQueries struct {
    *sqlx.DB
}

func (q *TimeFrameQueries)GetFeedRanges() ([]models.TimeFrame, error) {
    timeFrames := []models.TimeFrame{}

    query := `SELECT * FROM feed_range`

    err := q.Get(&timeFrames, query)
    if err != nil {
        return timeFrames, err
    }

    return timeFrames, nil
}

func (q *TimeFrameQueries) GetTimeFrame(id uuid.UUID) (models.TimeFrame, error) {
    timeFrame := models.TimeFrame{}

    query := `SELECT * FROM feed_range WHERE ID = $1`

    err := q.Get(&timeFrame, query, id)
    if err != nil {
        return timeFrame, err
    }

    return timeFrame, nil
}

func (q *TimeFrameQueries) CreateFeedRange(f *models.TimeFrame) error {
    query := `INSERT INTO feed_range VALUES ($1, $2, $3, $4, $5, $6)`

    _, err := q.Exec(query, f.ID, f.CreatedAt, f.UpdatedAt, f.Name, f.Description, f.IsEnabled)
    if err != nil {
        return err
    }

    return nil
}
