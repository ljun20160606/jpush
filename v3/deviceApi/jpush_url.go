package deviceApi

import (
	"github.com/ljun20160606/jpush"
)

type JPushUrl int

var (
	JPushUrls = [...]string{
		UrlDevice: VersionV3 + "/devices/",
		UrlTag:    VersionV3 + "/tags/",
		UrlAlias:  VersionV3 + "/alias/",
	}
)

const (
	VersionV3 = jpush.JPushDevice + "/v3"
)
const (
	UrlDevice JPushUrl = iota
	UrlTag
	UrlAlias
)

func (j JPushUrl) String() string {
	return JPushUrls[j]
}
