// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/internal/storage (interfaces: ImageServer,RuntimeServer)

// Package criostoragemock is a generated GoMock package.
package criostoragemock

import (
	reflect "reflect"

	types "github.com/containers/image/v5/types"
	storage "github.com/containers/storage"
	types0 "github.com/containers/storage/types"
	storage0 "github.com/cri-o/cri-o/internal/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockImageServer is a mock of ImageServer interface.
type MockImageServer struct {
	ctrl     *gomock.Controller
	recorder *MockImageServerMockRecorder
}

// MockImageServerMockRecorder is the mock recorder for MockImageServer.
type MockImageServerMockRecorder struct {
	mock *MockImageServer
}

// NewMockImageServer creates a new mock instance.
func NewMockImageServer(ctrl *gomock.Controller) *MockImageServer {
	mock := &MockImageServer{ctrl: ctrl}
	mock.recorder = &MockImageServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageServer) EXPECT() *MockImageServerMockRecorder {
	return m.recorder
}

// GetStore mocks base method.
func (m *MockImageServer) GetStore() storage.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore")
	ret0, _ := ret[0].(storage.Store)
	return ret0
}

// GetStore indicates an expected call of GetStore.
func (mr *MockImageServerMockRecorder) GetStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockImageServer)(nil).GetStore))
}

// ImageStatus mocks base method.
func (m *MockImageServer) ImageStatus(arg0 *types.SystemContext, arg1 string) (*storage0.ImageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageStatus", arg0, arg1)
	ret0, _ := ret[0].(*storage0.ImageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageStatus indicates an expected call of ImageStatus.
func (mr *MockImageServerMockRecorder) ImageStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageStatus", reflect.TypeOf((*MockImageServer)(nil).ImageStatus), arg0, arg1)
}

// ListImages mocks base method.
func (m *MockImageServer) ListImages(arg0 *types.SystemContext, arg1 string) ([]storage0.ImageResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]storage0.ImageResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockImageServerMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockImageServer)(nil).ListImages), arg0, arg1)
}

// PrepareImage mocks base method.
func (m *MockImageServer) PrepareImage(arg0 *types.SystemContext, arg1 string) (types.ImageCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareImage", arg0, arg1)
	ret0, _ := ret[0].(types.ImageCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareImage indicates an expected call of PrepareImage.
func (mr *MockImageServerMockRecorder) PrepareImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareImage", reflect.TypeOf((*MockImageServer)(nil).PrepareImage), arg0, arg1)
}

// PullImage mocks base method.
func (m *MockImageServer) PullImage(arg0 *types.SystemContext, arg1 string, arg2 *storage0.ImageCopyOptions) (types.ImageReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.ImageReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockImageServerMockRecorder) PullImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockImageServer)(nil).PullImage), arg0, arg1, arg2)
}

// ResolveNames mocks base method.
func (m *MockImageServer) ResolveNames(arg0 *types.SystemContext, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveNames", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveNames indicates an expected call of ResolveNames.
func (mr *MockImageServerMockRecorder) ResolveNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveNames", reflect.TypeOf((*MockImageServer)(nil).ResolveNames), arg0, arg1)
}

// UntagImage mocks base method.
func (m *MockImageServer) UntagImage(arg0 *types.SystemContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UntagImage indicates an expected call of UntagImage.
func (mr *MockImageServerMockRecorder) UntagImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagImage", reflect.TypeOf((*MockImageServer)(nil).UntagImage), arg0, arg1)
}

// MockRuntimeServer is a mock of RuntimeServer interface.
type MockRuntimeServer struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeServerMockRecorder
}

// MockRuntimeServerMockRecorder is the mock recorder for MockRuntimeServer.
type MockRuntimeServerMockRecorder struct {
	mock *MockRuntimeServer
}

