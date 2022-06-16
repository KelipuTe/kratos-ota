package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewDB(t *testing.T) {
	jwtStr := MakeJWT("secretkey", "keine")
	spew.Dump(jwtStr)
}
