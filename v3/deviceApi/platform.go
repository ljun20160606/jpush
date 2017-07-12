package deviceApi

type platform int

const (
	IOS platform = iota
	ANDROID
	ALL
)

var platformString = [...]string{
	IOS:     "ios",
	ANDROID: "android",
	ALL:     "android,ios",
}

func (p platform) String() string {
	return platformString[p]
}
