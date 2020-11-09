package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"time"
)

type ToolCredentials struct {
	Confluence struct {
		password string
		userid   string
		endpoint string
	}
	Jira struct {
		password string
		userid   string
		endpoint string
	}
	Crowd struct {
		password string
		userid   string
		endpoint string
	}
}

type vaultContext struct {
	userId   string
	password string
	endpoint string
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func FillToolCredentials(userId string, password string, endpoint string) *ToolCredentials {

	vaultContext := new(vaultContext)
	vaultContext.userId = userId
	vaultContext.password = password
	vaultContext.endpoint = endpoint

	creds := new(ToolCredentials)

	client := createClient(*vaultContext)

	getSecretConfluence(*client, creds)
	getSecretJira(*client, creds)

	return creds
}

func getSecretConfluence(client api.Client, creds *ToolCredentials) {

	secret, err := client.Logical().Read("ast/data/confluence")
	if err != nil {
		panic(err)
	}

	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}

	creds.Confluence.userid = fmt.Sprint(m["userid"])
	creds.Confluence.password = fmt.Sprint(m["password"])
	creds.Confluence.endpoint = fmt.Sprint(m["endpoint"])

}

func getSecretJira(client api.Client, creds *ToolCredentials) {

	secret, err := client.Logical().Read("ast/data/confluence")
	if err != nil {
		panic(err)
	}

	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}

	creds.Jira.userid = fmt.Sprint(m["userid"])
	creds.Jira.password = fmt.Sprint(m["password"])
	creds.Jira.endpoint = fmt.Sprint(m["endpoint"])

}

func createClient(ctx vaultContext) *api.Client {
	token, err := userpassLogin(ctx)
	if err != nil {
		panic(err)
	}

	client, err := api.NewClient(&api.Config{Address: ctx.endpoint, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(token)

	return client
}

func userpassLogin(ctx vaultContext) (string, error) {
	// create a vault client
	client, err := api.NewClient(&api.Config{Address: ctx.endpoint, HttpClient: httpClient})
	if err != nil {
		return "", err
	}

	// to pass the password
	options := map[string]interface{}{
		"password": ctx.password,
	}
	path := fmt.Sprintf("auth/ldap/login/%s", ctx.userId)

	// PUT call to get a token
	secret, err := client.Logical().Write(path, options)
	if err != nil {
		return "", err
	}

	token := secret.Auth.ClientToken

	return token, nil
}
