package config

import (
	"errors"
)

type serverCfg struct {
    Port int `yaml:"port"`
    AllowedOrigins []string `yaml:"allowedOrigins"`
}

type dbCfg struct {
    Host     string `yaml:"host"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Port     int    `yaml:"port"`
    Database string `yaml:"database"`
    Schema   string `yaml:"schema"`
    SSLMode  string `yaml:"sslmode"`
}

type authCfg struct {
    Secret string `yaml:"secret"`
}

type logCfg struct {
    Path string `yaml:"path"`
}

type config struct {
    Server *serverCfg `yaml:"server"`
    Auth *authCfg `yaml:"auth"`
    Db *dbCfg `yaml:"db"`
    Log *logCfg `yaml:"log"`
}

var Config = &config{
    Server: &serverCfg{},
    Auth: &authCfg{
        Secret: "17G0W9V6aMGJ",
    },
    Db: &dbCfg{
        Schema: "media",
        SSLMode: "disable",
    },
    Log: &logCfg{
        Path: "/tmp/verse.log",
    },
}

func Validate() error {
    server := Config.Server
    if server.Port == 0 {
        return errors.New("Invalid server port")
    }
    db := Config.Db
    if db.Host == "" || db.User == "" || db.Password == "" || db.Database == "" || db.Schema == "" || db.SSLMode == "" {
        return errors.New("Invalid db config")
    }
    auth := Config.Auth
    if auth.Secret == "" {
        return errors.New("Invalid jwt secret")
    }
    log := Config.Log
    if log.Path == "" {
        return errors.New("Invalid log path")
    }

    return nil
}

