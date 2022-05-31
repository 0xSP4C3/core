package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/jmoiron/sqlx"
)

type ExchangeQueries struct {
    *sqlx.DB
}


func (q *ExchangeQueries) GetExchanges() ([]models.Exchange, error) {
    exchange := []models.Exchange{}

    query := `SELECT * FROM exchanges`

    err := q.Get(&exchange, query)
    if err != nil {
        return exchange, err
    }

    return exchange, nil
}
