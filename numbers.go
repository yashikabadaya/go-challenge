package main

import (
	"net/http"
        "log"
        "fmt"
        "encoding/json"
        "strings"
)

func main() {
     http.HandleFunc("/numbers", NumberHandler)
     log.Fatal(http.ListenAndServe(":8080", nil))
}

// returns all the urls on which http GET is supposed to be done
func parseURL(path string) []string{
     urls := strings.Split(path, "u=")[1:]
     for i := 0; i < len(urls); i++ {
         if (urls[i][len(urls[i])-1] == '&') {
             urls[i] = urls[i][:len(urls[i])-1]
         }
     }
     return urls
}

func merge(left, right []int) []int {
     var result []int
     var remainingarr []int
     l := 0
     r := 0
     for l < len(left) && r < len(right) {
         if (left[l] < right[r]) {
             result = append(result, left[l])
             l++
         } else if (left[l] == right[r]) {
             result = append(result, right[r])
             r++
             l++
         } else {
             result = append(result, right[r])
             r++
         }
     }
     if (l != len(left)) {
        remainingarr = left[l:]
        result = append(result, remainingarr...)
     }
     if (r != len(right)) {
        remainingarr = right[r:]
        result = append(result, remainingarr...)
     }
     return result
}

func mergesort(arr []int) []int {
     if (len(arr) <= 1) {
         return arr
     }
     mid := len(arr) / 2
     left := arr[:mid]
     right := arr[mid:]

     left = mergesort(left)
     right = mergesort(right)
     return merge(left, right)
}

func NumberHandler(w http.ResponseWriter, r *http.Request) {
     // http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo
     // above req would enter NumberHandler
     // here we need to parse through each "u" of the req and then call blackbox server on each
     subreqs := parseURL(r.URL.RawQuery)
     var result []int
     for i := 0; i < len(subreqs); i++ {
             var m map[string][]int 
	     resp, err := http.Get(subreqs[i])
	     if err != nil {
		 fmt.Fprint(w, nil)
		 log.Fatal(err)
	     }
             if(resp.StatusCode == http.StatusOK) {
		     err = json.NewDecoder(resp.Body).Decode(&m)
		     if err != nil {
			log.Fatal(err)
		     }
		     result = append(result, m["numbers"]...)
             } 
     }
     
     fmt.Fprint(w, mergesort(result))
}



