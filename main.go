package gocolor

var defalt string = "\u001b[0m"

func Defalt(text string) string {
	result := "\u001b[0m" + text
	return result
}

func Red(text string) string {
	result := "\u001b[31m" + text + defalt
	return result
}

func Blue(text string) string {
	result := "\u001b[34m" + text + defalt
	return result
}

func Yellow(text string) string {
	result := "\u001b[33m" + text + defalt
	return result
}

func Green(text string) string {
	result := "\u001b[32m" + text + defalt
	return result
}

func Purple(text string) string {
	result := "\u001b[35m" + text + defalt
	return result
}

func Cayn(text string) string {
	result := "\u001b[36m" + text + defalt
	return result
}

func White(text string) string {
	result := "\u001b[37m" + text + defalt
	return result
}

func Black(text string) string {
	result := "\u001b[30m" + text + defalt
	return result
}
