package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
    adminPasswd string
    port int = 7900
    failAttempt int = 0
)

func adminPassword() {
    adminPasswd = os.Getenv("RYN_ADMIN_PASSWORD")
    if len(adminPasswd) < 16 {
        adminPasswd = RandomString(16)
    }
    log.Println("Admin Password: ", adminPasswd)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse the form data from the request body
		if err := r.ParseForm(); err != nil {
			log.Print("Error parsing form data: ", err)
			return
		}

		// Retrieve the form data from the parsed request body
		action := r.FormValue("action")
		username := r.FormValue("username")
		password := r.FormValue("password")
		expDay := r.FormValue("expday")

		if err := actionUser(action, username, password, expDay); err != nil {
			fmt.Fprint(w, err)
			return
		}

		// Print the form data to the console
		log.Printf("Received a POST request. Action: %s Username: %s, Password: %s, expDay: %s\n", action, username, password, expDay)
		fmt.Fprint(w, "Hey, there\n")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Retrieve query parameters from the URL
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		// Print the query parameters to the console
		log.Printf("Received a GET request. Username: %s, Password: %s\n", username, password)
	}
    fmt.Fprint(w, "RYN Management Service\n")
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
        w.Header().Set("RYN-FailAttempt", fmt.Sprintf("%d", failAttempt))
        if failAttempt > 3 {
            http.Error(w, "Service Locked You Have Reset it Manuelly!", http.StatusUnauthorized)
            failAttempt += 1
			return
        }
		if !ok || username != "rynadmin" || password != adminPasswd {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
            failAttempt += 1
			return
		}
        failAttempt = 0
		next.ServeHTTP(w, r)
	}
}

func main() {

    adminPassword()

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ryn/api/usr", authenticate(userHandler))
    log.Println("Service Running On Port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
