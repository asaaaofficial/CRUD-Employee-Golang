package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/asaaaofficial/CRUD-Employee-GO/database"
)

func NewClassEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			npwp := r.FormValue("npwp")
			address := r.FormValue("address")

			fmt.Println("Received Data:", name, npwp, address)

    		if name == "" || npwp == "" || address == "" {
        	http.Error(w, "All fields are required", http.StatusBadRequest)
        	return
    	}

			db, err := database.InitDatabase()
			if err != nil {
				log.Println("Error initializing database:", err)
				http.Error(w, "Database initialization error: "+err.Error(), http.StatusInternalServerError)
				return
			}
    		defer db.Close()

			result, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES (?,?,?)", name, npwp, address)
			if err != nil {
				log.Println("Error saat insert ke database:", err)
				http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				http.Error(w, "No rows inserted", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data inserted successfully"))
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "class.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if err = tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
