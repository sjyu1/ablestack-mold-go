//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/KubernetesService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKubernetesServiceIface is a mock of KubernetesServiceIface interface.
type MockKubernetesServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesServiceIfaceMockRecorder
}

// MockKubernetesServiceIfaceMockRecorder is the mock recorder for MockKubernetesServiceIface.
type MockKubernetesServiceIfaceMockRecorder struct {
	mock *MockKubernetesServiceIface
}

// NewMockKubernetesServiceIface creates a new mock instance.
func NewMockKubernetesServiceIface(ctrl *gomock.Controller) *MockKubernetesServiceIface {
	mock := &MockKubernetesServiceIface{ctrl: ctrl}
	mock.recorder = &MockKubernetesServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesServiceIface) EXPECT() *MockKubernetesServiceIfaceMockRecorder {
	return m.recorder
}

// AddKubernetesSupportedVersion mocks base method.
func (m *MockKubernetesServiceIface) AddKubernetesSupportedVersion(p *AddKubernetesSupportedVersionParams) (*AddKubernetesSupportedVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKubernetesSupportedVersion", p)
	ret0, _ := ret[0].(*AddKubernetesSupportedVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddKubernetesSupportedVersion indicates an expected call of AddKubernetesSupportedVersion.
func (mr *MockKubernetesServiceIfaceMockRecorder) AddKubernetesSupportedVersion(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKubernetesSupportedVersion", reflect.TypeOf((*MockKubernetesServiceIface)(nil).AddKubernetesSupportedVersion), p)
}

// AddVirtualMachinesToKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) AddVirtualMachinesToKubernetesCluster(p *AddVirtualMachinesToKubernetesClusterParams) (*AddVirtualMachinesToKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVirtualMachinesToKubernetesCluster", p)
	ret0, _ := ret[0].(*AddVirtualMachinesToKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddVirtualMachinesToKubernetesCluster indicates an expected call of AddVirtualMachinesToKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) AddVirtualMachinesToKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVirtualMachinesToKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).AddVirtualMachinesToKubernetesCluster), p)
}

// CreateKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) CreateKubernetesCluster(p *CreateKubernetesClusterParams) (*CreateKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKubernetesCluster", p)
	ret0, _ := ret[0].(*CreateKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKubernetesCluster indicates an expected call of CreateKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) CreateKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).CreateKubernetesCluster), p)
}

// DeleteKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) DeleteKubernetesCluster(p *DeleteKubernetesClusterParams) (*DeleteKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKubernetesCluster", p)
	ret0, _ := ret[0].(*DeleteKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKubernetesCluster indicates an expected call of DeleteKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) DeleteKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).DeleteKubernetesCluster), p)
}

// DeleteKubernetesSupportedVersion mocks base method.
func (m *MockKubernetesServiceIface) DeleteKubernetesSupportedVersion(p *DeleteKubernetesSupportedVersionParams) (*DeleteKubernetesSupportedVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKubernetesSupportedVersion", p)
	ret0, _ := ret[0].(*DeleteKubernetesSupportedVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKubernetesSupportedVersion indicates an expected call of DeleteKubernetesSupportedVersion.
func (mr *MockKubernetesServiceIfaceMockRecorder) DeleteKubernetesSupportedVersion(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKubernetesSupportedVersion", reflect.TypeOf((*MockKubernetesServiceIface)(nil).DeleteKubernetesSupportedVersion), p)
}

// GetKubernetesClusterByID mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesClusterByID(id string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesClusterByID", varargs...)
	ret0, _ := ret[0].(*KubernetesCluster)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesClusterByID indicates an expected call of GetKubernetesClusterByID.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesClusterByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesClusterByID", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesClusterByID), varargs...)
}

// GetKubernetesClusterByName mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesClusterByName(name string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesClusterByName", varargs...)
	ret0, _ := ret[0].(*KubernetesCluster)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesClusterByName indicates an expected call of GetKubernetesClusterByName.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesClusterByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesClusterByName", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesClusterByName), varargs...)
}

// GetKubernetesClusterConfig mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesClusterConfig(p *GetKubernetesClusterConfigParams) (*GetKubernetesClusterConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubernetesClusterConfig", p)
	ret0, _ := ret[0].(*GetKubernetesClusterConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubernetesClusterConfig indicates an expected call of GetKubernetesClusterConfig.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesClusterConfig(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesClusterConfig", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesClusterConfig), p)
}

// GetKubernetesClusterID mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesClusterID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesClusterID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesClusterID indicates an expected call of GetKubernetesClusterID.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesClusterID(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesClusterID", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesClusterID), varargs...)
}

