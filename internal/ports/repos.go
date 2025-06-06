package ports

import (
	"context"
	"io"

	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/pkg/dbx"
)

// Migration is a method that runs the migration.
type Migration interface {
	// Migrate is a method that runs the migration.
	Migrate(context.Context) error
}

// Datastore provides methods for transactional operations.
type Datastore interface {
	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error
	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error

	io.Closer
	Migration
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	// GetSession is a method that returns a session by ID
	GetSession(ctx context.Context, session *models.Session) error
	// ListSessions is a method that lists all sessions
	ListSessions(ctx context.Context, session *dbx.Results[models.Session]) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx
}
