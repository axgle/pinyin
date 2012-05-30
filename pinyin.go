package pinyin
import ( 
	"github.com/axgle/mahonia"
	"regexp" 
)
var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

var enc = mahonia.NewEncoder("gbk")
func utf8ToGBK(s string) string{   
   return enc.ConvertString(s)
}	

// convert chinese to pinyin
func Convert(s string) string {
	pyString := ""	
    var str string
	var chrAsc int
	var i1,i2 int
	
	for _, rune := range []rune(s) {
	    str = string(rune)
		if  hzRegexp.MatchString(str) {
		gbkString  := utf8ToGBK(str)
 
			i1 = int(gbkString[0])
			i2 = int(gbkString[1])
			chrAsc = i1*256 + i2 - 65536
			if chrAsc > 0 && chrAsc < 160 {
				pyString += str
			} else {
				if   v,ok:=tableMap[chrAsc];ok {
				   pyString += v
				} else {
				for i := (len(pyValue) - 1); i >= 0; i-- {

					if pyValue[i] <= int(chrAsc) {
						pyString += pyName[i]
						break
					}
				}
				}
			}		
		} else{
		  pyString += str
		}
	}	
	return pyString
}
