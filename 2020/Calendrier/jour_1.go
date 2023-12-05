package noel

import (
	"bufio"
	noel "noel/fonctions"
	"os"
)

// partie n°1
func Somme_nombre() int {
	txt, _ := os.Open("2020/files/jour_1.txt")
	scan := bufio.NewScanner(txt)
	resultat := 0
	tab_nombre := []int{}
	for scan.Scan() {
		tab_nombre = append(tab_nombre, noel.Atoi(scan.Text()))
	}
	for i := 0; i < len(tab_nombre)-1; i++ {
		for j := i + 1; j < len(tab_nombre)-i; j++ {
			if tab_nombre[i]+tab_nombre[j] == 2020 {
				resultat = tab_nombre[i] * tab_nombre[j]
				break
			}
		}
	}
	return resultat
}

// partie n°2
func Somme_3_nombre() int {
	txt, _ := os.Open("2020/files/jour_1.txt")
	scan := bufio.NewScanner(txt)
	resultat := 0
	tab_nombre := []int{}
	for scan.Scan() {
		tab_nombre = append(tab_nombre, noel.Atoi(scan.Text()))
	}
	for i := 0; i < len(tab_nombre); i++ {
		for j := 0; j < len(tab_nombre); j++ {
			for k := 0; k < len(tab_nombre); k++ {
				if tab_nombre[i]+tab_nombre[j]+tab_nombre[k] == 2020 {
					resultat = tab_nombre[i] * tab_nombre[j] * tab_nombre[k]
					break
				}
			}
		}
	}
	return resultat
}
