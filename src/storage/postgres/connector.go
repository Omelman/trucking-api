package postgres

import (
	"context"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Omelman/trucking-api/src/config"
)

// Postgres bun connection.
type Postgres struct {
	*pg.DB
	ctx context.Context
}

func New(ctx context.Context, wg *sync.WaitGroup, cfg *config.Postgres) (*Postgres, error) {
	writeTimeout, err := time.ParseDuration(cfg.WriteTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "write timeout")
	}

	readTimeout, err := time.ParseDuration(cfg.ReadTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "read timeout")
	}

	conn := pg.Connect(&pg.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		User:         cfg.User,
		Password:     cfg.Password,
		Database:     cfg.Name,
		PoolSize:     cfg.PoolSize,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		MaxRetries:   cfg.MaxRetries,
	})

	p := &Postgres{DB: conn, ctx: ctx}

	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		err := conn.Close()
		if err != nil {
			log.Error("close db connection error:", err.Error())

			return
		}

		log.Info("close db connection")
	}()

	return p, p.Ping(ctx)
}
