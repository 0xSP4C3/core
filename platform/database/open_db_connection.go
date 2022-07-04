package database

import "github.com/0xsp4c3/core/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.BookQueries // load queries from Book model
    *queries.CoinQueries // load queries from Coin model
    *queries.FeedQueries // load queries from Feed model
    *queries.FeedRangeQueries // load queries from FeedRange model
    *queries.ExchangeQueries // load queries from Exchange model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
		BookQueries: &queries.BookQueries{DB: db}, // from Book model
        CoinQueries: &queries.CoinQueries{DB: db}, // from Coin model
        FeedQueries: &queries.FeedQueries{DB: db},
        FeedRangeQueries: &queries.FeedRangeQueries{DB: db},
        ExchangeQueries: &queries.ExchangeQueries{DB: db},
	}, nil
}
