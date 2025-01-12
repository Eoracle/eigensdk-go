// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry (interfaces: AvsRegistryReader)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry AvsRegistryReader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	contractOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	types "github.com/Layr-Labs/eigensdk-go/types"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsRegistryReader is a mock of AvsRegistryReader interface.
type MockAvsRegistryReader struct {
	ctrl     *gomock.Controller
	recorder *MockAvsRegistryReaderMockRecorder
}

// MockAvsRegistryReaderMockRecorder is the mock recorder for MockAvsRegistryReader.
type MockAvsRegistryReaderMockRecorder struct {
	mock *MockAvsRegistryReader
}

// NewMockAvsRegistryReader creates a new mock instance.
func NewMockAvsRegistryReader(ctrl *gomock.Controller) *MockAvsRegistryReader {
	mock := &MockAvsRegistryReader{ctrl: ctrl}
	mock.recorder = &MockAvsRegistryReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsRegistryReader) EXPECT() *MockAvsRegistryReaderMockRecorder {
	return m.recorder
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAvsRegistryReader) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 []byte, arg3 [][32]byte) (contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAvsRegistryReaderMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetOperatorAddrsInQuorumsAtCurrentBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 []byte) ([][]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorAddrsInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorAddrsInQuorumsAtCurrentBlock indicates an expected call of GetOperatorAddrsInQuorumsAtCurrentBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorAddrsInQuorumsAtCurrentBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorAddrsInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorFromId mocks base method.
func (m *MockAvsRegistryReader) GetOperatorFromId(arg0 *bind.CallOpts, arg1 [32]byte) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorFromId", arg0, arg1)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorFromId indicates an expected call of GetOperatorFromId.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorFromId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorFromId", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorFromId), arg0, arg1)
}

// GetOperatorId mocks base method.
func (m *MockAvsRegistryReader) GetOperatorId(arg0 *bind.CallOpts, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorId), arg0, arg1)
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 [32]byte) (map[byte]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].(map[byte]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsAtBlock(arg0 *bind.CallOpts, arg1 []byte, arg2 uint32) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsAtCurrentBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 []byte) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsAtCurrentBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtCurrentBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0 *bind.CallOpts, arg1 [32]byte, arg2 uint32) ([]byte, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 [32]byte) ([]byte, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetQuorumCount mocks base method.
func (m *MockAvsRegistryReader) GetQuorumCount(arg0 *bind.CallOpts) (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuorumCount", arg0)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuorumCount indicates an expected call of GetQuorumCount.
func (mr *MockAvsRegistryReaderMockRecorder) GetQuorumCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuorumCount", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetQuorumCount), arg0)
}

// IsOperatorRegistered mocks base method.
func (m *MockAvsRegistryReader) IsOperatorRegistered(arg0 *bind.CallOpts, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockAvsRegistryReaderMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockAvsRegistryReader)(nil).IsOperatorRegistered), arg0, arg1)
}

// QueryExistingRegisteredOperatorPubKeys mocks base method.
func (m *MockAvsRegistryReader) QueryExistingRegisteredOperatorPubKeys(arg0 context.Context, arg1, arg2 *big.Int) ([]common.Address, []types.OperatorPubkeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExistingRegisteredOperatorPubKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].([]types.OperatorPubkeys)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryExistingRegisteredOperatorPubKeys indicates an expected call of QueryExistingRegisteredOperatorPubKeys.
func (mr *MockAvsRegistryReaderMockRecorder) QueryExistingRegisteredOperatorPubKeys(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExistingRegisteredOperatorPubKeys", reflect.TypeOf((*MockAvsRegistryReader)(nil).QueryExistingRegisteredOperatorPubKeys), arg0, arg1, arg2)
}
