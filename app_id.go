package main

import (
	"encoding/json"
	"io"

	"github.com/hashicorp/vault/api"
)

// AppIDAuth is used to perform token backend operations on Vault.
type AppIDAuth struct {
	c     *api.Client
	token string
}

type LoginResponse struct {
	LeaseID       string `json:"lease_id,omitempty"`
	Renewable     bool   `json:"renewable,omitempty"`
	LeaseDuration int    `json:"lease_durationm,omitempty"`
	Auth          struct {
		ClientToken string `json:"client_token,omitempty"`
	} `json:"auth,omitempty"`
}

func NewAppIDAuth(c *api.Client) *AppIDAuth {
	return &AppIDAuth{c: c}
}

func (c *AppIDAuth) Login(vars map[string]string) error {
	r := c.c.NewRequest("POST", "/v1/auth/app-id/login")
	if err := r.SetJSONBody(vars); err != nil {
		return err
	}

	resp, err := c.c.RawRequest(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	loginResponse, err := ParseLoginResponse(resp.Body)
	if err != nil {
		return err
	}
	c.token = loginResponse.Auth.ClientToken
	return nil
}

func (c *AppIDAuth) Token() string {
	return c.token
}

func ParseLoginResponse(r io.Reader) (*LoginResponse, error) {
	// First decode the JSON into a map[string]interface{}
	var loginResponse LoginResponse
	dec := json.NewDecoder(r)
	if err := dec.Decode(&loginResponse); err != nil {
		return nil, err
	}
	return &loginResponse, nil
}
