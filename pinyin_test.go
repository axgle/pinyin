package pinyin

import (
    "testing"
)

var pinyinTests = map[string] string {
    "hello,世界!":"hello,ShiJie!", 
    "囧":"Jiong",
    "中国":"ZhongGuo",
    "中a国":"ZhongaGuo",
    "中a1国":"Zhonga1Guo",
    "a中a1国1":"aZhonga1Guo1",
	"兲朝":"TianChao",
	"和谐":"HeXie",
	 
}

func TestConvert(t *testing.T) {
    for key, value := range pinyinTests {
        result := Convert(key)
        if result != value {
            t.Errorf("%s != %s", result, value)
        }
    }
}