package jwt_auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/at-syot/be_auth/pkg/httpx"
)

type (
	SignedUpUser struct {
		ID    any    `json:"id"`
		Uname string `json:"username"`
		Hash  string `json:"password"`
	}
)

// storage
var (
	signedUpUser  SignedUpUser
	signedInToken string
)

// ############

func MakeSignUpClient(done chan<- uint8) {
	time.Sleep(time.Second)
	client := http.Client{Timeout: time.Second}
	url := serverURL + "/signup"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Printf("create req err: %v", err)
	}

	// prepare cookie
	// c := &http.Cookie{Name: "uid", Value: "just uid value"}
	// req.AddCookie(c)

	body := SignUpReqBody{Uname: "aiosdev", Password: "password"}
	if err := httpx.ReqWithJSON(req, body); err != nil {
		log.Fatalf("marshal body err: %v", err)
	}

	httpResp, err := client.Do(req)
	if err != nil {
		log.Printf("do req err: %v", err)
	}
	defer httpResp.Body.Close()

	var resp httpx.Resp
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		log.Printf("decoding res.Body err: %v", err)
		return
	}

	// save signedUpUser in memory
	data, _ := resp.Data.(map[string]any)
	signedUpUser = SignedUpUser{
		ID:    data["id"],
		Uname: data["username"].(string),
		Hash:  data["password"].(string),
	}
	log.Printf("client:save signed user - %+v", signedUpUser)
	done <- 0
}

func MakeSigninClient(signupDone <-chan uint8, done chan<- uint8) {
	<-signupDone

	url := serverURL + "/signin"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Printf("create req err: %v", err)
	}

	body := SignUpReqBody{Uname: "aiosdev", Password: "password"}
	if err := httpx.ReqWithJSON(req, body); err != nil {
		log.Fatalf("marshal body err: %v", err)
	}

	httpResp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Printf("do req err: %v", err)
	}
	defer httpResp.Body.Close()

	resp := new(httpx.Resp)
	if err := httpx.DecodeStreamedV(httpResp.Body, resp); err != nil {
		log.Fatalf("decode body err: %v", err)
	}

	data, _ := resp.Data.(map[string]any)
	signedInToken = data["token"].(string)
	log.Printf("signedInToken %s", signedInToken)

	done <- 0
}
