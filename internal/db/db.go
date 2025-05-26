package db

import (
	"context"

	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/m/internal/ports"

	"github.com/katallaxie/pkg/dbx"
	"gorm.io/gorm"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() dbx.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

// GetSession is a method that returns a session by ID
func (r *readTxImpl) GetSession(ctx context.Context, session *models.Session) error {
	return r.conn.
		Where(session).
		First(session, session.ID).
		Error
}

// ListSessions is a method that lists all sessions
func (r *readTxImpl) ListSessions(ctx context.Context, pagination *dbx.Results[models.Session]) error {
	return r.conn.
		Scopes(dbx.PaginatedResults(&pagination.Rows, pagination, r.conn)).
		Find(&pagination.Rows).
		Error
}

type writeTxImpl struct {
	conn *gorm.DB
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx() dbx.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}
