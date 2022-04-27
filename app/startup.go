package app

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	html "pdf-test/html"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir(filepath.Join("./static")))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.HandleFunc("/close", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			html.HTML(w, 200, "close", nil)
		},
	)).Methods("GET")

	router.HandleFunc("/{Path}", PageHandler).Methods("GET")
	router.HandleFunc("/", PageHandler).Methods("GET")

	srv := &http.Server{
		Addr: "0.0.0.0:1234",

		// Slowloris is no friend of mine
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,

		Handler: router,
	}

	fmt.Println("Server started on 0.0.0.0:1234")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	filename := params["Path"]

	fmt.Println(filename)
	if len(filename) == 1 {
		filename = "index"
	}

	html.HTML(w, 200, strings.ToLower(filename), nil)
}
