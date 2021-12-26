package categories

import (
	"time"
)

type CategoryUseCase struct {
	contextTimeout time.Duration
	repository     CategoryRepository
}
