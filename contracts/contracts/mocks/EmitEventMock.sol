// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

contract EmitEventMock {
    event TestIndexed(
        address indexed user,
        bytes32 indexed hash,
        uint256 indexed id
    );

    event TestData(address user, uint256 amount, bool value, bytes data);

    event TestUint8(uint8 indexed value);

    bytes32 public constant sampleHash = keccak256("HASH");

    constructor() {}

    function test() public {
        emit TestIndexed(address(0x010203040506070809), sampleHash, 1234567890);
        emit TestData(
            address(0x010203040506070809),
            1234567890,
            true,
            abi.encodePacked(sampleHash)
        );
        emit TestUint8(13);
    }
}
