package jwt_auth

import (
	"github.com/uptrace/bun"
	"net/http"
)

type (
	SignInReqBody struct {
		Uname    string `json:"username"`
		Password string `json:"password"`
	}
	SignInResp struct {
		ID any `json:"id"`
		SignUpReqBody
	}
)

func makeSigninHandler(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
