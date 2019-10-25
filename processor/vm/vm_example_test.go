// Copyright 2018 The Fractal Team Authors
// This file is part of the fractal project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package vm_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/fractalplatform/fractal/common"
	"github.com/fractalplatform/fractal/plugin"
	"github.com/fractalplatform/fractal/processor/vm/runtime"
	"github.com/fractalplatform/fractal/rawdb"
	"github.com/fractalplatform/fractal/state"
	"github.com/fractalplatform/fractal/types"
	"github.com/fractalplatform/fractal/utils/abi"
	"github.com/fractalplatform/fractal/utils/rlp"
	"github.com/stretchr/testify/assert"
)

func input(abifile string, method string, params ...interface{}) ([]byte, error) {
	var abicode string

	hexcode, err := ioutil.ReadFile(abifile)
	if err != nil {
		fmt.Printf("Could not load code from file: %v\n", err)
		return nil, err
	}
	abicode = string(bytes.TrimRight(hexcode, "\n"))

	parsed, err := abi.JSON(strings.NewReader(abicode))
	if err != nil {
		fmt.Println("abi.json error ", err)
		return nil, err
	}

	input, err := parsed.Pack(method, params...)
	if err != nil {
		fmt.Println("parsed.pack error ", err)
		return nil, err
	}
	return input, nil
}

func createContract(abifile string, binfile string, contractName string, runtimeConfig runtime.Config) error {
	hexcode, err := ioutil.ReadFile(binfile)
	if err != nil {
		fmt.Printf("Could not load code from file: %v\n", err)
		os.Exit(1)
	}
	code := common.Hex2Bytes(string(bytes.TrimRight(hexcode, "\n")))

	createInput, err := input(abifile, "")
	if err != nil {
		fmt.Println("createInput error ", err)
		return err
	}

	createCode := append(code, createInput...)
	action := types.NewAction(types.CreateContract, runtimeConfig.Origin, contractName, 0, 1, runtimeConfig.GasLimit, runtimeConfig.Value, createCode, nil)
	_, _, err = runtime.Create(action, &runtimeConfig)
	if err != nil {
		fmt.Println("create error ", err)
		return err
	}
	return nil
}

func createAccount(pm plugin.IPM, name string) error {
	if _, err := pm.CreateAccount(string(name), common.HexToPubKey("12345"), ""); err != nil {
		fmt.Printf("create account %s err %s", name, err)
		return fmt.Errorf("create account %s err %s", name, err)
	}
	return nil
}

func issueAssetAction(ownerName, toName string) *types.Action {
	asset := plugin.IssueAsset{
		AssetName:  "bitcoin",
		Symbol:     "btc",
		Amount:     big.NewInt(1000000000000000000),
		Decimals:   10,
		Owner:      ownerName,
		UpperLimit: big.NewInt(2000000000000000000),
		Founder:    ownerName,
	}

	b, err := rlp.EncodeToBytes(asset)
	if err != nil {
		panic(err)
	}

	action := types.NewAction(types.IssueAsset, ownerName, string("fractal.asset"), 0, 0, 0, big.NewInt(0), b, nil)
	return action
}

// func TestAsset(t *testing.T) {
// 	state, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
// 	pm := plugin.NewPM(state)

// 	senderName := string("jacobwolf12345")
// 	senderPubkey := common.HexToPubKey("12345")

// 	receiverName := string("denverfolk12345")

// 	if err := createAccount(pm, "jacobwolf12345"); err != nil {
// 		return
// 	}

// 	if err := createAccount(pm, "denverfolk12345"); err != nil {
// 		return
// 	}

// 	if err := createAccount(pm, "fractal.asset"); err != nil {
// 		return
// 	}

// 	//action := issueAssetAction(senderName, receiverName)
// 	if _, err := pm.IssueAsset(senderName, "bitcoin", "btc", big.NewInt(1000000000000000000), 10, senderName, senderName, big.NewInt(2000000000000000000), ""); err != nil {
// 		fmt.Println("issue asset error", err)
// 		return
// 	}

// 	runtimeConfig := runtime.Config{
// 		Origin:      senderName,
// 		FromPubkey:  senderPubkey,
// 		State:       state,
// 		Account:     account,
// 		AssetID:     0,
// 		GasLimit:    10000000000,
// 		GasPrice:    big.NewInt(0),
// 		Value:       big.NewInt(0),
// 		BlockNumber: new(big.Int).SetUint64(0),
// 	}

// 	binfile := "./runtime/contract/Asset/Asset.bin"
// 	abifile := "./runtime/contract/Asset/Asset.abi"
// 	contractName := string("assetcontract")
// 	if err := createAccount(pm, "assetcontract"); err != nil {
// 		return
// 	}

// 	err := createContract(abifile, binfile, contractName, runtimeConfig)
// 	if err != nil {
// 		fmt.Println("create calledcontract error", err)
// 		return
// 	}

