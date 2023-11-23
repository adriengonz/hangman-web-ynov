package main

type GameData struct {
	originalWord     string
	hidden_word      []string
	used_letters     []string
	try              int
	pseudo           string
	Error            string
}

var currentdatagame *GameData