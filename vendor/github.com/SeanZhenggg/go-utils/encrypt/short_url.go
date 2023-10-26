package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	charList = [62]string{"V", "0", "S", "1", "s", "6", "F", "v", "j", "D", "C", "O", "c", "2", "G", "u", "X", "P", "K", "3", "i", "J", "f", "t", "y", "k", "l", "q", "7", "n", "b", "a", "A", "H", "B", "e", "W", "d", "h", "o", "m", "Z", "I", "E", "N", "M", "g", "R", "x", "p", "4", "Y", "r", "9", "w", "U", "L", "Q", "8", "5", "T", "z"}
)

func GenShortUrlById(id int64) string {
	result := base10ToBase62(id)

	return strings.Join(shuffle(strings.Split(result, "")), "")
}

func GenShortUrlByInput(input string) {
	hashByte := md5.Sum([]byte(input))
	inputMd5 := hex.EncodeToString(hashByte[:])
	reUrl := [4]string{}
	for i := 0; i < 4; i++ {
		str := inputMd5[i*8 : (i*8)+8]
		decimal, err := strconv.ParseInt(str, 16, 64)
		if err != nil {
			return
		}

		strint := decimal & 0x3fffffff
		outStr := ""

		for j := 0; j < 6; j++ {
			idx := 0x00003D & strint
			outStr += charList[idx]
			strint >>= 5
		}

		reUrl[i] = outStr
	}

	for _, v := range reUrl {
		fmt.Printf("hash value : %s\n", v)
	}
}

// Fisher-Yates shuffle algorithm
func shuffle(input []string) []string {
	for i := 0; i < len(input); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		j := r.Intn(i + 1)
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func base10ToBase62(input int64) string {
	var result string
	radix := int64(len(charList))
	quotient := input
	for quotient > 0 {
		residual := quotient % radix
		result = charList[residual] + result
		quotient /= radix
	}

	return result
}
