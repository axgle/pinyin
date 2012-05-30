package pinyin
import (	
	"github.com/axgle/mahonia"
)

//define func UTF8ToGBK by using mahonia package
var enc = mahonia.NewEncoder("gbk")
func UTF8ToGBK(s string) string {
	return enc.ConvertString(s)
}