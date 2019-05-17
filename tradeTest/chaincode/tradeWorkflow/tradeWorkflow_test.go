package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestTradeWrokflowInit(t *testing.T) {
	twf := new(TradeWorkflowChaincode)
	twf.testMode = true
	mstub := shim.NewMockStub("TradeWorkflow", twf)

	args := [][]byte{[]byte("Init"),
		[]byte("Exporter Inc"),
		[]byte("Exporter Bank"),
		[]byte("1000000"),
		[]byte("Importer Inc"),
		[]byte("Importer Bank"),
		[]byte("1000000"),
		[]byte("KAL"),
		[]byte("Authority Department")}

	//check Init
	res := mstub.MockInit("1", args)

	if res.Status != shim.OK {
		fmt.Println("[Test] Init Failed", string(res.Message))
		t.FailNow()
	}
}
