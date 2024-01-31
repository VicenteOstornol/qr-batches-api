package qr

import "github.com/skip2/go-qrcode"

func GenerateQR(i int, url string) ([]byte, error) {
	bytes, err := qrcode.Encode(url, qrcode.Medium, 128)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
