// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/venus/venus-shared/api/gateway/v1 (interfaces: IGateway)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	network "github.com/filecoin-project/go-state-types/network"
	proof "github.com/filecoin-project/go-state-types/proof"
	storage "github.com/filecoin-project/specs-storage/storage"
	types "github.com/filecoin-project/venus/venus-shared/types"
	gateway "github.com/filecoin-project/venus/venus-shared/types/gateway"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
)

// MockIGateway is a mock of IGateway interface.
type MockIGateway struct {
	ctrl     *gomock.Controller
	recorder *MockIGatewayMockRecorder
}

// MockIGatewayMockRecorder is the mock recorder for MockIGateway.
type MockIGatewayMockRecorder struct {
	mock *MockIGateway
}

// NewMockIGateway creates a new mock instance.
func NewMockIGateway(ctrl *gomock.Controller) *MockIGateway {
	mock := &MockIGateway{ctrl: ctrl}
	mock.recorder = &MockIGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGateway) EXPECT() *MockIGatewayMockRecorder {
	return m.recorder
}

// AddNewAddress mocks base method.
func (m *MockIGateway) AddNewAddress(arg0 context.Context, arg1 types.UUID, arg2 []address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewAddress indicates an expected call of AddNewAddress.
func (mr *MockIGatewayMockRecorder) AddNewAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewAddress", reflect.TypeOf((*MockIGateway)(nil).AddNewAddress), arg0, arg1, arg2)
}

