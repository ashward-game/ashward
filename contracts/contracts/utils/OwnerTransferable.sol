// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";

abstract contract OwnerTransferable is Ownable {
    function transferToOwner() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "OwnerTransferable: empty balance");

        address payable owner = payable(owner());
        (bool success, ) = owner.call{value: balance}("");
        require(success);
    }
}
