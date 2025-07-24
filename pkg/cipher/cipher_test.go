package cipher

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

type (
	RepoErr struct {
		e error
	}
	SvsErr struct {
		e error
	}
	HandlerErr struct {
		e error
	}
)

func (re *RepoErr) Error() string {
	return fmt.Sprintf("#repo err -> %v", re.e)
}

func (re *RepoErr) Is(e error) bool {
	fmt.Println("repo.Is is called")
	// e.(*RepoErr)
	return false
}

// SvsErr -----------
func (s *SvsErr) Error() string {
	return fmt.Sprintf("#service err -> %v", s.e)
}

func (s *SvsErr) Is(e error) bool {
	fmt.Println("svsErr.is is called?")
	return s.e == e
}

func (s *SvsErr) Unwrap() error {
	fmt.Println("svsErr unwrap is called?")
	return s.e
}

// -------------

func (s *HandlerErr) Error() string {
	return fmt.Sprintf("#Handler err -> %v", s.e)
}

func (s *HandlerErr) Is(e error) bool {
	fmt.Println("HandlerErr.is is called?")
	return s.e == e
}

func (s *HandlerErr) Unwrap() error {
	fmt.Println("HandlerErr unwrap is called?")
	return s.e
}

func TestErrors(t *testing.T) {
	dbErr := sql.ErrNoRows
	repErr := &RepoErr{dbErr}
	svsErr := &SvsErr{repErr}
	hErr := &HandlerErr{svsErr}
	isOk := errors.Is(hErr, repErr)
	t.Logf("isOk: %v", isOk)

	// if !isDbErr0 {
	// 	t.Fatal()
	// }
	// t.Log(svsErr)
	//
	// unwrapToRoot := func(e error) error {
	// 	for {
	// 		err := errors.Unwrap(e)
	// 		if err == nil {
	// 			break
	// 		}
	// 		e = err
	// 	}
	// 	return e
	// }
	//
	// rootE := unwrapToRoot(svsErr)
	// t.Logf("rootE of svsE: %v", rootE)

	// target := sql.ErrNoRows
	// isAs := errors.As(svsErr, &target)
	// t.Logf("isAs %v\n", isAs)
	// t.Logf("now svsErr %v", svsErr)
}

func TestJWTSign(t *testing.T) {
	ss, err := JWTSign()
	if err != nil {
		t.Fatal(err)
	}

	claims, err := JWTParse(ss)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("claims %+v", claims)
}
