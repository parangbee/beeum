package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// TradeWorkflowChaincode implementation
type TradeWorkflowChaincode struct {
	testMode bool
}

// Init implementation
func (t *TradeWorkflowChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("[TradeWorkflowChaincode.Init()] Init TradeWorkflowChaincode")
	fname, args := stub.GetFunctionAndParameters()
	var err error

	fmt.Printf("[TradeWorkflowChaincode.Init()] %s", fname)

	if len(args) == 0 {
		return shim.Success(nil)
	}

	// Upgrade mode 2: change all the names and account balances
	if len(args) != 8 {
		err = fmt.Errorf("Incorrect number of arguments. Expecting 8: {"+
			"Exporter, Exporter's Bank, Exporter's Account Balance, "+
			"Importer, Importer's Bank, Importer's Account Balance, "+
			"Carrier, Regulatory Authority}. Found %d", len(args))
		return shim.Error(err.Error())
	}

	// check : Exporter's Balance
	_, err = strconv.Atoi(string(args[2]))
	if err != nil {
		fmt.Printf("[TradeWorkflowChaincode.Init()] Exporter's Balnace must be an integer.(%s)", args[2])
		return shim.Error(err.Error())
	}

	// check : Importer's Balance
	_, err = strconv.Atoi(string(args[5]))
	if err != nil {
		fmt.Printf("[TradeWorkflowChaincode.Init()] Importer's Balnace must be an integer.(%s)", args[5])
		return shim.Error(err.Error())
	}

	fmt.Printf("[TradeWorkflowChaincode.Init()] Exporter: %s\n", args[0])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Exporter's Bank: %s\n", args[1])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Exporter's Account Balance: %s\n", args[2])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Importer: %s\n", args[3])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Importer's Bank: %s\n", args[4])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Importer's Account Balance: %s\n", args[5])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Carrier: %s\n", args[6])
	fmt.Printf("[TradeWorkflowChaincode.Init()] Regulatory Authority: %s\n", args[7])

	// map participant identities to roles on the ledger
	roleKeys := []string{expKey, exbKey, exaKey, impKey, imbKey, imaKey, carKey, rgaKey}
	for i, roleKey := range roleKeys {
		err = stub.PutState(roleKey, []byte(args[i]))
		if err != nil {
			fmt.Printf("[TradeWorkflowChaincode.Init()] %s:%s\n", roleKey, err.Error())
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

// Invoke implementation
func (t *TradeWorkflowChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke TradeWorkflowChaincode")

	function, args := stub.GetFunctionAndParameters()

	if function == "requestTrade" {
		fmt.Println("[TradeWorkflowChaincode.Invoke()] requestTrade was called")
		return shim.Success(nil)
	} else if function == "acceptTrade" {
		fmt.Println("[TradeWorkflowChaincode.Invoke()] acceptTrade was called")
		return shim.Success(nil)
	} else if function == "requestLC" {
		fmt.Println("[TradeWorkflowChaincode.Invoke()] requestLC was called")
		return shim.Success(nil)
	} else if function == "issueLC" {
		fmt.Println("[TradeWorkflowChaincode.Invoke()] issueLC was called")
		return shim.Success(nil)
	} else if function == "accpetLC" {
		fmt.Printf("[TradeWorkflowChaincode.Invoke()] accpetLC was called with args(%d)", len(args))
		return shim.Success(nil)
	}

	return shim.Error("[TradeWorkflowChaincode.Invoke()] Invalid funtion name")

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
