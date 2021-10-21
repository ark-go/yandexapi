package appconf

//https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

var kode []byte

func init() {
	if ser, err := diskSerial(); err != nil {
		log.Println(err.Error())
	} else {
		kode = append(ser, "1249863Yandex5632954"...)

	}
}

func criptConfig(conf []byte) []byte {
	secretKeyBytes, err := scrypt.Key(kode, []byte("salt 798798"), 32768, 8, 1, 32)
	if err != nil {
		log.Println("err:", err)
	}
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	// Вы должны использовать разные одноразовые номера для каждого сообщения, которое вы зашифровываете с помощью
	// тот же ключ. Поскольку nonce здесь имеет длину 192 бита, случайное значение
	// обеспечивает достаточно малую вероятность повторов.
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	// This encrypts "hello world" and appends the result to the nonce.
	encrypted := secretbox.Seal(nonce[:], conf, &nonce, &secretKey)

	// decriptConfig(encrypted, secretKeyBytes)
	return encrypted
}

func decriptConfig(encrypted []byte) []byte {

	secretKeyBytes, err := scrypt.Key(kode, []byte("salt 798798"), 32768, 8, 1, 32)
	if err != nil {
		log.Println("err:", err)
	}
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	// При расшифровке вы должны использовать тот же одноразовый номер и ключ, которые вы использовали для
	// зашифровать сообщение. Один из способов добиться этого - сохранить одноразовый номер
	// вместе с зашифрованным сообщением. Выше мы сохранили одноразовый номер в первом
	// 24 байта зашифрованного текста.

	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error Ошибка декодирования")
	}

	//fmt.Println(string(decrypted))
	return decrypted
}
