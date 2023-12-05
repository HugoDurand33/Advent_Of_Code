package noel

func Atoi(s string) int {
	nombre := []rune(s)
	n := len(s)
	if n == 0 {
		return 0
	}
	text := 0
	if nombre[0] == rune(45) {
		for i := 1; i < n; i++ {
			if nombre[i] < '0' || nombre[i] > '9' {
				return 0
			} else {
				text *= 10
				text += int(nombre[i]) - '0'
			}
		}
		text = text * -1
	} else if nombre[0] == rune(43) {
		for i := 1; i < n; i++ {
			if nombre[i] < '0' || nombre[i] > '9' {
				return 0
			} else {
				text *= 10
				text += int(nombre[i]) - '0'
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if nombre[i] < '0' || nombre[i] > '9' {
				return 0
			} else {
				text *= 10
				text += int(nombre[i]) - '0'
			}
		}
	}
	return text
}
