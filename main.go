package main

import (
	"fmt"
	//"time"
	"net/http"
	"html/template"
)

type GameData struct {
	OriginalWord     string
	hidden_word      []string
	Attempts         int
	used_letters     []string
	Error            string
	Pseudo           string
}

var currentdatagame *GameData

func handler(w http. ResponseWriter, r *http. Request) {
	switch r.URL.Path { // Execute un code block différent selon l'URL demandée
	case "/home":
		fmt.Fprintf(w, "tkt")
	case "/test":
		fmt.Printf("test page")
	default:
		renderTemplate(w, "default")
		currentdatagame.Pseudo = r.FormValue("name")
		fmt.Println(currentdatagame.Pseudo)
	}
}

func renderTemplate (w http.ResponseWriter, htmlfile string) { // Permet d'afficher une page web lors de l'appel de la fonction
	t, err := template.ParseFiles("./templates/" + htmlfile + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
	}
	t.Execute(w, nil)
}

func main() {
	currentdatagame = &GameData{} // Initialisation d'une nouvelle instance de GameData
	fmt.Println("Server running on port 8080")
	fmt.Println("Access: http://localhost:8080")
	http.HandleFunc("/", handler) // Initilaise la page par défaut
	http.HandleFunc("/home", handler) // Initialise la page /home
	http.Handle("./templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("static"))))
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