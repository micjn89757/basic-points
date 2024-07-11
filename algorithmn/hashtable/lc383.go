package hashtable

func canConstruct(ransomNote string, magazine string) bool {
    tmp := make(map[rune]int, len(ransomNote))

    for _, r := range ransomNote {
        if _, ok :=  tmp[r]; ok {
            tmp[r]++
        } else {
            tmp[r] = 1
        }
    }


    for _, r := range magazine {
        if _, ok := tmp[r]; ok {
            tmp[r]--
            if tmp[r] == 0 {
                delete(tmp, r)
            }
        } 
    }

    for _, v := range tmp {
        if v != 0 {
            return false
        }
    }

    return true
}