// 	issuseAssetInput, err := input(abifile, "reg", "ethnewfromname2,ethereum,10000000000,0,jacobwolf12345,20000000000,jacobwolf12345,assetcontract,this is contracgt asset")
// 	if err != nil {
// 		fmt.Println("issuseAssetInput error ", err)
// 		return
// 	}
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, 0, runtimeConfig.GasLimit, runtimeConfig.Value, issuseAssetInput, nil)

// 	ret, _, err := runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("call error ", err)
// 		return
// 	}
// 	num := new(big.Int).SetBytes(ret)
// 	if num.Cmp(big.NewInt(1)) != 0 {
// 		t.Error("getBalance fail, want 1, get ", num)
// 	}

// 	senderAcc, err := pm.GetAccountByName(senderName)
// 	if err != nil {
// 		fmt.Println("GetAccountByName sender account error", err)
// 		return
// 	}

// 	result := senderAcc.GetBalancesList()
// 	if err != nil {
// 		fmt.Println("GetAllAccountBalancesset error", err)
// 		return
// 	}
// 	assert.Equal(t, result[0], &accountmanager.AssetBalance{AssetID: 0, Balance: big.NewInt(1000000000000000000)})
// 	assert.Equal(t, result[1], &accountmanager.AssetBalance{AssetID: 1, Balance: big.NewInt(10000000000)})

// 	addAssetInput, err := input(abifile, "add", big.NewInt(1), common.BigToAddress(big.NewInt(4097)), big.NewInt(210000))
// 	if err != nil {
// 		fmt.Println("addAssetInput error ", err)
// 		return
// 	}
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, 0, runtimeConfig.GasLimit, runtimeConfig.Value, addAssetInput, nil)

// 	_, _, err = runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("add call error ", err)
// 		return
// 	}

// 	senderAcc, err = account.GetAccountByName(senderName)
// 	if err != nil {
// 		fmt.Println("GetAccountByName sender account error", err)
// 		return
// 	}

// 	result = senderAcc.GetBalancesList()
// 	for _, b := range result {
// 		fmt.Println("asset result ", b)
// 	}

// 	transferExAssetInput, err := input(abifile, "transAsset", common.BigToAddress(big.NewInt(4098)), big.NewInt(1), big.NewInt(10000))
// 	if err != nil {
// 		fmt.Println("transferExAssetInput error ", err)
// 		return
// 	}
// 	runtimeConfig.Value = big.NewInt(100000)
// 	runtimeConfig.AssetID = 1
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, transferExAssetInput, nil)

// 	_, _, err = runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("call error ", err)
// 		return
// 	}

// 	senderAcc, err = account.GetAccountByName(senderName)
// 	if err != nil {
// 		fmt.Println("GetAccountByName sender account error", err)
// 		return
// 	}

// 	receiverAcc, err := account.GetAccountByName(receiverName)
// 	if err != nil {
// 		fmt.Println("GetAccountByName receiver account error", err)
// 		return
// 	}

// 	result = senderAcc.GetBalancesList()
// 	assert.Equal(t, result[0], &accountmanager.AssetBalance{AssetID: 0, Balance: big.NewInt(1000000000000000000)})
// 	assert.Equal(t, result[1], &accountmanager.AssetBalance{AssetID: 1, Balance: big.NewInt(9999900000)})

// 	result = receiverAcc.GetBalancesList()
// 	assert.Equal(t, result[0], &accountmanager.AssetBalance{AssetID: 1, Balance: big.NewInt(10000)})

// 	setOwnerInput, err := input(abifile, "setname", common.BigToAddress(big.NewInt(4098)), big.NewInt(1))
// 	if err != nil {
// 		fmt.Println("setOwnerInput error ", err)
// 		return
// 	}
// 	runtimeConfig.Value = big.NewInt(0)
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, setOwnerInput, nil)

// 	_, _, err = runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("call error ", err)
// 		return
// 	}

// 	getBalanceInput, err := input(abifile, "getbalance", common.BigToAddress(big.NewInt(4098)), big.NewInt(1))
// 	if err != nil {
// 		fmt.Println("getBalanceInput error ", err)
// 		return
// 	}
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, getBalanceInput, nil)

// 	ret, _, err = runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("call error ", err)
// 		return
// 	}
// 	num = new(big.Int).SetBytes(ret)
// 	if num.Cmp(big.NewInt(10000)) != 0 {
// 		t.Error("getBalance fail, want 10000, get ", num)
// 	}

// 	getAssetIDInput, err := input(abifile, "getAssetId")
// 	if err != nil {
// 		fmt.Println("getBalanceInput error ", err)
// 		return
// 	}
// 	action = types.NewAction(types.CallContract, runtimeConfig.Origin, contractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, getAssetIDInput, nil)

