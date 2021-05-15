package main

import (
	"testing"
        "io/ioutil"
        "fmt"
        "net/http"
        "net/http/httptest"
)

func TestNumberHandler(t *testing.T) {
     t.Run("test to call NumberHandler function and stdout the response", func(t *testing.T) {
          //http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo
          request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/numbers?u=http://localhost:8090/primes&u=http://localhost:8090/fibo", nil)
          response := httptest.NewRecorder()
          NumberHandler(response, request)
          r, _ := ioutil.ReadAll(response.Body)
          fmt.Println(string(r))
     })
}

