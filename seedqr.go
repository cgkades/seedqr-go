package main

// import qrcode "github.com/skip2/go-qrcode"
// import bib39 "github.com/tyler-smith/go-bip39"
import (
	"flag"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	bip39 "github.com/tyler-smith/go-bip39"
	"os"
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
	outFile := flag.String("o", "", "out PNG file prefix, empty for stdout")
	size := flag.Int("s", 256, "image size (pixel)")
	textArt := flag.Bool("t", false, "print as text-art on stdout")
	negative := flag.Bool("i", false, "invert black and white")
	disableBorder := flag.Bool("d", false, "disable QR Code border")
	mnemonicSize := flag.Int("m", 24, "Size of Mnemonic (12 or 24)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `seedqr -- SeedQR Generator for air-gapped systems
Flags:
`)
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, `seedqr -o myseed.png -t`)
	}
	flag.Parse()
	var entropyBitSize int
	if *mnemonicSize == 12 {
		entropyBitSize = 128
	} else if *mnemonicSize == 24 {
		entropyBitSize = 256
	} else {
		fmt.Errorf("Invalid mnemonic size: %s. Only 12 or 24 are valid", mnemonicSize)
	}

	words := getWordListToMap()
	mnemonic := getMnemonic(entropyBitSize)
	mString := createMnemonicIntString(mnemonic, words)
	// fmt.Println(mString)
	q := generateQRCode(mString)
	if *disableBorder {
		q.DisableBorder = true
	}

	if *textArt {
		art := q.ToSmallString(*negative)
		fmt.Println(art)
	}

	if *negative {
		q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	}
	var png []byte
	png, err := q.PNG(*size)
	checkError(err)

	if *outFile == "" {
		os.Stdout.Write(png)
	} else {
		var fh *os.File
		fh, err := os.Create(*outFile + ".png")
		checkError(err)
		defer fh.Close()
		fh.Write(png)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
