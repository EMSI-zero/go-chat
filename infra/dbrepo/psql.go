package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type psqlDB struct {
	GormConnectionPool *gorm.DB
}

func (cfg dbConfig) toPSQLConnStr() string {
	var connectionString strings.Builder

	connectionType := os.Getenv("PSQL_CONNNECTION_TYPE")

	connectionString.WriteString("postgres://")

	// Add username and password if provided
	if cfg.User != "" && cfg.Password != "" {
		connectionString.WriteString(cfg.User)
		connectionString.WriteByte(':')
		connectionString.WriteString(cfg.Password)
		connectionString.WriteByte('@')
	}

	// Add host and port
	connectionString.WriteString(cfg.Host)

	if cfg.Port != 0 && connectionType != "SRV" {
		connectionString.WriteByte(':')
		connectionString.WriteString(strconv.Itoa(int(cfg.Port)))
	}

	// Add database name
	connectionString.WriteByte('/')
	connectionString.WriteString(cfg.Database)

	// Add SSL mode if provided
	if connectionType != "SRV" {
		connectionString.WriteString("?sslmode=disable")
	}
	log.Print(connectionString.String())
	return connectionString.String()
}

func NewPSQLConn() (db *psqlDB, err error) {
	cfg := &dbConfig{}
	if err := cfg.buildConfigFromEnv(); err != nil {
		return nil, fmt.Errorf("could not load config, %w", err)
	}

	gormLogEnv := os.Getenv(LogGormEnv)
	if gormLogEnv != "" {
		if GormLog, err = strconv.ParseBool(gormLogEnv); err != nil {
			return nil, fmt.Errorf("couldn't parse %v env value '%s': %w",
				LogGormEnv, gormLogEnv, err)
		}
	} else {
		GormLog = false
	}

	connStr := cfg.toPSQLConnStr()

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(int(cfg.MaxIdleConnection))
	conn.SetMaxOpenConns(int(cfg.MaxOpenConnection))

	if err := conn.Ping(); err != nil {
		log.Print(err)
		return nil, err
	}

	GormConnectionPool, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	GormDebugLogger = logger.New(log.Default(), logger.Config{
		SlowThreshold:             365 * 24 * time.Hour,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	db = &psqlDB{
		GormConnectionPool: GormConnectionPool,
	}

	return db, nil
}

func (db *psqlDB) GetConn(ctx context.Context) (*gorm.DB, error) {
	if db.GormConnectionPool == nil {
		return nil, fmt.Errorf("no database connection found")
	}
	return db.GormConnectionPool.WithContext(ctx), nil
}
