package mySqlExt

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func (m *mySqlExt) ExecTx(ctx context.Context, fn func(tx IMySqlExt) error) error {
	var (
		start = time.Now()

		span     trace.Span
		duration time.Duration
		err      error
	)

	ctx, span = m.OtelTracer().Start(ctx, "MySqlExt.ExecTx")

	defer func(ctx context.Context, query string, tableName string, start *time.Time) {
		duration = time.Since(*start)

		m.InstrumentMetric(ctx, query, tableName, &duration)
		span.SetAttributes(
			semconv.DBSystemMySQL,
			semconv.DBNameKey.String(m.DBName()),
			semconv.DBSQLTableKey.String(tableName),
			semconv.DBStatementKey.String(query),
			attribute.Float64("duration", duration.Seconds()),
		)
		span.End()
	}(ctx, "ExecTx", m.TableName(ctx), &start)

	ctx, err = m.BeginTxx(ctx)
	if err != nil {
		return err
	}

	err = fn(m)
	if err != nil {
		if rbErr := m.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return m.Commit(ctx)
}

func (m *mySqlExt) GetSchema() string {
	return m.schemaName
}
