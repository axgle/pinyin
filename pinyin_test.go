package pinyin

import (
	"testing"
)

var pinyinTests = map[string]string{
	"hello,世界!": "hello,ShiJie!",
	"中文":        "ZhongWen",
	"汉字":        "HanZi",
	"拼音":        "PinYin",
	"简体字":       "JianTiZi",
	
}

func TestConvert(t *testing.T) {
	for key, value := range pinyinTests {
		result := Convert(key)
		if result != value {
			t.Errorf("%s != %s", result, value)
		}
	}
}
