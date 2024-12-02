package screen

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	//"time"
)

func ChoiceMenu() {
	var choice string
	fmt.Println("Choisissez une map :")
	fmt.Scan(&choice)

	switch choice {
	case "1":
		showMap(1)
	case "2":
		showMap(2)
	case "3":
		showMap(3)
	default:
		fmt.Println("Choix invalide")
	}
}

// Menu affiche le menu
func Menu() {
	clearScreen()
	file, err := os.Open("assets/menu.txt")
	if err != nil {
		fmt.Println("Problème de chargement du menu")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
