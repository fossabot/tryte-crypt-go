package main

import (
	"bufio"
	"fmt"
	"github.com/iotaledger/iota.go/trinary"
	tryteCipher "github.com/yegamble/tryte-crypt-go/tryte-cipher"
	"os"
	"strconv"
	"strings"
)

func main() {
	var defaultOptions tryteCipher.ScryptOptions
	buf := bufio.NewReader(os.Stdin)

	seed := promptSeed(buf)

	passphrase := promptPassphrase(buf)

	toughnessInt := promptDifficulty(buf)

	encrypt, err := tryteCipher.Encrypt(seed, passphrase, defaultOptions, toughnessInt)
	if err != nil {
		fmt.Println(err)
		main()
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Encrypted Seed: " + encrypt)
	}
}

func promptSeed(buf *bufio.Reader) string {

	fmt.Print("Enter IOTA Seed: ")
	seed, err := buf.ReadBytes('\n')
	if err != nil {
		return ""
	}

	seedString := strings.TrimSuffix(string(seed), "\n")

	if seedString == "" {
		fmt.Println("Seed is Empty")
		promptSeed(buf)
	}

	err = trinary.ValidTrytes(seedString)
	if err != nil {
		fmt.Println(err)
		return promptSeed(buf)
	}

	return seedString
}

func promptPassphrase(buf *bufio.Reader) string {

	fmt.Print("Enter Passprahse: ")
	passphraseBytes, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	passphraseString := strings.TrimSuffix(string(passphraseBytes), "\n")

	if passphraseString == "" {
		fmt.Println("Passphrase is Empty")
		return promptPassphrase(buf)
	}

	return passphraseString
}

func promptDifficulty(buf *bufio.Reader) int {
	fmt.Print("Enter Encryption Difficulty (0-9): ")
	toughness, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}

	toughnessString := strings.TrimSuffix(string(toughness), "\n")

	toughnessInt, err := strconv.Atoi(toughnessString)
	if err != nil {
		fmt.Println(err)
		return promptDifficulty(buf)
	}

	return toughnessInt
}