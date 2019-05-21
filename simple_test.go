package UprinoAuthClient

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
    token := TokenCheck{
    	Token: "Test",
    	PublicKey: "Test",
    	PrivateKey: "Test",
	}
	res, err := token.Check()
	fmt.Println(err)
	fmt.Println(res)
}
