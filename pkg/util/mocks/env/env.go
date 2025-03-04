// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/env (interfaces: Core,Interface)

// Package mock_env is a generated GoMock package.
package mock_env

import (
	context "context"
	rsa "crypto/rsa"
	x509 "crypto/x509"
	net "net"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	azure "github.com/Azure/go-autorest/autorest/azure"
	gomock "github.com/golang/mock/gomock"

	clientauthorizer "github.com/Azure/ARO-RP/pkg/util/clientauthorizer"
	deployment "github.com/Azure/ARO-RP/pkg/util/deployment"
	keyvault "github.com/Azure/ARO-RP/pkg/util/keyvault"
	refreshable "github.com/Azure/ARO-RP/pkg/util/refreshable"
)

// MockCore is a mock of Core interface
type MockCore struct {
	ctrl     *gomock.Controller
	recorder *MockCoreMockRecorder
}

// MockCoreMockRecorder is the mock recorder for MockCore
type MockCoreMockRecorder struct {
	mock *MockCore
}

// NewMockCore creates a new mock instance
func NewMockCore(ctrl *gomock.Controller) *MockCore {
	mock := &MockCore{ctrl: ctrl}
	mock.recorder = &MockCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCore) EXPECT() *MockCoreMockRecorder {
	return m.recorder
}

// DeploymentMode mocks base method
func (m *MockCore) DeploymentMode() deployment.Mode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentMode")
	ret0, _ := ret[0].(deployment.Mode)
	return ret0
}

// DeploymentMode indicates an expected call of DeploymentMode
func (mr *MockCoreMockRecorder) DeploymentMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentMode", reflect.TypeOf((*MockCore)(nil).DeploymentMode))
}

// Environment mocks base method
func (m *MockCore) Environment() *azure.Environment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environment")
	ret0, _ := ret[0].(*azure.Environment)
	return ret0
}

// Environment indicates an expected call of Environment
func (mr *MockCoreMockRecorder) Environment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environment", reflect.TypeOf((*MockCore)(nil).Environment))
}

// Hostname mocks base method
func (m *MockCore) Hostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hostname indicates an expected call of Hostname
func (mr *MockCoreMockRecorder) Hostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hostname", reflect.TypeOf((*MockCore)(nil).Hostname))
}

// Location mocks base method
func (m *MockCore) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location
func (mr *MockCoreMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockCore)(nil).Location))
}

// NewRPAuthorizer mocks base method
func (m *MockCore) NewRPAuthorizer(arg0 string) (autorest.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRPAuthorizer", arg0)
	ret0, _ := ret[0].(autorest.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRPAuthorizer indicates an expected call of NewRPAuthorizer
func (mr *MockCoreMockRecorder) NewRPAuthorizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRPAuthorizer", reflect.TypeOf((*MockCore)(nil).NewRPAuthorizer), arg0)
}

// ResourceGroup mocks base method
func (m *MockCore) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup
func (mr *MockCoreMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockCore)(nil).ResourceGroup))
}

// SubscriptionID mocks base method
func (m *MockCore) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID
func (mr *MockCoreMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockCore)(nil).SubscriptionID))
}

// TenantID mocks base method
func (m *MockCore) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID
func (mr *MockCoreMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockCore)(nil).TenantID))
}

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ACRDomain mocks base method
func (m *MockInterface) ACRDomain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRDomain")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRDomain indicates an expected call of ACRDomain
func (mr *MockInterfaceMockRecorder) ACRDomain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRDomain", reflect.TypeOf((*MockInterface)(nil).ACRDomain))
}

// ACRResourceID mocks base method
func (m *MockInterface) ACRResourceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRResourceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRResourceID indicates an expected call of ACRResourceID
func (mr *MockInterfaceMockRecorder) ACRResourceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRResourceID", reflect.TypeOf((*MockInterface)(nil).ACRResourceID))
}

// AROOperatorImage mocks base method
func (m *MockInterface) AROOperatorImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AROOperatorImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// AROOperatorImage indicates an expected call of AROOperatorImage
func (mr *MockInterfaceMockRecorder) AROOperatorImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AROOperatorImage", reflect.TypeOf((*MockInterface)(nil).AROOperatorImage))
}

