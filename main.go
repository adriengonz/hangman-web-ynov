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
		if Currentdatagame.Running { // Si une partie est en cours
		Currentdatagame.Pseudo = r.FormValue("name")
		Currentdatagame.CurrentLetter = r.FormValue("letter")
		RunHangman()
		startmainpage(w)
		} else { // Si la partie est terminée
			if Currentdatagame.WordRevealed { // Si le joueur a gagné, on redirige vers la page de victoire
				http.Redirect(w,r, "/win", http.StatusSeeOther)
				fmt.Println("[VERBOSE WEB] Changement de la page en win page")
			} else { // Si le joueur a perdu, on redirige vers la page de défaite
				http.Redirect(w,r, "/lose", http.StatusSeeOther)
				fmt.Println("[VERBOSE WEB] Changement de la page en lose page")
			}
		}
		// Selon si perdu ou gagné, lancer une autre page
	case "/win":
		fmt.Fprintf(w, "t'as gagné chef")
	case "/lose":
		fmt.Fprintf(w, "t'as perdu mdr")
	default:
		fmt.Println(Currentdatagame.Pseudo)
		//Currentdatagame.OriginalWord = WordPicker(RandomNumber()) // Initalisation du mot aléatoire a faire deviner
		//Hidden() // Modificiton du mot généré en underscore
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
	Currentdatagame.Running = true
	Currentdatagame.OriginalWord = "OUIFI"
	Currentdatagame.Hidden_word = []string{"_", "_", "_", "_", "_"}
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
