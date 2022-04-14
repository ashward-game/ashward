#!/usr/bin/env python3
import errno
import os.path
from os.path import isfile, join, splitext
from os import listdir
import json
from typing import Tuple
import shutil

contractsDir = "../contracts/contracts/"
jsonDir = "../contracts/build/contracts/"
outABIDir = "../backend/contract/abi/"
outEventDir = "../backend/contract/event/"
addressFile = "../common/address.json"

outBinDir = "../backend/testutil/contract/"


typeConvert = {
    "address": "common.Address",
    "uint256": "*big.Int",
    "bool": "bool",
    "string": "string",
    "bytes32": "common.Hash",
    "bytes": "[]byte",
    "uint8": "uint8"
}


def is_deployed_contract(f):
    name = os.path.basename(f)
    base = os.path.splitext(name)[0]
    data = []
    with open(addressFile) as file:
        data = json.load(file)
    return isfile(f) and f.endswith('.sol') and name != "Migrations.sol" and (base in data)


def rmMkDir(dirName):
    if os.path.exists(os.path.dirname(dirName)):
        shutil.rmtree(os.path.dirname(dirName))
    try:
        os.makedirs(os.path.dirname(dirName))
    except OSError as exc:  # Guard against race condition
        if exc.errno != errno.EEXIST:
            raise


def cap(s):
    return ''.join((c.upper() if i == 0 or s[i-1] == ' ' else c) for i, c in enumerate(s))


def generateImport(contract):
    return """
import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/{}"
    orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

""".format(contract.lower())


def generateLogNameSig(item):
    s = "const Log{}Name string = \"{}\"\n".format(
        item['name'], item['name'])
    flag = True
    s += "const Log{}Sig string = \"{}".format(
        item['name'], item['name'])
    s += "("
    for input in item['inputs']:
        if flag:
            s += input['type']
            flag = False
        else:
            s += ",{}".format(input['type'])
    s += ")\"\n\n"
    return s


def generateLogStruct(name, info):
    s = "type Log{} struct {{\n".format(name)
    for field in info['inputs']:
        s += "\t{} {}".format(cap(field['name']),
                              typeConvert[field['type']])
        s += "\n"
    s += "}\n"
    return s


def generateParseFunction(name, info):
    s = """
func ParseLog{}(contractAbi *abi.ABI, vLog *types.Log) (*Log{}, error) {{
	var evt Log{}
	if len(vLog.Data) > 0 {{
		err := contractAbi.UnpackIntoInterface(&evt, Log{}Name, vLog.Data)
		if err != nil {{
			return nil, err
		}}
	}}
""".format(name, name, name, name)

    # generate topics, ie, indexed fields
    count = 1
    for field in info['inputs']:
        if field['indexed'] and field['type'] == 'address':
            s += """
	evt.{} = common.HexToAddress(vLog.Topics[{}].Hex())
""".format(cap(field['name']), count)
            count += 1
        elif field['indexed'] and field['type'] == 'bytes32':
            s += """
	evt.{} = common.HexToHash(vLog.Topics[{}].Hex())
""".format(cap(field['name']), count)
            count += 1
        elif field['indexed'] and field['type'] == 'uint256':
            s += """
	evt.{} = util.HexToBigInt(vLog.Topics[{}].Hex())
""".format(cap(field['name']), count)
        elif field['indexed'] and field['type'] == 'uint8':
            s += """
    evt.{} = uint8(util.HexToBigInt(vLog.Topics[{}].Hex()).Uint64())
    """.format(cap(field['name']), count)
        elif field['indexed']:
            raise Exception(
                'Unsupported indexed type: {}'.format(field['type']))

    s += """
	return &evt, nil
}

"""
    return s


def generateSwitchCaseLog(name):
    s = """\t\t\tcase Log{}Name:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*Log{}, error))(&contractAbi, vLog)
				if err != nil {{
					return err
				}}
				return handler.(func(context.Context, *Log{}) error)(ctx, evt)
""".format(name, name, name)
    return s


def generateParserHandlerStruct(name, logNames):
    s = """type parserHandler struct {{
	name string
}}

var _ event.ParserHandler = (*parserHandler)(nil)

func (p *parserHandler) Name() string {{
	return p.name
}}

func (p *parserHandler) ParseHandle(ctx context.Context, vLog *types.Log) error {{
	contractAbi, err := abi.JSON(strings.NewReader(string({}.ABI)))
	if err != nil {{
		return err
	}}

    ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {{
		if k == vLog.Topics[0].Hex() {{
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {{
""".format(name.lower())
    for log in logNames:
        s += generateSwitchCaseLog(log)
    s += """\t\t\tdefault:
				return errors.New("{}: event type is not supported")
\t\t\t}}
		}}
	}}
	return errors.New("{}: topic is not supported")
}}

""".format(name.lower(), name.lower())
    return s