// AdminClientAuthorizer mocks base method
func (m *MockInterface) AdminClientAuthorizer() clientauthorizer.ClientAuthorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminClientAuthorizer")
	ret0, _ := ret[0].(clientauthorizer.ClientAuthorizer)
	return ret0
}

// AdminClientAuthorizer indicates an expected call of AdminClientAuthorizer
func (mr *MockInterfaceMockRecorder) AdminClientAuthorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminClientAuthorizer", reflect.TypeOf((*MockInterface)(nil).AdminClientAuthorizer))
}

// ArmClientAuthorizer mocks base method
func (m *MockInterface) ArmClientAuthorizer() clientauthorizer.ClientAuthorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArmClientAuthorizer")
	ret0, _ := ret[0].(clientauthorizer.ClientAuthorizer)
	return ret0
}

// ArmClientAuthorizer indicates an expected call of ArmClientAuthorizer
func (mr *MockInterfaceMockRecorder) ArmClientAuthorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArmClientAuthorizer", reflect.TypeOf((*MockInterface)(nil).ArmClientAuthorizer))
}

// ClusterGenevaLoggingConfigVersion mocks base method
func (m *MockInterface) ClusterGenevaLoggingConfigVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGenevaLoggingConfigVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterGenevaLoggingConfigVersion indicates an expected call of ClusterGenevaLoggingConfigVersion
func (mr *MockInterfaceMockRecorder) ClusterGenevaLoggingConfigVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGenevaLoggingConfigVersion", reflect.TypeOf((*MockInterface)(nil).ClusterGenevaLoggingConfigVersion))
}

// ClusterGenevaLoggingEnvironment mocks base method
func (m *MockInterface) ClusterGenevaLoggingEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGenevaLoggingEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterGenevaLoggingEnvironment indicates an expected call of ClusterGenevaLoggingEnvironment
func (mr *MockInterfaceMockRecorder) ClusterGenevaLoggingEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGenevaLoggingEnvironment", reflect.TypeOf((*MockInterface)(nil).ClusterGenevaLoggingEnvironment))
}

// ClusterGenevaLoggingSecret mocks base method
func (m *MockInterface) ClusterGenevaLoggingSecret() (*rsa.PrivateKey, *x509.Certificate) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGenevaLoggingSecret")
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].(*x509.Certificate)
	return ret0, ret1
}

// ClusterGenevaLoggingSecret indicates an expected call of ClusterGenevaLoggingSecret
func (mr *MockInterfaceMockRecorder) ClusterGenevaLoggingSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGenevaLoggingSecret", reflect.TypeOf((*MockInterface)(nil).ClusterGenevaLoggingSecret))
}

// ClusterKeyvault mocks base method
func (m *MockInterface) ClusterKeyvault() keyvault.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterKeyvault")
	ret0, _ := ret[0].(keyvault.Manager)
	return ret0
}

// ClusterKeyvault indicates an expected call of ClusterKeyvault
func (mr *MockInterfaceMockRecorder) ClusterKeyvault() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterKeyvault", reflect.TypeOf((*MockInterface)(nil).ClusterKeyvault))
}

// DeploymentMode mocks base method
func (m *MockInterface) DeploymentMode() deployment.Mode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentMode")
	ret0, _ := ret[0].(deployment.Mode)
	return ret0
}

// DeploymentMode indicates an expected call of DeploymentMode
func (mr *MockInterfaceMockRecorder) DeploymentMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentMode", reflect.TypeOf((*MockInterface)(nil).DeploymentMode))
}

// DialContext mocks base method
func (m *MockInterface) DialContext(arg0 context.Context, arg1, arg2 string) (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialContext indicates an expected call of DialContext
func (mr *MockInterfaceMockRecorder) DialContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialContext", reflect.TypeOf((*MockInterface)(nil).DialContext), arg0, arg1, arg2)
}

// Domain mocks base method
func (m *MockInterface) Domain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Domain")
	ret0, _ := ret[0].(string)
	return ret0
}

// Domain indicates an expected call of Domain
func (mr *MockInterfaceMockRecorder) Domain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Domain", reflect.TypeOf((*MockInterface)(nil).Domain))
}

