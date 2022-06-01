package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ExchangeQueries struct {
    *sqlx.DB
}


func (q *ExchangeQueries) GetExchanges() ([]models.Exchange, error) {
    exchanges := []models.Exchange{}

    query := `SELECT * FROM exchanges`

    err := q.Get(&exchanges, query)
    if err != nil {
        return exchanges, err
    }

    return exchanges, nil
}

func (q *ExchangeQueries) GetExchange(id uuid.UUID) (models.Exchange, error) {
    exchange := models.Exchange{}

    query := `SELECT * FROM exchanges WHERE ID = $1`

    err := q.Get(&exchange, query, id)
    if err != nil {
        return exchange, err
    }

    return exchange, nil
}

func (q *ExchangeQueries) CreateExchange(e *models.Exchange) error {
    query := `INSERT INTO exchanges VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

    _, err := q.Exec(query, e.ID, e.CreatedAt, e.UpdatedAt, e.Name, e.Description, e.Uri, e.IsEnabled, e.IsBlocked) 

    if err != nil {
        return err
    }

    return nil
}

func (q *ExchangeQueries) UpdateExchange(id uuid.UUID, e *models.Exchange) error {
    query := `UPDATE exchanges SET updated_at = $2, name = $3, description = $4, uri = $5, is_enabled = $6, is_blocked = $7 WHERE id = $1`

    _, err := q.Exec(query, e.ID, e.UpdatedAt, e.Name, e.Description, e.Uri, e.IsEnabled, e.IsBlocked)
    if err != nil {
        return err
    }

    return nil
}

func (q *ExchangeQueries) DeleteExchange(id uuid.UUID) error {
    query := `UPDATE exchanges SET is_deleted = 1 WHERE id = $1`

    _, err := q.Exec(query, id)
    if err != nil {
        return err
    }

    return nil
}
