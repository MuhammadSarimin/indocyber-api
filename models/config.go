package models

import "strings"

type Config struct {
	AppName string `env:"APP_NAME"`
	AppHost string `env:"APP_HOST" envDefault:"0.0.0.0"`
	AppPort string `env:"APP_PORT" envDefault:"9999"`
	PathLog string `env:"PATH_LOG"  envDefault:"log"`
	DB      DBConfig
}

type DBConfig struct {
	Host        string `env:"DB_HOST"     envDefault:"localhost"`
	Port        string `env:"DB_PORT"     envDefault:"5432"`
	User        string `env:"DB_USER"     envDefault:"postgres"`
	Pass        string `env:"DB_PASSWORD" envDefault:""`
	Name        string `env:"DB_NAME"     envDefault:"postgres"`
	SSL         string `env:"DB_SSL"      envDefault:"disable"`
	AutoMigrate bool   `env:"DB_AUTO_MIGRATE" envDefault:"true"`
}

func (c *Config) Address() string {
	return c.AppHost + ":" + c.AppPort
}

func (c *DBConfig) DSN() string {
	var s strings.Builder

	if c.Host != "" {
		c.writeParam(&s, "host", c.Host)
	}

	if c.Port != "" {
		c.writeParam(&s, "port", c.Port)
	}

	if c.User != "" {
		c.writeParam(&s, "user", c.User)
	}

	if c.Pass != "" {
		c.writeParam(&s, "password", c.Pass)
	}

	if c.Name != "" {
		c.writeParam(&s, "dbname", c.Name)
	}

	if c.SSL == "" {
		c.SSL = "disable"
	}

	c.writeParam(&s, "sslmode", c.SSL)

	return s.String()
}

func (c *DBConfig) writeParam(s *strings.Builder, key, value string) {
	s.WriteByte(' ')
	s.WriteString(key)
	s.WriteByte('=')
	s.WriteString(value)
}
