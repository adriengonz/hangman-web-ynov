package main

import (
	"fmt"
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
			if Currentdatagame.Pseudo == "" { // Si il n'y a pas de pseudo, on l'initialise
				Currentdatagame.Pseudo = r.FormValue("name")
			}
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
	case "/tuto":
		starttutopage(w)
	case "/init":
		startinitpage(w)
	case "/win":
		startwinpage(w)
	case "/lose":
		startlosepage(w)
	default:
		Reset() // Reinitialise les valeurs
		fmt.Println("[DEBUG START] Valeurs réinitialisées")
		Currentdatagame.OriginalWord = WordPicker(RandomNumber()) // Initalisation du mot aléatoire a faire deviner
		Hidden() // Modificiton du mot généré en underscore
		fmt.Println("[DEBUG START] Mot :", Currentdatagame.OriginalWord)
		startpage(w)
	}
}

func main() {
	running := true
	Currentdatagame = &GameData{} // Initialisation d'une nouvelle instance de GameData
	Currentdatagame.Running = true
	fmt.Println("Server running on port 8080")
	fmt.Println("Access: http://localhost:8080")
	if running {
		fs := http.FileServer(http.Dir("templates"))
		http.Handle("/templates/", http.StripPrefix("/templates/", fs))
		http.HandleFunc("/", handler)     // Initilaise la page par défaut
		http.ListenAndServe(":8080", nil) // Lance le serveur sur le port 8080
	}
}
