package route

import (
	"database/sql"
	"net/http"

	"github.com/asaaaofficial/CRUD-Employee-GO/controller"
)

func MapRoute(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployee(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployee(db))
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/class", controller.NewClassEmployee(db))
	server.HandleFunc("/", controller.NewHelloWorld())
}