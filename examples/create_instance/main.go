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

	// more about instance https://vrchatapi.github.io/tutorials/instances/
	err = vrc.InstanceLaunch("wrld_84f5858a-180b-42d8-9404-f71a99854354", vrcapi.InstanceId{
		Name:    "1234",
		Privacy: vrcapi.Friends,
		Region:  vrcapi.JP,
		OwnerID: vrc.User.Id,
		UUID:    "cf2fbae7-fadc-4a96-bdc3-485b1a5959e9",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}
