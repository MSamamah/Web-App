package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func containersPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/containers.html")
}

func cicdPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/cicd.html")
}

func kubernetesPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/kubernetes.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/containers", containersPage)
	http.HandleFunc("/cicd", cicdPage)
	http.HandleFunc("/kubernetes", kubernetesPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
