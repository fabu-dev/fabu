package qrcode

import (
	"testing"

	"github.com/boombuler/barcode/qr"
)

func TestQrCode(t *testing.T) {
	qrc := NewQrCode("http://127.0.0.1:8803/v1BED", 300, 300, qr.M, qr.Auto)
	path := GetQrCodeFullPath()
	name, path, err := qrc.Encode("/static/app/")

	t.Log("name", name)
	t.Log("path", path)
	if err != nil {
		t.Log(err)
	}
}
