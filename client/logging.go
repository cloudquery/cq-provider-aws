package client

import (
	"fmt"

	"github.com/aws/smithy-go/logging"
	"github.com/rs/zerolog"
)

type awsLoggerAdapter struct {
	l zerolog.Logger
}

func (a awsLoggerAdapter) Logf(classification logging.Classification, format string, v ...interface{}) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}
