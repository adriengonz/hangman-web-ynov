package main

import (
	"fmt"
)

// Si y a une lettre en double, on fait rien
// Le dessin Hangman se basera sur les lettres deja utilisées (essais)
// Stocker le resultat (gagné ou perdu) dans la variable gameengine et selon le resultat, 
// rediriger vers la page web perdu ou gagné

func RunHangman() { // Fonction majeure du jeu
	if Currentdatagame.CurrentLetter !="" {
		CheckUsedLetter()
		fmt.Println("[VERBOSE] Le check de la lettre utilisée a été effectué")
		LetterPresent()
		fmt.Println("[VERBOSE] Le check de la lettre dans le mot a été effectué")
		CheckWordRevealed()
		fmt.Println("[VERBOSE] Le check du mot révélé a été effectué")
	}
	Currentdatagame.CurrentLetter = ""
	fmt.Println("[VERBOSE] La lettre temporaire a été vidée")
/*
	presentInWord := false // Valeur qui vérifie si la lettre est présente dans le mot caché
	draw_hangman(Currentdatagame.Try) // Appel de la fonction qui affiche le pendu
	
	for _, char := range Currentdatagame.Used_letters {
		if Currentdatagame.CurrentLetter == string(char) {

		} else {
			Currentdatagame.Used_letters = append(Currentdatagame.Used_letters, Currentdatagame.CurrentLetter)
		}
	}
	
	presentInWord = false // Reinitialisation de la variable
	for i := 0; i < len(Currentdatagame.OriginalWord); i++ { // Boucle qui vérifie si la lettre entrée est présente dans le mot
		if Currentdatagame.CurrentLetter == string(Currentdatagame.OriginalWord[i]) {
		presentInWord = true // Si la lettre est présente, la variable passe vraie
		}
	}
	if presentInWord { // Si la lettre est présente dans le mot
		for index, char := range Currentdatagame.OriginalWord {
			if Currentdatagame.CurrentLetter == string(char) {
				Currentdatagame.Hidden_word[index] = Currentdatagame.CurrentLetter // Remplacement de _ avec la lettre devinée précedemment
			}
		}
		} else { // Si la lettre n'est pas dans le mot
			if Currentdatagame.Try < 9 { // S'il reste encore des essais possibles
				Currentdatagame.Try++ // Incrémentation du compteur d'essais
			} else { // Si tous les essais ont été utilisées
				Currentdatagame.Running = false
			}
		}
		fmt.Println(Currentdatagame.Hidden_word) 

	for index, _ := range Currentdatagame.OriginalWord { // Boucle qui vérifie s'il y a encore des underscore dans le mot caché
		if Currentdatagame.Hidden_word[index] == "_" {
			Currentdatagame.WordRevealed = false
			RunHangman() // Reprise du jeu
			break
		}
	}
	if Currentdatagame.WordRevealed == true { // Si le mot est deviné entièrement
		fmt.Println("Felicitations ! Vous avez déviné !")
		time.Sleep(3 * time.Second)
		Exit() // Sortie du jeu
	}
*/
}

func CheckUsedLetter() { // Verifie si la lettre entrée a déja été utilisée, sinon elle est ajoutée a la liste 
	for _, char := range Currentdatagame.Used_letters {
		if Currentdatagame.CurrentLetter == string(char) {

		} else {
			Currentdatagame.Used_letters = append(Currentdatagame.Used_letters, Currentdatagame.CurrentLetter)
		}
	}
}

func LetterPresent() { // Verifie si la lettre est présente dans le mot
	presentInWord := false
	presentInWord = false // Reinitialisation de la variable
	for i := 0; i < len(Currentdatagame.OriginalWord); i++ { // Boucle qui vérifie si la lettre entrée est présente dans le mot
		if Currentdatagame.CurrentLetter == string(Currentdatagame.OriginalWord[i]) {
		presentInWord = true // Si la lettre est présente, la variable passe vraie
		}
	}
	if presentInWord { // Si la lettre est présente dans le mot
		for index, char := range Currentdatagame.OriginalWord {
			if Currentdatagame.CurrentLetter == string(char) {
				Currentdatagame.Hidden_word[index] = Currentdatagame.CurrentLetter // Remplacement de _ avec la lettre devinée précedemment
			}
		}
		} else { // Si la lettre n'est pas dans le mot
			if Currentdatagame.Try < 9 { // S'il reste encore des essais possibles
				Currentdatagame.Try++ // Incrémentation du compteur d'essais
			} else { // Si tous les essais ont été utilisées
				Currentdatagame.Running = false
			}
		}
}

func CheckWordRevealed() { // Fonction qui vérifie si le mot à été entièrement deviné
	for index, _ := range Currentdatagame.OriginalWord { // Boucle qui vérifie s'il y a encore des underscore dans le mot caché
		if Currentdatagame.Hidden_word[index] == "_" {
			Currentdatagame.WordRevealed = false
			break
		} else {
			Currentdatagame.WordRevealed = true
		}
	}
}