package conexion

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "log"
	// User Go pq driver
	_ "github.com/lib/pq"
)

// DBConnect create a connection db
func DBConnect() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://usersnake01@localhost:26257/truora?sslmode=disable")
	if err != nil {
		// log.Fatal(err)
		fmt.Println("ERROR connecting to the database: ", err)
	}
	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS infoUsuario (id INT PRIMARY KEY DEFAULT , usuario STRING, puntaje INT)"); err != nil {
		// log.Fatal(err)
		fmt.Println("ERROR create table info_usuario: ", err)
	}
	return db
}

func GetUsuario(domain string, db *sql.DB) infoUsuario.ListaUsuario {

	rows, err := db.Query("SELECT * FROM info WHERE nombre = '" + name + "'")
	if err != nil {
		// log.Fatal(err)
		fmt.Println("ERROR select info_usuario: ", err)
	}
	defer rows.Close()

	var infoUsuario infoUsuario.ListaUsuario
	var id string
	var nombre, puntaje string
	for rows.Next() {
		if err := rows.Scan(&id, &nombre); err != nil {
			// log.Fatal(err)
			fmt.Println("ERROR get data for result set: ", err)
		}
		err := json.Unmarshal(&ListaUsuario)
		if err != nil {
			// panic(err)
			fmt.Println("ERROR create data infoUsuario: ", err)
		}
		ListaUsuario.ID = id
	}

	return puntaje
}

func CreateUsuario(nombre string, infoUsuario infoUsuario.ListaUsuario, db *sql.DB) bool {
	infoServerStr, err := json.Marshal(infoUsuario)
	if err != nil {
		// panic(err)
		fmt.Println("ERROR parse data infoServer to string: ", err)
		return false
	}

	if _, err := db.Exec(
		"INSERT INTO infoservers (domain, data_infoserver, last_updated) VALUES ('" + nombre + "', '" + string(infoServerStr) + "')"); err != nil {
		// log.Fatal(err)
		fmt.Println("ERROR Insert one row infousuarios table: ", err)
		return false
	}
	return true
}
