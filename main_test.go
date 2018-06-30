package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	setup()
	router := ConfigureRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestSuccessHandshake(t *testing.T) {
	setup()
	router := ConfigureRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/handshake?hub.mode=subscribe&hub.challenge=x1&hub.verify_token=aaa", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "x1", w.Body.String())
}

func TestSuccessHandshakeFailByToken(t *testing.T) {
	setup()
	router := ConfigureRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/handshake?hub.mode=subscribe&hub.challenge=x1&hub.verify_token=aaba", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
	assert.Equal(t, "Unnautorized", w.Body.String())
}

func TestSuccessHandshakeFailByMode(t *testing.T) {
	setup()
	router := ConfigureRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/handshake?hub.mode=message&hub.challenge=x1&hub.verify_token=aaa", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
	assert.Equal(t, "Unnautorized", w.Body.String())
}

func TestSuccessProcessMessage(t *testing.T) {
	router := ConfigureRouter()

	w := httptest.NewRecorder()
	body := []byte(`
		{
		"object":"page",
		"entry":[
		  {
			"id":"111",
			"time":1458692752478,
			"messaging":[
			  {
				"sender":{
				  "id":"1111"
				},
				"recipient":{
				  "id":"2222"
				}
			  },
			  {
				"sender":{
				  "id":"3333"
				},
				"recipient":{
				  "id":"4444"
				}
			  }
			]
		  }
		]
	  }`)
	req, _ := http.NewRequest("POST", "/processMessage", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func setup() {
	os.Setenv("VERIFY_TOKEN", "aaa")
}
