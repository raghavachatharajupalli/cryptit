package encryption

func EncryptString(str string) string {
	encryptedString := ""

	for _, schar := range str {
		asciicode := int(schar)
		echar := string(asciicode + 3)
		encryptedString += echar
	}

	return encryptedString
}
