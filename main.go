package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

var (
	// the redirect URL must be an exact match of a URL you've registered for your application
	// scopes determine which permissions the user is prompted to authorize
	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)

	ch = make(chan *spotify.Client)

	state = "abc123" // thefuq is this?? ToDo: Find out
)

// the user will eventually be redirected back to your redirect URL
// typically you'll have a handler set up like the following:
//func redirectHandler(w http.ResponseWriter, r *http.Request) {
//	// use the same state string here that you used to generate the URL
//	token, err := auth.Token(state, r)
//	if err != nil {
//		http.Error(w, "Couldn't get token", http.StatusNotFound)
//		return
//	}
//	// create a client using the specified token
//	// client := auth.NewClient(token)
//
//	// the client can now be used to make authenticated requests
//}

func main() {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll.FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		fmt.Println("Failed to open authentication URI:", err)
	}

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
