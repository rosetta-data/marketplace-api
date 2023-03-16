// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	context "github.com/evgeniy-dammer/marketplace-api/pkg/context"
	mock "github.com/stretchr/testify/mock"

	organization "github.com/evgeniy-dammer/marketplace-api/internal/domain/organization"

	query "github.com/evgeniy-dammer/marketplace-api/pkg/query"

	queryparameter "github.com/evgeniy-dammer/marketplace-api/pkg/queryparameter"
)

// Organization is an autogenerated mock type for the Organization type
type Organization struct {
	mock.Mock
}

// OrganizationCreate provides a mock function with given fields: ctx, meta, input
func (_m *Organization) OrganizationCreate(ctx context.Context, meta query.MetaData, input organization.CreateOrganizationInput) (string, error) {
	ret := _m.Called(ctx, meta, input)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, organization.CreateOrganizationInput) (string, error)); ok {
		return rf(ctx, meta, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, organization.CreateOrganizationInput) string); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, organization.CreateOrganizationInput) error); ok {
		r1 = rf(ctx, meta, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationDelete provides a mock function with given fields: ctx, meta, organizationID
func (_m *Organization) OrganizationDelete(ctx context.Context, meta query.MetaData, organizationID string) error {
	ret := _m.Called(ctx, meta, organizationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) error); ok {
		r0 = rf(ctx, meta, organizationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrganizationGetAll provides a mock function with given fields: ctx, meta, params
func (_m *Organization) OrganizationGetAll(ctx context.Context, meta query.MetaData, params queryparameter.QueryParameter) ([]organization.Organization, error) {
	ret := _m.Called(ctx, meta, params)

	var r0 []organization.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) ([]organization.Organization, error)); ok {
		return rf(ctx, meta, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, queryparameter.QueryParameter) []organization.Organization); ok {
		r0 = rf(ctx, meta, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]organization.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, meta, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationGetOne provides a mock function with given fields: ctx, meta, organizationID
func (_m *Organization) OrganizationGetOne(ctx context.Context, meta query.MetaData, organizationID string) (organization.Organization, error) {
	ret := _m.Called(ctx, meta, organizationID)

	var r0 organization.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) (organization.Organization, error)); ok {
		return rf(ctx, meta, organizationID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, string) organization.Organization); ok {
		r0 = rf(ctx, meta, organizationID)
	} else {
		r0 = ret.Get(0).(organization.Organization)
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.MetaData, string) error); ok {
		r1 = rf(ctx, meta, organizationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrganizationUpdate provides a mock function with given fields: ctx, meta, input
func (_m *Organization) OrganizationUpdate(ctx context.Context, meta query.MetaData, input organization.UpdateOrganizationInput) error {
	ret := _m.Called(ctx, meta, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, query.MetaData, organization.UpdateOrganizationInput) error); ok {
		r0 = rf(ctx, meta, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewOrganization interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrganization creates a new instance of Organization. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrganization(t mockConstructorTestingTNewOrganization) *Organization {
	mock := &Organization{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
