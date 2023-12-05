package noel

import (
	"bufio"
	"fmt"
	"os"
)

/* code de la 1ere partie */
func Nombre_Fichier() int {
	txt, _ := os.Open("files/jour_1.txt")
	scan := bufio.NewScanner(txt)
	resultat := 0
	for scan.Scan() {
		ligne := scan.Text()
		tab_nombre := []int{}
		tab_rune := []rune(ligne)
		for i := 0; i < len(tab_rune); i++ {
			if tab_rune[i] >= 48 && tab_rune[i] <= 57 {
				tab_nombre = append(tab_nombre, int(tab_rune[i]-48))
			}
		}
		if len(tab_nombre) == 1 {
			resultat += tab_nombre[0]*10 + tab_nombre[0]
		} else {
			resultat += tab_nombre[0]*10 + tab_nombre[len(tab_nombre)-1]
		}
	}
	return resultat
}

// merde ce putain de code
func Nombre_mot_Fichier() int {
	tab_mot := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	txt, _ := os.Open("jour_1.txt")
	scan := bufio.NewScanner(txt)
	resultat := 0
	for scan.Scan() {
		ligne := scan.Text()
		tab_nombre := []int{}
		tab_rune := []rune(ligne)
		tab_lettre := ""
		for i := 0; i < len(tab_rune); i++ {
			if tab_rune[i] >= 48 && tab_rune[i] <= 57 {
				tab_nombre = append(tab_nombre, int(tab_rune[i]-48))
				tab_lettre = ""
			} else {
				tab_lettre += string(tab_rune[i])
				bool := false
				for j := 0; j < len(tab_mot); j++ {
					if tab_mot[j] == tab_lettre {
						tab_nombre = append(tab_nombre, j+1)
						tab_lettre = ""
						break
					}
					if len(tab_lettre) <= len(tab_mot[j]) {
						if tab_lettre == tab_mot[j][:len(tab_lettre)] {
							bool = true
						}
					}
				}
				if !bool {
					if len(tab_lettre) == 1 || len(tab_lettre) == 0 {
						tab_lettre = ""
						tab_lettre += string(tab_rune[i])
					} else {
						tab_lettre = tab_lettre[len(tab_lettre)-1:]
						if i < len(tab_rune)-1 {
							i++
						}
					}
				}
			}
		}
		fmt.Println(tab_nombre)
		if len(tab_nombre) == 1 {
			resultat += tab_nombre[0]*10 + tab_nombre[0]
		} else {
			resultat += tab_nombre[0]*10 + tab_nombre[len(tab_nombre)-1]
		}
	}
	return resultat
}
