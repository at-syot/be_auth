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
	go jwt_auth.MakeSignUpClient()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	log.Println("server is closing..")
	time.Sleep(time.Millisecond * 500)
	s.Close()
}
