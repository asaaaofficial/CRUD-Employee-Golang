package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// NewDeleteEmployee menangani penghapusan data karyawan berdasarkan ID
func NewDeleteEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("==> DELETE handler called")
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			log.Println("==> ERROR: ID kosong")
			return
		}

		log.Println("==> Mencoba menghapus employee dengan ID:", id)

		// Cek apakah ID ada di database sebelum delete
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM employee WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			log.Println("==> ERROR: Gagal cek ID di database:", err)
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if !exists {
			log.Println("==> ERROR: Employee dengan ID", id, "tidak ditemukan")
			http.Error(w, fmt.Sprintf("Employee dengan ID %s tidak ditemukan", id), http.StatusNotFound)
			return
		}

		// Jalankan query DELETE
		res, err := db.Exec("DELETE FROM employee WHERE id = ?", id)
		if err != nil {
			log.Println("==> ERROR: Masalah database saat delete:", err)
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			log.Println("==> ERROR: Tidak ada baris yang dihapus")
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		log.Println("==> SUCCESS: Employee dengan ID", id, "berhasil dihapus")
		http.Redirect(w, r, "/employee/list", http.StatusSeeOther)
	}
}
