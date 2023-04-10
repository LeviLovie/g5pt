package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

func pad(input []byte, blockSize int) []byte {
	padding := blockSize - len(input)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padText...)
}

func unpad(input []byte, blockSize int) ([]byte, error) {
	paddingSize := int(input[len(input)-1])
	if paddingSize > blockSize || paddingSize > len(input) {
		return nil, errors.New("invalid padding size")
	}
	for i := len(input) - 1; i >= len(input)-paddingSize; i-- {
		if int(input[i]) != paddingSize {
			return nil, errors.New("invalid padding")
		}
	}
	return input[:len(input)-paddingSize], nil
}

func EncryptHex(input string, key string) string {
	inputBytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	paddedInputBytes := pad(inputBytes, block.BlockSize())

	iv := make([]byte, block.BlockSize())
	stream := cipher.NewCBCEncrypter(block, iv)

	encryptedBytes := make([]byte, len(paddedInputBytes))
	stream.CryptBlocks(encryptedBytes, paddedInputBytes)

	outputHex := hex.EncodeToString(encryptedBytes)

	return outputHex
}

func DecryptHex(inputHex string, key string) string {
	inputBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	iv := make([]byte, block.BlockSize())
	stream := cipher.NewCBCDecrypter(block, iv)

	decryptedBytes := make([]byte, len(inputBytes))
	stream.CryptBlocks(decryptedBytes, inputBytes)

	unpaddedBytes, err := unpad(decryptedBytes, block.BlockSize())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := string(unpaddedBytes)

	outputHex := hex.EncodeToString([]byte(output))

	return outputHex
}

func hexToASCII(hexString string) string {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	asciiString := string(hexBytes)

	return asciiString
}

func asciiToHex(asciiString string) string {
	asciiBytes := []byte(asciiString)

	hexString := hex.EncodeToString(asciiBytes)

	return hexString
}

func encrypt(inputFile, outputFile *os.File) {
	outputFile.WriteString("0x67 0x35 0x70 0x74 0x20 0x64 0x61 0x74 0x61 0x20 0x66 0x6F 0x72 0x6D 0x61 0x74 \n")

	scanner := bufio.NewScanner(inputFile)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		stringToEncrypt := "3a" + asciiToHex(fmt.Sprintf("%06d", i+1)) + "3a" + asciiToHex(line) + "3a"
		result := EncryptHex(stringToEncrypt, "7529437302566106")
		outputFile.WriteString(result + "\n")
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func decrypt(inputFile, outputFile *os.File) {
	scanner := bufio.NewScanner(inputFile)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		if i != 0 {
			decryptedString := DecryptHex(line, "7529437302566106")
			asciiString := hexToASCII(decryptedString)
			resultString := asciiString[8 : len(asciiString)-1]
			outputFile.WriteString(resultString + "\n")
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("Usage: go run main.go <mode> ('e' or 'encrypt', for encrypt from input file to outpt file. Or 'd' or 'decrypt', for decrypt from input file to output file)) <input file> <output file>")
		return
	}
	mode := args[1]
	input := args[2]
	output := args[3]

	inputFile, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	if mode == "encrypt" || mode == "e" {
		encrypt(inputFile, outputFile)
	} else if mode == "decrypt" || mode == "d" {
		decrypt(inputFile, outputFile)
	} else {
		fmt.Println("Invalid mode")
	}
}
