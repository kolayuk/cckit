// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: debug/debug_state.proto

/*
Package debug contains
  *   chaincode methods names {service_name}Chaincode_{method_name}
  *   chaincode interface definition {service_name}Chaincode
  *   chaincode gateway definition {service_name}}Gateway
  *   chaincode service to cckit router registration func
*/
package debug

import (
	context "context"
	_ "embed"

	cckit_gateway "github.com/s7techlab/cckit/gateway"
	cckit_router "github.com/s7techlab/cckit/router"
	cckit_defparam "github.com/s7techlab/cckit/router/param/defparam"
	cckit_sdk "github.com/s7techlab/cckit/sdk"
)

// DebugStateServiceChaincode method names
const (

	// DebugStateServiceChaincodeMethodPrefix allows to use multiple services with same method names in one chaincode
	DebugStateServiceChaincodeMethodPrefix = "DebugStateService."

	DebugStateServiceChaincode_ListKeys = DebugStateServiceChaincodeMethodPrefix + "ListKeys"

	DebugStateServiceChaincode_GetState = DebugStateServiceChaincodeMethodPrefix + "GetState"

	DebugStateServiceChaincode_PutState = DebugStateServiceChaincodeMethodPrefix + "PutState"

	DebugStateServiceChaincode_DeleteState = DebugStateServiceChaincodeMethodPrefix + "DeleteState"

	DebugStateServiceChaincode_DeleteStates = DebugStateServiceChaincodeMethodPrefix + "DeleteStates"
)

// DebugStateServiceChaincode chaincode methods interface
type DebugStateServiceChaincode interface {
	ListKeys(cckit_router.Context, *Prefix) (*CompositeKeys, error)

	GetState(cckit_router.Context, *CompositeKey) (*Value, error)

	PutState(cckit_router.Context, *Value) (*Value, error)

	DeleteState(cckit_router.Context, *CompositeKey) (*Value, error)

	DeleteStates(cckit_router.Context, *Prefixes) (*PrefixesMatchCount, error)
}

// RegisterDebugStateServiceChaincode registers service methods as chaincode router handlers
func RegisterDebugStateServiceChaincode(r *cckit_router.Group, cc DebugStateServiceChaincode) error {

	r.Query(DebugStateServiceChaincode_ListKeys,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.ListKeys(ctx, ctx.Param().(*Prefix))
		},
		cckit_defparam.Proto(&Prefix{}))

	r.Query(DebugStateServiceChaincode_GetState,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetState(ctx, ctx.Param().(*CompositeKey))
		},
		cckit_defparam.Proto(&CompositeKey{}))

	r.Invoke(DebugStateServiceChaincode_PutState,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.PutState(ctx, ctx.Param().(*Value))
		},
		cckit_defparam.Proto(&Value{}))

	r.Invoke(DebugStateServiceChaincode_DeleteState,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.DeleteState(ctx, ctx.Param().(*CompositeKey))
		},
		cckit_defparam.Proto(&CompositeKey{}))

	r.Invoke(DebugStateServiceChaincode_DeleteStates,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.DeleteStates(ctx, ctx.Param().(*Prefixes))
		},
		cckit_defparam.Proto(&Prefixes{}))

	return nil
}

//go:embed debug_state.swagger.json
var DebugStateServiceSwagger []byte

// NewDebugStateServiceGateway creates gateway to access chaincode method via chaincode service
func NewDebugStateServiceGateway(sdk cckit_sdk.SDK, channel, chaincode string, opts ...cckit_gateway.Opt) *DebugStateServiceGateway {
	return &DebugStateServiceGateway{
		Invoker: &cckit_gateway.ChaincodeInstanceServiceInvoker{
			ChaincodeInstance: cckit_gateway.NewChaincodeInstanceService(
				sdk,
				&cckit_gateway.ChaincodeLocator{Channel: channel, Chaincode: chaincode},
				opts...),
		},
	}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type DebugStateServiceGateway struct {
	Invoker cckit_gateway.ChaincodeInstanceInvoker
}

// ServiceDef returns service definition
func (c *DebugStateServiceGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.NewServiceDef(
		_DebugStateService_serviceDesc.ServiceName,
		DebugStateServiceSwagger,
		&_DebugStateService_serviceDesc,
		c,
		RegisterDebugStateServiceHandlerFromEndpoint,
	)
}

func (c *DebugStateServiceGateway) ListKeys(ctx context.Context, in *Prefix) (*CompositeKeys, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx, DebugStateServiceChaincode_ListKeys, []interface{}{in}, &CompositeKeys{}); err != nil {
		return nil, err
	} else {
		return res.(*CompositeKeys), nil
	}
}

func (c *DebugStateServiceGateway) GetState(ctx context.Context, in *CompositeKey) (*Value, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx, DebugStateServiceChaincode_GetState, []interface{}{in}, &Value{}); err != nil {
		return nil, err
	} else {
		return res.(*Value), nil
	}
}

func (c *DebugStateServiceGateway) PutState(ctx context.Context, in *Value) (*Value, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Invoke(ctx, DebugStateServiceChaincode_PutState, []interface{}{in}, &Value{}); err != nil {
		return nil, err
	} else {
		return res.(*Value), nil
	}
}

func (c *DebugStateServiceGateway) DeleteState(ctx context.Context, in *CompositeKey) (*Value, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Invoke(ctx, DebugStateServiceChaincode_DeleteState, []interface{}{in}, &Value{}); err != nil {
		return nil, err
	} else {
		return res.(*Value), nil
	}
}

func (c *DebugStateServiceGateway) DeleteStates(ctx context.Context, in *Prefixes) (*PrefixesMatchCount, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Invoke(ctx, DebugStateServiceChaincode_DeleteStates, []interface{}{in}, &PrefixesMatchCount{}); err != nil {
		return nil, err
	} else {
		return res.(*PrefixesMatchCount), nil
	}
}
