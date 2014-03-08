package pinyin

import (
	"testing"
)

/*
Add following test case for last version: 204-03-08
str:莞, code:-8776
str:濮, code:-6745
str:泸, code:-7182
str:漯, code:-6928
str:亳, code:-9743
str:儋, code:-9767
*/
var pinyinTests = map[string]string{
	"hello,世界!": "hello,ShiJie!",
	"中文":        "ZhongWen",
	"汉字":        "HanZi",
	"拼音":        "PinYin",
	"简体字":       "JianTiZi",
	"莞":         "Wan",
	"濮":         "Pu",
	"泸":         "Lu",
	"漯":         "Luo",
	"亳":         "Bo",
	"儋":         "Dan",
}

func TestConvert(t *testing.T) {
	for key, value := range pinyinTests {
		result := Convert(key)
		if result != value {
			t.Errorf("%s != %s", result, value)
		}
	}
}
