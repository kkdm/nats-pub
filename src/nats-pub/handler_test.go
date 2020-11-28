package main

import (
    "testing"
    "net/http"
    "net/url"
    "strings"
    "net/http/httptest"
    "io/ioutil"

    "github.com/stretchr/testify/assert"
)

func TestPublishHandlerFails(t *testing.T) {
    val := url.Values{}
    val.Set("subject", "somesubject")
    val.Set("message", "some-test")

    req, _ := http.NewRequest("POST", "/publish", strings.NewReader(val.Encode())) 
    req.Header.Set("Content-Type", "application/json")

    res := httptest.NewRecorder()

    r := getRouter()
    r.ServeHTTP(res, req)

    body, _ := ioutil.ReadAll(res.Body)

    assert.Equal(t, res.Code, http.StatusInternalServerError)
    assert.Equal(
        t,
        string(body),
        "{\"message\":\"nats: no servers available for connection\",\"ok\":false}\n",
    )
}