// EnsureARMResourceGroupRoleAssignment mocks base method
func (m *MockInterface) EnsureARMResourceGroupRoleAssignment(arg0 context.Context, arg1 refreshable.Authorizer, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureARMResourceGroupRoleAssignment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureARMResourceGroupRoleAssignment indicates an expected call of EnsureARMResourceGroupRoleAssignment
func (mr *MockInterfaceMockRecorder) EnsureARMResourceGroupRoleAssignment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureARMResourceGroupRoleAssignment", reflect.TypeOf((*MockInterface)(nil).EnsureARMResourceGroupRoleAssignment), arg0, arg1, arg2)
}

// Environment mocks base method
func (m *MockInterface) Environment() *azure.Environment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environment")
	ret0, _ := ret[0].(*azure.Environment)
	return ret0
}

// Environment indicates an expected call of Environment
func (mr *MockInterfaceMockRecorder) Environment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environment", reflect.TypeOf((*MockInterface)(nil).Environment))
}

// FPAuthorizer mocks base method
func (m *MockInterface) FPAuthorizer(arg0, arg1 string) (refreshable.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FPAuthorizer", arg0, arg1)
	ret0, _ := ret[0].(refreshable.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FPAuthorizer indicates an expected call of FPAuthorizer
func (mr *MockInterfaceMockRecorder) FPAuthorizer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FPAuthorizer", reflect.TypeOf((*MockInterface)(nil).FPAuthorizer), arg0, arg1)
}

// Hostname mocks base method
func (m *MockInterface) Hostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hostname indicates an expected call of Hostname
func (mr *MockInterfaceMockRecorder) Hostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hostname", reflect.TypeOf((*MockInterface)(nil).Hostname))
}

// InitializeAuthorizers mocks base method
func (m *MockInterface) InitializeAuthorizers() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeAuthorizers")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeAuthorizers indicates an expected call of InitializeAuthorizers
func (mr *MockInterfaceMockRecorder) InitializeAuthorizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeAuthorizers", reflect.TypeOf((*MockInterface)(nil).InitializeAuthorizers))
}

// Listen mocks base method
func (m *MockInterface) Listen() (net.Listener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen")
	ret0, _ := ret[0].(net.Listener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listen indicates an expected call of Listen
func (mr *MockInterfaceMockRecorder) Listen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockInterface)(nil).Listen))
}

// Location mocks base method
func (m *MockInterface) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location
func (mr *MockInterfaceMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockInterface)(nil).Location))
}

// NewRPAuthorizer mocks base method
func (m *MockInterface) NewRPAuthorizer(arg0 string) (autorest.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRPAuthorizer", arg0)
	ret0, _ := ret[0].(autorest.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRPAuthorizer indicates an expected call of NewRPAuthorizer
func (mr *MockInterfaceMockRecorder) NewRPAuthorizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRPAuthorizer", reflect.TypeOf((*MockInterface)(nil).NewRPAuthorizer), arg0)
}

// ResourceGroup mocks base method
func (m *MockInterface) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup
func (mr *MockInterfaceMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockInterface)(nil).ResourceGroup))
}

// ServiceKeyvault mocks base method
func (m *MockInterface) ServiceKeyvault() keyvault.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceKeyvault")
	ret0, _ := ret[0].(keyvault.Manager)
	return ret0
}

// ServiceKeyvault indicates an expected call of ServiceKeyvault
func (mr *MockInterfaceMockRecorder) ServiceKeyvault() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceKeyvault", reflect.TypeOf((*MockInterface)(nil).ServiceKeyvault))
}

// SubscriptionID mocks base method
func (m *MockInterface) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID
func (mr *MockInterfaceMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockInterface)(nil).SubscriptionID))
}

// TenantID mocks base method
func (m *MockInterface) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID
func (mr *MockInterfaceMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockInterface)(nil).TenantID))
}

// Zones mocks base method
func (m *MockInterface) Zones(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Zones", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Zones indicates an expected call of Zones
func (mr *MockInterfaceMockRecorder) Zones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Zones", reflect.TypeOf((*MockInterface)(nil).Zones), arg0)
}
