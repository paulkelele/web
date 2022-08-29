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

//  site pour mysql documentation => https://zetcode.com/golang/mysql/

// les commandes: 
//				db.QueryRow().Scan() => query une seule ligne => renvoit *Row
//				db.Query => query sur toutes les lignes => renvoit *Row, err
// 				db.Exec(sql) => INSERT || DELETE => renvoit *Row, err
// prepared statement : db.Exec("...?", varOf?)



func main(){
 
	db, err := sql.Open("mysql","root:cerise@tcp(localhost:3306)/toto")
	defer db.Close()
	if err != nil{
		log.Fatal((err))
	}

	sql :="insert into user(nom,prenom,age) values(?,?,?)"
	 
res, err2 :=db.Exec(sql,"MARTIN","PAUL",44)
if err2 != nil{
	log.Fatal(err2)
}
 LastId, err3 := res.LastInsertId()

if err3 != nil{
	log.Fatal((err3.Error()))
}
 fmt.Println("Les tid inseert: ",LastId)
}

 