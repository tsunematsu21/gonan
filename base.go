package gonan

import "fmt"

func Say(speeker Character, messages ...string) {
	fmt.Printf("💬%s\n", speeker.DisplayName())
	for _, message := range messages {
		fmt.Printf("  %s\n", message)
	}
	fmt.Printf("\n")
}

func PlaySound(who string, effect string) {
	fmt.Printf("📣%s\n  %s\n\n", who, effect)
}
