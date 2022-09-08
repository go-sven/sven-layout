package data

import (
	"context"
	"gorm.io/gorm"
)

type Transaction interface {
	ExecTx (context.Context,func(c context.Context) error ) error
}
type contextTxKey struct {}

type Transact struct {
	db *gorm.DB
}

func NewTract(db *gorm.DB) *Transact {
	return &Transact{db: db}
}

func NewTransaction(trans *Transact) Transaction  {
	return trans
}

func (t *Transact) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (t *Transact) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return t.db
}
