//_test.goなら、パッケージ名は以下のものが適切(標準ライブラリより)
package mypkg_test

import (
	"fmt"
	"testing"

	"github.com/ShunyaNagashige/is-even/mypkg"
)

func TestIsEven(t *testing.T){
	expect:=true
	actual,_:=mypkg.IsEven(2)
	if actual!=expect{
		//失敗理由を出力してテストを失敗させる
		t.Errorf(`expect="%t" actual="%t"`,expect,actual)
	}
}

func ExampleIsEven(){
	fmt.Println(mypkg.IsEven(2))
	// Output: a
}