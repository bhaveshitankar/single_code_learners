package learngo

func checkMap(key string, m1 map[string]int) bool {
	// normal way to solve
	/*for k := range m1 {
		if k == key {
			return true
		}
	}*/
	// go way to find if key is present in map
	/*_, ok := m1[key]
	if ok {
		return true
	}*/
	// using shorthand.
	if _, ok := m1[key]; ok {
		return true
	}
	return false
}
