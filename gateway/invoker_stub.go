package gateway

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"

	"github.com/s7techlab/cckit/router"
)

var ErrInvokeMethodNotAllowed = errors.New(`invoke method not allowed`)

type (
	ChaincodeLocatorResolver func(ctx router.Context, serviceName string) (*ChaincodeLocator, error)
	// ChaincodeStubInvoker for cross chaincode calls - only query (read) methods
	// context argument is router.Context, not context.Context
	ChaincodeStubInvoker interface {
		Query(stub shim.ChaincodeStubInterface, fn string, args []interface{}, target interface{}) (interface{}, error)
	}

	LocatorChaincodeStubInvoker struct {
		Locator *ChaincodeLocator
	}
)

func (c *LocatorChaincodeStubInvoker) Query(
	stub shim.ChaincodeStubInterface, fn string, args []interface{}, target interface{}) (interface{}, error) {

	argsBytes, err := InvokerArgs(fn, args)
	if err != nil {
		return nil, err
	}

	response := stub.InvokeChaincode(c.Locator.Chaincode, argsBytes, c.Locator.Channel)
	if response.Status != shim.OK {
		return nil, fmt.Errorf(`cross chaincode=%s, channel=%s invoke: %w`,
			c.Locator.Chaincode, c.Locator.Channel, errors.New(response.Message))
	}

	return ссOutput(&response, target)
}
