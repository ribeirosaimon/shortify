package base62

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Encode(input string) string {
	num := 0
	for _, char := range input {
		num = num*256 + int(char)
	}

	var result string
	for num > 0 {
		remainder := num % 62
		result = string(base62Chars[remainder]) + result
		num /= 62
	}

	return result
}
