package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
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

type Cercle struct {
	radius float64
}

type Carre struct {
	cote int32
}

// on definiiune interface
type ShapeInterface interface {
	Shape() float64
}

func (c Cercle) Shape() float64 {
	i := rand.Intn(100)
	fmt.Println(i)
	return math.Pi * c.radius * c.radius
}

func getArea(s ShapeInterface) {
	fmt.Println(s.Shape())
}

type User struct{
	id int
	nom string
	prenom string
	age int
}


func main() {
	// var bo bool
	// compteur := 0
	// fmt.Printf("BOOOl %v\n",bo)
	// for bo == false{
	// 	fmt.Print("GGG compteur: ")
	// 	fmt.Println( compteur)
	// 	compteur +=1
	// 	if compteur > 10{
	// 		bo = true
	// 	}
	// }
  	// fmt.Printf("BOOOlBOOOlBOOOlBOOOl%v\n",bo)



	c := Cercle{radius: 5.2}

	getArea(c)

	db, err := sql.Open("mysql", "root:cerise@tcp(localhost:3306)/toto")
	defer db.Close()
	if err != nil {
		log.Fatal((err))
	}

	// INSERTION _________________--------------------___________
	// sql := "insert into user(nom,prenom,age) values(?,?,?)"
	// res, err2 := db.Exec(sql, "MARTIN", "PAULéàî", 44)

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// LastId, err3 := res.LastInsertId()

	// if err3 != nil {
	// 	log.Fatal((err3.Error()))
	// }
	// fmt.Println("Les tid inseert: ", LastId)


	// SELECT ___________--------------______________


	sliceOfUsers := make([]User,0)

	sql := "select * from user"

	res, err2 := db.Query(sql)
	
	defer res.Close()

	if err2 != nil{
		log.Fatal(err2.Error())
	}

	for res.Next(){
		var u User
		err := res.Scan( &u.id, &u.age, &u.nom, &u.prenom)
		if err != nil{
			log.Fatal(err)
		}
		 
		sliceOfUsers = append(sliceOfUsers, u) 
	}
	cpt := 0
	for _,user := range sliceOfUsers{
		cpt += 1
		fmt.Printf("each user: %v nbr d'utilisateurs: %d\n",user,cpt)
	}
	 
}
