package noel

import (
	"bufio"
	"os"
)

func Nombre_Fichier() int {
	txt, _ := os.Open("jour_1.txt")
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
