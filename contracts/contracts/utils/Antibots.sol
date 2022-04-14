// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Antibots is AccessControl {
    bool public antibotsEnabled = true;
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");
    bytes32 public constant WHITELISTED_ROLE = keccak256("WHITELISTED_ROLE");

    event AntibotsDisabled();
    event TransferBurned(address indexed wallet, uint256 amount);

    constructor() {
        antibotsEnabled = true;
    }

    function isAntibotsEnabled() internal view returns (bool) {
        if (hasRole(OPERATOR_ROLE, msg.sender)) {
            return false;
        }
        if (hasRole(WHITELISTED_ROLE, msg.sender)) {
            return false;
        }
        return antibotsEnabled;
    }

    function disableAntibots() external onlyRole(OPERATOR_ROLE) {
        require(!antibotsEnabled, "Antibots: anti bots has been disabled");
        antibotsEnabled = false;

        emit AntibotsDisabled();
    }

    function addWhiteListed(address addr) external onlyRole(OPERATOR_ROLE) {
        grantRole(WHITELISTED_ROLE, addr);
    }

    function dropWhiteListed(address addr) external onlyRole(OPERATOR_ROLE) {
        revokeRole(WHITELISTED_ROLE, addr);
    }

    function setOperator(address addr) internal {
        grantRole(OPERATOR_ROLE, addr);
    }
}
