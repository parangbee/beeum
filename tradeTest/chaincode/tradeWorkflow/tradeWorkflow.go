package main

import (
	"encoding/json"
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

	fmt.Printf("[TradeWorkflowChaincode.Init()] %s\n", fname)

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
		return requestTrade(stub)
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
}

func (t *TradeWorkflowChaincode) requestTrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var creatorOrg string
	var creatorCertIssuer string

	var tradeKey string
	var tradeAgreement *TradeAgreement
	var tradeAgreementBytes []byte
	var amount int
	var err error

	creatorOrg, creatorCertIssuer, _ = getTxCreatorInfo(stub)
	if !t.testMode && !authenticateImporterOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller is not a member of Importer Org. Access denied.")
	}

	if len(args) != 3 {
		err = fmt.Errorf("Incorrect number of arguments. Expecting 3:{ID, Amount, Description Of Goods}, Found %d", len(args))
		return shim.Error(err.Error())
	}

	amount, err = strconv.Atoi(string(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}

	tradeAgreement = &TradeAgreement{amount, args[2], REQUESTED, 0}
	tradeAgreementBytes, err = json.Marshal(tradeAgreement)
	if err != nil {
		return shim.Error("Error marshaling trade agreement structure")
	}

	// Write the state to the ledger
	tradeKey, err = stub.CreateCompositeKey("Trade", []string{args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(tradeKey, tradeAgreementBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Trade %s request recorded", args[0])

	return shim.Success(nil)
}

func (t *TradeWorkflowChaincode) acceptTrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var creatorOrg string
	var creatorCertIssuer string

	var tradeKey string
	var tradeAgreement *TradeAgreement
	var tradeAgreementBytes []byte
	var amount int
	var err error

	creatorOrg, creatorCertIssuer, _ = getTxCreatorInfo(stub)
	if !t.testMode && !authenticateExporterOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller is not a member of Exporter Org. Access denied.")
	}

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expecting 1:{ID}, Found %d", len(args))
		return shim.Error(err.Error())
	}

	tradeKey, err = stub.CreateCompositeKey("Trade", []string{args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}

	tradeAgreementBytes, err = stub.GetState(tradeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(tradeAgreementByte) == 0 {
		err = fmt.Errorf("No record found for trade ID %s", args[0])
		return shim.Error(err.Error())
	}

	err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
	if err != nil {
		return shim.Error(err.Error())
	}

	if tradeAgreement.Status == ACCEPTED {
		fmt.Printf("Trade %s is already accepted", args[0])
	}

}


func main() {
	twc := new(TradeWorkflowChaincode)
	twc.testMode = false

	err := shim.Start(twc)
	if err != nil {
		fmt.Printf(" Error starting Trade Workflow chaincode: %s", err)
	}
}
