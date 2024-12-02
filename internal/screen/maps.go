package screen

import ()

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func showMap (mapNumber int) {
	clearScreen()

	mapPath := fmt.Sprintf("assets/maps/map%d.txt", mapNumber)
	// Ouvrir le fichier contenant la carte
	file, err := os.Open(mapPath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	// Lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	var rows, cols int
	var matrix [][]string

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineCount == 0 {
			// Lire les dimensions (première ligne)
			fmt.Sscanf(line, "%d %d", &rows, &cols)
			matrix = make([][]string, rows)
			for i := range matrix {
				matrix[i] = make([]string, cols)
			}
		} else {
			// Lire les lignes de la carte
			for i, char := range line {
				matrix[lineCount-1][i] = string(char)
			}
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Afficher la matrice pour vérifier
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}
