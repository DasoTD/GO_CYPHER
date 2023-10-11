package util

// // package main

// import (
// 	// "crypto/aes"
// 	// "crypto/cipher"
// 	// "crypto/rand"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	// "io"
// )

// // Example struct to encrypt
// type MyStruct struct {
// 	Name  string
// 	Value int
// }

// func main() {
// 	// Define a secret key (32 bytes for AES-256)
// 	// Replace this with a securely generated key in a real application.
// 	key := []byte("this-is-a-32-byte-secret-key-123456")

// 	// Create an instance of MyStruct
// 	obj := MyStruct{Name: "John Doe", Value: 42}

// 	// Encrypt the object
// 	encryptedObj, err := encryptObject(obj, key)
// 	if err != nil {
// 		fmt.Println("Encryption error:", err)
// 		return
// 	}

// 	// Print the encrypted object (in base64)
// 	fmt.Println("Encrypted:", base64.StdEncoding.EncodeToString(encryptedObj))

// 	// Decrypt the object
// 	decryptedObj, err := decryptObject(encryptedObj, key)
// 	if err != nil {
// 		fmt.Println("Decryption error:", err)
// 		return
// 	}

// 	// Print the decrypted object
// 	fmt.Printf("Decrypted: %+v\n", decryptedObj)
// }

// // EncryptObject serializes and encrypts an object using AES-GCM.
// func encryptObject(obj interface{}, key []byte) ([]byte, error) {
// 	// Serialize the object to JSON
// 	objJSON, err := json.Marshal(obj)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Encrypt the JSON data
// 	encryptedData, err := encrypt(objJSON, key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return encryptedData, nil
// }

// // DecryptObject decrypts and deserializes an object using AES-GCM.
// func decryptObject(data []byte, key []byte, obj interface{}) error {
// 	// Decrypt the data
// 	decryptedData, err := Decrypt(data, key)
// 	if err != nil {
// 		return err
// 	}

// 	// Deserialize the JSON data into the object
// 	if err := json.Unmarshal(decryptedData, obj); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // The encrypt and decrypt functions are the same as in the previous response.

