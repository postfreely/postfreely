package postfreely

import (
	"context"
	"errors"
	"fmt"
	"github.com/postfreely/web-core/log"
	"net/http"
	"net/url"
	"strings"
)

type giteaOauthClient struct {
	ClientID         string
	ClientSecret     string
	AuthLocation     string
	ExchangeLocation string
	InspectLocation  string
	CallbackLocation string
	Scope            string
	MapUserID        string
	MapUsername      string
	MapDisplayName   string
	MapEmail         string
	HttpClient       HttpClient
}

var _ oauthClient = giteaOauthClient{}

const (
	giteaDisplayName = "Gitea"
)

func (c giteaOauthClient) GetProvider() string {
	return "gitea"
}

func (c giteaOauthClient) GetClientID() string {
	return c.ClientID
}

func (c giteaOauthClient) GetCallbackLocation() string {
	return c.CallbackLocation
}

func (c giteaOauthClient) buildLoginURL(state string) (string, error) {
	u, err := url.Parse(c.AuthLocation)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("client_id", c.ClientID)
	q.Set("redirect_uri", c.CallbackLocation)
	q.Set("response_type", "code")
	q.Set("state", state)
	q.Set("scope", c.Scope)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c giteaOauthClient) exchangeOauthCode(ctx context.Context, code string) (*TokenResponse, error) {
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", c.CallbackLocation)
	form.Add("scope", c.Scope)
	form.Add("code", code)
	req, err := http.NewRequest("POST", c.ExchangeLocation, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("User-Agent", ServerUserAgent(""))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ClientID, c.ClientSecret)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unable to exchange code for access token")
	}

	var tokenResponse TokenResponse
	if err := limitedJsonUnmarshal(resp.Body, tokenRequestMaxLen, &tokenResponse); err != nil {
		return nil, err
	}
	if tokenResponse.Error != "" {
		return nil, errors.New(tokenResponse.Error)
	}
	return &tokenResponse, nil
}

func (c giteaOauthClient) inspectOauthAccessToken(ctx context.Context, accessToken string) (*InspectResponse, error) {
	req, err := http.NewRequest("GET", c.InspectLocation, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("User-Agent", ServerUserAgent(""))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unable to inspect access token")
	}

	// since we don't know what the JSON from the server will look like, we create a
	// generic interface and then map manually to values set in the config
	var genericInterface map[string]interface{}
	if err := limitedJsonUnmarshal(resp.Body, infoRequestMaxLen, &genericInterface); err != nil {
		return nil, err
	}

	// map each relevant field in inspectResponse to the mapped field from the config
	var inspectResponse InspectResponse
	inspectResponse.UserID, _ = genericInterface[c.MapUserID].(string)
	// log.Info("Userid from Gitea: %s", inspectResponse.UserID)
	if inspectResponse.UserID == "" {
		log.Error("[CONFIGURATION ERROR] Gitea OAuth provider returned empty UserID value (`%s`).\n  Do you need to configure a different `map_user_id` value for this provider?", c.MapUserID)
		return nil, fmt.Errorf("no UserID (`%s`) value returned", c.MapUserID)
	}
	inspectResponse.Username, _ = genericInterface[c.MapUsername].(string)
	inspectResponse.DisplayName, _ = genericInterface[c.MapDisplayName].(string)
	inspectResponse.Email, _ = genericInterface[c.MapEmail].(string)

	return &inspectResponse, nil
}
