package strings

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestDelExtaSpace(t *testing.T) {
	testStr := `hello     world   !`
	Convey("去掉字符串中的多余空格",t, func() {
	 s := DelExtaSpace(testStr)
	 So(s,ShouldResemble,"hello world !")
	})
}
