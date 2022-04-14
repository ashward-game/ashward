// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "./helpers/TypeConversion.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract MakeRand {
    bytes32 public immutable publicKey;

    /**
     * @dev `key` serverHash
     * @dev `value` clientRandom
     */
    mapping(bytes32 => bytes32) internal _commitments;

    constructor(bytes32 _publicKey) {
        publicKey = _publicKey;
    }

    function _commit(
        bytes32 _serverHash,
        bytes memory _signature,
        bytes32 _clientRandom
    ) internal isFreshCommit(_serverHash) verified(_serverHash, _signature) {
        require(
            _clientRandom != 0,
            "MakeRand: client random must be not equal 0"
        );
        _commitments[_serverHash] = _clientRandom;
    }

    modifier isFreshCommit(bytes32 hashBytes) {
        require(
            _commitments[hashBytes] == 0,
            "MakeRand: hash value already exists"
        );
        _;
    }

    modifier verified(bytes32 hashBytes, bytes memory signature) {
        require(
            ECDSA.recover(hashBytes, signature) ==
                TypeConversion.bytes32ToAddress(publicKey),
            "MakeRand: signature is invalid"
        );
        _;
    }
}
