package main

func mergeAlternately(word1 string, word2 string) string {
    i, j := 0, 0
    m, n := len(word1), len(word2)

    res := make([]byte, 0, m + n)

    for i < m || j < n {
        if i < m {
            res = append(res, word1[i])
            i++
        }

        if j < n {
            res = append(res, word2[j])
            j++
        }
    }

    return string(res)
}
