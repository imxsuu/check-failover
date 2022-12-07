package main

import (
    "fmt"
    // "io"
    "log"
    "net/http"
	"io/ioutil"
    "os"
    // "runtime"
    "time"
)

var redirect_url = "http://localhost:9010"

func main() {

    srv := startHttpServer()

    // Backup to gracefully shutdown the server
    time.Sleep(10 * time.Second)
    if err := srv.Shutdown(nil); err != nil {
        panic(err) // failure/timeout shutting down the server gracefully
    }
}

func startHttpServer() *http.Server {
    srv := &http.Server{Addr: ":9010"}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Method : ", r.Method)
    	fmt.Println("URL : ", r.URL)
   	 	fmt.Println("Header : ", r.Header)
 
    	b, _ := ioutil.ReadAll(r.Body)
		body := string(b)
    	defer r.Body.Close()
    	fmt.Println("Body : ", body)
 

		if body == "a" {
			print("Gedge-platform is running")
		} else if body == "b" {
			print("Gdege-platform is not working")
			os.Exit(3)
    	}
	})

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            // cannot panic, because this probably is an intentional close
            log.Printf("Httpserver: ListenAndServe() error: %s", err)
        }
    }()

    // returning reference so caller can call Shutdown()
    return srv
}
