// internal/components/questions/questions.go
package questions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Question struct {
	Text         string
	Type         string          // "yn", "select", "check", "dyn"
	Options      []string        // static choices
	Dynamic      func() []string // dynamic source
	AllowNone    bool
	AllowOther   bool
	FreeText     bool
	InputType    string // "email", "phone", "name", "password", etc.
	DisplayStyle string // "bracket", "dot", "numbered", etc.
	DefaultKey   string
}

var styleMap = map[string]func(key, label string) string{
	"bracket":  func(k, l string) string { return fmt.Sprintf("[%s] %s", k, l) },
	"dot":      func(k, l string) string { return fmt.Sprintf("%s. %s", k, l) },
	"paren":    func(k, l string) string { return fmt.Sprintf("%s) %s", k, l) },
	"numbered": func(k, l string) string { return fmt.Sprintf("%s) %s", k, l) },
	"bold":     func(k, l string) string { return fmt.Sprintf("**%s**%s", k, l[1:]) },
	"minimal":  func(_, l string) string { return l },
}

func Ask(q Question) string {
	fmt.Println(q.Text)

	opts := q.Options
	if q.Type == "dyn" && q.Dynamic != nil {
		opts = q.Dynamic()
	}

FilterLoop:
	for {
		displayOpts := opts

		// Offer filter option if more than 5 choices
		if len(displayOpts) > 5 {
			fmt.Println("[f] Filter list")
		}

		hintKeys := []string{"y", "n", "s", "h", "m", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		choiceMap := make(map[string]string)
		hintIdx := 0

		// Show options with hotkeys
		for _, opt := range displayOpts {
			if hintIdx >= len(hintKeys) {
				break
			}
			k := hintKeys[hintIdx]
			label := formatOption(q.DisplayStyle, k, opt)
			fmt.Println(label)
			choiceMap[k] = opt
			hintIdx++
		}

		// Extras
		if q.AllowNone {
			fmt.Println(formatOption(q.DisplayStyle, "0", "None"))
			choiceMap["0"] = "none"
		}
		if q.AllowOther {
			fmt.Println(formatOption(q.DisplayStyle, "x", "Other"))
			choiceMap["x"] = "other"
		}

		// Read input
		fmt.Print("Your choice: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Handle filter
		if input == "f" && len(opts) > 5 {
			fmt.Print("Enter filter text: ")
			filter, _ := reader.ReadString('\n')
			filter = strings.ToLower(strings.TrimSpace(filter))

			// Apply filter
			var filtered []string
			for _, o := range opts {
				if strings.Contains(strings.ToLower(o), filter) {
					filtered = append(filtered, o)
				}
			}

			if len(filtered) == 0 {
				fmt.Println("No matches found. Showing full list again.")
			} else {
				opts = filtered
			}
			continue FilterLoop
		}

		// Return matching mapped value
		if val, ok := choiceMap[input]; ok {
			if val == "other" && q.FreeText {
				return promptFreeText(q.InputType)
			}
			return val
		}

		// Return raw if unmatched
		return input
	}
}

func formatOption(style, key, label string) string {
	if fn, ok := styleMap[style]; ok {
		return fn(key, label)
	}
	return fmt.Sprintf("[%s] %s", key, label) // default
}

func promptFreeText(inputType string) string {
	fmt.Printf("Enter %s: ", inputType)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch inputType {
	case "email":
		matched, _ := regexp.MatchString(`^[^@]+@[^@]+\.[^@]+$`, input)
		if !matched {
			fmt.Println("Invalid email format.")
			return promptFreeText(inputType)
		}
	case "phone":
		matched, _ := regexp.MatchString(`^[0-9\-\+\(\) ]+$`, input)
		if !matched {
			fmt.Println("Invalid phone number format.")
			return promptFreeText(inputType)
		}
	}
	return "other:" + input
}
