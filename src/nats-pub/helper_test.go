package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "io/ioutil"

    "github.com/stretchr/testify/assert"
)

func TestResponseJsonNotFound(t *testing.T) {
    w := httptest.NewRecorder()
    responseJson(w,
        http.StatusOK,
        Result{
            Message: "some error",
            Ok: false,
        })
    res := w.Result()
    body, _ := ioutil.ReadAll(res.Body)

    assert.Equal(t, res.StatusCode, http.StatusOK)
    assert.Equal(t, res.Header.Get("Content-Type"), "application/json")
    assert.Equal(t, string(body), "{\"message\":\"some error\",\"ok\":false}\n")
}
