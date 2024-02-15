package service

import (
	"context"

	"github.com/mustthink/microblog/grpc/gen/microblog"
)

func (m *Microblog) GetPostComments(context.Context, *microblog.PostCommentsRequest) (*microblog.PostCommentsResponse, error) {
	panic("implement me")
}
func (m *Microblog) CreateComment(context.Context, *microblog.CreateCommentRequest) (*microblog.CreateCommentResponse, error) {
	panic("implement me")
}
func (m *Microblog) GetComment(context.Context, *microblog.GetCommentRequest) (*microblog.GetCommentResponse, error) {
	panic("implement me")
}
func (m *Microblog) DeleteComment(context.Context, *microblog.DeleteCommentRequest) (*microblog.DeleteCommentResponse, error) {
	panic("implement me")
}
func (m *Microblog) UpdateComment(context.Context, *microblog.UpdateCommentRequest) (*microblog.UpdateCommentResponse, error) {
	panic("implement me")
}