// NewMockRuntimeServer creates a new mock instance.
func NewMockRuntimeServer(ctrl *gomock.Controller) *MockRuntimeServer {
	mock := &MockRuntimeServer{ctrl: ctrl}
	mock.recorder = &MockRuntimeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeServer) EXPECT() *MockRuntimeServerMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method.
func (m *MockRuntimeServer) CreateContainer(arg0 *types.SystemContext, arg1, arg2, arg3, arg4, arg5, arg6, arg7 string, arg8 uint32, arg9 *types0.IDMappingOptions, arg10 []string, arg11 bool) (storage0.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	ret0, _ := ret[0].(storage0.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockRuntimeServerMockRecorder) CreateContainer(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockRuntimeServer)(nil).CreateContainer), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

// CreatePodSandbox mocks base method.
func (m *MockRuntimeServer) CreatePodSandbox(arg0 *types.SystemContext, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 string, arg10 uint32, arg11 *types0.IDMappingOptions, arg12 []string, arg13 bool) (storage0.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePodSandbox", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
	ret0, _ := ret[0].(storage0.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePodSandbox indicates an expected call of CreatePodSandbox.
func (mr *MockRuntimeServerMockRecorder) CreatePodSandbox(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePodSandbox", reflect.TypeOf((*MockRuntimeServer)(nil).CreatePodSandbox), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
}

// DeleteContainer mocks base method.
func (m *MockRuntimeServer) DeleteContainer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockRuntimeServerMockRecorder) DeleteContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockRuntimeServer)(nil).DeleteContainer), arg0)
}

// GetContainerMetadata mocks base method.
func (m *MockRuntimeServer) GetContainerMetadata(arg0 string) (storage0.RuntimeContainerMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerMetadata", arg0)
	ret0, _ := ret[0].(storage0.RuntimeContainerMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerMetadata indicates an expected call of GetContainerMetadata.
func (mr *MockRuntimeServerMockRecorder) GetContainerMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerMetadata", reflect.TypeOf((*MockRuntimeServer)(nil).GetContainerMetadata), arg0)
}

// GetRunDir mocks base method.
func (m *MockRuntimeServer) GetRunDir(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunDir", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunDir indicates an expected call of GetRunDir.
func (mr *MockRuntimeServerMockRecorder) GetRunDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunDir", reflect.TypeOf((*MockRuntimeServer)(nil).GetRunDir), arg0)
}

// GetWorkDir mocks base method.
func (m *MockRuntimeServer) GetWorkDir(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkDir", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkDir indicates an expected call of GetWorkDir.
func (mr *MockRuntimeServerMockRecorder) GetWorkDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkDir", reflect.TypeOf((*MockRuntimeServer)(nil).GetWorkDir), arg0)
}

// RemovePodSandbox mocks base method.
func (m *MockRuntimeServer) RemovePodSandbox(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePodSandbox", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePodSandbox indicates an expected call of RemovePodSandbox.
func (mr *MockRuntimeServerMockRecorder) RemovePodSandbox(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePodSandbox", reflect.TypeOf((*MockRuntimeServer)(nil).RemovePodSandbox), arg0)
}

// SetContainerMetadata mocks base method.
func (m *MockRuntimeServer) SetContainerMetadata(arg0 string, arg1 *storage0.RuntimeContainerMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContainerMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContainerMetadata indicates an expected call of SetContainerMetadata.
func (mr *MockRuntimeServerMockRecorder) SetContainerMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerMetadata", reflect.TypeOf((*MockRuntimeServer)(nil).SetContainerMetadata), arg0, arg1)
}

// StartContainer mocks base method.
func (m *MockRuntimeServer) StartContainer(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartContainer indicates an expected call of StartContainer.
func (mr *MockRuntimeServerMockRecorder) StartContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockRuntimeServer)(nil).StartContainer), arg0)
}

// StopContainer mocks base method.
func (m *MockRuntimeServer) StopContainer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer.
func (mr *MockRuntimeServerMockRecorder) StopContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockRuntimeServer)(nil).StopContainer), arg0)
}
