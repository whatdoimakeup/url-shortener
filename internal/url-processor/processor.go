package urlprocessor

func Encode(id int) string {
	const base62Chars = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"
	shortURL := ""

	for id > 0 {
		shortURL = string(base62Chars[id%58]) + shortURL
		id /= 58
	}

	return shortURL
}