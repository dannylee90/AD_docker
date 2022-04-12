package main

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
)



func main() {

        timeSource := rand.NewSource(time.Now().UnixNano())
        random := rand.New(timeSource)
        no := random.Intn(100000)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, no)
	})


	http.ListenAndServe(":80", nil)
}
