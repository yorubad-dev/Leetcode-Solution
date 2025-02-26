package main

func reverseVowels1(s string) string {
	vowels := "aeiouAEIOU"
	isVowel := make(map[byte]bool, len(vowels))

	for _, vowel := range vowels {
		isVowel[byte(vowel)] = true
	}

	sByte := []byte(s)

	left, right := 0, len(s)-1

	for left < right {
		if !isVowel[s[left]] {
			left++
			continue
		}

		if !isVowel[s[right]] {
			right--
			continue
		}

		sByte[right], sByte[left] = sByte[left], sByte[right]
		left++
		right--
	}

	return string(sByte)
}

func reverseVowels(s string) string {
    sByte := []byte(s)
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	vowelMap := make(map[byte]bool, len(vowels))

    for _, vowel := range vowels {
        vowelMap[vowel] = true
    }

    left, right := 0, len(sByte)-1

    for left <= right {
        if !vowelMap[sByte[left]] {
            left++
            continue
        }

        if !vowelMap[sByte[right]] {
            right--
            continue
        }

        sByte[left], sByte[right] = sByte[right], sByte[left]
        left++
        right--
    }

    return string(sByte)
}

func main() {
	reverseVowels1("IceCreAm")
	reverseVowels("leetcode")
}

