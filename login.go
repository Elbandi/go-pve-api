package pve

import (
	"time"
)

type loginParams struct {
	Username string   `url:"username"`
	Password string   `url:"password"`
	Realm    string   `url:"realm"`
}

type Login struct {
	Ticket              string `json:"ticket"`
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	Username            string `json:"username"`
//TODO: cap
}

func (client *PveClient) Login() (*Login, error) {
	loginParams := &loginParams{Username: client.username, Password:client.password, Realm:client.realm}
	login := &struct {
		Data Login `json:"data"`
	}{}
	_, err := client.sling.New().Post("access/ticket").BodyForm(loginParams).ReceiveSuccess(&login)
	if err != nil {
		return nil, err
	}
	client.httpClient.ticket = login.Data.Ticket
	client.httpClient.token = login.Data.CSRFPreventionToken
	client.httpClient.timestamp = time.Now()
	return &login.Data, nil
}

func (client *PveClient) CheckLogin() (error) {
	if client.httpClient.ticket == "" || client.httpClient.timestamp.Before(time.Now().Add(-time.Hour)) {
		_, err := client.Login()
		return err
	}
	return nil
}