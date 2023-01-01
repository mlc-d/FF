package protected

import (
	"gitlab.com/mlc-d/ff/pkg/comment"
	"gitlab.com/mlc-d/ff/pkg/thread"
	"gitlab.com/mlc-d/ff/pkg/topic"
)

var (
	commentService = comment.NewService()
	threadService  = thread.NewService()
	topicService   = topic.NewService()
)
