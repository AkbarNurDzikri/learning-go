package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic: ", error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// fmt.Fprint(w, "Hello world")
		panic("uppss")
	})

	request := httptest.NewRequest("GET", "http://localhost:4000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Panic: uppss", string(body))
}