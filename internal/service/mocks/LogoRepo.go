// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/himmel520/uoffer/mediaAd/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// LogoRepo is an autogenerated mock type for the LogoRepo type
type LogoRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, logo
func (_m *LogoRepo) Add(ctx context.Context, logo *models.Logo) (*models.LogoResp, error) {
	ret := _m.Called(ctx, logo)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 *models.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Logo) (*models.LogoResp, error)); ok {
		return rf(ctx, logo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Logo) *models.LogoResp); ok {
		r0 = rf(ctx, logo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Logo) error); ok {
		r1 = rf(ctx, logo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields: ctx
func (_m *LogoRepo) Count(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *LogoRepo) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *LogoRepo) GetAll(ctx context.Context) ([]*models.Logo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*models.Logo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*models.Logo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Logo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Logo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWithPagination provides a mock function with given fields: ctx, limit, offset
func (_m *LogoRepo) GetAllWithPagination(ctx context.Context, limit int, offset int) (map[int]*models.Logo, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAllWithPagination")
	}

	var r0 map[int]*models.Logo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (map[int]*models.Logo, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) map[int]*models.Logo); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]*models.Logo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, logoID
func (_m *LogoRepo) GetByID(ctx context.Context, logoID int) (*models.LogoResp, error) {
	ret := _m.Called(ctx, logoID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.LogoResp, error)); ok {
		return rf(ctx, logoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.LogoResp); ok {
		r0 = rf(ctx, logoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, logoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, logo
func (_m *LogoRepo) Update(ctx context.Context, id int, logo *models.LogoUpdate) (*models.LogoResp, error) {
	ret := _m.Called(ctx, id, logo)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.LogoResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *models.LogoUpdate) (*models.LogoResp, error)); ok {
		return rf(ctx, id, logo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *models.LogoUpdate) *models.LogoResp); ok {
		r0 = rf(ctx, id, logo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.LogoResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *models.LogoUpdate) error); ok {
		r1 = rf(ctx, id, logo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogoRepo creates a new instance of LogoRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogoRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogoRepo {
	mock := &LogoRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
