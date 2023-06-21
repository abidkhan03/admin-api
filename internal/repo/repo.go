package repo

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"

	"github.com/spongeling/admin-api/internal/dao"
)

type DB *gorm.DB

type Repo struct {
	conn *gorm.DB
}

func (r *Repo) db(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return r.conn
}

func (r *Repo) Close() error {
	sql, err := r.conn.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func New(dsn string) (*Repo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		FullSaveAssociations: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond)
		},
		PrepareStmt:     true,
		DryRun:          false,
		CreateBatchSize: 10000,
		Logger:          logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &Repo{
		conn: db,
	}, nil
}

func (r *Repo) GetByID(ctx context.Context, m dao.Model, id uint64) error {
	return wrap(r.db(ctx).First(m, id).Error)
}

func (r *Repo) GetWithAssociations(ctx context.Context, m dao.Model, id uint64) error {
	return wrap(r.db(ctx).Preload(clause.Associations).First(m, id).Error)
}

func (r *Repo) DeleteByID(ctx context.Context, m dao.Model, id uint64) error {
	return wrap(r.db(ctx).Delete(m, id).Error)
}

func (r *Repo) Create(ctx context.Context, m dao.Model) error {
	if err := m.Validate(); err != nil {
		return err
	}

	return wrap(r.db(ctx).Create(m).Error)
}

func (r *Repo) Update(ctx context.Context, m dao.Model) error {
	if err := m.Validate(); err != nil {
		return err
	}

	return wrap(r.db(ctx).Save(m).Error)
}

func wrap(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return err
}

func wrapf(format string, a ...interface{}) error {
	return wrap(fmt.Errorf(format, a...))
}

type transactionKey string

const (
	txKey transactionKey = `tx`
)

func (r *Repo) RunTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx := r.db(ctx).Begin()

	rollback := func() {
		err := tx.Rollback().Error
		if err != nil {
			log.Printf("error rolling back transaction: %s", err)
		}
	}

	err := f(context.WithValue(ctx, txKey, tx))
	if err != nil {
		rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		rollback()
		return wrapf("error committing transaction: %s", err)
	}

	return nil
}
