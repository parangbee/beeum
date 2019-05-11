package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// TradeWorkflowChaincode implementation
type TradeWorkflowChaincode struct {
	testMode bool
}

// Init implementation
func (t *TradeWorkflowChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init TradeWorkflowChaincode")
	fname, args := stub.GetFunctionAndParameters)
	var err error

	if len(args) == 0 {
		return shim.Success(nil)
	}

	if len(args) != 8 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments, Expecting 8:" +
			"Exporter, Exporter's Bank, Exporter Account Balance, Importer, Importer's Bank, " +
			"Importer Account Balance, Carrier, Requlatory Authority"))
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Invoke implementation
func (t *TradeWorkflowChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke TradeWorkflowChaincode")
	return shim.Success(nil)
}

func main() {
	twc := new(TradeWorkflowChaincode)
	twc.testMode = false
	err := shim.Start(twc)
	if err != nil {
		fmt.Printf(" Error starting Trade Workflow chaincode: %s", err)
	}
}
