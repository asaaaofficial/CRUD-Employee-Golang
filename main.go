package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asaaaofficial/CRUD-Employee-GO/database"
	"github.com/asaaaofficial/CRUD-Employee-GO/route"
)

func main() {
	fmt.Println("Server starting on port 8080...") 
	
	db, err := database.InitDatabase() // Tangani error
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close() // Pastikan database ditutup saat aplikasi berhenti

	server := http.NewServeMux()
	route.MapRoute(server, db)

	err = http.ListenAndServe(":8080", server)
	if err != nil {  
		log.Fatal("Server error:", err)  
	}
}

