package queries

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CoinQueries struct {
    *sqlx.DB
}

// GetCoins method for getting all coins.
func (q *CoinQueries) GetCoins() ([]models.Coin, error) {
    // Define coins variable
    coins := []models.Coin{}

    // Define query string
    query := `SELECT * FROM coins`

    // Send query to database
    err := q.Get(&coins, query)
    if err != nil {
        // return empty object and error.
        return coins, err
    }
    
    // Return query result
    return coins, nil
}

func (q *CoinQueries) GetCoinByExchangeID(exchangeID int) ([]models.Coin, error) {
    coins := []models.Coin{}

    query := `SELECT * FROM coins WHERE exchange_id = $1`

    err := q.Get(&coins, query, exchangeID)
    if err != nil {
        return coins, err
    }

    return coins, nil
} 

func (q *CoinQueries) GetCoin(id uuid.UUID) (models.Coin, error) {
    coin := models.Coin{}

    query := `SELECT * FROM coins WHERE id = $1`

    err := q.Get(&coin, query, id)
    if err != nil {
        return coin, err
    }

    return coin, nil
}

// CreateCoin method for creating coin by given Coin object
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

// Delete coin method for change is_deleted = 1 by given ID.
func (q *CoinQueries) DeleteCoin(id uuid.UUID) error {
    query := `UPDATE coins SET is_deleted = 1 WHERE id = $1`

    // Send query to database
    _, err := q.Exec(query, id)
    if err != nil {
        return err
    }

    return nil
}
