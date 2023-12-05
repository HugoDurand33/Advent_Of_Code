package noel

import (
	"bufio"
	"os"
)

func Nombre_syboles() [][]rune {
	txt, _ := os.Open("files/jour_3.txt")
	scan := bufio.NewScanner(txt)
	matrice := [][]rune{}
	for scan.Scan() {
		ligne := scan.Text()
		tab_rune := []rune(ligne)
		matrice = append(matrice, tab_rune)
	}
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[0]); j++ {

		}
	}
	return matrice
}
