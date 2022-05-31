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
