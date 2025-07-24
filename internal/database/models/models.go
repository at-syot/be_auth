package models

import "github.com/uptrace/bun"

type (
	User struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		ID       int    `bun:",pk,autoincrement"`
		Uname    string `bun:",unique,notnull"`
		Password string `bun:",notnull"`
	}
)
