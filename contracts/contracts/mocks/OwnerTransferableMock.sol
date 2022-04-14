// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "../utils/OwnerTransferable.sol";

contract OwnerTransferableMock is OwnerTransferable {
    constructor() {}

    function send() public payable {
        require(msg.value > 0, "balance must be greater than 0");
    }
}
