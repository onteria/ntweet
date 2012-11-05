package ntweet

import (
  "github.com/garyburd/go-oauth/oauth"
)

const (
  RequestUri = "https://api.twitter.com/oauth/request_token"
  AuthorizationUri = "https://api.twitter.com/oauth/authorize"
  TokenUri = "https://api.twitter.com/oauth/access_token"
)

type TwitterClient struct {
  OauthClient oauth.Client
  TokenCredentials oauth.Credentials
}

func NewClient(ckey string, csecret string, token string, secret string) *TwitterClient {
  oauthClient := oauth.Client{
    TemporaryCredentialRequestURI: RequestUri,
    ResourceOwnerAuthorizationURI: AuthorizationUri,
    TokenRequestURI: TokenUri,
    Credentials: oauth.Credentials {
      Token: ckey,
      Secret: csecret,
    },
  }

  tokenInfo := oauth.Credentials{
    Token: token,
    Secret: secret,
  }

  return &TwitterClient {
    OauthClient: oauthClient,
    TokenCredentials: tokenInfo,
  }
}
