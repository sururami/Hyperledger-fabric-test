/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the fido structure, with 4 properties.  Structure tags are used by encoding/json library
type Fido struct {
/*
 *	Make   string `json:"make"`
 *	Model  string `json:"model"`
 *	Colour string `json:"colour"`
 *	Owner  string `json:"owner"`
 */
	
/*
 *	BcID                    string `json:"bcID"`
 *	AaguID                  string `json:"aaguID"`
 *	CredentialID            string `json:"credentialID"`
 *	RegResp                 string `json:"regResp"`
 *	SignCounter             string `json:"signCounter"`
 *	RegistrationTime        string `json:"registrationTime"`
 *	LastAuthenticationTime  string `json:"lastAuthenticationTime"`
 */
 
	BcID                    string `json:"bcID"`
 	AaguID                  string `json:"aaguID"`
 	CredentialID            string `json:"credentialID"`
 	RegResp                 string `json:"regResp"`
}

/*
 * The Init method is called when the Smart Contract "fabfido" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabfido"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryFido" {
		return s.queryFido(APIstub, args)
	} else if function == "initFidoLedger" {
		return s.initFidoLedger(APIstub)
	} else if function == "createFido" {
		return s.createFido(APIstub, args)
	} else if function == "queryAllFidos" {
		return s.queryAllFidos(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryFido(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fidoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(fidoAsBytes)
}

func (s *SmartContract) initFidoLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	fidos := []Fido{
	/*
	 *	Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
	 *	Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
	 *	Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
	 *	Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
	 *	Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
	 *	Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
	 *	Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
	 *	Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
	 *	Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
	 *	Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	 */
		
		Fido{BcID: "aaa", AaguID: "aa", CredentialID: "credentialID1", RegResp: "a"},
	 	Fido{BcID: "bbb", AaguID: "bb", CredentialID: "credentialID2", RegResp: "b"},
	 	Fido{BcID: "ccc", AaguID: "cc", CredentialID: "credentialID3", RegResp: "c"},
	 	Fido{BcID: "ddd", AaguID: "dd", CredentialID: "credentialID4", RegResp: "d"},
	 	Fido{BcID: "eee", AaguID: "ee", CredentialID: "credentialID5", RegResp: "e"},
	 	Fido{BcID: "fff", AaguID: "ff", CredentialID: "credentialID6", RegResp: "f"},
	 	Fido{BcID: "ggg", AaguID: "gg", CredentialID: "credentialID7", RegResp: "g"},
	 	Fido{BcID: "hhh", AaguID: "hh", CredentialID: "credentialID8", RegResp: "h"},
	 	Fido{BcID: "iii", AaguID: "ii", CredentialID: "credentialID9", RegResp: "i"},
	 	Fido{BcID: "jjj", AaguID: "jj", CredentialID: "credentialID10", RegResp: "j"},
	
	/*
	 *	Fido{BcID: "aaa", AaguID: "aa", CredentialID: "credentialID1", RegResp: "a", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "bbb", AaguID: "bb", CredentialID: "credentialID2", RegResp: "b", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "ccc", AaguID: "cc", CredentialID: "credentialID3", RegResp: "c", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "ddd", AaguID: "dd", CredentialID: "credentialID4", RegResp: "d", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "eee", AaguID: "ee", CredentialID: "credentialID5", RegResp: "e", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "fff", AaguID: "ff", CredentialID: "credentialID6", RegResp: "f", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "ggg", AaguID: "gg", CredentialID: "credentialID7", RegResp: "g", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "hhh", AaguID: "hh", CredentialID: "credentialID8", RegResp: "h", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "iii", AaguID: "ii", CredentialID: "credentialID9", RegResp: "i", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 *	Fido{BcID: "jjj", AaguID: "jj", CredentialID: "credentialID10", RegResp: "j", SignCounter: "1", RegistrationTime: "123", LastAuthenticationTime: "123"},
	 */
	}

	i := 0
	for i < len(fidos) {
		fmt.Println("i is ", i)
		fidoAsBytes, _ := json.Marshal(fidos[i])
		APIstub.PutState("FIDO"+strconv.Itoa(i), fidoAsBytes)
		fmt.Println("Added", fidos[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createFido(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	// var fido = Fido{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
	var fido = Fido{BcID: args[1], AaguID: args[2], CredentialID: args[3], RegResp: args[4], SignCounter: args[5], RegistrationTime: args[6], LastAuthenticationTime: args[7]}

	fidoAsBytes, _ := json.Marshal(fido)
	APIstub.PutState(args[0], fidoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllFidos(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "FIDO0"
	endKey := "FIDO999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllFidos:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}



// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
