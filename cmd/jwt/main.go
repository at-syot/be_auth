package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/at-syot/be_auth/internal/jwt_auth"
	_ "modernc.org/sqlite"
)

func main() {
	log.SetPrefix("jwt_base - ")
	log.SetFlags(0)

	s := jwt_auth.NewServer()
	go s.Listen()

	signupDone := make(chan uint8)
	signinDone := make(chan uint8)
	defer close(signupDone)
	defer close(signinDone)

	go jwt_auth.MakeSignUpClient(signupDone)
	// after client is signup -> do signin
	go jwt_auth.MakeSigninClient(signupDone, signinDone)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	log.Println("server is closing..")
	time.Sleep(time.Millisecond * 500)
	s.Close()
}
