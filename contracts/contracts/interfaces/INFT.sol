// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

interface INFT {
    function mintAndTransfer(string memory uri, address owner) external returns (uint256);
}
