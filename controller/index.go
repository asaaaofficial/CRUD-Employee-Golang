package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id      string
	Name    string
	NPWP    string
	Address string
}

func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")
		if err != nil {
			log.Println("Error executing query:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error retrieving employees"))
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var employee Employee

			err = rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)
			if err != nil {
				log.Println("Error scanning row:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error processing employee data"))
				return
			}

			employees = append(employees, employee)
		}

		log.Printf("%d data karyawan ditemukan.\n", len(employees))

		// Pastikan file template ada
		fp := filepath.Join("views", "main.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			log.Println("Error parsing template:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error loading template"))
			return
		}

		// Kirim data ke template
		data := map[string]any{
			"employees": employees,
		}

		log.Printf("Mengirim %d data ke template", len(employees))

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error executing template:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error rendering page"))
			return
		}

		log.Println("Template berhasil dirender!")
	}
}
