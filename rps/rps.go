package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //Piedra vence a las tijeras (tijeras + 1 )%3 =0
	PAPER    = 1 //Papel vence a la piedra (piedra + 1 ) % 3 = 1
	SCISSORS = 2 //Tijeras vence a papel (papel + 1 ) %3 = 2
)

// Estructura para dar respuesta
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Sigue asi",
	"Bien hecho",
	"Eres de otro planeta",
}

var loseMessages = []string{
	"Intentalo de nuevo",
	"No te rindas",
	"Animo, tu puedes",
}

var drawMessages = []string{
	"Nadie gana, intentalo de nuevo",
	"Por poco",
	"Las grandes mentes piensan igual",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {

	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	//Mensaje dependiendo lo que elija la computadora
	switch computerValue{
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora ha elegido PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora ha elegido PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora ha elegido TIJERAS"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue{
		roundResult = "Empate"
		message = drawMessages[messageInt]
	}else if playerValue == (computerValue+1)%3{
		PlayerScore++
		roundResult = "Has ganado"
		message = winMessages[messageInt]
	}else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:	    	message,
		ComputerChoice: 	computerChoice,
		RoundResult:    	roundResult,
		ComputerChoiceInt:  computerChoiceInt,
		ComputerScore: 		strconv.Itoa(ComputerScore),
		PlayerScore: 		strconv.Itoa(PlayerScore),
	}

}