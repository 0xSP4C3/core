package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/jmoiron/sqlx"
)

type CoinQueries struct {
    *sqlx.DB
}


func (q *CoinQueries) CreateCoin(c *models.Coin) error {
    // Define query string.
    query := `INSERT INTO coins VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

    // Send query to database
    _, err := q.Exec(query, c.ID, c.CreatedAt, c.UpdatedAt, c.ExchangeId, c.Name, c.Code, c.Description, c.IsDeleted, c.ImageUri, c.CoinUri)

    if err != nil {
        // throw err to upper level
        return err
    }

    // This query should return nothing.
    return nil
}
