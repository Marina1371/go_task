// Генерируем ключ
key := []byte("MySuperSecretKey")

// Генерируем случайный вектор инициализации
iv := make([]byte, aes.BlockSize)
if _, err := rand.Read(iv); err != nil {
	panic(err)
}

// Создаём шифратор
block, err := aes.NewCipher(key)
if err != nil {
	panic(err)
}
mode := cipher.NewCBCEncrypter(block, iv)

// Шифруем данные
plaintext := []byte("Hello, World!")
ciphertext := make([]byte, len(plaintext))
mode.CryptBlocks(ciphertext, plaintext)

   



