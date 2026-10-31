package infomaniak

import (
 "context"
 "net/http"
 "net/http/httptest"
 "strings"
 "testing"
)

func TestProfile(t *testing.T) {
 server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/profile" { t.Errorf("path = %q", r.URL.Path) }
  if r.Header.Get("Authorization") != "Bearer test-token" { t.Error("missing bearer token") }
  w.Header().Set("Content-Type", "application/json")
  _, _ = w.Write([]byte(`{"ok":true}`))
 }))
 defer server.Close()
 c := NewClient("test-token")
 c.BaseURL = server.URL
 data, err := c.Profile(context.Background())
 if err != nil || !strings.Contains(string(data), `"ok":true`) { t.Fatalf("profile = %s, %v", data, err) }
}

func TestProfileMissingToken(t *testing.T) {
 _, err := NewClient("").Profile(context.Background())
 if err == nil { t.Fatal("expected missing token error") }
}
