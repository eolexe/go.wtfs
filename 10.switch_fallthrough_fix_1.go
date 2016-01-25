

import "fmt"

func main() {
	fmt.Println(selection("y"))
}

// Usually fallthrough is a default behaviour in a language.
// But Go automatically breaks after each case

func selection(in string) (choice string) {
	switch in {
	case "y":
		choice = "why"
		fallthrough
	case "yeah":
		choice = "yes"
		break
	default:
		choice = "no"
	}
	return
}
