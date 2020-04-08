package korean

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
ko-KP,Korean (North Korea)
ko-KR,Korean (Korea)
*/

func init() {
	tag.Register("ko-KP", "Korean (North Korea)")
	tag.Register("ko-KR", "(Korea)")
}
