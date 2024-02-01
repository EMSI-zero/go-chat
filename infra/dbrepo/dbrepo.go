package dbrepo

import (
	"context"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbConfig struct {
	Host              string
	Port              uint16
	Database          string
	User              string
	Password          string
	LogLevel          uint16
	MaxOpenConnection uint16
	MaxIdleConnection uint16
}

type DBConn interface {
	GetConn(context.Context) (*gorm.DB, error)
}

var DbHost = "DB_HOST"
var DbDatabase = "DB_DATABASE"
var DbPort = "DB_PORT"
var DbUser = "DB_USER"
var DbPass = "DB_PASS"
var GormLogMODE = "GORM_LOG_MODE"
var LogGormEnv = "LOG_GORM_ENV"
var GormLog bool
var GormMaxOpenConn = "MAX_OPEN_CONN"
var GormMaxIdleConn = "MAX_IDLE_CONN"
var GormDebugLogger logger.Interface

func (cfg *dbConfig) buildConfigFromEnv() error {
	cfg.Host = os.Getenv(DbHost)
	if cfg.Host == "" {
		return fmt.Errorf("%v env must be specified: db host", DbHost)
	}

	cfg.Database = os.Getenv(DbDatabase)
	if cfg.Database == "" {
		return fmt.Errorf("%v env must be specified: database", DbDatabase)
	}

	portStr := os.Getenv(DbPort)
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: db port", DbPort)
	}

	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", DbPort, portStr, err)
	}

	cfg.Port = uint16(port)

	cfg.User = os.Getenv(DbUser)
	if cfg.User == "" {
		return fmt.Errorf("%v env must be specified: db user", DbUser)
	}

	cfg.Password = os.Getenv(DbPass)
	if cfg.Password == "" {
		return fmt.Errorf("%v env must be specified: db password", DbPass)
	}

	logModeStr := os.Getenv(GormLogMODE)
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: gorm log mode", GormLogMODE)
	}

	logMode, err := strconv.ParseUint(logModeStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", GormLogMODE, logModeStr, err)
	}

	cfg.LogLevel = uint16(logMode)

	maxOpenConnStr := os.Getenv(GormMaxOpenConn)
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: max open connection", GormMaxOpenConn)
	}

	maxOpenConn, err := strconv.ParseUint(maxOpenConnStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", GormMaxOpenConn, maxOpenConnStr, err)
	}

	cfg.MaxOpenConnection = uint16(maxOpenConn)

	maxIdleConnStr := os.Getenv(GormMaxIdleConn)
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: max idle connection", GormMaxIdleConn)
	}

	maxIdleConn, err := strconv.ParseUint(maxIdleConnStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", GormMaxIdleConn, logModeStr, err)
	}

	cfg.MaxIdleConnection = uint16(maxIdleConn)

	return nil
}

func NewDBConn() (DBConn, error) {
	return NewPSQLConn()
}
