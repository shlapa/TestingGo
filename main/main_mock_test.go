package main

import (
	gomock "go.uber.org/mock/gomock"
)

type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex.
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance.
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}
