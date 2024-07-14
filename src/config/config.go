package config

import (
	"errors"
	"time"
)

type serverCfg struct {
    Port int `yaml:"port"`
    AllowedOrigins []string `yaml:"allowedOrigins"`
}

type dbPoolCfg struct {
    MaxConns int `yaml:"maxConns"`
    MinConns int `yaml:"minConns"`
    MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
    MaxConnIdleTime time.Duration `yaml:"maxConnIdleTime"`
    HealthCheckPeriod time.Duration `yaml:"healthCheckPeriod"`
    ConnTimeout time.Duration `yaml:"connTimeout"`
}

type dbCfg struct {
    Host     string `yaml:"host"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Port     int    `yaml:"port"`
    Database string `yaml:"database"`
    Schema   string `yaml:"schema"`
    SSLMode  string `yaml:"sslmode"`
    Pool     *dbPoolCfg `yaml:"pool"`
}

type fileCfg struct {
    MediaRoot string `yaml:"mediaRoot"`
    ThumbRoot string `yaml:"thumbRoot"`
}

type authCfg struct {
    Secret string `yaml:"secret"`
    MediaSecret string `yaml:"mediaSecret"`
}

type logCfg struct {
    Path string `yaml:"path"`
}

type config struct {
    Server *serverCfg `yaml:"server"`
    Auth *authCfg `yaml:"auth"`
    Db *dbCfg `yaml:"db"`
    File *fileCfg `yaml:"file"`
    Log *logCfg `yaml:"log"`
}

var Config = &config{
    Server: &serverCfg{},
    Auth: &authCfg{
        Secret: "17G0W9V6aMGJ",
        MediaSecret: "eat5yUa4j2zC",
    },
    Db: &dbCfg{
        Schema: "media",
        SSLMode: "disable",
        Pool: &dbPoolCfg{
            MaxConns: 10,
            MinConns: 1,
            MaxConnLifetime: time.Hour,
            MaxConnIdleTime: time.Minute * 30,
            HealthCheckPeriod: time.Minute,
            ConnTimeout: time.Second * 5,
        },
    },
    File: &fileCfg{},
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

