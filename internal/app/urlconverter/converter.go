package urlconverter

import (
    "math/rand"
    "strings"
)

const (
    alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
    length   = len(alphabet)
)

func GenerateShortUrl(urlLength int) string {
    var shortUrl strings.Builder

    for i := 0; i < urlLength; i++ {
        shortUrl.WriteByte(alphabet[rand.Intn(length)%length])
    }

    return shortUrl.String()
}
