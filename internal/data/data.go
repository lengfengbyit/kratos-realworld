package data

import (
	"context"
	"errors"
	"fmt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/data/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo, NewProfileRepo, NewFollowRepo, NewArticleRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	lh := log.NewHelper(logger)

	drv, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Errorf("failed opening connection to db: %v", err)
		return nil, nil, err
	}

	// 链路追踪，打印 SQL
	sqlDriver := dialect.DebugWithContext(drv, func(ctx context.Context, a ...any) {
		if c.Database.Debug {
			lh.Info(a...)
		}
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx, "Query",
			trace.WithAttributes(attribute.String("sql", fmt.Sprint(a...))),
			trace.WithSpanKind(kind))
		span.End()
	})

	client := ent.NewClient(ent.Driver(sqlDriver))

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{db: client}

	cleanup := func() {
		lh.Info("message", "closing the data resources")
		if err = d.db.Close(); err != nil {
			lh.Error(err)
		}
	}

	return d, cleanup, nil
}

// WithTx 事务，自动提交或回滚
func (d *Data) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := d.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Join(err, rerr)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
