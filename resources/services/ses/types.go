package ses

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type Template struct {
	*types.EmailTemplateContent
	TemplateName     *string
	CreatedTimestamp *time.Time
	Tags             []types.Tag
}
