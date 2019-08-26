package helper

func Padding(str string) string {
	maxStrLenght := 16
	space := 6
	totalSpace := maxStrLenght - len(str) + space
	res := ""
	for i := 0; i < totalSpace; i++ {
		res += " "
	}
	return res
}

func Line() string {
	return "-------------------------------------------------------"
}
