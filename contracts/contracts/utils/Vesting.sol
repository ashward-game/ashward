// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * Vesting base contract.
 * Each vesting instance has its own schedule, specifying its milestones.
 * Milestone's index starts from 1 (corresponding to the first month).
 */
abstract contract Vesting is AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    enum Status {
        NonEmployed,
        Employed
    }

    struct Beneficiary {
        uint256 totalAmount; // total amount of vesting tokens
        uint256 lastMilestone; // total claimed milestones
        bool tgeClaimed;
        Status status; // current status
    }

    // timing constant
    uint256 public constant TGE_MILESTONE = 1647621000; // 18-03-2022 16:30:00 UTC

    bytes32 private constant GRANTOR_ROLE = keccak256("GRANTOR_ROLE");
    bytes32 private constant BENEFICIARY_ROLE = keccak256("BENEFICIARY_ROLE");

    // percent constant
    uint256 private DENOMINATOR = 10000; // e.g., 8.33% = 833/DENOMINATOR

    IERC20 private immutable _token;

    mapping(address => Beneficiary) internal _beneficiaries;

    event BeneficiaryRegistered(address indexed beneficiary, uint256 amount);
    event BeneficiarySuspended(address indexed beneficiary);
    event Claimed(address indexed beneficiary, uint256 amount);
    event TGEClaimed(address indexed beneficiary, uint256 amount);

    constructor(address token) {
        require(token != address(0), "Vesting: token address must not be 0");
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(GRANTOR_ROLE, msg.sender);
        _token = IERC20(token);
    }

    function _addBeneficiaries(address beneficiary, uint256 totalAmount)
        internal
    {
        require(
            beneficiary != address(0),
            "Vesting: beneficiary address must not be 0"
        );
        require(
            totalAmount > 0,
            "Vesting: total amount must be greater than 0"
        );
        require(
            _beneficiaries[beneficiary].status == Status.NonEmployed,
            "Vesting: beneficiary is already in pool"
        );

        _beneficiaries[beneficiary] = Beneficiary(
            totalAmount,
            0,
            false,
            Status.Employed
        );
        _grantRole(BENEFICIARY_ROLE, beneficiary);
        emit BeneficiaryRegistered(beneficiary, totalAmount);
    }

    function addBeneficiaries(
        address[] memory beneficiaries,
        uint256[] memory amounts
    ) external onlyRole(GRANTOR_ROLE) {
        require(
            beneficiaries.length == amounts.length,
            "Vesting: beneficiaries and amounts' length should be equal"
        );

        for (uint256 i = 0; i < beneficiaries.length; i++) {
            _addBeneficiaries(beneficiaries[i], amounts[i]);
        }
    }

    function suspendBeneficiary(address beneficiary)
        external
        onlyRole(GRANTOR_ROLE)
    {
        require(
            _beneficiaries[beneficiary].status == Status.Employed,
            "Vesting: beneficiary must be employed"
        );
        _revokeRole(BENEFICIARY_ROLE, beneficiary);

        emit BeneficiarySuspended(beneficiary);
    }

    function vestingOf(address beneficiary)
        external
        view
        returns (uint256, uint256)
    {
        require(
            hasRole(BENEFICIARY_ROLE, beneficiary),
            "Vesting: beneficiary is not in pool"
        );
        return (
            _beneficiaries[beneficiary].totalAmount,
            _beneficiaries[beneficiary].lastMilestone
        );
    }

    function _tgePercent() internal virtual returns (uint256);

    function _percent(uint256) internal virtual returns (uint256, uint256);

    function claim() external onlyRole(BENEFICIARY_ROLE) nonReentrant {
        uint256 lastMilestone = _beneficiaries[msg.sender].lastMilestone;
        (uint256 percent, uint256 currentMilestone) = _percent(lastMilestone);
        require(
            percent > 0 && currentMilestone > lastMilestone,
            "Vesting: cannot unlock tokens for this milestone or already claimed tokens for current milestone"
        );

        uint256 transferAmount = (_beneficiaries[msg.sender].totalAmount *
            percent) / DENOMINATOR;
        _beneficiaries[msg.sender].lastMilestone = currentMilestone;
        _token.safeTransfer(msg.sender, transferAmount);

        emit Claimed(msg.sender, transferAmount);
    }

    function claimTGE() external onlyRole(BENEFICIARY_ROLE) nonReentrant {
        uint256 percent = _tgePercent();
        require(percent > 0, "Vesting: no TGE tokens");

        require(
            block.timestamp >= TGE_MILESTONE,
            "Vesting: need to wait for 30 minutes before unlocking TGE tokens"
        );
        require(
            _beneficiaries[msg.sender].tgeClaimed == false,
            "Vesting: already claimed all TGE tokens"
        );

        _beneficiaries[msg.sender].tgeClaimed = true;

        uint256 transferAmount = (_beneficiaries[msg.sender].totalAmount *
            percent) / DENOMINATOR;
        _token.safeTransfer(msg.sender, transferAmount);

        emit TGEClaimed(msg.sender, transferAmount);
    }

    function hasClaimedTGE(address addr) external view returns (bool) {
        require(
            hasRole(BENEFICIARY_ROLE, addr),
            "Vesting: address is not a beneficiary"
        );
        return _beneficiaries[addr].tgeClaimed;
    }

    function setGrantor(address _addr) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(GRANTOR_ROLE, _addr);
    }

    function removeGrantor(address _addr)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _revokeRole(GRANTOR_ROLE, _addr);
    }

    function collectToken() external onlyRole(GRANTOR_ROLE) {
        uint256 balance = _token.balanceOf(address(this));
        require(balance > 0, "Vesting: current balance is zero");
        _token.safeTransfer(msg.sender, balance);
    }
}
