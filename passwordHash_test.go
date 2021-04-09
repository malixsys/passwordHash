package passwordHash

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name         string
		password     string
		hashSegments int
		hashLength int
		options		 *GenerateOptions
	}{
		{name: "Should Work 1", password: "test1234", hashSegments: 4, hashLength: 175},
		{name: "Should Work 2", password: "4321test", hashSegments: 4, hashLength: 175},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := Generate(tt.password, tt.options)
			hashSegments := strings.Split(hash, "$")
			if len(hashSegments) != tt.hashSegments {
				t.Errorf("Generate() had %d segments. Want %d", len(hashSegments), tt.hashSegments)
			}
			if len(hash) != tt.hashLength {
				t.Errorf("Generate() hash length = %v, want %v", len(hash), tt.hashLength)
			}
		})
	}
}

func TestCompareHashWithPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		isValid  bool
	}{
		{name: "Should Work 1", hash: `whirlpool$2852a60fa97b97e02baa38e2c3ff5f79$100$ef2862f7e9eacf6202c12cbd5813bd1bdc8c4b51e936f70c705e76605f433046e93027783ece0ac02c7c4a5b31346a55717d3ef1dfb110136760d4a62fd60297`, password: `test1234`, isValid: true},
		{name: "Should Not Work 1", hash: `whirlpool$2852a60fa97b97e02baa38e2c3ff5f79$100$ef2862f7e9eacf6202c12cbd5813bd1bdc8c4b51e936f70c705e76605f433046e93027783ece0ac02c7c4a5b31346a55717d3ef1dfb110136760d4a62fd60297`, password: `test12345`, isValid: false},
		{name: "Should Not Work 2", hash: `other$2852a60fa97b97e02baa38e2c3ff5f79$100$ef2862f7e9eacf6202c12cbd5813bd1bdc8c4b51e936f70c705e76605f433046e93027783ece0ac02c7c4a5b31346a55717d3ef1dfb110136760d4a62fd60297`, password: `test12345`, isValid: false},
		{name: "Should Not Work 3", hash: `whirlpool$2852a60fa97b97e02baa38e2c3ff5f79$99$ef2862f7e9eacf6202c12cbd5813bd1bdc8c4b51e936f70c705e76605f433046e93027783ece0ac02c7c4a5b31346a55717d3ef1dfb110136760d4a62fd60297`, password: `test1234`, isValid: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Verify(tt.password, tt.hash)
			if got != tt.isValid {
				t.Errorf("Verify() = %v, want %v", got, tt.isValid)
			}
		})
	}
}
