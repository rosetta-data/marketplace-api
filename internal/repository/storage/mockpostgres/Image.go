// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	image "github.com/evgeniy-dammer/marketplace-api/internal/domain/image"
	context "github.com/evgeniy-dammer/marketplace-api/pkg/context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/evgeniy-dammer/marketplace-api/pkg/query"

	queryparameter "github.com/evgeniy-dammer/marketplace-api/pkg/queryparameter"
)

// Image is an autogenerated mock type for the Image type
type Image struct {
	mock.Mock
}

// ImageCreate provides a mock function with given fields: ctx, meta, input
func (_m *Image) ImageCreate(ctx context.Context, meta query.MetaData, input image.CreateImageInput) (string, error) {
	ret := _m.Called(ctx, meta, input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, image.CreateImageInput) (string, error)); ok {
		return rf(ctx, meta, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, image.CreateImageInput) string); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, image.CreateImageInput) error); ok {
		r1 = rf(ctx, meta, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageDelete provides a mock function with given fields: ctx, meta, imageID
func (_m *Image) ImageDelete(ctx context.Context, meta query.MetaData, imageID string) error {
	ret := _m.Called(ctx, meta, imageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) error); ok {
		r0 = rf(ctx, meta, imageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageGetAll provides a mock function with given fields: ctx, meta, params
func (_m *Image) ImageGetAll(ctx context.Context, meta query.MetaData, params queryparameter.QueryParameter) ([]image.Image, error) {
	ret := _m.Called(ctx, meta, params)

	var r0 []image.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) ([]image.Image, error)); ok {
		return rf(ctx, meta, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) []image.Image); ok {
		r0 = rf(ctx, meta, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]image.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, meta, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageGetOne provides a mock function with given fields: ctx, meta, imageID
func (_m *Image) ImageGetOne(ctx context.Context, meta query.MetaData, imageID string) (image.Image, error) {
	ret := _m.Called(ctx, meta, imageID)

	var r0 image.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) (image.Image, error)); ok {
		return rf(ctx, meta, imageID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) image.Image); ok {
		r0 = rf(ctx, meta, imageID)
	} else {
		r0 = ret.Get(0).(image.Image)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, string) error); ok {
		r1 = rf(ctx, meta, imageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageUpdate provides a mock function with given fields: ctx, meta, input
func (_m *Image) ImageUpdate(ctx context.Context, meta query.MetaData, input image.UpdateImageInput) error {
	ret := _m.Called(ctx, meta, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, image.UpdateImageInput) error); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImage interface {
	mock.TestingT
	Cleanup(func())
}

// NewImage creates a new instance of Image. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImage(t mockConstructorTestingTNewImage) *Image {
	mock := &Image{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}