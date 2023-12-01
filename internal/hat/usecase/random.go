package usecase

import (
	"github.com/f1xend/focus-grpc/pkg/util"
)

type Random struct {
}

func (r *Random) Random(min, max int) int {
	return int(util.RandomInt(1, 5))
}
