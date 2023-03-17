// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockCache

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

// CommentCreate provides a mock function with given fields: ctx, _a1
func (_m *Comment) CommentCreate(ctx context.Context, _a1 comment.Comment) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, comment.Comment) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommentDelete provides a mock function with given fields: ctx, commentID
func (_m *Comment) CommentDelete(ctx context.Context, commentID string) error {
	ret := _m.Called(ctx, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, commentID)
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

// CommentGetOne provides a mock function with given fields: ctx, commentID
func (_m *Comment) CommentGetOne(ctx context.Context, commentID string) (comment.Comment, error) {
	ret := _m.Called(ctx, commentID)

	var r0 comment.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (comment.Comment, error)); ok {
		return rf(ctx, commentID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) comment.Comment); ok {
		r0 = rf(ctx, commentID)
	} else {
		r0 = ret.Get(0).(comment.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, commentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommentInvalidate provides a mock function with given fields: ctx
func (_m *Comment) CommentInvalidate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommentSetAll provides a mock function with given fields: ctx, meta, params, comments
func (_m *Comment) CommentSetAll(ctx context.Context, meta query.MetaData, params queryparameter.QueryParameter, comments []comment.Comment) error {
	ret := _m.Called(ctx, meta, params, comments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter, []comment.Comment) error); ok {
		r0 = rf(ctx, meta, params, comments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommentUpdate provides a mock function with given fields: ctx, _a1
func (_m *Comment) CommentUpdate(ctx context.Context, _a1 comment.Comment) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, comment.Comment) error); ok {
		r0 = rf(ctx, _a1)
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