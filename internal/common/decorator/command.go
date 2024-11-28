package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func ApplyCommandDecorators[C any, R any](
	handler CommandHandler[C, R],
	logger *logrus.Entry,
) CommandHandler[C, R] {
	return commandLoggingDecorator[C, R]{
		base:   handler,
		logger: logger,
	}
}
