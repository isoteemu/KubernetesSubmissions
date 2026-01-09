package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const indexHTML = `<!DOCTYPE html>
<html>
<head>
	<title>Todo App</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			max-width: 800px;
			margin: 50px auto;
			padding: 20px;
			background-color: #f5f5f5;
		}
		.container {
			background-color: white;
			padding: 30px;
			border-radius: 8px;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}
		h1 {
			color: #333;
		}
		.status {
			color: #666;
			font-size: 14px;
		}
	</style>
</head>
<body>
	<div class="container">
		<h1>Welcome to Todo App</h1>
		<p class="status">Server is running and ready to handle requests.</p>
	</div>
</body>
</html>`

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(indexHTML))
	})

	fmt.Printf("Server started in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
