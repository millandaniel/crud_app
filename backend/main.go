package main

import (
	"database/sql"
	"encoding/xml"

	//"encoding/json"
	//"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
)

// Estructura de Usuario
type Usuario struct {
	ID     int    `xml:"id"`
	Nombre string `xml:"nombre"`
}

var db *sql.DB

func main() {
	var err error

	// Abrir la conexión a la base de datos
	db, err = sql.Open("mysql", "root:Mariadb0221*st21I@tcp(localhost:3306)/crud_db")
	if err != nil {
		log.Fatal("Error al abrir la conexión a la base de datos:", err)
	}
	defer db.Close()

	// Definir manejadores para las rutas
	http.HandleFunc("/items", obteneritems)

	// Iniciar el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Manejador para la ruta "/items"
func obteneritems(w http.ResponseWriter, r *http.Request) {
	// Consultar la base de datos para obtener la lista de items
	filas, err := db.Query("SELECT id, nombre FROM items")
	if err != nil {
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer filas.Close()

	// Iterar sobre las filas y construir una lista de items
	var items []Usuario
	for filas.Next() {
		var usuario Usuario
		if err := filas.Scan(&usuario.ID, &usuario.Nombre); err != nil {
			http.Error(w, "Error al leer los datos de la base de datos", http.StatusInternalServerError)
			return
		}
		items = append(items, usuario)
	}

	// Configurar la cabecera de la respuesta con el tipo de contenido XML
	w.Header().Set("Content-Type", "application/xml")

	// Serializar la lista de items a XML
	xmlitems, err := xml.Marshal(items)
	if err != nil {
		http.Error(w, "Error al serializar items a XML", http.StatusInternalServerError)
		return
	}

	// Escribir los datos XML en el cuerpo de la respuesta
	_, err = w.Write(xmlitems)
	if err != nil {
		http.Error(w, "Error al escribir datos XML en la respuesta", http.StatusInternalServerError)
		return
	}
}
