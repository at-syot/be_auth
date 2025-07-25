package jwt_auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/at-syot/be_auth/internal/database/models"
	"github.com/at-syot/be_auth/pkg/cipher"
	"github.com/at-syot/be_auth/pkg/httpx"
	"github.com/uptrace/bun"
)

type (
	SignUpReqBody struct {
		Uname    string `json:"username"`
		Password string `json:"password"`
	}
	SignUpResp struct {
		ID any `json:"id"`
		SignUpReqBody
	}
)

func makeSignupHandler(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		defer r.Body.Close()
		var reqBody SignUpReqBody
		if err := decoder.Decode(&reqBody); err != nil {
			log.Printf("decode req's body err: %v", err)
			httpx.WriteInternalErrResp(w)
			return
		}

		hash, _ := cipher.HashPassword(reqBody.Password)

		// save to db
		u := &models.User{Uname: reqBody.Uname, Password: hash}
		_, err := db.NewInsert().Model(u).Exec(ctx)
		if err != nil {
			log.Printf("handler:insert user err - %v", err)
			httpx.WriteInternalErrResp(w)
			return
		}

		signinResp := SignUpResp{
			ID:            u.ID,
			SignUpReqBody: reqBody,
		}
		httpx.WriteOKResp(w, signinResp)
	}
}
