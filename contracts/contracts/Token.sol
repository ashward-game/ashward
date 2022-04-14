// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "erc-payable-token/contracts/token/ERC1363/ERC1363.sol";

contract Token is ERC1363, Pausable, AccessControl {
    uint256 public constant cap = 1_000_000_000 * 1e18;

    // address of wallet that has right to collect taxes on transfer transactions
    address public _governor;
    uint256 public _tax;

    uint256 public constant PRECISION = 10000;

    // timing & sell tax
    uint256 public constant TGE_PLUS_1_HOUR = 1647622800; // 2022-03-18 17:00:00UTC
    uint256 public constant TAX_TGE = 800; // 8%
    uint256 public constant TAX_NORMAL = 200; // 2% = 200 / PRECISION

    bytes32 public constant FREETAX_ROLE = keccak256("FREETAX_ROLE");

    // an address is selling_role if it receives Token and transfers some other tokens back to the sender
    // for example when a user swaps Token for BNB on pancakeswap
    // example of selling_role: pancakeswap V2 contract's address
    bytes32 public constant SELLING_ROLE = keccak256("SELLING_ROLE");

    bool private isTransferable;

    event TaxUpdated(uint256 tax);

    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(FREETAX_ROLE, msg.sender);

        _governor = msg.sender;
        _tax = TAX_NORMAL;

        isTransferable = false;

        _mint(msg.sender, cap);
    }

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) whenNotPaused {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) whenPaused {
        _unpause();
    }

    function setGovernor(address addr) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _governor = addr;
    }

    function setTax(uint256 tax) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _tax = tax;
        emit TaxUpdated(_tax);
    }

    function addSellingAddress(address addr)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        isTransferable = true;
        grantRole(SELLING_ROLE, addr);
    }

    function removeSellingAddress(address addr)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        revokeRole(SELLING_ROLE, addr);
    }

    function addNoTaxAddress(address addr)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        grantRole(FREETAX_ROLE, addr);
    }

    function removeNoTaxAddress(address addr)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        revokeRole(FREETAX_ROLE, addr);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }

    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override {
        if (hasRole(FREETAX_ROLE, from)) {
            super._transfer(from, to, amount);
            return;
        }
        require(isTransferable, "Token: cannot yet transferring");
        if (hasRole(SELLING_ROLE, to) && _governor != address(0x00)) {
            uint256 taxAmount = (amount * _taxCalculation()) / PRECISION;
            amount = amount - taxAmount;

            super._transfer(from, _governor, taxAmount);
            super._transfer(from, to, amount);
            return;
        }

        super._transfer(from, to, amount);
    }

    function _taxCalculation() private view returns (uint256) {
        if (block.timestamp <= TGE_PLUS_1_HOUR) {
            return TAX_TGE;
        }
        return _tax;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC1363, AccessControl)
        returns (bool)
    {
        return
            interfaceId == type(IERC1363).interfaceId ||
            interfaceId == type(IAccessControl).interfaceId ||
            interfaceId == type(IERC165).interfaceId ||
            super.supportsInterface(interfaceId);
    }
}
