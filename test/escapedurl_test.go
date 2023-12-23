package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestEscapeUrl(t *testing.T) {
	rawURL := "https://storage.pusatrumahlelang.com/prl-devel/photo_house/Rumah Tambun Utara (Karang Satria)/0.jpeg"

	encodedURL := strings.ReplaceAll(rawURL, " ", "%20")

	fmt.Println("Encoded URL:", encodedURL)
}
