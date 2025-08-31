package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    SOCKS struct {
        Listen string `json:"listen"`
    } `json:"socks"`
}

func LoadConfig(path string) (Config, error) {
    var cfg Config
    data, err := os.ReadFile(path)
    if err != nil {
        return cfg, err
    }
    err = json.Unmarshal(data, &cfg)
    return cfg, err
}
