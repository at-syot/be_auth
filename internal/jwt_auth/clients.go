package jwt_auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type SignedUpUser struct {
	ID    any    `json:"id"`
	Uname string `json:"username"`
	Hash  string `json:"password"`
}

// stored signedUser
var signedUpUser SignedUpUser

// ############

func MakeSignUpClient() {
	time.Sleep(time.Second)
	client := http.Client{Timeout: time.Second}
	url := serverURL + "/signup"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Printf("create req err: %v", err)
	}

	// prepare cookie
	c := &http.Cookie{Name: "uid", Value: "just uid value"}
	req.AddCookie(c)

	// preparing body
	req.Header.Add("Content-Type", "application/json")
	body := SignUpReqBody{
		Uname:    "aiosdev",
		Password: "password",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("marshal body err: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	httpResp, err := client.Do(req)
	if err != nil {
		log.Printf("do req err: %v", err)
	}
	defer httpResp.Body.Close()

	var resp Resp
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
}
