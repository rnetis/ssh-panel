package main

import (
    "fmt"
    "log"
    "net/http"
)

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

        // Print the form data to the console
        log.Printf("Received a POST request. Action: %s Username: %s, Password: %s, expDay: %s\n", action, name, age, expDay)
		fmt.Fprint(w, "Hey, there\n")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
        // Retrieve query parameters from the URL
        name := r.URL.Query().Get("name")
        age := r.URL.Query().Get("age")

        // Print the query parameters to the console
        log.Printf("Received a GET request. Name: %s, Age: %s\n", name, age)
		fmt.Fprint(w, "Hey, there\n")
    }else if r.Method == "POST" {
        // Parse the form data from the request body
        if err := r.ParseForm(); err != nil {
            log.Print("Error parsing form data: ", err)
            return
        }

        // Retrieve the form data from the parsed request body
        name := r.FormValue("name")
        age := r.FormValue("age")

        // Print the form data to the console
        log.Printf("Received a POST request. Name: %s, Age: %s\n", name, age)
		fmt.Fprint(w, "Hey, there\n")
    }else {
		fmt.Fprint(w, "Hey, there\n")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact us at contact@example.com")
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "rynadmin" || password != "r7a1y9a3n@8n2a4y6a5r/" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", authenticate(aboutHandler))
	mux.HandleFunc("/contact", authenticate(contactHandler))
	mux.HandleFunc("/ryn/api/usr", authenticate(userHandler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}