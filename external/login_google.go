package external

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/cryptowatch_challenge/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type LoginGoogleClient struct {
	config *config.Config
}

func NewLoginGoogleClient(config *config.Config) *LoginGoogleClient {
	return &LoginGoogleClient{
		config: config,
	}
}

var googleOauthConfig *oauth2.Config

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
func (s *LoginGoogleClient) initAuth() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  s.config.GoogleOauthRedirectUrl,
		ClientID:     s.config.GoogleOauthClientID,
		ClientSecret: s.config.GoogleOauthClientSecret,
		Scopes:       []string{s.config.GoogleOauthScope},
		Endpoint:     google.Endpoint,
	}
}

func (s *LoginGoogleClient) OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	s.initAuth()
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (s *LoginGoogleClient) OauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := s.getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (s *LoginGoogleClient) getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(s.config.GoogleOauthApiUrl + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
