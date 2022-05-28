package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Article struct {
	Title		string	`json:"title"`
	Desc		string	`json:"description"`
	Content	string	`json:"content"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Articles := []Article {
		{ Title: "Hello", Desc: "Article Description", Content: "Article Content" },
		{ Title: "Hello 2", Desc: "Article Description", Content: "Article Content" },
	}

	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Add("X-Frame-Options", "DENY")

	fmt.Fprintf(w, "Category: %v | %v \n", vars["key"], r.URL.Query().Get("key"))

	// json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Custom404Handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("lalal")

	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]bool { "ok": false })
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet, http.MethodOptions)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)

	router.HandleFunc("/products/{key}", ProductHandler)// .Queries("key", "value")
	// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	router.Use(mux.CORSMethodMiddleware(router))

	router.NotFoundHandler = http.HandlerFunc(Custom404Handler)

	// http.Handle("/", router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:5005",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	log.Fatal(srv.ListenAndServe())

	/*
	go func() {
		if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
		}
	}()
	*/
}
