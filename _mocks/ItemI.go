package mocks

import helpers "github.com/lcollin/warehouse/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/lcollin/warehouse/models"
import multipart "mime/multipart"
import sql "github.com/ghmeier/bloodlines/gateways/sql"

// ItemI is an autogenerated mock type for the ItemI type
type ItemI struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *ItemI) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: _a0, _a1, _a2
func (_m *ItemI) GetAll(_a0 int, _a1 int, _a2 sql.Search) ([]*models.Item, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*models.Item
	if rf, ok := ret.Get(0).(func(int, int, sql.Search) []*models.Item); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, sql.Search) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllInStock provides a mock function with given fields: _a0, _a1
func (_m *ItemI) GetAllInStock(_a0 int, _a1 int) ([]*models.Item, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.Item
	if rf, ok := ret.Get(0).(func(int, int) []*models.Item); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *ItemI) GetByID(_a0 string) (*models.Item, error) {
	ret := _m.Called(_a0)

	var r0 *models.Item
	if rf, ok := ret.Get(0).(func(string) *models.Item); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRoasterID provides a mock function with given fields: _a0, _a1, _a2
func (_m *ItemI) GetByRoasterID(_a0 int, _a1 int, _a2 string) ([]*models.Item, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*models.Item
	if rf, ok := ret.Get(0).(func(int, int, string) []*models.Item); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ItemI) Insert(_a0 *models.Item) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Item) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ItemI) Update(_a0 *models.Item, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Item, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: _a0, _a1, _a2
func (_m *ItemI) Upload(_a0 string, _a1 string, _a2 multipart.File) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, multipart.File) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ helpers.ItemI = (*ItemI)(nil)