// GetKubernetesSupportedVersionByID mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesSupportedVersionByID(id string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesSupportedVersionByID", varargs...)
	ret0, _ := ret[0].(*KubernetesSupportedVersion)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesSupportedVersionByID indicates an expected call of GetKubernetesSupportedVersionByID.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesSupportedVersionByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesSupportedVersionByID", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesSupportedVersionByID), varargs...)
}

// GetKubernetesSupportedVersionByName mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesSupportedVersionByName(name string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesSupportedVersionByName", varargs...)
	ret0, _ := ret[0].(*KubernetesSupportedVersion)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesSupportedVersionByName indicates an expected call of GetKubernetesSupportedVersionByName.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesSupportedVersionByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesSupportedVersionByName", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesSupportedVersionByName), varargs...)
}

// GetKubernetesSupportedVersionID mocks base method.
func (m *MockKubernetesServiceIface) GetKubernetesSupportedVersionID(keyword string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{keyword}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubernetesSupportedVersionID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKubernetesSupportedVersionID indicates an expected call of GetKubernetesSupportedVersionID.
func (mr *MockKubernetesServiceIfaceMockRecorder) GetKubernetesSupportedVersionID(keyword interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{keyword}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesSupportedVersionID", reflect.TypeOf((*MockKubernetesServiceIface)(nil).GetKubernetesSupportedVersionID), varargs...)
}

// ListKubernetesClusters mocks base method.
func (m *MockKubernetesServiceIface) ListKubernetesClusters(p *ListKubernetesClustersParams) (*ListKubernetesClustersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesClusters", p)
	ret0, _ := ret[0].(*ListKubernetesClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKubernetesClusters indicates an expected call of ListKubernetesClusters.
func (mr *MockKubernetesServiceIfaceMockRecorder) ListKubernetesClusters(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesClusters", reflect.TypeOf((*MockKubernetesServiceIface)(nil).ListKubernetesClusters), p)
}

// ListKubernetesSupportedVersions mocks base method.
func (m *MockKubernetesServiceIface) ListKubernetesSupportedVersions(p *ListKubernetesSupportedVersionsParams) (*ListKubernetesSupportedVersionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesSupportedVersions", p)
	ret0, _ := ret[0].(*ListKubernetesSupportedVersionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKubernetesSupportedVersions indicates an expected call of ListKubernetesSupportedVersions.
func (mr *MockKubernetesServiceIfaceMockRecorder) ListKubernetesSupportedVersions(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesSupportedVersions", reflect.TypeOf((*MockKubernetesServiceIface)(nil).ListKubernetesSupportedVersions), p)
}

// NewAddKubernetesSupportedVersionParams mocks base method.
func (m *MockKubernetesServiceIface) NewAddKubernetesSupportedVersionParams(mincpunumber, minmemory int, semanticversion string) *AddKubernetesSupportedVersionParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddKubernetesSupportedVersionParams", mincpunumber, minmemory, semanticversion)
	ret0, _ := ret[0].(*AddKubernetesSupportedVersionParams)
	return ret0
}

// NewAddKubernetesSupportedVersionParams indicates an expected call of NewAddKubernetesSupportedVersionParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewAddKubernetesSupportedVersionParams(mincpunumber, minmemory, semanticversion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddKubernetesSupportedVersionParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewAddKubernetesSupportedVersionParams), mincpunumber, minmemory, semanticversion)
}

// NewAddVirtualMachinesToKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewAddVirtualMachinesToKubernetesClusterParams(id string, virtualmachineids []string) *AddVirtualMachinesToKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddVirtualMachinesToKubernetesClusterParams", id, virtualmachineids)
	ret0, _ := ret[0].(*AddVirtualMachinesToKubernetesClusterParams)
	return ret0
}

// NewAddVirtualMachinesToKubernetesClusterParams indicates an expected call of NewAddVirtualMachinesToKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewAddVirtualMachinesToKubernetesClusterParams(id, virtualmachineids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddVirtualMachinesToKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewAddVirtualMachinesToKubernetesClusterParams), id, virtualmachineids)
}

// NewCreateKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewCreateKubernetesClusterParams(clustertype, name, zoneid string) *CreateKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateKubernetesClusterParams", clustertype, name, zoneid)
	ret0, _ := ret[0].(*CreateKubernetesClusterParams)
	return ret0
}

// NewCreateKubernetesClusterParams indicates an expected call of NewCreateKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewCreateKubernetesClusterParams(clustertype, name, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewCreateKubernetesClusterParams), clustertype, name, zoneid)
}

// NewDeleteKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewDeleteKubernetesClusterParams(id string) *DeleteKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteKubernetesClusterParams", id)
	ret0, _ := ret[0].(*DeleteKubernetesClusterParams)
	return ret0
}

// NewDeleteKubernetesClusterParams indicates an expected call of NewDeleteKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewDeleteKubernetesClusterParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewDeleteKubernetesClusterParams), id)
}

// NewDeleteKubernetesSupportedVersionParams mocks base method.
func (m *MockKubernetesServiceIface) NewDeleteKubernetesSupportedVersionParams(id string) *DeleteKubernetesSupportedVersionParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteKubernetesSupportedVersionParams", id)
	ret0, _ := ret[0].(*DeleteKubernetesSupportedVersionParams)
	return ret0
}

// NewDeleteKubernetesSupportedVersionParams indicates an expected call of NewDeleteKubernetesSupportedVersionParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewDeleteKubernetesSupportedVersionParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteKubernetesSupportedVersionParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewDeleteKubernetesSupportedVersionParams), id)
}

// NewGetKubernetesClusterConfigParams mocks base method.
func (m *MockKubernetesServiceIface) NewGetKubernetesClusterConfigParams() *GetKubernetesClusterConfigParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGetKubernetesClusterConfigParams")
	ret0, _ := ret[0].(*GetKubernetesClusterConfigParams)
	return ret0
}

// NewGetKubernetesClusterConfigParams indicates an expected call of NewGetKubernetesClusterConfigParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewGetKubernetesClusterConfigParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetKubernetesClusterConfigParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewGetKubernetesClusterConfigParams))
}

// NewListKubernetesClustersParams mocks base method.
func (m *MockKubernetesServiceIface) NewListKubernetesClustersParams() *ListKubernetesClustersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListKubernetesClustersParams")
	ret0, _ := ret[0].(*ListKubernetesClustersParams)
	return ret0
}

// NewListKubernetesClustersParams indicates an expected call of NewListKubernetesClustersParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewListKubernetesClustersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListKubernetesClustersParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewListKubernetesClustersParams))
}

// NewListKubernetesSupportedVersionsParams mocks base method.
func (m *MockKubernetesServiceIface) NewListKubernetesSupportedVersionsParams() *ListKubernetesSupportedVersionsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListKubernetesSupportedVersionsParams")
	ret0, _ := ret[0].(*ListKubernetesSupportedVersionsParams)
	return ret0
}

// NewListKubernetesSupportedVersionsParams indicates an expected call of NewListKubernetesSupportedVersionsParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewListKubernetesSupportedVersionsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListKubernetesSupportedVersionsParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewListKubernetesSupportedVersionsParams))
}

// NewRemoveVirtualMachinesFromKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewRemoveVirtualMachinesFromKubernetesClusterParams(id string, virtualmachineids []string) *RemoveVirtualMachinesFromKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoveVirtualMachinesFromKubernetesClusterParams", id, virtualmachineids)
	ret0, _ := ret[0].(*RemoveVirtualMachinesFromKubernetesClusterParams)
	return ret0
}

// NewRemoveVirtualMachinesFromKubernetesClusterParams indicates an expected call of NewRemoveVirtualMachinesFromKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewRemoveVirtualMachinesFromKubernetesClusterParams(id, virtualmachineids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoveVirtualMachinesFromKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewRemoveVirtualMachinesFromKubernetesClusterParams), id, virtualmachineids)
}

// NewScaleKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewScaleKubernetesClusterParams(id string) *ScaleKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewScaleKubernetesClusterParams", id)
	ret0, _ := ret[0].(*ScaleKubernetesClusterParams)
	return ret0
}

// NewScaleKubernetesClusterParams indicates an expected call of NewScaleKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewScaleKubernetesClusterParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewScaleKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewScaleKubernetesClusterParams), id)
}

// NewStartKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewStartKubernetesClusterParams(id string) *StartKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStartKubernetesClusterParams", id)
	ret0, _ := ret[0].(*StartKubernetesClusterParams)
	return ret0
}

// NewStartKubernetesClusterParams indicates an expected call of NewStartKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewStartKubernetesClusterParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStartKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewStartKubernetesClusterParams), id)
}

// NewStopKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewStopKubernetesClusterParams(id string) *StopKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStopKubernetesClusterParams", id)
	ret0, _ := ret[0].(*StopKubernetesClusterParams)
	return ret0
}

