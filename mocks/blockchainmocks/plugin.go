// Code generated by mockery v2.46.0. DO NOT EDIT.

package blockchainmocks

import (
	cache "github.com/hyperledger/firefly/internal/cache"
	blockchain "github.com/hyperledger/firefly/pkg/blockchain"

	config "github.com/hyperledger/firefly-common/pkg/config"

	context "context"

	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	metrics "github.com/hyperledger/firefly/internal/metrics"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// AddContractListener provides a mock function with given fields: ctx, subscription, lastProtocolID
func (_m *Plugin) AddContractListener(ctx context.Context, subscription *core.ContractListener, lastProtocolID string) error {
	ret := _m.Called(ctx, subscription, lastProtocolID)

	if len(ret) == 0 {
		panic("no return value specified for AddContractListener")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListener, string) error); ok {
		r0 = rf(ctx, subscription, lastProtocolID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddFireflySubscription provides a mock function with given fields: ctx, namespace, contract, lastProtocolID
func (_m *Plugin) AddFireflySubscription(ctx context.Context, namespace *core.Namespace, contract *blockchain.MultipartyContract, lastProtocolID string) (string, error) {
	ret := _m.Called(ctx, namespace, contract, lastProtocolID)

	if len(ret) == 0 {
		panic("no return value specified for AddFireflySubscription")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Namespace, *blockchain.MultipartyContract, string) (string, error)); ok {
		return rf(ctx, namespace, contract, lastProtocolID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Namespace, *blockchain.MultipartyContract, string) string); ok {
		r0 = rf(ctx, namespace, contract, lastProtocolID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Namespace, *blockchain.MultipartyContract, string) error); ok {
		r1 = rf(ctx, namespace, contract, lastProtocolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *blockchain.Capabilities {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Capabilities")
	}

	var r0 *blockchain.Capabilities
	if rf, ok := ret.Get(0).(func() *blockchain.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*blockchain.Capabilities)
		}
	}

	return r0
}

// CheckOverlappingLocations provides a mock function with given fields: ctx, left, right
func (_m *Plugin) CheckOverlappingLocations(ctx context.Context, left *fftypes.JSONAny, right *fftypes.JSONAny) (bool, error) {
	ret := _m.Called(ctx, left, right)

	if len(ret) == 0 {
		panic("no return value specified for CheckOverlappingLocations")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.JSONAny, *fftypes.JSONAny) (bool, error)); ok {
		return rf(ctx, left, right)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.JSONAny, *fftypes.JSONAny) bool); ok {
		r0 = rf(ctx, left, right)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.JSONAny, *fftypes.JSONAny) error); ok {
		r1 = rf(ctx, left, right)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContractListener provides a mock function with given fields: ctx, subscription, okNotFound
func (_m *Plugin) DeleteContractListener(ctx context.Context, subscription *core.ContractListener, okNotFound bool) error {
	ret := _m.Called(ctx, subscription, okNotFound)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContractListener")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.ContractListener, bool) error); ok {
		r0 = rf(ctx, subscription, okNotFound)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployContract provides a mock function with given fields: ctx, nsOpID, signingKey, definition, contract, input, options
func (_m *Plugin) DeployContract(ctx context.Context, nsOpID string, signingKey string, definition *fftypes.JSONAny, contract *fftypes.JSONAny, input []interface{}, options map[string]interface{}) (bool, error) {
	ret := _m.Called(ctx, nsOpID, signingKey, definition, contract, input, options)

	if len(ret) == 0 {
		panic("no return value specified for DeployContract")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *fftypes.JSONAny, *fftypes.JSONAny, []interface{}, map[string]interface{}) (bool, error)); ok {
		return rf(ctx, nsOpID, signingKey, definition, contract, input, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *fftypes.JSONAny, *fftypes.JSONAny, []interface{}, map[string]interface{}) bool); ok {
		r0 = rf(ctx, nsOpID, signingKey, definition, contract, input, options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *fftypes.JSONAny, *fftypes.JSONAny, []interface{}, map[string]interface{}) error); ok {
		r1 = rf(ctx, nsOpID, signingKey, definition, contract, input, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateErrorSignature provides a mock function with given fields: ctx, errorDef
func (_m *Plugin) GenerateErrorSignature(ctx context.Context, errorDef *fftypes.FFIErrorDefinition) string {
	ret := _m.Called(ctx, errorDef)

	if len(ret) == 0 {
		panic("no return value specified for GenerateErrorSignature")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIErrorDefinition) string); ok {
		r0 = rf(ctx, errorDef)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateEventSignature provides a mock function with given fields: ctx, event
func (_m *Plugin) GenerateEventSignature(ctx context.Context, event *fftypes.FFIEventDefinition) (string, error) {
	ret := _m.Called(ctx, event)

	if len(ret) == 0 {
		panic("no return value specified for GenerateEventSignature")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIEventDefinition) (string, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIEventDefinition) string); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.FFIEventDefinition) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateEventSignatureWithLocation provides a mock function with given fields: ctx, event, location
func (_m *Plugin) GenerateEventSignatureWithLocation(ctx context.Context, event *fftypes.FFIEventDefinition, location *fftypes.JSONAny) (string, error) {
	ret := _m.Called(ctx, event, location)

	if len(ret) == 0 {
		panic("no return value specified for GenerateEventSignatureWithLocation")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIEventDefinition, *fftypes.JSONAny) (string, error)); ok {
		return rf(ctx, event, location)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIEventDefinition, *fftypes.JSONAny) string); ok {
		r0 = rf(ctx, event, location)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.FFIEventDefinition, *fftypes.JSONAny) error); ok {
		r1 = rf(ctx, event, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateFFI provides a mock function with given fields: ctx, generationRequest
func (_m *Plugin) GenerateFFI(ctx context.Context, generationRequest *fftypes.FFIGenerationRequest) (*fftypes.FFI, error) {
	ret := _m.Called(ctx, generationRequest)

	if len(ret) == 0 {
		panic("no return value specified for GenerateFFI")
	}

	var r0 *fftypes.FFI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIGenerationRequest) (*fftypes.FFI, error)); ok {
		return rf(ctx, generationRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIGenerationRequest) *fftypes.FFI); ok {
		r0 = rf(ctx, generationRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.FFI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.FFIGenerationRequest) error); ok {
		r1 = rf(ctx, generationRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAndConvertDeprecatedContractConfig provides a mock function with given fields: ctx
func (_m *Plugin) GetAndConvertDeprecatedContractConfig(ctx context.Context) (*fftypes.JSONAny, string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAndConvertDeprecatedContractConfig")
	}

	var r0 *fftypes.JSONAny
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (*fftypes.JSONAny, string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *fftypes.JSONAny); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.JSONAny)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetContractListenerStatus provides a mock function with given fields: ctx, namespace, subID, okNotFound
func (_m *Plugin) GetContractListenerStatus(ctx context.Context, namespace string, subID string, okNotFound bool) (bool, interface{}, fftypes.FFEnum, error) {
	ret := _m.Called(ctx, namespace, subID, okNotFound)

	if len(ret) == 0 {
		panic("no return value specified for GetContractListenerStatus")
	}

	var r0 bool
	var r1 interface{}
	var r2 fftypes.FFEnum
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) (bool, interface{}, fftypes.FFEnum, error)); ok {
		return rf(ctx, namespace, subID, okNotFound)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) bool); ok {
		r0 = rf(ctx, namespace, subID, okNotFound)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) interface{}); ok {
		r1 = rf(ctx, namespace, subID, okNotFound)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, bool) fftypes.FFEnum); ok {
		r2 = rf(ctx, namespace, subID, okNotFound)
	} else {
		r2 = ret.Get(2).(fftypes.FFEnum)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, string, bool) error); ok {
		r3 = rf(ctx, namespace, subID, okNotFound)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetFFIParamValidator provides a mock function with given fields: ctx
func (_m *Plugin) GetFFIParamValidator(ctx context.Context) (fftypes.FFIParamValidator, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFFIParamValidator")
	}

	var r0 fftypes.FFIParamValidator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (fftypes.FFIParamValidator, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) fftypes.FFIParamValidator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.FFIParamValidator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkVersion provides a mock function with given fields: ctx, location
func (_m *Plugin) GetNetworkVersion(ctx context.Context, location *fftypes.JSONAny) (int, error) {
	ret := _m.Called(ctx, location)

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkVersion")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.JSONAny) (int, error)); ok {
		return rf(ctx, location)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.JSONAny) int); ok {
		r0 = rf(ctx, location)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.JSONAny) error); ok {
		r1 = rf(ctx, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionStatus provides a mock function with given fields: ctx, operation
func (_m *Plugin) GetTransactionStatus(ctx context.Context, operation *core.Operation) (interface{}, error) {
	ret := _m.Called(ctx, operation)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionStatus")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) (interface{}, error)); ok {
		return rf(ctx, operation)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) interface{}); ok {
		r0 = rf(ctx, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Operation) error); ok {
		r1 = rf(ctx, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx, cancelCtx, _a2, _a3, cacheManager
func (_m *Plugin) Init(ctx context.Context, cancelCtx context.CancelFunc, _a2 config.Section, _a3 metrics.Manager, cacheManager cache.Manager) error {
	ret := _m.Called(ctx, cancelCtx, _a2, _a3, cacheManager)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, context.CancelFunc, config.Section, metrics.Manager, cache.Manager) error); ok {
		r0 = rf(ctx, cancelCtx, _a2, _a3, cacheManager)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitConfig provides a mock function with given fields: _a0
func (_m *Plugin) InitConfig(_a0 config.Section) {
	_m.Called(_a0)
}

// InvokeContract provides a mock function with given fields: ctx, nsOpID, signingKey, location, parsedMethod, input, options, batch
func (_m *Plugin) InvokeContract(ctx context.Context, nsOpID string, signingKey string, location *fftypes.JSONAny, parsedMethod interface{}, input map[string]interface{}, options map[string]interface{}, batch *blockchain.BatchPin) (bool, error) {
	ret := _m.Called(ctx, nsOpID, signingKey, location, parsedMethod, input, options, batch)

	if len(ret) == 0 {
		panic("no return value specified for InvokeContract")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}, *blockchain.BatchPin) (bool, error)); ok {
		return rf(ctx, nsOpID, signingKey, location, parsedMethod, input, options, batch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}, *blockchain.BatchPin) bool); ok {
		r0 = rf(ctx, nsOpID, signingKey, location, parsedMethod, input, options, batch)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}, *blockchain.BatchPin) error); ok {
		r1 = rf(ctx, nsOpID, signingKey, location, parsedMethod, input, options, batch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NormalizeContractLocation provides a mock function with given fields: ctx, ntype, location
func (_m *Plugin) NormalizeContractLocation(ctx context.Context, ntype blockchain.NormalizeType, location *fftypes.JSONAny) (*fftypes.JSONAny, error) {
	ret := _m.Called(ctx, ntype, location)

	if len(ret) == 0 {
		panic("no return value specified for NormalizeContractLocation")
	}

	var r0 *fftypes.JSONAny
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, blockchain.NormalizeType, *fftypes.JSONAny) (*fftypes.JSONAny, error)); ok {
		return rf(ctx, ntype, location)
	}
	if rf, ok := ret.Get(0).(func(context.Context, blockchain.NormalizeType, *fftypes.JSONAny) *fftypes.JSONAny); ok {
		r0 = rf(ctx, ntype, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.JSONAny)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, blockchain.NormalizeType, *fftypes.JSONAny) error); ok {
		r1 = rf(ctx, ntype, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInterface provides a mock function with given fields: ctx, method, errors
func (_m *Plugin) ParseInterface(ctx context.Context, method *fftypes.FFIMethod, errors []*fftypes.FFIError) (interface{}, error) {
	ret := _m.Called(ctx, method, errors)

	if len(ret) == 0 {
		panic("no return value specified for ParseInterface")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIMethod, []*fftypes.FFIError) (interface{}, error)); ok {
		return rf(ctx, method, errors)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.FFIMethod, []*fftypes.FFIError) interface{}); ok {
		r0 = rf(ctx, method, errors)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.FFIMethod, []*fftypes.FFIError) error); ok {
		r1 = rf(ctx, method, errors)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryContract provides a mock function with given fields: ctx, signingKey, location, parsedMethod, input, options
func (_m *Plugin) QueryContract(ctx context.Context, signingKey string, location *fftypes.JSONAny, parsedMethod interface{}, input map[string]interface{}, options map[string]interface{}) (interface{}, error) {
	ret := _m.Called(ctx, signingKey, location, parsedMethod, input, options)

	if len(ret) == 0 {
		panic("no return value specified for QueryContract")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}) (interface{}, error)); ok {
		return rf(ctx, signingKey, location, parsedMethod, input, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}) interface{}); ok {
		r0 = rf(ctx, signingKey, location, parsedMethod, input, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *fftypes.JSONAny, interface{}, map[string]interface{}, map[string]interface{}) error); ok {
		r1 = rf(ctx, signingKey, location, parsedMethod, input, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFireflySubscription provides a mock function with given fields: ctx, subID
func (_m *Plugin) RemoveFireflySubscription(ctx context.Context, subID string) {
	_m.Called(ctx, subID)
}

// ResolveSigningKey provides a mock function with given fields: ctx, keyRef, intent
func (_m *Plugin) ResolveSigningKey(ctx context.Context, keyRef string, intent blockchain.ResolveKeyIntent) (string, error) {
	ret := _m.Called(ctx, keyRef, intent)

	if len(ret) == 0 {
		panic("no return value specified for ResolveSigningKey")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, blockchain.ResolveKeyIntent) (string, error)); ok {
		return rf(ctx, keyRef, intent)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, blockchain.ResolveKeyIntent) string); ok {
		r0 = rf(ctx, keyRef, intent)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, blockchain.ResolveKeyIntent) error); ok {
		r1 = rf(ctx, keyRef, intent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetHandler provides a mock function with given fields: namespace, handler
func (_m *Plugin) SetHandler(namespace string, handler blockchain.Callbacks) {
	_m.Called(namespace, handler)
}

// SetOperationHandler provides a mock function with given fields: namespace, handler
func (_m *Plugin) SetOperationHandler(namespace string, handler core.OperationCallbacks) {
	_m.Called(namespace, handler)
}

// StartNamespace provides a mock function with given fields: ctx, namespace
func (_m *Plugin) StartNamespace(ctx context.Context, namespace string) error {
	ret := _m.Called(ctx, namespace)

	if len(ret) == 0 {
		panic("no return value specified for StartNamespace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopNamespace provides a mock function with given fields: ctx, namespace
func (_m *Plugin) StopNamespace(ctx context.Context, namespace string) error {
	ret := _m.Called(ctx, namespace)

	if len(ret) == 0 {
		panic("no return value specified for StopNamespace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitBatchPin provides a mock function with given fields: ctx, nsOpID, networkNamespace, signingKey, batch, location
func (_m *Plugin) SubmitBatchPin(ctx context.Context, nsOpID string, networkNamespace string, signingKey string, batch *blockchain.BatchPin, location *fftypes.JSONAny) error {
	ret := _m.Called(ctx, nsOpID, networkNamespace, signingKey, batch, location)

	if len(ret) == 0 {
		panic("no return value specified for SubmitBatchPin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *blockchain.BatchPin, *fftypes.JSONAny) error); ok {
		r0 = rf(ctx, nsOpID, networkNamespace, signingKey, batch, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitNetworkAction provides a mock function with given fields: ctx, nsOpID, signingKey, action, location
func (_m *Plugin) SubmitNetworkAction(ctx context.Context, nsOpID string, signingKey string, action fftypes.FFEnum, location *fftypes.JSONAny) error {
	ret := _m.Called(ctx, nsOpID, signingKey, action, location)

	if len(ret) == 0 {
		panic("no return value specified for SubmitNetworkAction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, fftypes.FFEnum, *fftypes.JSONAny) error); ok {
		r0 = rf(ctx, nsOpID, signingKey, action, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateInvokeRequest provides a mock function with given fields: ctx, parsedMethod, input, hasMessage
func (_m *Plugin) ValidateInvokeRequest(ctx context.Context, parsedMethod interface{}, input map[string]interface{}, hasMessage bool) error {
	ret := _m.Called(ctx, parsedMethod, input, hasMessage)

	if len(ret) == 0 {
		panic("no return value specified for ValidateInvokeRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, map[string]interface{}, bool) error); ok {
		r0 = rf(ctx, parsedMethod, input, hasMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifierType provides a mock function with given fields:
func (_m *Plugin) VerifierType() fftypes.FFEnum {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VerifierType")
	}

	var r0 fftypes.FFEnum
	if rf, ok := ret.Get(0).(func() fftypes.FFEnum); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(fftypes.FFEnum)
	}

	return r0
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
