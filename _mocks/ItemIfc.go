package mocks

import gin "gopkg.in/gin-gonic/gin.v1"
import handlers "github.com/lcollin/warehouse/handlers"
import mock "github.com/stretchr/testify/mock"

// ItemIfc is an autogenerated mock type for the ItemIfc type
type ItemIfc struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx
func (_m *ItemIfc) Delete(ctx *gin.Context) {
	_m.Called(ctx)
}

// GetJWT provides a mock function with given fields:
func (_m *ItemIfc) GetJWT() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// New provides a mock function with given fields: ctx
func (_m *ItemIfc) New(ctx *gin.Context) {
	_m.Called(ctx)
}

// Time provides a mock function with given fields:
func (_m *ItemIfc) Time() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx
func (_m *ItemIfc) Update(ctx *gin.Context) {
	_m.Called(ctx)
}

// Upload provides a mock function with given fields: ctx
func (_m *ItemIfc) Upload(ctx *gin.Context) {
	_m.Called(ctx)
}

// View provides a mock function with given fields: ctx
func (_m *ItemIfc) View(ctx *gin.Context) {
	_m.Called(ctx)
}

// ViewAll provides a mock function with given fields: ctx
func (_m *ItemIfc) ViewAll(ctx *gin.Context) {
	_m.Called(ctx)
}

// ViewByRoasterID provides a mock function with given fields: ctx
func (_m *ItemIfc) ViewByRoasterID(ctx *gin.Context) {
	_m.Called(ctx)
}

var _ handlers.ItemIfc = (*ItemIfc)(nil)
