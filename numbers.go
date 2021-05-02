package main

import (
	"net/http"
        "log"
        "fmt"
        "io/ioutil"
        "strings"
)

func main() {
     http.HandleFunc("/numbers", NumberHandler)
     log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseURL(path string) []string{
     urls := strings.Split(path, "u=")[1:]
     for i := 0; i < len(urls); i++ {
         if (urls[i][len(urls[i])-1] == '&') {
             urls[i] = urls[i][:len(urls[i])-1]
         }
     }
     return urls
}

func NumberHandler(w http.ResponseWriter, r *http.Request) {
     // http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo
     // above req would enter NumberHandler
     // here we need to parse through each "u" of the req and then call blackbox server on each
     subreqs := parseURL(r.URL.RawQuery)
     fmt.Println(subreqs)
     for i := 0; i < len(subreqs); i++ {
	     resp, err := http.Get(subreqs[i])
	     if err != nil {
		 fmt.Fprint(w, nil)
		 log.Fatal(err)
	     }     
	     strresp, _ := ioutil.ReadAll(resp.Body)
	     fmt.Fprint(w, string(strresp))
     }
     
}



