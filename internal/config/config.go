package config

import (
	"os"

	"github.com/creamdog/gonfig"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	TIMESTAMP_FORMAT = "2006-01-02 15:04:05"
)

type Config struct {
	ServerConfig ServerConfig
	LogLevel     string
	GinMode      string
}

type ServerConfig struct {
	Port    string
	Host    string
	Version string
}

func InitConfig() (*Config, error) {
	cp := os.Getenv("CONFIG_PATH")
	f, err := os.Open(cp)
	c := defaultLocalConf()

	if err != nil {
		if os.IsNotExist(err) {
			// if configmap not mounted
			// init config file with default values for local docker run
			SetLogLevel(c.LogLevel)
			return c, nil
		}
		return c, err
	}

	defer f.Close()

	cfg, err := gonfig.FromJson(f)
	if err != nil {
		return c, err
	}

	var sc ServerConfig
	if err := cfg.GetAs("server", &sc); err != nil {
		return c, err
	}
	os.Setenv("VERSION", sc.Version)

	ll, err := cfg.GetString("loglevel", "info")

	if err != nil {
		return c, err
	}

	SetLogLevel(ll)

	gm, err := cfg.GetString("ginmode", gin.DebugMode)

	if err != nil {
		return c, err
	}

	c = &Config{
		ServerConfig: sc,
		LogLevel:     ll,
		GinMode:      gm,
	}

	return c, nil
}

func SetLogLevel(lvl string) {
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.InfoLevel
	}
	old := log.GetLevel()
	// set global log level
	if ll != old {
		log.SetLevel(ll)
		log.Debugf("Switch log level from %s to %s", old, ll)
	}

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = TIMESTAMP_FORMAT
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

// test local only
func defaultLocalConf() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			Port:    "8080",
			Host:    "0.0.0.0",
			Version: "local",
		},
		LogLevel: "debug",
		GinMode:  gin.ReleaseMode,
	}
}
