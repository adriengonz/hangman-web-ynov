package main

import (
	"fmt"
	"strings"
)

func RunHangman() { // Fonction majeure du jeu
	if Currentdatagame.CurrentLetter !="" {
		fmt.Println("Lettre utilisée :", Currentdatagame.CurrentLetter)
		CheckUsedLetter()
		fmt.Println("[VERBOSE] Le check de la lettre utilisée a été effectué")
		fmt.Println("[VERBOSE]", Currentdatagame.Used_letters)

		Currentdatagame.Used_letters_string = ConvertListToString(Currentdatagame.Used_letters) // Met à jour les lettres utilisés sur la page web

		LetterPresent()
		fmt.Println("[VERBOSE] Le check de la lettre dans le mot a été effectué")

		Currentdatagame.Hidden_word_string = ConvertListToString(Currentdatagame.Hidden_word) // Met à jour le mot sur la page web

		CheckWordRevealed()
		fmt.Println("[VERBOSE] Le check du mot révélé a été effectué")

		Currentdatagame.CurrentLetter = ""
		fmt.Println("[VERBOSE] La lettre temporaire a été vidée")
	}
}

func CheckUsedLetter() { // Verifie si la lettre entrée a déja été utilisée, sinon elle est ajoutée a la liste
	letterused := false // On initialise la valeur en faux
	for _, char := range Currentdatagame.Used_letters { // Si la lettre à deja été utilisé (si elle est trouvé dans la liste), on modifie la valeur de letterused
		if Currentdatagame.CurrentLetter == string(char) {
			letterused = true
			break
		}
	}
	if !letterused { // Si la lettre n'est pas dans la liste, on l'ajoute
		Currentdatagame.Used_letters = append(Currentdatagame.Used_letters, Currentdatagame.CurrentLetter)
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
				fmt.Println("[VERBOSE] Il reste encore ", 10-Currentdatagame.Try, "essais")
			} else { // Si tous les essais ont été utilisées
				Currentdatagame.Running = false // Arret du jeu
				Currentdatagame.Try++ // Incrémentation du compteur d'essais
				fmt.Println("[VERBOSE] Le mot n'a pas été deviné et il ne reste aucune tentative")
			}
		}
}

func CheckWordRevealed() { // Fonction qui vérifie si le mot à été entièrement deviné
	Currentdatagame.WordRevealed = true // On initialise la valeur comme si le mot etait deviné, et si il y a encore des "_", la valeur reviendra en false
	for index, _ := range Currentdatagame.OriginalWord { // Boucle qui vérifie s'il y a encore des underscore dans le mot caché
		if Currentdatagame.Hidden_word[index] == "_" {
			Currentdatagame.WordRevealed = false
			break
		}
	}
	// à des fins de tests
	if Currentdatagame.WordRevealed {
		Currentdatagame.Running = false // Arret du jeu
		fmt.Println("[VERBOSE] Le mot a été deviné !")
	}
}

func ConvertListToString(char []string) string { // Fonction qui convertit une liste de caractères en chaîne de caractères (utile pour afficher les lettres et le mot sur la page web)
	return strings.Join(char, " ")
}