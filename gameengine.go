package main

type GameData struct {
	OriginalWord     string // Mot Original selectionné
	Hidden_word      []string // Mot caché sous forme de tableau de chaine de caractètres
	Used_letters     []string // Lettres utilisés sous forme de tableau de chaine de caractètres
	Try              int // Essais
	Pseudo           string // Pseudo du joueur utilisé sur les pages de win et loose
	CurrentLetter    string // Lettre selectionné par le joueur (stocké temporairement)
	Error            string
	Running          bool // Si une partie est en cours (utilisé pour la boucle du jeu dans le main)
	WordRevealed     bool // Boolen pour connaitre si le mot est entièrement révélé ou non (Si true, le joueur sera redirigé vers la page de win)
}

var Currentdatagame *GameData