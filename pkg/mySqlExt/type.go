package mySqlExt

import (
	"context"

	pdkMySql "github.com/paper-indonesia/pdk/v2/mySqlExt"
)

type IMySqlExt interface {
	pdkMySql.IMySqlExt

	ExecTx(ctx context.Context, fn func(tx IMySqlExt) error) error
	GetSchema() string
}

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxIdleTime  int
	MaxLifeTime  int
	MaxOpenConns int
}

type mySqlExt struct {
	pdkMySql.IMySqlExt

	schemaName string
}

type IMySqlRows interface {
	Next() bool
	Close() error
	Scan(dest ...any) error
	Err() error
}

func New(config pdkMySql.Config, opts ...pdkMySql.OptionFunc) (IMySqlExt, error) {
	pdkIMySql, err := pdkMySql.New(config, opts...)
	if err != nil {
		return nil, err
	}

	return &mySqlExt{pdkIMySql, config.DBName}, nil
}
