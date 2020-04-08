package chinese

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
zh-CN,Chinese (People's Republic of China)
zh-HK,Chinese (Hong Kong)
zh-SG,Chinese (Singapore)
zh-TW,Chinese (Taiwan)
*/

func init() {
	tag.Register("zh-CN", "Chinese (People's Republic of China)")
	tag.Register("zh-HK", "Chinese (Hong Kong)")
	tag.Register("zh-SG", "Chinese (Singapore)")
	tag.Register("zh-TW", "Chinese (Taiwan)")
}
