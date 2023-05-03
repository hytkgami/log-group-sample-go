package internal

import (
	"context"
	"fmt"

	"github.com/glassonion1/logz"
	"go.opentelemetry.io/otel/trace"
)

func LogInfof(ctx context.Context, format string, args ...any) {
	go func() {
		sc := trace.SpanContextFromContext(ctx)
		fmt.Printf("traceID: %s, spanID: %s\n", sc.TraceID().String(), sc.SpanID().String())
	}()
	logz.Infof(ctx, format, args...)
}