// NewStopKubernetesClusterParams indicates an expected call of NewStopKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewStopKubernetesClusterParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStopKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewStopKubernetesClusterParams), id)
}

// NewUpdateKubernetesSupportedVersionParams mocks base method.
func (m *MockKubernetesServiceIface) NewUpdateKubernetesSupportedVersionParams(id, state string) *UpdateKubernetesSupportedVersionParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateKubernetesSupportedVersionParams", id, state)
	ret0, _ := ret[0].(*UpdateKubernetesSupportedVersionParams)
	return ret0
}

// NewUpdateKubernetesSupportedVersionParams indicates an expected call of NewUpdateKubernetesSupportedVersionParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewUpdateKubernetesSupportedVersionParams(id, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateKubernetesSupportedVersionParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewUpdateKubernetesSupportedVersionParams), id, state)
}

// NewUpgradeKubernetesClusterParams mocks base method.
func (m *MockKubernetesServiceIface) NewUpgradeKubernetesClusterParams(id, kubernetesversionid string) *UpgradeKubernetesClusterParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpgradeKubernetesClusterParams", id, kubernetesversionid)
	ret0, _ := ret[0].(*UpgradeKubernetesClusterParams)
	return ret0
}

// NewUpgradeKubernetesClusterParams indicates an expected call of NewUpgradeKubernetesClusterParams.
func (mr *MockKubernetesServiceIfaceMockRecorder) NewUpgradeKubernetesClusterParams(id, kubernetesversionid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpgradeKubernetesClusterParams", reflect.TypeOf((*MockKubernetesServiceIface)(nil).NewUpgradeKubernetesClusterParams), id, kubernetesversionid)
}

// RemoveVirtualMachinesFromKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) RemoveVirtualMachinesFromKubernetesCluster(p *RemoveVirtualMachinesFromKubernetesClusterParams) (*RemoveVirtualMachinesFromKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVirtualMachinesFromKubernetesCluster", p)
	ret0, _ := ret[0].(*RemoveVirtualMachinesFromKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVirtualMachinesFromKubernetesCluster indicates an expected call of RemoveVirtualMachinesFromKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) RemoveVirtualMachinesFromKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVirtualMachinesFromKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).RemoveVirtualMachinesFromKubernetesCluster), p)
}

// ScaleKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) ScaleKubernetesCluster(p *ScaleKubernetesClusterParams) (*ScaleKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleKubernetesCluster", p)
	ret0, _ := ret[0].(*ScaleKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleKubernetesCluster indicates an expected call of ScaleKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) ScaleKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).ScaleKubernetesCluster), p)
}

// StartKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) StartKubernetesCluster(p *StartKubernetesClusterParams) (*StartKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartKubernetesCluster", p)
	ret0, _ := ret[0].(*StartKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartKubernetesCluster indicates an expected call of StartKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) StartKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).StartKubernetesCluster), p)
}

// StopKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) StopKubernetesCluster(p *StopKubernetesClusterParams) (*StopKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopKubernetesCluster", p)
	ret0, _ := ret[0].(*StopKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopKubernetesCluster indicates an expected call of StopKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) StopKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).StopKubernetesCluster), p)
}

// UpdateKubernetesSupportedVersion mocks base method.
func (m *MockKubernetesServiceIface) UpdateKubernetesSupportedVersion(p *UpdateKubernetesSupportedVersionParams) (*UpdateKubernetesSupportedVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKubernetesSupportedVersion", p)
	ret0, _ := ret[0].(*UpdateKubernetesSupportedVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateKubernetesSupportedVersion indicates an expected call of UpdateKubernetesSupportedVersion.
func (mr *MockKubernetesServiceIfaceMockRecorder) UpdateKubernetesSupportedVersion(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKubernetesSupportedVersion", reflect.TypeOf((*MockKubernetesServiceIface)(nil).UpdateKubernetesSupportedVersion), p)
}

// UpgradeKubernetesCluster mocks base method.
func (m *MockKubernetesServiceIface) UpgradeKubernetesCluster(p *UpgradeKubernetesClusterParams) (*UpgradeKubernetesClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeKubernetesCluster", p)
	ret0, _ := ret[0].(*UpgradeKubernetesClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeKubernetesCluster indicates an expected call of UpgradeKubernetesCluster.
func (mr *MockKubernetesServiceIfaceMockRecorder) UpgradeKubernetesCluster(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeKubernetesCluster", reflect.TypeOf((*MockKubernetesServiceIface)(nil).UpgradeKubernetesCluster), p)
}
