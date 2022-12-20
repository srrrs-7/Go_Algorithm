package lang

import (
	"fmt"

	"github.com/sRRRs-7/Go_Algorithm.git/lang/hebon"
	"github.com/sRRRs-7/Go_Algorithm.git/lang/mecab"
)

func ConvertLanguage() {
	en := hebon.ToHebon("こんにちは")
	fmt.Println(en)

	strs := mecab.NewMecab("私はアニメが大好きです")
	fmt.Println(strs)

}
