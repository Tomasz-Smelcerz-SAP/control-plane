// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"

// ScenariosDefService is an autogenerated mock type for the ScenariosDefService type
type ScenariosDefService struct {
	mock.Mock
}

// EnsureScenariosLabelDefinitionExists provides a mock function with given fields: ctx, tenant
func (_m *ScenariosDefService) EnsureScenariosLabelDefinitionExists(ctx context.Context, tenant string) error {
	ret := _m.Called(ctx, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAvailableScenarios provides a mock function with given fields: ctx, tenant
func (_m *ScenariosDefService) GetAvailableScenarios(ctx context.Context, tenant string) ([]string, error) {
	ret := _m.Called(ctx, tenant)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
