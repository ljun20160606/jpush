package v3

import (
	"encoding/base64"
	"github.com/ljun20160606/jpush/v3/deviceApi"
)

type JPusher struct {
	Authorization string
	deviceApi.Device
}

func NewJPusher(appKey, masterSecret string) *JPusher {
	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+masterSecret))
	return &JPusher{
		Authorization: authorization,
		Device:        deviceApi.NewDevice(authorization),
	}
}
