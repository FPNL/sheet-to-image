package tableimage

var DefaultConfig = Config{
	fontSize:     14,
	hPadding:     0,
	wrapWordsLen: 12,
}

type Config struct {
	fontSize     int // 14
	hPadding     int // 0
	vPadding     int // 2
	wrapWordsLen int // 12
}

func (receiver Config) Height() int {
	return receiver.fontSize + receiver.vPadding
}

func (receiver Config) Width() int {
	return receiver.hPadding + receiver.LineLen()
}

func (receiver Config) LineLen() int {
	return receiver.wrapWordsLen * receiver.fontSize
}
