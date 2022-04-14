// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "../MakeRand.sol";

// mock class using MakeRand
contract MakeRandMock is MakeRand {
    event Committed(bytes32 serverHash, bytes32 clientRandom);

    constructor(bytes32 _publicKey) MakeRand(_publicKey) {}

    function commit(
        bytes32 _serverHash,
        bytes memory _signature,
        bytes32 _clientRandom
    ) external {
        _commit(_serverHash, _signature, _clientRandom);
        emit Committed(_serverHash, _clientRandom);
    }

    function getClientRandom(bytes32 hashBytes)
        external
        view
        returns (bytes32)
    {
        require(
            _commitments[hashBytes] > 0,
            "MakeRandMock: server hash does not exist"
        );
        return _commitments[hashBytes];
    }
}