def generateHashMap(name):
    s = """
	log{}SigHash := crypto.Keccak256Hash([]byte(Log{}Sig))
	hashMap[log{}SigHash.Hex()] = Log{}Name
""".format(name, name, name, name)
    return s


def generateInit(contract, names, hashMap):
    s = """var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
"""
    for name in names:
        s += "\tlogParserMap[Log{}Name] = ParseLog{}\n".format(name, name)
    s += "\n"
    for name in names:
        s += "\tlogHandlerMap[Log{}Name] = HandleLog{}\n".format(name, name)
    s += hashMap
    s += """
	event.Register(&parserHandler{{
		name: {}.Name,
	}})
""".format(contract.lower())
    s += "}"
    return s


def generateParser(contracts):
    print("[+] Generating Parser ...")
    for contract in contracts:
        print("[++] Generating for {}.sol... ".format(contract), end="")
        with open(join(jsonDir, contract + ".json")) as compiled_file:
            data = json.load(compiled_file)
            rmMkDir(join(outEventDir, contract.lower()) + "/")
            with open(join(outEventDir, contract.lower(), "parser.go"), "w") as f:
                s = "package {}\n".format(contract.lower())
                s += generateImport(contract)

                # generate signature for events and event's names
                names = []
                hashMap = ""
                for item in data['abi']:
                    if item['type'] == "event":
                        s += generateLogNameSig(item)
                        s += generateLogStruct(item['name'], item)
                        s += generateParseFunction(item['name'], item)
                        names.append(item['name'])
                        hashMap += generateHashMap(item['name'])
                s += generateParserHandlerStruct(contract, names)
                s += generateInit(contract, names, hashMap)
                s += "\n"
                f.write(s)
        print("Done")
    print("[+] Done.")


def generateHandleFunction(name, info):
    s = """
func HandleLog{}(ctx context.Context, evt *Log{}) error {{
	return nil
}}
""".format(name, name)
    return s


def generateHandler(contracts):
    print("[+] Generating Handler ...")
    for contract in contracts:
        print("[++] Generating for {}.sol... ".format(contract), end="")
        with open(join(jsonDir, contract + ".json")) as compiled_file:
            data = json.load(compiled_file)
            with open(join(outEventDir, contract.lower(), "handler.go"), "w") as f:
                s = "package {}\n\n".format(contract.lower())
                s += 'import "context"\n'
                for item in data['abi']:
                    if item['type'] == "event":
                        s += generateHandleFunction(item['name'], item)
                f.write(s)
            print("Done")
    print("[+] Done.")


def generateImportFile():
    print("[+] Generating import.go... ", end="")
    with open(join("../backend/import.go"), "w") as f:
        s = "package main\n\n"
        s += "import (\n"
        for contract in sorted(contracts):
            s += "\t_ \"orbit_nft/contract/event/{}\"\n".format(
                contract.lower())
        s += ")\n"
        f.write(s)
    print("Done")


def generateABI(contracts):
    print("[+] Generating ABI ...")
    for contract in contracts:
        print("[++] Generating for {}.sol... ".format(contract), end="")
        rmMkDir(join(outABIDir, contract.lower()) + "/")
        with open(join(jsonDir, contract + ".json")) as compiled_file:
            data = json.load(compiled_file)
            with open(join(outABIDir, contract.lower(), contract + ".go"), "w") as f:
                s = "package {}\n\n".format(contract.lower())
                abi = json.dumps(data['abi']).replace('"', '\\"')
                s += "var ABI = \"{}\"\n\n".format(abi)
                s += "const Name = \"{}\"".format(contract)
                s += "\n"
                f.write(s)
        print("Done")
    print("[+] Done.")


def generateBin(contracts):
    print("[+] Generating Bin ...")
    for contract in contracts:
        print("[++] Generating for {}.sol... ".format(contract), end="")
        with open(join(jsonDir, contract + ".json")) as compiled_file:
            data = json.load(compiled_file)
            with open(join(outBinDir, contract + ".go"), "w") as f:
                s = "package contract\n\n"
                s += "import \"github.com/ethereum/go-ethereum/accounts/abi/bind\"\n\n"
                bytecode = json.dumps(data['bytecode'])
                abi = json.dumps(data['abi']).replace('"', '\\"')
                s += """var {}MetaData = &bind.MetaData{{
\tABI: \"{}\",
\tBin: {},
}}
""".format(
                    contract, abi, bytecode)
                f.write(s)
        print("Done")
    print("[+] Done.")


if __name__ == "__main__":
    rmMkDir(outABIDir)
    rmMkDir(outBinDir)

    contracts = [splitext(f)[0] for f in listdir(
        contractsDir) if is_deployed_contract(join(contractsDir, f))]

    generateImportFile()
    generateABI(contracts)
    generateParser(contracts)
    generateHandler(contracts)
    generateBin(contracts)
