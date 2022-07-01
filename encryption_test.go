package encryption

import (
	"strings"
	"testing"
)

var sentence = "hello, this is a simple message from somewhere"
var paragraph = strings.Repeat(sentence, 10)
var key = []byte("passwordpasswordpasswordpassword")

func BenchmarkEncryptAES_WholeContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptAES(key, paragraph)
	}
}

func BenchmarkEncryptAES_Separate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
		EncryptAES(key, sentence)
	}
}
