package database

import (
	"context"
	"log"

	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/at-syot/be_auth/internal/database/models"
)

func NewDB(ctx context.Context) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, ":memory:")
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv(),
	))

	_, err = db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user table created.")

	return db, nil
}
