package main

import (
    "database/sql"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/denieryd/url-shorter/internal/app/apiserver"
    "github.com/denieryd/url-shorter/internal/app/model"
    "github.com/denieryd/url-shorter/internal/app/store"
    "github.com/denieryd/url-shorter/internal/app/urlconverter"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    log "github.com/sirupsen/logrus"
    "math/rand"
    "net/http"
    "time"
)

var repo = store.Repository{}

func main() {
    rand.Seed(time.Now().UnixNano())

    config := apiserver.NewConfig()
    if _, err := toml.DecodeFile("configs/config.toml", config); err != nil {
        log.Fatalf("config cannot be read %v", err)
    }

    if _, err := log.ParseLevel(config.LogLevel); err != nil {
        log.Errorf("logging level cannot be properly set %v", err)
    }

    dbStore := store.GetNewStore()

    if err := dbStore.Open(); err != nil {
        log.Fatalf("Database cannot be opened %v", err)
    }

    if err := dbStore.InitDB(); err != nil {
        log.Fatalf("db cannot be initialized %v", err)
    }

    repo.Store = dbStore

    log.Infof("server setup completed. Config is %+v", config)

    router := gin.Default()
    router.GET("/original_url/:ShortenedURL", GetOriginalUrl)
    router.POST("/", CreateShortenedUrl)

    if err := router.Run("localhost:8080"); err != nil {
        log.Fatalf("server cannot be run %v", err)
    }
}

func GetOriginalUrl(c *gin.Context) {
    ShortenedURL := c.Param("ShortenedURL")
    originalUrl, err := repo.GetOriginalURL(ShortenedURL)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusBadRequest, gin.H{"message": "no results"})
            return
        }
        log.Warning(err)
        c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
        return
    }
    c.JSON(http.StatusOK, originalUrl)
}

func CreateShortenedUrl(c *gin.Context) {
    url := model.URL{}

    if err := c.BindJSON(&url); err != nil {
        c.String(http.StatusBadRequest, "Not found originalUrl parameter")
        return
    }

    uuid_val, err := uuid.NewUUID()
    if err != nil {
        fmt.Printf("uuid is generated with err, %s", err)
        c.String(http.StatusInternalServerError, "Internal server error")
        return
    }

    shortenedUrl := urlconverter.GenerateShortUrl(10)
    url.ShortenedURL = shortenedUrl
    url.UUID = uuid_val

    repo.CreateURL(&url)

    c.JSON(http.StatusCreated, gin.H{"short_url": url.ShortenedURL})
}
