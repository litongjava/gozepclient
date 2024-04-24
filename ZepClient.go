package gozepclient

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type ZepClient struct {
  BaseURL    string
  Token      string
  HTTPClient *http.Client
}

func NewZepClient(baseURL string, token string) *ZepClient {
  return &ZepClient{
    BaseURL:    baseURL,
    Token:      token,
    HTTPClient: &http.Client{},
  }
}

func (c *ZepClient) SearchSession(sessionID string, requestMap map[string]interface{}) (string, error) {
  url := fmt.Sprintf("%s/api/v1/sessions/%s/search", c.BaseURL, sessionID)
  requestBody, err := json.Marshal(requestMap)
  if err != nil {
    return "", err
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
  if err != nil {
    return "", err
  }
  req.Header.Set("Content-Type", "application/json")
  if c.Token != "" {
    req.Header.Set("Authorization", "Bearer "+c.Token)
  }

  resp, err := c.HTTPClient.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}

func (c *ZepClient) PostSessionMemory(sessionID string, postSesionReqeustBody map[string]interface{}) (string, error) {
  url := fmt.Sprintf("%s/api/v1/sessions/%s/memory", c.BaseURL, sessionID)
  requestBody, err := json.Marshal(postSesionReqeustBody)
  if err != nil {
    return "", err
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
  if err != nil {
    return "", err
  }
  if c.Token != "" {
    req.Header.Set("Authorization", "Bearer "+c.Token)
  }
  req.Header.Set("Content-Type", "application/json")

  resp, err := c.HTTPClient.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}
