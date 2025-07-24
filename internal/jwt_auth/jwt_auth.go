package jwt_auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/at-syot/be_auth/internal/database"
	"log"
	"net/http"
)

// simulate server
// with routes [register, login, get protected resource]
const (
	serverPort = "8080"
	serverAddr = "localhost:" + serverPort
	serverURL  = "http://" + serverAddr
)

type (
	Server struct {
		s *http.Server
	}

	ReqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Age      int    `json:"age"`
	}
)

func (s *Server) Listen() {
	log.Println("serv addr :8080")
	if err := s.s.ListenAndServe(); err != nil {
		log.Fatalf("server closed err: %v", err)
	}
}

func (s *Server) Close() error {
	return s.s.Close()
}

func NewServer() *Server {
	ctx := context.Background()
	db, err := database.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("POST /signup", makeSignupHandler(db))
	mux.HandleFunc("POST /signin", makeSigninHandler(db))

	return &Server{s: &http.Server{Addr: serverAddr, Handler: mux}}
}

// handlers
type RespStatus string

const (
	RespOk   RespStatus = "ok"
	RespFail RespStatus = "fails"
)

type Resp struct {
	Status RespStatus `json:"status"`
	Data   any        `json:"data"`
}

func newOkResp(data any) Resp {
	return Resp{Status: RespOk, Data: data}
}
func newFailResp() Resp {
	return Resp{Status: RespFail, Data: nil}
}

func (r Resp) Bytes() []byte {
	w := new(bytes.Buffer)
	json.NewEncoder(w).Encode(r)
	return w.Bytes()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("header: %+v\n", r.Header)
	log.Printf("body: %v\n", r.Body)
	for _, c := range r.Cookies() {
		log.Println(c)
	}

	var reqBody ReqBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		log.Printf("decode req's body err: %v", err)
	}

	fmt.Printf("reqBody	%+v \n", reqBody)
	fmt.Fprintf(w, "hello world")
}