// ComputeProof mocks base method.
func (m *MockIGateway) ComputeProof(arg0 context.Context, arg1 address.Address, arg2 []proof.ExtendedSectorInfo, arg3 abi.PoStRandomness, arg4 abi.ChainEpoch, arg5 network.Version) ([]proof.PoStProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeProof", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]proof.PoStProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeProof indicates an expected call of ComputeProof.
func (mr *MockIGatewayMockRecorder) ComputeProof(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeProof", reflect.TypeOf((*MockIGateway)(nil).ComputeProof), arg0, arg1, arg2, arg3, arg4, arg5)
}

// IsUnsealed mocks base method.
func (m *MockIGateway) IsUnsealed(arg0 context.Context, arg1 address.Address, arg2 cid.Cid, arg3 storage.SectorRef, arg4 types.PaddedByteIndex, arg5 abi.PaddedPieceSize) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUnsealed", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUnsealed indicates an expected call of IsUnsealed.
func (mr *MockIGatewayMockRecorder) IsUnsealed(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnsealed", reflect.TypeOf((*MockIGateway)(nil).IsUnsealed), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListConnectedMiners mocks base method.
func (m *MockIGateway) ListConnectedMiners(arg0 context.Context) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnectedMiners", arg0)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectedMiners indicates an expected call of ListConnectedMiners.
func (mr *MockIGatewayMockRecorder) ListConnectedMiners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectedMiners", reflect.TypeOf((*MockIGateway)(nil).ListConnectedMiners), arg0)
}

// ListMarketConnectionsState mocks base method.
func (m *MockIGateway) ListMarketConnectionsState(arg0 context.Context) ([]gateway.MarketConnectionState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMarketConnectionsState", arg0)
	ret0, _ := ret[0].([]gateway.MarketConnectionState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMarketConnectionsState indicates an expected call of ListMarketConnectionsState.
func (mr *MockIGatewayMockRecorder) ListMarketConnectionsState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMarketConnectionsState", reflect.TypeOf((*MockIGateway)(nil).ListMarketConnectionsState), arg0)
}

// ListMinerConnection mocks base method.
func (m *MockIGateway) ListMinerConnection(arg0 context.Context, arg1 address.Address) (*gateway.MinerState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMinerConnection", arg0, arg1)
	ret0, _ := ret[0].(*gateway.MinerState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMinerConnection indicates an expected call of ListMinerConnection.
func (mr *MockIGatewayMockRecorder) ListMinerConnection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMinerConnection", reflect.TypeOf((*MockIGateway)(nil).ListMinerConnection), arg0, arg1)
}

// ListWalletInfo mocks base method.
func (m *MockIGateway) ListWalletInfo(arg0 context.Context) ([]*gateway.WalletDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWalletInfo", arg0)
	ret0, _ := ret[0].([]*gateway.WalletDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWalletInfo indicates an expected call of ListWalletInfo.
func (mr *MockIGatewayMockRecorder) ListWalletInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWalletInfo", reflect.TypeOf((*MockIGateway)(nil).ListWalletInfo), arg0)
}

// ListWalletInfoByWallet mocks base method.
func (m *MockIGateway) ListWalletInfoByWallet(arg0 context.Context, arg1 string) (*gateway.WalletDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWalletInfoByWallet", arg0, arg1)
	ret0, _ := ret[0].(*gateway.WalletDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWalletInfoByWallet indicates an expected call of ListWalletInfoByWallet.
func (mr *MockIGatewayMockRecorder) ListWalletInfoByWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWalletInfoByWallet", reflect.TypeOf((*MockIGateway)(nil).ListWalletInfoByWallet), arg0, arg1)
}

// ListenMarketEvent mocks base method.
func (m *MockIGateway) ListenMarketEvent(arg0 context.Context, arg1 *gateway.MarketRegisterPolicy) (<-chan *gateway.RequestEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketEvent", arg0, arg1)
	ret0, _ := ret[0].(<-chan *gateway.RequestEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketEvent indicates an expected call of ListenMarketEvent.
func (mr *MockIGatewayMockRecorder) ListenMarketEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketEvent", reflect.TypeOf((*MockIGateway)(nil).ListenMarketEvent), arg0, arg1)
}

// ListenProofEvent mocks base method.
func (m *MockIGateway) ListenProofEvent(arg0 context.Context, arg1 *gateway.ProofRegisterPolicy) (<-chan *gateway.RequestEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenProofEvent", arg0, arg1)
	ret0, _ := ret[0].(<-chan *gateway.RequestEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenProofEvent indicates an expected call of ListenProofEvent.
func (mr *MockIGatewayMockRecorder) ListenProofEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenProofEvent", reflect.TypeOf((*MockIGateway)(nil).ListenProofEvent), arg0, arg1)
}

// ListenWalletEvent mocks base method.
func (m *MockIGateway) ListenWalletEvent(arg0 context.Context, arg1 *gateway.WalletRegisterPolicy) (<-chan *gateway.RequestEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenWalletEvent", arg0, arg1)
	ret0, _ := ret[0].(<-chan *gateway.RequestEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenWalletEvent indicates an expected call of ListenWalletEvent.
func (mr *MockIGatewayMockRecorder) ListenWalletEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenWalletEvent", reflect.TypeOf((*MockIGateway)(nil).ListenWalletEvent), arg0, arg1)
}

// RemoveAddress mocks base method.
func (m *MockIGateway) RemoveAddress(arg0 context.Context, arg1 types.UUID, arg2 []address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAddress indicates an expected call of RemoveAddress.
func (mr *MockIGatewayMockRecorder) RemoveAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAddress", reflect.TypeOf((*MockIGateway)(nil).RemoveAddress), arg0, arg1, arg2)
}

// ResponseMarketEvent mocks base method.
func (m *MockIGateway) ResponseMarketEvent(arg0 context.Context, arg1 *gateway.ResponseEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseMarketEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseMarketEvent indicates an expected call of ResponseMarketEvent.
func (mr *MockIGatewayMockRecorder) ResponseMarketEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseMarketEvent", reflect.TypeOf((*MockIGateway)(nil).ResponseMarketEvent), arg0, arg1)
}

// ResponseProofEvent mocks base method.
func (m *MockIGateway) ResponseProofEvent(arg0 context.Context, arg1 *gateway.ResponseEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseProofEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseProofEvent indicates an expected call of ResponseProofEvent.
func (mr *MockIGatewayMockRecorder) ResponseProofEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseProofEvent", reflect.TypeOf((*MockIGateway)(nil).ResponseProofEvent), arg0, arg1)
}

// ResponseWalletEvent mocks base method.
func (m *MockIGateway) ResponseWalletEvent(arg0 context.Context, arg1 *gateway.ResponseEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseWalletEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseWalletEvent indicates an expected call of ResponseWalletEvent.
func (mr *MockIGatewayMockRecorder) ResponseWalletEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseWalletEvent", reflect.TypeOf((*MockIGateway)(nil).ResponseWalletEvent), arg0, arg1)
}

// SectorsUnsealPiece mocks base method.
func (m *MockIGateway) SectorsUnsealPiece(arg0 context.Context, arg1 address.Address, arg2 cid.Cid, arg3 storage.SectorRef, arg4 types.PaddedByteIndex, arg5 abi.PaddedPieceSize, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SectorsUnsealPiece", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// SectorsUnsealPiece indicates an expected call of SectorsUnsealPiece.
func (mr *MockIGatewayMockRecorder) SectorsUnsealPiece(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SectorsUnsealPiece", reflect.TypeOf((*MockIGateway)(nil).SectorsUnsealPiece), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SupportNewAccount mocks base method.
func (m *MockIGateway) SupportNewAccount(arg0 context.Context, arg1 types.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportNewAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SupportNewAccount indicates an expected call of SupportNewAccount.
func (mr *MockIGatewayMockRecorder) SupportNewAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportNewAccount", reflect.TypeOf((*MockIGateway)(nil).SupportNewAccount), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockIGateway) Version(arg0 context.Context) (types.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockIGatewayMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockIGateway)(nil).Version), arg0)
}

// WalletHas mocks base method.
func (m *MockIGateway) WalletHas(arg0 context.Context, arg1 string, arg2 address.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletHas", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletHas indicates an expected call of WalletHas.
func (mr *MockIGatewayMockRecorder) WalletHas(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletHas", reflect.TypeOf((*MockIGateway)(nil).WalletHas), arg0, arg1, arg2)
}

// WalletSign mocks base method.
func (m *MockIGateway) WalletSign(arg0 context.Context, arg1 string, arg2 address.Address, arg3 []byte, arg4 types.MsgMeta) (*crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletSign", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletSign indicates an expected call of WalletSign.
func (mr *MockIGatewayMockRecorder) WalletSign(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletSign", reflect.TypeOf((*MockIGateway)(nil).WalletSign), arg0, arg1, arg2, arg3, arg4)
}
