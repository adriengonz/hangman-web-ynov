package main

import (
	"fmt"
	//"time"
	"net/http"
	"html/template"
)

type GameData struct {
	OriginalWord     string
	ToFind           string
	Attempts         int
	LettersSuggested []string
	Error            string
	HangmanPositions [10]string
}

func handler(w http. ResponseWriter, r *http. Request) {
	switch r.URL.Path { // Execute un code block différent selon l'URL demandée
	case "/home":
		fmt.Fprintf(w, "tkt")
	case "/test":
		renderTemplate(w, "default")
	default:
		fmt.Fprintf(w, "Default Page")
	}
}

func renderTemplate (w http.ResponseWriter, htmlfile string) {
	t, err := template.ParseFiles("./templates/" + htmlfile + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

var hidden_word []string // Variable sous forme de liste qui contient le mot caché
var used_letters []string // Variable sous forme de liste qui contient les lettres utilisées

func main() {
	fmt.Println("Server running on port 8080")
	fmt.Println("Access: http://localhost:8080")
	http.HandleFunc("/", handler) // Initilaise la page par défaut
	http.HandleFunc("/home", handler) // Initialise la page /home
	http.ListenAndServe(":8080", nil) // Lance le serveur sur le port 8080
	/*
	PrintHangmanAscii() // Appel de la fonction pour afficher "Hangman" en ascii
	fmt.Println("Bienvenue sur Hangman !")
	fmt.Println("La partie va bientôt commencer..")
	time.Sleep(4 * time.Second)
	Clear()
	tries := 0 // Initialisation de la variable des essais
	originalWord := WordPicker(RandomNumber()) // Initalisation du mot aléatoire a faire deviner
	Hidden(originalWord) // Modificiton du mot généré en underscore
	RunHangman(originalWord, tries) // Lancement du jeu
	*/
}