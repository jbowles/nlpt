package tkz

func DeDupBytes(a []byte) []byte {
	result := []byte{}
	seen := map[byte]byte{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
