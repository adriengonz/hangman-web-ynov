package main

import (
	"fmt"
	"time"
)

func RunHangman() { // Fonction majeure du jeu
	presentInWord := false // Valeur qui vérifie si la lettre est présente dans le mot caché
	wordNotRevealed := false // Valeur qui vérifie si le mot entier a été deviné
	draw_hangman(currentdatagame.try) // Appel de la fonction qui affiche le pendu
	letter := ""
	fmt.Println(letter)

	fmt.Println("Déjà utilisé(s) : ", currentdatagame.used_letters)
	fmt.Println("\nEntrez votre choix :")
	fmt.Scanln(&letter) // Récupération de l'entrée utilisateur (lettre)
	
	for _, char := range currentdatagame.used_letters {
		if letter == string(char) {
			fmt.Println("Essayez autre chose, vous avez déjà essayé ceci !")
			time.Sleep(1 * time.Second)
			Clear()
			RunHangman()
		}
	}
	currentdatagame.used_letters = append(currentdatagame.used_letters, letter)
	
	if len(letter) < 1 || len(letter) > 1 { // Condition qui vérifie s'il n'y a qu'une lettre entrée
	fmt.Println("Impossible, entrez une seule lettre.")
	time.Sleep(1 * time.Second)
	RunHangman()
	} else {
		presentInWord = false // Reinitialisation de la variable
		for i := 0; i < len(currentdatagame.originalWord); i++ { // Boucle qui vérifie si la lettre entrée est présente dans le mot
			if letter == string(currentdatagame.originalWord[i]) {
			presentInWord = true // Si la lettre est présente, la variable passe vraie
		}
	}
		if presentInWord { // Si la lettre est présente dans le mot
			fmt.Println("La lettre", letter, "est dans le mot")
			for index, char := range currentdatagame.originalWord {
				if letter == string(char) {
					currentdatagame.hidden_word[index] = letter // Remplacement de _ avec la lettre devinée précedemment
				}
			}
			time.Sleep(1 * time.Second)
		} else { // Si la lettre n'est pas dans le mot
			fmt.Println("La lettre", letter, "n'est pas dans le mot")
			time.Sleep(1 * time.Second)
			if currentdatagame.try < 9 { // S'il reste encore des essais possibles
				currentdatagame.try++ // Incrémentation du compteur d'essais
				fmt.Println("Il ne vous reste plus que", 10-currentdatagame.try, "essais")
			} else { // Si tous les essais ont été utilisées
				fmt.Println("Perdu ! Vous n'avez plus d'essais disponibles")
				fmt.Println("Le mot était", currentdatagame.originalWord)
				time.Sleep(3 * time.Second)
				Exit() // Sortie du jeu
			}
		}
		Clear()
		fmt.Println(currentdatagame.hidden_word) 
	}
	for index, _ := range currentdatagame.originalWord { // Boucle qui vérifie s'il y a encore des underscore dans le mot caché
		if currentdatagame.hidden_word[index] == "_" {
			wordNotRevealed = true
			RunHangman() // Reprise du jeu
			break
		}
	}
	if wordNotRevealed == false { // Si le mot est deviné entièrement
		fmt.Println("Felicitations ! Vous avez déviné !")
		time.Sleep(3 * time.Second)
		Exit() // Sortie du jeu
	}
}