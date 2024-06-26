// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "course/internal/service/dto"

	mock "github.com/stretchr/testify/mock"

	model "course/internal/model"
)

// CheckpointStorage is an autogenerated mock type for the CheckpointStorage type
type CheckpointStorage struct {
	mock.Mock
}

// CreateCheckpoint provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) CreateCheckpoint(ctx context.Context, request *dto.CreateCheckpointRequest) (*model.Checkpoint, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateCheckpoint")
	}

	var r0 *model.Checkpoint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateCheckpointRequest) (*model.Checkpoint, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateCheckpointRequest) *model.Checkpoint); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Checkpoint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateCheckpointRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePassage provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreatePassage")
	}

	var r0 *model.Passage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreatePassageRequest) (*model.Passage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreatePassageRequest) *model.Passage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Passage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreatePassageRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCheckpoint provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) DeleteCheckpoint(ctx context.Context, request *dto.DeleteCheckpointRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCheckpoint")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteCheckpointRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePassage provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeletePassage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeletePassageRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCheckpoint provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) GetCheckpoint(ctx context.Context, request *dto.GetCheckpointRequest) (*model.Checkpoint, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetCheckpoint")
	}

	var r0 *model.Checkpoint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetCheckpointRequest) (*model.Checkpoint, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetCheckpointRequest) *model.Checkpoint); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Checkpoint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetCheckpointRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPassage provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetPassage")
	}

	var r0 *model.Passage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetPassageRequest) (*model.Passage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetPassageRequest) *model.Passage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Passage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetPassageRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPassages provides a mock function with given fields: ctx, request
func (_m *CheckpointStorage) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListPassages")
	}

	var r0 []*model.Passage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListPassagesRequest) ([]*model.Passage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListPassagesRequest) []*model.Passage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Passage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListPassagesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCheckpointStorage creates a new instance of CheckpointStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCheckpointStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *CheckpointStorage {
	mock := &CheckpointStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
