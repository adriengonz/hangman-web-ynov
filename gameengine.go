package main

type GameData struct {
	OriginalWord     string
	Hidden_word      []string
	Used_letters     []string
	Try              int
	Pseudo           string
	Error            string
}

var Currentdatagame *GameData