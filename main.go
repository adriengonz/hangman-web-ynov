package main

import (
	"fmt"
	//"time"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path { // Execute un code block différent selon l'URL demandée
	case "/home":
		fmt.Fprintf(w, "tkt")
	case "/test":
		fmt.Printf("test page")
	case "/main":
		Currentdatagame.Pseudo = r.FormValue("name")
		Currentdatagame.CurrentLetter = r.FormValue("letter")
		startmainpage(w)
	default:
		fmt.Println(Currentdatagame.Pseudo)
		startpage(w)
	}
}

/*
func renderTemplate (w http.ResponseWriter, htmlfile string) { // Permet d'afficher une page web lors de l'appel de la fonction
	t, err := template.ParseFiles("./templates/" + htmlfile + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retourne une erreur dans les logs si une erreur est produite
	}
	t.Execute(w, nil)
}
*/

func main() {
	running := true
	Currentdatagame = &GameData{} // Initialisation d'une nouvelle instance de GameData
	Currentdatagame.OriginalWord = "ouifi"
	Currentdatagame.Hidden_word = []string{"_", "_", "_", "_", "_"}
	//Currentdatagame.Pseudo = "test"
	fmt.Println("Server running on port 8080")
	fmt.Println("Access: http://localhost:8080")
	if running {
		fs := http.FileServer(http.Dir("templates"))
		http.Handle("/templates/", http.StripPrefix("/templates/", fs))
		http.HandleFunc("/", handler)     // Initilaise la page par défaut
		http.ListenAndServe(":8080", nil) // Lance le serveur sur le port 8080
	}

	/*
		PrintHangmanAscii() // Appel de la fonction pour afficher "Hangman" en ascii
		fmt.Println("Bienvenue sur Hangman !")
		fmt.Println("La partie va bientôt commencer..")
		time.Sleep(4 * time.Second)
		Clear()
		currentdatagame.originalWord = WordPicker(RandomNumber()) // Initalisation du mot aléatoire a faire deviner
		Hidden() // Modificiton du mot généré en underscore
		RunHangman() // Lancement du jeu
	*/
}
