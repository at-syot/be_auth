package jwt_auth

import (
	"encoding/json"
	"github.com/at-syot/be_auth/internal/database/models"
	"github.com/at-syot/be_auth/pkg/cipher"
	"github.com/at-syot/be_auth/pkg/httpx"
	"github.com/uptrace/bun"
	"log"
	"net/http"
)

type (
	SignInReqBody struct {
		Uname    string `json:"username"`
		Password string `json:"password"`
	}
	SignInResp struct {
		Token string `json:"token"`
	}
)

func makeSigninHandler(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer r.Body.Close()
		var reqBody SignInReqBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			log.Printf("signin.handler:decode req's body err: %v", err)
			httpx.WriteInternalErrResp(w)
			return
		}

		u := new(models.User)
		if err := bun.NewSelectQuery(db).
			Model(u).
			Where("uname = ?", reqBody.Uname).
			Limit(1).
			Scan(ctx); err != nil {
			httpx.WriteNotfoundResp(w)
			return
		}

		if err := cipher.CheckHashWithPassword(u.Password, reqBody.Password); err != nil {
			httpx.WriteUnauthResp(w)
			return
		}

		token, err := cipher.JWTSign()
		if err != nil {
			httpx.WriteInternalErrResp(w)
			return
		}

		httpx.WriteOKResp(w, SignInResp{token})
	}
}
