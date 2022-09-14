package Exercism

func IsIsogram(word string) bool {
	seen := make([]bool, 1<<8)

	for _, char := range word {
		// Deal only with uppercase letters.
		if 97 <= char && char <= 122 {
			char -= 32
		}

		// Make sure we only add uppercase letters.
		if 65 <= char && char <= 90 {
			if seen[char] {
				return false
			}

			seen[char] = true
		}
	}

	return true
}
