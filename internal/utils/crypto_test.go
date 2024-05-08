package utils

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	data, err := RSAEncrypt([]byte("HELLO,WORLD!"))
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
}
