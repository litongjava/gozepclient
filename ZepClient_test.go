package gozepclient

import (
  "fmt"
  "testing"
)

func TestZepClient_PostSessionMemory(t *testing.T) {

  sessionID := "434ada9362dc4404b895db3e81d2487c"

  baseUrl := "https://zep-api.fly.dev/"
  var token = ""
  client := NewZepClient(baseUrl, token)
  // Example of searching in a session
  searchRequestMap := map[string]interface{}{
    "mmr_lambda":   0,
    "search_scope": "messages",
    //"search_type":  "similarity",
    "text": "What am I",
  }

  result, err := client.SearchSession(sessionID, searchRequestMap)
  if err != nil {
    fmt.Println("Search error:", err)
  } else {
    fmt.Println("Search result:", result)
  }

  // Example of posting to a session's memory
  messages := []map[string]interface{}{
    {
      "content": "thi is the test user input",
      "role":    "user",
    },
  }
  summary := map[string]interface{}{
    "content": "summary content",
  }
  postSesionReqeustBody := map[string]interface{}{
    "messages": messages,
    "summary":  summary,
  }
  result, err = client.PostSessionMemory(sessionID, postSesionReqeustBody)
  if err != nil {
    fmt.Println("Post error:", err)
  } else {
    fmt.Println("Post result:", result)
  }

}
