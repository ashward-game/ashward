// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

library TypeConversion {
    function uint256ToBytes(uint256 x) internal pure returns (bytes memory b) {
        b = new bytes(32);
        assembly {
            mstore(add(b, 32), x)
        }
    }

    function bytesToUint256(bytes memory b) internal pure returns (uint256) {
        require(
            b.length == 32,
            "TypeConversion: bytesToUint256 requires input of type bytes with length 32"
        );
        uint256 tempUint;
        assembly {
            tempUint := mload(add(add(b, 0x20), 0))
        }
        return tempUint;
    }

    // https://docs.soliditylang.org/en/v0.5.3/types.html#address
    function bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(bytes20(b)));
    }
}
