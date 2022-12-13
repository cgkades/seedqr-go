package main

// import qrcode "github.com/skip2/go-qrcode"
// import bib39 "github.com/tyler-smith/go-bip39"
import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	bip39 "github.com/tyler-smith/go-bip39"
	"strings"
)

func getWordListToMap() map[string]string {
	wordMap := make(map[string]string)
	wordlist := bip39.GetWordList()
	for index, word := range wordlist {
		wordMap[word] = fmt.Sprintf("%04d", index)
	}

	return wordMap
}

func getMnemonic(entropyBitSize int) string {
	entropy, err := bip39.NewEntropy(entropyBitSize)
	if err != nil {
		return ""
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return ""
	}
	return mnemonic
}

func createMnemonicIntString(mnemonic string, wordMap map[string]string) string {
	mnemonicList := strings.Fields(mnemonic)
	mnemonicString := ""
	for _, word := range mnemonicList {
		mnemonicString += wordMap[word]
	}
	return mnemonicString

}

func generateQRCode(encodeString string) *qrcode.QRCode {
	qrCode, _ := qrcode.New(encodeString, qrcode.Medium)
	return qrCode
}

func main() {
	words := getWordListToMap()
	mnemonic := getMnemonic(256)
	mString := createMnemonicIntString(mnemonic, words)
	// fmt.Println(mString)
	qrCode := generateQRCode(mString)
	// if file is specified
	qrCode.WriteFile(125, "out.png")

	// if printing is specified
	fmt.Println(qrCode.ToSmallString(false))

}
