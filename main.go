package main

import (
	"fmt"
	// "WebApp/toto"
)
// pour importer "APP/NOMPACKAGE"
// puis NOMPACKAGE.nomMethode ...Le nom du fichier dans
// le package on s'en fou

// liste des build possibles => go tool dist list
// pour le build => go build .  => pour l'os ou on build
// pour le build => GOOS=windows GOARCH=386 go build .  => pour Windows

func main(){
	// toto.Imprime()
	// toto.Reimprime()
	fmt.Println("eeeee")
}

 