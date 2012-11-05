package ntweet

import (
  "github.com/garyburd/go-oauth/oauth"
  "fmt"
  "encoding/json"
  "net/url"
  "net/http"
  "io/ioutil"
)

const (
  RequestUri = "https://api.twitter.com/oauth/request_token"
  AuthorizationUri = "https://api.twitter.com/oauth/authorize"
  TokenUri = "https://api.twitter.com/oauth/access_token"
  ApiUri = "https://api.twitter.com/1.1/"
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

func (c *TwitterClient) apiGetCall(resource string, params url.Values, data interface{}) error{
  if params == nil {
    params = make(url.Values)
  }

  urlStr := ApiUri + resource

  c.OauthClient.SignParam(&c.TokenCredentials, "GET", urlStr, params)
  resp, err := http.Get( urlStr + "?" + params.Encode())
  if err != nil {
    return err
  }

  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    p, _ := ioutil.ReadAll(resp.Body)
    return fmt.Errorf("Get %s returned status %d, %s", urlStr, resp.StatusCode, p)
  }

  return json.NewDecoder(resp.Body).Decode(data)
}
