package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/vault/api"
)

var vaultClient *api.Client

func client() *api.Client {
	if vaultClient == nil {
		fmt.Println(config.VaultAddress)
		vaultConfig := api.DefaultConfig()
		vaultConfig.Address = config.VaultAddress
		vaultClient, _ = api.NewClient(vaultConfig)
	}
	return vaultClient
}

func login() {
	loginMap := make(map[string]string)
	loginMap["app_id"] = config.AppID
	loginMap["user_id"] = config.UserID
	appAuth := NewAppIDAuth(client())
	appAuth.Login(loginMap)
	vaultClient.SetToken(appAuth.Token())
	fmt.Println("logged in")
}

func writeFile(key string, outputFile string) {
	path := fmt.Sprintf("auth/app-id/map/%s/%s", config.UserID, config.Key)
	secret, err := client().Logical().Read(key)
	if err != nil {
		log.Fatalf("Error reading path %s\n", path)
	}
	decodedFile, err := base64.StdEncoding.DecodeString(secret.Data["value"].(string))
	err = ioutil.WriteFile(outputFile, decodedFile, 0644)
}

func main() {
	login()
	writeFile(config.Key, config.OutputFile)
}
