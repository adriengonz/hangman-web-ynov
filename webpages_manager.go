package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func startpage(w http.ResponseWriter) {
	var fileName = "./templates/default.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
		return
	}
	t.Execute(w, Currentdatagame)
}

func startmainpage(w http.ResponseWriter) {
	var fileName = "./templates/main_game.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
		return
	}
	t.Execute(w, Currentdatagame)
}

func startwinpage(w http.ResponseWriter) {
	var fileName = "./templates/win.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
		return
	}
	t.Execute(w, Currentdatagame)
}

func startlosepage(w http.ResponseWriter) {
	var fileName = "./templates/lose.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
		return
	}
	t.Execute(w, Currentdatagame)
}