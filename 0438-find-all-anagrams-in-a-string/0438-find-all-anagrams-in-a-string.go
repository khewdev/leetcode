import "reflect"

func findAnagrams(s string, p string) []int {
	result := []int{}
	pCount := map[int]int{}
	sCount := map[int]int{}
	l := 0

	if len(p) > len(s) {
		return result
	}

	for i := range p {
		_, pExist := pCount[int(p[i])]
		if pExist {
			pCount[int(p[i])] = 1 + pCount[int(p[i])]
        } else {
            pCount[int(p[i])] = 1
        }

		_, sExist := sCount[int(s[i])]
		if sExist {
			sCount[int(s[i])] = 1 + sCount[int(s[i])]
        } else {
            sCount[int(s[i])] = 1
        }
	}

	if reflect.DeepEqual(pCount, sCount) {
		result = append(result, 0)
	}

	for r := len(p); r < len(s); r++ {
		_, sExist := sCount[int(s[r])]
		if sExist {
			sCount[int(s[r])] = 1 + sCount[int(s[r])]
        } else {
            sCount[int(s[r])] = 1
        }
		sCount[int(s[l])] = sCount[int(s[l])] - 1

		if sCount[int(s[l])] == 0 {
			delete(sCount, int(s[l]))
		}
		l++
		if reflect.DeepEqual(pCount, sCount) {
			result = append(result, l)
		}
	}

	return result
}