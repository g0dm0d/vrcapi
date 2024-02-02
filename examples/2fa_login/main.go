package main

import (
	"fmt"

	"github.com/g0dm0d/vrcapi"
)

func main() {
	vrc, err := vrcapi.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = vrc.Auth("username", "pass", "2fa secret")
	if err != nil {
		fmt.Println(err)
		return
	}
}
