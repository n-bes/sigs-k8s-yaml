package yaml

import "testing"

type SomeStruct struct {
	data string `yaml:"data"`
}

func TestUUnmarshal(f *testing.T) {
	someStruct := SomeStruct{}
	_ = Unmarshal([]byte{0x26, 0x30, 0x30}, someStruct)
}
