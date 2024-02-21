// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/azureclients/vmssclient/interface.go
//
// Generated by this command:
//
//	mockgen -copyright_file=/home/runner/work/cloud-provider-azure/cloud-provider-azure/hack/boilerplate/boilerplate.generatego.txt -source=pkg/azureclients/vmssclient/interface.go -package=mockvmssclient Interface
//

// Package mockvmssclient is a generated GoMock package.
package mockvmssclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-08-01/compute"
	azure "github.com/Azure/go-autorest/autorest/azure"
	gomock "go.uber.org/mock/gomock"
	retry "sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, VMScaleSetName string, parameters compute.VirtualMachineScaleSet) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, VMScaleSetName, parameters)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, VMScaleSetName, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, VMScaleSetName, parameters)
}

// CreateOrUpdateAsync mocks base method.
func (m *MockInterface) CreateOrUpdateAsync(ctx context.Context, resourceGroupName, VMScaleSetName string, parameters compute.VirtualMachineScaleSet) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAsync", ctx, resourceGroupName, VMScaleSetName, parameters)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// CreateOrUpdateAsync indicates an expected call of CreateOrUpdateAsync.
func (mr *MockInterfaceMockRecorder) CreateOrUpdateAsync(ctx, resourceGroupName, VMScaleSetName, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAsync", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdateAsync), ctx, resourceGroupName, VMScaleSetName, parameters)
}

// DeallocateInstancesAsync mocks base method.
func (m *MockInterface) DeallocateInstancesAsync(ctx context.Context, resourceGroupName, vmScaleSetName string, vmInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeallocateInstancesAsync", ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// DeallocateInstancesAsync indicates an expected call of DeallocateInstancesAsync.
func (mr *MockInterfaceMockRecorder) DeallocateInstancesAsync(ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeallocateInstancesAsync", reflect.TypeOf((*MockInterface)(nil).DeallocateInstancesAsync), ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
}

// DeleteInstances mocks base method.
func (m *MockInterface) DeleteInstances(ctx context.Context, resourceGroupName, vmScaleSetName string, vmInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstances", ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// DeleteInstances indicates an expected call of DeleteInstances.
func (mr *MockInterfaceMockRecorder) DeleteInstances(ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstances", reflect.TypeOf((*MockInterface)(nil).DeleteInstances), ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
}

// DeleteInstancesAsync mocks base method.
func (m *MockInterface) DeleteInstancesAsync(ctx context.Context, resourceGroupName, vmScaleSetName string, vmInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs, forceDelete bool) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstancesAsync", ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs, forceDelete)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// DeleteInstancesAsync indicates an expected call of DeleteInstancesAsync.
func (mr *MockInterfaceMockRecorder) DeleteInstancesAsync(ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs, forceDelete any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstancesAsync", reflect.TypeOf((*MockInterface)(nil).DeleteInstancesAsync), ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs, forceDelete)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, VMScaleSetName string) (compute.VirtualMachineScaleSet, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, VMScaleSetName)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, VMScaleSetName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, VMScaleSetName)
}

// List mocks base method.
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]compute.VirtualMachineScaleSet, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
}

// StartInstancesAsync mocks base method.
func (m *MockInterface) StartInstancesAsync(ctx context.Context, resourceGroupName, vmScaleSetName string, vmInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstancesAsync", ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// StartInstancesAsync indicates an expected call of StartInstancesAsync.
func (mr *MockInterfaceMockRecorder) StartInstancesAsync(ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstancesAsync", reflect.TypeOf((*MockInterface)(nil).StartInstancesAsync), ctx, resourceGroupName, vmScaleSetName, vmInstanceIDs)
}

// WaitForAsyncOperationResult mocks base method.
func (m *MockInterface) WaitForAsyncOperationResult(ctx context.Context, future *azure.Future, resourceGroupName, request, asyncOpName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAsyncOperationResult", ctx, future, resourceGroupName, request, asyncOpName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForAsyncOperationResult indicates an expected call of WaitForAsyncOperationResult.
func (mr *MockInterfaceMockRecorder) WaitForAsyncOperationResult(ctx, future, resourceGroupName, request, asyncOpName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAsyncOperationResult", reflect.TypeOf((*MockInterface)(nil).WaitForAsyncOperationResult), ctx, future, resourceGroupName, request, asyncOpName)
}

// WaitForCreateOrUpdateResult mocks base method.
func (m *MockInterface) WaitForCreateOrUpdateResult(ctx context.Context, future *azure.Future, resourceGroupName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCreateOrUpdateResult", ctx, future, resourceGroupName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForCreateOrUpdateResult indicates an expected call of WaitForCreateOrUpdateResult.
func (mr *MockInterfaceMockRecorder) WaitForCreateOrUpdateResult(ctx, future, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCreateOrUpdateResult", reflect.TypeOf((*MockInterface)(nil).WaitForCreateOrUpdateResult), ctx, future, resourceGroupName)
}

// WaitForDeallocateInstancesResult mocks base method.
func (m *MockInterface) WaitForDeallocateInstancesResult(ctx context.Context, future *azure.Future, resourceGroupName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDeallocateInstancesResult", ctx, future, resourceGroupName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForDeallocateInstancesResult indicates an expected call of WaitForDeallocateInstancesResult.
func (mr *MockInterfaceMockRecorder) WaitForDeallocateInstancesResult(ctx, future, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDeallocateInstancesResult", reflect.TypeOf((*MockInterface)(nil).WaitForDeallocateInstancesResult), ctx, future, resourceGroupName)
}

// WaitForDeleteInstancesResult mocks base method.
func (m *MockInterface) WaitForDeleteInstancesResult(ctx context.Context, future *azure.Future, resourceGroupName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDeleteInstancesResult", ctx, future, resourceGroupName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForDeleteInstancesResult indicates an expected call of WaitForDeleteInstancesResult.
func (mr *MockInterfaceMockRecorder) WaitForDeleteInstancesResult(ctx, future, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDeleteInstancesResult", reflect.TypeOf((*MockInterface)(nil).WaitForDeleteInstancesResult), ctx, future, resourceGroupName)
}

// WaitForStartInstancesResult mocks base method.
func (m *MockInterface) WaitForStartInstancesResult(ctx context.Context, future *azure.Future, resourceGroupName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForStartInstancesResult", ctx, future, resourceGroupName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForStartInstancesResult indicates an expected call of WaitForStartInstancesResult.
func (mr *MockInterfaceMockRecorder) WaitForStartInstancesResult(ctx, future, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForStartInstancesResult", reflect.TypeOf((*MockInterface)(nil).WaitForStartInstancesResult), ctx, future, resourceGroupName)
}
