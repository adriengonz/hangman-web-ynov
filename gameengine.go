package main

type GameData struct {
	originalWord     string
	hidden_word      []string
	try              int
	used_letters     []string
	Error            string
	pseudo           string
}

var currentdatagame *GameData