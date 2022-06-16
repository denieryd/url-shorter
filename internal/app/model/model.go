package model

import (
    "github.com/google/uuid"
)

type URL struct {
    UUID         uuid.UUID
    OriginalURL  string
    ShortenedURL string
}
