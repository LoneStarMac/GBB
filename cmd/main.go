// cmd/gbb/main.go
package main

import (
	"fmt"

	q "github.com/lonestarmac/gbb/internal/components/questions"
	str "github.com/lonestarmac/gbb/internal/components/strings"
)

func main() {
	selection := q.Ask(q.Question{
		Text: str.MainMenuPrompt,
		Type: "select",
		Options: []string{
			str.MenuCheckOnline,
			str.MenuChangeInstall,
			str.MenuChangeFolder,
			str.MenuChangeBuilder,
			str.MenuOptions,
			str.MenuBuildNow,
		},
		DisplayStyle: "numbered",
	})

	switch selection {
	case str.MenuCheckOnline:
		fmt.Println("Checking plugins...")
	case str.MenuBuildNow:
		fmt.Println("Building now...")
	default:
		fmt.Println("TODO:", selection)
	}
}
