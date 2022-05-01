package sqlite

import (
	"context"
	"fmt"

	"github.com/e4t4g/URL_shortener_GB-/internal/app/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

type repoURL struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.Repository {
	return &repoURL{
		db: db,
	}
}

func (rdb *repoURL) Create(url *repository.URLData) (*repository.URLData, error) {
	var URLData repository.URLData
	ctx := context.Background()
	query := "INSERT INTO url (full_url, short_url, counter) VALUES (?, ?, ?)"
	statement, err := rdb.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.QueryRow(url.FullURL, url.ShortURL, 0)

	query = "SELECT * FROM url WHERE short_url = ?"

	err = rdb.db.GetContext(ctx, &URLData, query, url.ShortURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &URLData, nil
}

func (rdb *repoURL) FindByToken(ctx context.Context, token string) (*repository.URLData, error) {
	var URLData repository.URLData

	statQuery := "SELECT * FROM url WHERE short_url = ?"

	err := rdb.db.GetContext(ctx, &URLData, statQuery, token)
	if err != nil {
		return nil, err
	}

	return &URLData, nil
}

func (rdb *repoURL) FindByID(id int) (*repository.URLData, error) {
	var URLData repository.URLData

	ctx := context.Background()

	statQuery := "SELECT * FROM url WHERE id = ?"

	err := rdb.db.GetContext(ctx, &URLData, statQuery, id)
	if err != nil {
		return nil, err
	}

	return &URLData, err
}

func (rdb *repoURL) UpdateCounter(counter int64, shortURL string) {
	query := "UPDATE url SET counter = ? WHERE short_url = ? "
	statement, _ := rdb.db.Prepare(query)
	statement.QueryRow(counter, shortURL)
}
