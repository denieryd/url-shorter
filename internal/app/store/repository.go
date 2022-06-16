package store

import (
    "database/sql"
    "github.com/denieryd/url-shorter/internal/app/model"
    log "github.com/sirupsen/logrus"
)

type Repository struct {
    Store *Store
}

func (r *Repository) CreateURL(u *model.URL) {
    r.Store.db.QueryRow("INSERT INTO urls (uuid, originalurl, shortenedurl) values ($1, $2, $3) RETURNING uuid",
        u.UUID, u.OriginalURL, u.ShortenedURL)

}

func (r *Repository) CheckURLExisting(shortUrl string) bool {
    var result string
    err := r.Store.db.QueryRow("SELECT shortenedurl from urls where shortenedurl = $1", shortUrl).Scan(&result)
    if err != nil {
        if err == sql.ErrNoRows {
            return false
        }
        log.Fatalf("URL exists checking ended with error, %v", err)
    }

    return true
}

func (r *Repository) GetOriginalURL(shortenedUrl string) (string, error) {
    var originalURL string
    err := r.Store.db.QueryRow("SELECT originalurl from urls where shortenedurl = $1", shortenedUrl).Scan(&originalURL)
    return originalURL, err
}
