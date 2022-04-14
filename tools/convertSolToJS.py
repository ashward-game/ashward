#!/usr/bin/env python3
import errno
import os.path
from os.path import isfile, join, splitext
from os import listdir
import json
from typing import Tuple

contractsDir = "../contracts/contracts/"
jsonDir = "../contracts/build/contracts/"
outDir = "../web/plugins/contracts/"
commonDir = "../common/"
ADDRESS = "address.json"


def is_contract(f):
    return isfile(f) and f.endswith('.sol')


if not os.path.exists(os.path.dirname(outDir)):
    try:
        os.makedirs(os.path.dirname(outDir))
    except OSError as exc:  # Guard against race condition
        if exc.errno != errno.EEXIST:
            raise

contracts = [splitext(f)[0] for f in listdir(
    contractsDir) if is_contract(join(contractsDir, f))]


def cap(s):
    return ''.join((c.upper() if i == 0 or s[i-1] == ' ' else c) for i, c in enumerate(s))


def generateHeader():
    return """const contract = require('./address.js');

"""


def generateAddress(name):
    s = ""
    with open(join(commonDir, ADDRESS)) as addressFile:
        address = json.load(addressFile)
        if name in address:
            s += "export const ADDRESS_{} = contract.loadContractAddress(\'{}\');\n\n".format(
                name.upper(), name)
    return s


def generateABI(name, abi):
    return "export const ABI_{} = {};".format(name.upper(), abi)


for contract in contracts:
    print("[+] Generating for {}.sol... ".format(contract), end="")
    with open(join(jsonDir, contract + ".json")) as compiled_file:
        data = json.load(compiled_file)
        os.makedirs(os.path.dirname(outDir), exist_ok=True)
        with open(join(outDir, contract + ".js"), "w") as f:
            s = generateHeader()
            s += generateAddress(contract)

            # generate ABI
            abi = json.dumps(data['abi'])
            s += generateABI(contract, abi)
            s += "\n"
            f.write(s)
    print("Done")
