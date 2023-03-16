// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	comment "github.com/evgeniy-dammer/marketplace-api/internal/domain/comment"
	context "github.com/evgeniy-dammer/marketplace-api/pkg/context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/evgeniy-dammer/marketplace-api/pkg/query"

	queryparameter "github.com/evgeniy-dammer/marketplace-api/pkg/queryparameter"
)

// Comment is an autogenerated mock type for the Comment type
type Comment struct {
	mock.Mock
}

// CommentCreate provides a mock function with given fields: ctx, meta, input
func (_m *Comment) CommentCreate(ctx context.Context, meta query.MetaData, input comment.CreateCommentInput) (string, error) {
	ret := _m.Called(ctx, meta, input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, comment.CreateCommentInput) (string, error)); ok {
		return rf(ctx, meta, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, comment.CreateCommentInput) string); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, comment.CreateCommentInput) error); ok {
		r1 = rf(ctx, meta, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommentDelete provides a mock function with given fields: ctx, meta, commentID
func (_m *Comment) CommentDelete(ctx context.Context, meta query.MetaData, commentID string) error {
	ret := _m.Called(ctx, meta, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) error); ok {
		r0 = rf(ctx, meta, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommentGetAll provides a mock function with given fields: ctx, meta, params
func (_m *Comment) CommentGetAll(ctx context.Context, meta query.MetaData, params queryparameter.QueryParameter) ([]comment.Comment, error) {
	ret := _m.Called(ctx, meta, params)

	var r0 []comment.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) ([]comment.Comment, error)); ok {
		return rf(ctx, meta, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) []comment.Comment); ok {
		r0 = rf(ctx, meta, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, meta, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommentGetOne provides a mock function with given fields: ctx, meta, commentID
func (_m *Comment) CommentGetOne(ctx context.Context, meta query.MetaData, commentID string) (comment.Comment, error) {
	ret := _m.Called(ctx, meta, commentID)

	var r0 comment.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) (comment.Comment, error)); ok {
		return rf(ctx, meta, commentID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) comment.Comment); ok {
		r0 = rf(ctx, meta, commentID)
	} else {
		r0 = ret.Get(0).(comment.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, string) error); ok {
		r1 = rf(ctx, meta, commentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommentUpdate provides a mock function with given fields: ctx, meta, input
func (_m *Comment) CommentUpdate(ctx context.Context, meta query.MetaData, input comment.UpdateCommentInput) error {
	ret := _m.Called(ctx, meta, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, comment.UpdateCommentInput) error); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewComment interface {
	mock.TestingT
	Cleanup(func())
}

// NewComment creates a new instance of Comment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComment(t mockConstructorTestingTNewComment) *Comment {
	mock := &Comment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
