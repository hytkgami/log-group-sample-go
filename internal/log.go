package internal

import (
	"context"

	"github.com/glassonion1/logz"
)

func LogInfof(ctx context.Context, format string, args ...any) {
	logz.Infof(ctx, format, args...)
}
