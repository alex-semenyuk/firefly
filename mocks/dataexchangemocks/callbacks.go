// Code generated by mockery v1.0.0. DO NOT EDIT.

package dataexchangemocks

import (
	fftypes "github.com/hyperledger-labs/firefly/pkg/fftypes"
	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// BLOBReceived provides a mock function with given fields: peerID, hash, payloadRef
func (_m *Callbacks) BLOBReceived(peerID string, hash fftypes.Bytes32, payloadRef string) error {
	ret := _m.Called(peerID, hash, payloadRef)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fftypes.Bytes32, string) error); ok {
		r0 = rf(peerID, hash, payloadRef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MessageReceived provides a mock function with given fields: peerID, data
func (_m *Callbacks) MessageReceived(peerID string, data []byte) error {
	ret := _m.Called(peerID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(peerID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferResult provides a mock function with given fields: trackingID, status, info, additionalInfo
func (_m *Callbacks) TransferResult(trackingID string, status fftypes.OpStatus, info string, additionalInfo fftypes.JSONObject) error {
	ret := _m.Called(trackingID, status, info, additionalInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fftypes.OpStatus, string, fftypes.JSONObject) error); ok {
		r0 = rf(trackingID, status, info, additionalInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}