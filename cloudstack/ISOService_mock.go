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
// Source: ./cloudstack/ISOService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISOServiceIface is a mock of ISOServiceIface interface.
type MockISOServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockISOServiceIfaceMockRecorder
}

// MockISOServiceIfaceMockRecorder is the mock recorder for MockISOServiceIface.
type MockISOServiceIfaceMockRecorder struct {
	mock *MockISOServiceIface
}

// NewMockISOServiceIface creates a new mock instance.
func NewMockISOServiceIface(ctrl *gomock.Controller) *MockISOServiceIface {
	mock := &MockISOServiceIface{ctrl: ctrl}
	mock.recorder = &MockISOServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISOServiceIface) EXPECT() *MockISOServiceIfaceMockRecorder {
	return m.recorder
}

// AttachIso mocks base method.
func (m *MockISOServiceIface) AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachIso", p)
	ret0, _ := ret[0].(*AttachIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachIso indicates an expected call of AttachIso.
func (mr *MockISOServiceIfaceMockRecorder) AttachIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachIso", reflect.TypeOf((*MockISOServiceIface)(nil).AttachIso), p)
}

// CopyIso mocks base method.
func (m *MockISOServiceIface) CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyIso", p)
	ret0, _ := ret[0].(*CopyIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyIso indicates an expected call of CopyIso.
func (mr *MockISOServiceIfaceMockRecorder) CopyIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyIso", reflect.TypeOf((*MockISOServiceIface)(nil).CopyIso), p)
}

// DeleteIso mocks base method.
func (m *MockISOServiceIface) DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIso", p)
	ret0, _ := ret[0].(*DeleteIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIso indicates an expected call of DeleteIso.
func (mr *MockISOServiceIfaceMockRecorder) DeleteIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIso", reflect.TypeOf((*MockISOServiceIface)(nil).DeleteIso), p)
}

// DetachIso mocks base method.
func (m *MockISOServiceIface) DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachIso", p)
	ret0, _ := ret[0].(*DetachIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachIso indicates an expected call of DetachIso.
func (mr *MockISOServiceIfaceMockRecorder) DetachIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachIso", reflect.TypeOf((*MockISOServiceIface)(nil).DetachIso), p)
}

// ExtractIso mocks base method.
func (m *MockISOServiceIface) ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractIso", p)
	ret0, _ := ret[0].(*ExtractIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractIso indicates an expected call of ExtractIso.
func (mr *MockISOServiceIfaceMockRecorder) ExtractIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractIso", reflect.TypeOf((*MockISOServiceIface)(nil).ExtractIso), p)
}

// GetIsoByID mocks base method.
func (m *MockISOServiceIface) GetIsoByID(id string, opts ...OptionFunc) (*Iso, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIsoByID", varargs...)
	ret0, _ := ret[0].(*Iso)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIsoByID indicates an expected call of GetIsoByID.
func (mr *MockISOServiceIfaceMockRecorder) GetIsoByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsoByID", reflect.TypeOf((*MockISOServiceIface)(nil).GetIsoByID), varargs...)
}

// GetIsoByName mocks base method.
func (m *MockISOServiceIface) GetIsoByName(name, isofilter, zoneid string, opts ...OptionFunc) (*Iso, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, isofilter, zoneid}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIsoByName", varargs...)
	ret0, _ := ret[0].(*Iso)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIsoByName indicates an expected call of GetIsoByName.
func (mr *MockISOServiceIfaceMockRecorder) GetIsoByName(name, isofilter, zoneid interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, isofilter, zoneid}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsoByName", reflect.TypeOf((*MockISOServiceIface)(nil).GetIsoByName), varargs...)
}

// GetIsoID mocks base method.
func (m *MockISOServiceIface) GetIsoID(name, isofilter, zoneid string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, isofilter, zoneid}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIsoID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIsoID indicates an expected call of GetIsoID.
func (mr *MockISOServiceIfaceMockRecorder) GetIsoID(name, isofilter, zoneid interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, isofilter, zoneid}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsoID", reflect.TypeOf((*MockISOServiceIface)(nil).GetIsoID), varargs...)
}

// GetIsoPermissionByID mocks base method.
func (m *MockISOServiceIface) GetIsoPermissionByID(id string, opts ...OptionFunc) (*IsoPermission, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIsoPermissionByID", varargs...)
	ret0, _ := ret[0].(*IsoPermission)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIsoPermissionByID indicates an expected call of GetIsoPermissionByID.
func (mr *MockISOServiceIfaceMockRecorder) GetIsoPermissionByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsoPermissionByID", reflect.TypeOf((*MockISOServiceIface)(nil).GetIsoPermissionByID), varargs...)
}

// ListIsoPermissions mocks base method.
func (m *MockISOServiceIface) ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIsoPermissions", p)
	ret0, _ := ret[0].(*ListIsoPermissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIsoPermissions indicates an expected call of ListIsoPermissions.
func (mr *MockISOServiceIfaceMockRecorder) ListIsoPermissions(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIsoPermissions", reflect.TypeOf((*MockISOServiceIface)(nil).ListIsoPermissions), p)
}

// ListIsos mocks base method.
func (m *MockISOServiceIface) ListIsos(p *ListIsosParams) (*ListIsosResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIsos", p)
	ret0, _ := ret[0].(*ListIsosResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIsos indicates an expected call of ListIsos.
func (mr *MockISOServiceIfaceMockRecorder) ListIsos(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIsos", reflect.TypeOf((*MockISOServiceIface)(nil).ListIsos), p)
}

// NewAttachIsoParams mocks base method.
func (m *MockISOServiceIface) NewAttachIsoParams(id, virtualmachineid string) *AttachIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAttachIsoParams", id, virtualmachineid)
	ret0, _ := ret[0].(*AttachIsoParams)
	return ret0
}

// NewAttachIsoParams indicates an expected call of NewAttachIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewAttachIsoParams(id, virtualmachineid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAttachIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewAttachIsoParams), id, virtualmachineid)
}

// NewCopyIsoParams mocks base method.
func (m *MockISOServiceIface) NewCopyIsoParams(id string) *CopyIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCopyIsoParams", id)
	ret0, _ := ret[0].(*CopyIsoParams)
	return ret0
}

