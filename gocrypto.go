// gocrypto - basic implementation of cryptography in Go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	//"log"
)

func main() {

	//Generate 16 Byte Random AES(128) Key
	key := make([]byte, 16)
	rand.Read(key)

	//Static 16 Byte Key Example
	//key := []byte("example key 1234")

	fmt.Printf("Whatup Doggy!  --  Let's Encyrpt some shit!\n\n\tUsing this key:  %s\n\n", base64.StdEncoding.EncodeToString(key))

	//Create an Instance of the AES Cipher & Wrap it with GCM
	block, _ := aes.NewCipher(key)
	gcmAES, _ := cipher.NewGCM(block)
	fmt.Printf("The Size of the nonce...%d\n\n----------------\n\n", gcmAES.NonceSize())

	plaintext := []byte("Punch it Chewy!")
	fmt.Printf("Plaintext:\n\t%s\n\n", plaintext)
	var ciphertext []byte

	//Generate a Random Nonce
	nonce := make([]byte, gcmAES.NonceSize())
	rand.Read(nonce)

	//Encrypt the plaintext data
	ciphertext_data := gcmAES.Seal(ciphertext, nonce, plaintext, nil)

	fmt.Printf("Ciphertext:\n\t%s\n\n", base64.StdEncoding.EncodeToString(ciphertext_data))

	//Decrypt the ciphertext data
	decryptedPlaintext := make([]byte, len(plaintext))
	plaintext_data, _ := gcmAES.Open(decryptedPlaintext, nonce, ciphertext_data, nil)

	//log.Println(err)
	fmt.Printf("Decrypted Text:\n\t%s\n\n", plaintext_data)
}
