package main

import (
	"context"
	"fmt"
	"github.com/12storeez/auth-go/keycloak"
)

func main() {
	hostname := "https:....com"
	realm := "autorization-project"
	accessToken := "xxx"

	ctx := context.Background()
	client := keycloak.NewClient(hostname)
	info, err := client.GetUserInfo(ctx, accessToken, realm)
	if err != nil {
		fmt.Println("GetUserInfo failed:" + err.Error())
	}
	fmt.Println(info)

	roles, err := keycloak.GetUserRolesByClientID(accessToken, "logistics")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(roles)
}