// NewCopyIsoParams indicates an expected call of NewCopyIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewCopyIsoParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCopyIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewCopyIsoParams), id)
}

// NewDeleteIsoParams mocks base method.
func (m *MockISOServiceIface) NewDeleteIsoParams(id string) *DeleteIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteIsoParams", id)
	ret0, _ := ret[0].(*DeleteIsoParams)
	return ret0
}

// NewDeleteIsoParams indicates an expected call of NewDeleteIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewDeleteIsoParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewDeleteIsoParams), id)
}

// NewDetachIsoParams mocks base method.
func (m *MockISOServiceIface) NewDetachIsoParams(virtualmachineid string) *DetachIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDetachIsoParams", virtualmachineid)
	ret0, _ := ret[0].(*DetachIsoParams)
	return ret0
}

// NewDetachIsoParams indicates an expected call of NewDetachIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewDetachIsoParams(virtualmachineid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDetachIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewDetachIsoParams), virtualmachineid)
}

// NewExtractIsoParams mocks base method.
func (m *MockISOServiceIface) NewExtractIsoParams(id, mode string) *ExtractIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewExtractIsoParams", id, mode)
	ret0, _ := ret[0].(*ExtractIsoParams)
	return ret0
}

// NewExtractIsoParams indicates an expected call of NewExtractIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewExtractIsoParams(id, mode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewExtractIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewExtractIsoParams), id, mode)
}

// NewListIsoPermissionsParams mocks base method.
func (m *MockISOServiceIface) NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListIsoPermissionsParams", id)
	ret0, _ := ret[0].(*ListIsoPermissionsParams)
	return ret0
}

// NewListIsoPermissionsParams indicates an expected call of NewListIsoPermissionsParams.
func (mr *MockISOServiceIfaceMockRecorder) NewListIsoPermissionsParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListIsoPermissionsParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewListIsoPermissionsParams), id)
}

// NewListIsosParams mocks base method.
func (m *MockISOServiceIface) NewListIsosParams() *ListIsosParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListIsosParams")
	ret0, _ := ret[0].(*ListIsosParams)
	return ret0
}

// NewListIsosParams indicates an expected call of NewListIsosParams.
func (mr *MockISOServiceIfaceMockRecorder) NewListIsosParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListIsosParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewListIsosParams))
}

// NewRegisterIsoParams mocks base method.
func (m *MockISOServiceIface) NewRegisterIsoParams(name, url, zoneid string) *RegisterIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegisterIsoParams", name, url, zoneid)
	ret0, _ := ret[0].(*RegisterIsoParams)
	return ret0
}

// NewRegisterIsoParams indicates an expected call of NewRegisterIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewRegisterIsoParams(name, url, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegisterIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewRegisterIsoParams), name, url, zoneid)
}

// NewUpdateIsoParams mocks base method.
func (m *MockISOServiceIface) NewUpdateIsoParams(id string) *UpdateIsoParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateIsoParams", id)
	ret0, _ := ret[0].(*UpdateIsoParams)
	return ret0
}

// NewUpdateIsoParams indicates an expected call of NewUpdateIsoParams.
func (mr *MockISOServiceIfaceMockRecorder) NewUpdateIsoParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateIsoParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewUpdateIsoParams), id)
}

// NewUpdateIsoPermissionsParams mocks base method.
func (m *MockISOServiceIface) NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateIsoPermissionsParams", id)
	ret0, _ := ret[0].(*UpdateIsoPermissionsParams)
	return ret0
}

// NewUpdateIsoPermissionsParams indicates an expected call of NewUpdateIsoPermissionsParams.
func (mr *MockISOServiceIfaceMockRecorder) NewUpdateIsoPermissionsParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateIsoPermissionsParams", reflect.TypeOf((*MockISOServiceIface)(nil).NewUpdateIsoPermissionsParams), id)
}

// RegisterIso mocks base method.
func (m *MockISOServiceIface) RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterIso", p)
	ret0, _ := ret[0].(*RegisterIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterIso indicates an expected call of RegisterIso.
func (mr *MockISOServiceIfaceMockRecorder) RegisterIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterIso", reflect.TypeOf((*MockISOServiceIface)(nil).RegisterIso), p)
}

// UpdateIso mocks base method.
func (m *MockISOServiceIface) UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIso", p)
	ret0, _ := ret[0].(*UpdateIsoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIso indicates an expected call of UpdateIso.
func (mr *MockISOServiceIfaceMockRecorder) UpdateIso(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIso", reflect.TypeOf((*MockISOServiceIface)(nil).UpdateIso), p)
}

// UpdateIsoPermissions mocks base method.
func (m *MockISOServiceIface) UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIsoPermissions", p)
	ret0, _ := ret[0].(*UpdateIsoPermissionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIsoPermissions indicates an expected call of UpdateIsoPermissions.
func (mr *MockISOServiceIfaceMockRecorder) UpdateIsoPermissions(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIsoPermissions", reflect.TypeOf((*MockISOServiceIface)(nil).UpdateIsoPermissions), p)
}
