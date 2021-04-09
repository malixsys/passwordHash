package passwordHash

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/hex"
	"github.com/jzelinskie/whirlpool"
	"strconv"
	"strings"
)

const ALGO = "whirlpool"

func GenerateSalt(saltSize int) string {
	var salt = make([]byte, saltSize / 2)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(salt)
}

func GenerateHash(pw string, salt string, iterations int) string {
	current := pw

	for n := 1; n <= iterations; n++ {
		// Create a new HMAC by defining the hash type and the key (as byte array)
		h := hmac.New(whirlpool.New, []byte(salt))

		// Write Data to it
		h.Write([]byte(current))

		current = hex.EncodeToString(h.Sum(nil))
	}
	return strings.Join([]string{ALGO, salt, strconv.Itoa(iterations), current}, "$");
}

func Verify(password string, hashedPassword string) bool {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 4 {
		return false;
	}
	if parts[0] != ALGO {
		return false;
	}
	iterations, err := strconv.Atoi(parts[2])
	if err != nil {
		return false;
	}
	hash := GenerateHash(password, parts[1], iterations)
	return hash == hashedPassword
}

type GenerateOptions struct {
	saltLength int
	iterations int
}

func Generate(pw string, options *GenerateOptions) string {
	if options == nil{
		options = &GenerateOptions{
			saltLength: 32,
			iterations: 100,
		}
	}
	if options.saltLength == 0 {
		options.saltLength = 32
	}

	if options.iterations == 0 {
		options.iterations = 100
	}
	salt := GenerateSalt(options.saltLength)
	return GenerateHash(pw, salt, options.iterations);
}

func IsHashed(password string) bool {
	parts := strings.Split(password, "$")
	return len(parts) == 4
}

