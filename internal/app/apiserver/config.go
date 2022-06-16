package apiserver

type Config struct {
    BindAddr       string `toml:"bind_addr"`
    LogLevel       string `toml:"log_level"`
    ShortUrlLength int    `toml:"short_url_length"`
}

func NewConfig() *Config {
    return &Config{BindAddr: ":8080", LogLevel: "debug", ShortUrlLength: 10}
}
