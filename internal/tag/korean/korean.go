package korean

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
ko-KP,Korean (North Korea)
ko-KR,Korean (Korea)
*/

func init() {
	tag.Register("ko-KP", "Korean (North Korea)")
	tag.Register("ko-KR", "Korean (Korea)")
}
