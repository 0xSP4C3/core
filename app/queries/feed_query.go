package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


type FeedQueries struct {
    *sqlx.DB
}

func (q *FeedQueries) GetFeeds() ([]models.Feed, error) {
    feeds := []models.Feed{}

    query := `SELECT * FROM feeds`

    err := q.Get(&feeds, query)
    if err != nil {
        return feeds, err
    }

    return feeds, nil
}
// TODO: Probably should check foreign keys exists or not but it will return empty result no matter what.
func (q *FeedQueries) QueryFeedsByTimeRange(range_id uuid.UUID) ([]models.Feed, error) {
    feeds := []models.Feed{}

    query := `SELECT * FROM feeds f
    INNER JOIN feed_range fr (nolock)
    ON fr.id = f.id
    WHERE fr.id = $1`

    err := q.Get(&feeds, query, range_id)
    if err != nil {
        return feeds, err
    }

    return feeds, nil
}

func (q *FeedQueries) QueryFeedsByCoinId(coin_id uuid.UUID) ([]models.Feed, error) {
    feeds := []models.Feed{}

    query := `SELECT * FROM feeds f
    INNER JOIN coin c (nolock)
    ON f.coin_id = c.id
    where c.id = $1`

    err := q.Get(&feeds, query, coin_id) 
    if err != nil {
        return feeds, err
    }

    return feeds, nil
}

func (q *FeedQueries) QueryFeedsByCoinAndRangeId(range_id, coin_id uuid.UUID) ([]models.Feed, error) {
    feeds := []models.Feed{}

    query := `SELECT * FROM feeds f
    INNER JOIN coin c (nolock)
    ON f.coin_id = c.id
    INNER JOIN feed_range fr (nolock)
    ON f.range_id = fr.id
    WHERE c.id = $1 AND fr.id = $2`

    err := q.Get(&feeds, query, coin_id, range_id)
    if err != nil {
        return feeds, err
    }

    return feeds, nil
}

