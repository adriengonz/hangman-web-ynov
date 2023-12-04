package main

import (
	"fmt"
	//"time"
	"net/http"
)

func handler(w http. ResponseWriter, r *http. Request) {
	startpage(w)
	/*switch r.URL.Path { // Execute un code block différent selon l'URL demandée
	case "/home":
		fmt.Fprintf(w, "tkt")
	case "/test":
		fmt.Printf("test page")
	default:
		startpage(w, r)
		if currentdatagame.pseudo != ""  {
		} else {
			currentdatagame.pseudo = r.FormValue("name")
		}
		fmt.Println(currentdatagame.pseudo)
	}*/
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
	Currentdatagame = &GameData{} // Initialisation d'une nouvelle instance de GameData
	Currentdatagame.Used_letters = append(Currentdatagame.Used_letters, "s")
	Currentdatagame.OriginalWord = "tkt"
	fmt.Println("Server running on port 8080")
	fmt.Println("Access: http://localhost:8080")
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/", handler) // Initilaise la page par défaut
	//http.Handle("./templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil) // Lance le serveur sur le port 8080

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