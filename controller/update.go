package controller

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Remove duplicate Employee struct and use the one from index.go or a shared package.

func NewUpdateEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("==> Handler NewUpdateEmployee DIPANGGIL")
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
				return
			}

			id := r.FormValue("id")
			name := r.FormValue("name")
			npwp := r.FormValue("npwp")
			address := r.FormValue("address")

			if id == "" || name == "" || npwp == "" || address == "" {
				http.Error(w, "All fields are required", http.StatusBadRequest)
				return
			}

			// Update database
			result, err := db.Exec("UPDATE employee SET name=?, npwp=?, address=? WHERE id=?", name, npwp, address, id)
			if err != nil {
				log.Println("Error saat update database:", err)
				http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				http.Error(w, "No rows updated", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee/list", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")
			if id == "" {
				http.Error(w, "ID is required", http.StatusBadRequest)
				return
			}

			var employee Employee
			err := db.QueryRow("SELECT id, name, npwp, address FROM employee WHERE id = ?", id).
				Scan(&employee.Id, &employee.Name, &employee.NPWP, &employee.Address)

			if err != nil {
				if err == sql.ErrNoRows {
					log.Println("==> ERROR: Employee tidak ditemukan")
					http.Error(w, "Employee not found", http.StatusNotFound)
				} else {
					log.Println("==> ERROR: Masalah database:", err)
					http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
				}
				return
			}

			log.Println("==> Data Karyawan berhasil diambil:", employee)

		// **DEBUG**: Sementara hentikan eksekusi sebelum template dijalankan
		// return 

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				log.Println("==> ERROR: Template error:", err)
				http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			log.Println("==> Rendering Template update.html")

			if err = tmpl.Execute(w, map[string]any{"employee": employee}); err != nil {
				log.Println("==> ERROR: Template execution error:", err)
				http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			}
		}
	}
}