package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

	// "WebApp/toto"
)
// pour importer "APP/NOMPACKAGE"
// puis NOMPACKAGE.nomMethode  qui est dans un fichier...Le nom du fichier dans
// le package on s'en fou

// liste des build possibles => go tool dist list
// pour le build => go build .  => pour l'os ou on build
// pour le build => GOOS=windows GOARCH=386 go build .  => pour Windows

func main(){
 
	db, err := sql.Open("mysql","root:cerise@tcp(localhost:3306)/toto")
	defer db.Close()
	if err != nil{
		log.Fatal((err))
	}
	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan((&version))

	if err2 != nil{
		log.Fatal(err2)
	}

	fmt.Println(version)

	 
}

 