package deviceApi

import (
	"github.com/LFZJun/jpush"
)

const (
	VersionV3 = jpush.JPushDevice + "/v3"
	UrlDevice = VersionV3 + "/devices/"
	UrlTag    = VersionV3 + "/tags/"
)
