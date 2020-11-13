package qrcode

import (
	"testing"

	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/boombuler/barcode/qr"
)

func TestQrCode(t *testing.T) {
	qrc := qrcode.NewQrCode("http://127.0.0.1:8803/v1BED", 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(path)

	if err != nil {
		t.Log(err)
	}
}
