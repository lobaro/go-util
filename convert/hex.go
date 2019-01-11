package convert

import (
	"encoding/hex"
	"strings"
)

func HexToBin(hexData string) []byte {
	hexData = strings.Replace(hexData, " ", "", -1)
	bin, err := hex.DecodeString(hexData)
	if err != nil {
		panic(err)
	}
	return bin
}

func BinToHex(bin []byte) string {
	return hex.EncodeToString(bin)
}