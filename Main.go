package main

func main() {
	println(isIsomorphic("badc", "baba"))
	//println(isIsomorphic("", ""))
	//println(isIsomorphic("a", "b"))
	//println(isIsomorphic("paper", "title"))
	//println(isIsomorphic("egg", "add"))
	//println(isIsomorphic("foo", "bar"))
}

func isIsomorphic(s string, t string) bool {
	mirrorL := map[uint8]uint8{}
	mirrorR := map[uint8]uint8{}
	for i := 0; i < len(s); i++ {
		left := s[i]
		right := t[i]

		if mirrorValue, isExist := mirrorL[left]; !isExist {
			mirrorL[left] = right
		} else {
			if mirrorValue != right {
				return false
			}
		}

		if mirrorValue, isExist := mirrorR[right]; !isExist {
			mirrorR[right] = left
		} else {
			if mirrorValue != left {
				return false
			}
		}
	}
	return true
}