// 	ret, _, err = runtime.Call(action, &runtimeConfig)
// 	if err != nil {
// 		fmt.Println("call error ", err)
// 		return
// 	}
// 	num = new(big.Int).SetBytes(ret)
// 	if num.Cmp(big.NewInt(1)) != 0 {
// 		t.Error("getBalance fail, want 1, get ", num)
// 	}
// }
func TestVEN(t *testing.T) {
	state, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))
	pm := plugin.NewPM(state)

	senderName := string("jacobwolf12345")
	senderPubkey := common.HexToPubKey("12345")

	if err := createAccount(pm, "jacobwolf12345"); err != nil {
		return
	}

	if err := createAccount(pm, "denverfolk12345"); err != nil {
		return
	}

	if err := createAccount(pm, "fractal.asset"); err != nil {
		return
	}

	if _, err := pm.IssueAsset(senderName, "bitcoin", "btc", big.NewInt(1000000000000000000), 10, senderName, senderName, big.NewInt(2000000000000000000), "", pm); err != nil {
		fmt.Println("issue asset error", err)
		return
	}

	runtimeConfig := runtime.Config{
		Origin:      senderName,
		FromPubkey:  senderPubkey,
		State:       state,
		PM:          pm,
		AssetID:     0,
		GasLimit:    10000000000,
		GasPrice:    big.NewInt(0),
		Value:       big.NewInt(0),
		BlockNumber: new(big.Int).SetUint64(0),
	}

	VenBinfile := "./runtime/contract/Ven/VEN.bin"
	VenAbifile := "./runtime/contract/Ven/VEN.abi"
	VenSaleBinfile := "./runtime/contract/Ven/VENSale.bin"
	VenSaleAbifile := "./runtime/contract/Ven/VENSale.abi"
	venContractName := string("vencontract12345")
	venSaleContractName := string("vensalevontract")

	if err := createAccount(pm, "vencontract12345"); err != nil {
		return
	}

	if err := createAccount(pm, "vensalevontract"); err != nil {
		return
	}

	if err := createAccount(pm, "ethvault12345"); err != nil {
		return
	}

	if err := createAccount(pm, "venvault12345"); err != nil {
		return
	}

	err := createContract(VenSaleAbifile, VenSaleBinfile, venSaleContractName, runtimeConfig)
	if err != nil {
		fmt.Println("create venSaleContractAddress error")
		return
	}

	err = createContract(VenAbifile, VenBinfile, venContractName, runtimeConfig)
	if err != nil {
		fmt.Println("create venContractAddress error")
		return
	}

	pm.AddBalanceByID(venContractName, 0, big.NewInt(1))
	pm.AddBalanceByID(venSaleContractName, 0, big.NewInt(1))

	setVenOwnerInput, err := input(VenAbifile, "setOwner", common.BytesToAddress([]byte("vensalevontract")))
	if err != nil {
		fmt.Println("initializeVenSaleInput error ", err)
		return
	}

	action := types.NewAction(types.CallContract, runtimeConfig.Origin, venContractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, setVenOwnerInput, nil)

	_, _, err = runtime.Call(action, &runtimeConfig)
	if err != nil {
		fmt.Println("call set ven owner error ", err)
		return
	}

	initializeVenSaleInput, err := input(VenSaleAbifile, "initialize", common.BytesToAddress([]byte("vencontract12345")), common.BytesToAddress([]byte("ethvault12345")), common.BytesToAddress([]byte("venvault12345")))
	if err != nil {
		fmt.Println("initializeVenSaleInput error ", err)
		return
	}
	action = types.NewAction(types.CallContract, runtimeConfig.Origin, venSaleContractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, initializeVenSaleInput, nil)
	runtimeConfig.Time = big.NewInt(1504180700)

	_, _, err = runtime.Call(action, &runtimeConfig)
	if err != nil {
		fmt.Println("call initialize vensale error ", err)
		return
	}

	runtimeConfig.Value = big.NewInt(100000000000000000)
	action = types.NewAction(types.CallContract, runtimeConfig.Origin, venSaleContractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, nil, nil)

	fmt.Println("start")
	_, _, err = runtime.Call(action, &runtimeConfig)
	fmt.Println("finish")
	if err != nil {
		fmt.Println("call buy ven sale error ", err)
		return
	}

	runtimeConfig.Value = big.NewInt(0)
	getBalanceInput, err := input(VenAbifile, "balanceOf", common.BytesToAddress([]byte("jacobwolf12345")))
	if err != nil {
		fmt.Println("getBalanceInput error ", err)
		return
	}
	action = types.NewAction(types.CallContract, runtimeConfig.Origin, venContractName, 0, runtimeConfig.AssetID, runtimeConfig.GasLimit, runtimeConfig.Value, getBalanceInput, nil)

	ret, _, err := runtime.Call(action, &runtimeConfig)
	if err != nil {
		fmt.Println("call get balance error ", err)
		return
	}

	num := new(big.Int).SetBytes(ret)
	assert.Equal(t, num, new(big.Int).Mul(big.NewInt(3500000000), big.NewInt(100000000000)))
}
