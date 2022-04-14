// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "erc-payable-token/contracts/payment/ERC1363Payable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract StakingRewards is
    AccessControl,
    ERC1363Payable,
    ReentrancyGuard,
    Pausable
{
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    // percent constant
    /**
     * Used math percent of date
     * e.g., 8.33% = 833/DENOMINATOR_DAY
     */
    uint256 internal DENOMINATOR_DAY = 1e4;

    /**
     * Use math percent of seconds
     * e.g., 0.000083333% = 83333/DENOMINATOR_SECONDS
     */
    uint256 internal DENOMINATOR_SECONDS = 1e11;

    /**
     * Stake struct is used to represent the way we store stakes.
     * A stake will contain the amount staked and a timestamp, since which is
     * when the stake was made (or the last time when the rewards was withdrawn).
     */
    struct Stake {
        uint256 amount;
        uint256 since;
    }

    /**
     * This mapping is used to keep track of the index for the stakers.
     */
    mapping(address => Stake[]) internal _stakes;

    /**
     * This mapping was used to save last index `Stake[]`
     */
    mapping(address => uint256) internal _lastIndex;

    /**
     * This mapping keeps track the total stakes of all stakers.
     * Its main purpose is to avoid iterating the whole array `_stakes` in `Stakeholders` when a staker wants to see all token staked.
     * @dev
     * Always use this mapping with value > 0 when checking the existence of a staker.
     */
    mapping(address => uint256) internal _stakesSummary;

    /**
     * Metadata: current number of staked tokens of all stakers.
     * `totalStalked` is increased when a `stake` call succeeded, and
     * decreased when a `withdraw` or `exit` call succeeded.
     */
    uint256 public totalStaking;

    /**
     * Metadata: current number of stakers.
     */
    uint256 public totalStakers;

    /**
     * Min token amount each stake
     */
    uint256 public immutable minStake;

    /**
     * Max total token amount of staker staked
     */
    uint256 public immutable maxStake;

    /**
     * Percent reward daily
     */
    uint256 public immutable percentPerDay;

    /**
     * Percent reward each seconds
     */
    uint256 public immutable percentPerSecond;

    /**
     * Limit total token staked in pool
     */
    uint256 public immutable totalTokenInPool;

    /**
     * Date start event
     */
    uint256 public immutable startDate;

    /**
     * Date end event
     */
    uint256 public immutable endDate;

    event Staked(address indexed staker, uint256 amount);
    event Rewarded(
        address indexed staker,
        uint256 stakeAmount,
        uint256 rewardsAmount
    );

    constructor(
        IERC1363 acceptedToken_,
        uint256 minStake_,
        uint256 maxStake_,
        uint256 percentPerDay_,
        uint256 percentPerSecond_,
        uint256 totalTokenInPool_,
        uint256 startDate_,
        uint256 endDate_
    ) ERC1363Payable(acceptedToken_) {
        require(
            minStake_ <= maxStake_,
            "StakingRewards: min stake must be less than or equal max stake"
        );
        require(
            totalTokenInPool_ > 0,
            "StakingRewards: total token in pool must be greater than 0"
        );
        require(
            endDate_ > startDate_,
            "StakingRewards: end date must be after start date"
        );
        minStake = minStake_;
        maxStake = maxStake_;
        percentPerDay = percentPerDay_;
        percentPerSecond = percentPerSecond_;
        totalTokenInPool = totalTokenInPool_;
        startDate = startDate_;
        endDate = endDate_;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function stakingToken() public view returns (IERC20) {
        return IERC20(acceptedToken());
    }

    /**
     * Pause pool
     */
    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    /**
     * Unpause pool
     */
    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    /**
     * @dev This method is called after `onTransferReceived`.
     *  Note: remember that the token contract address is always the message sender.
     * @param operator The address which called `transferAndCall` or `transferFromAndCall` function
     * @param sender Address performing the token purchase
     * @param amount The amount of tokens transferred
     * @param data Additional data with no specified format
     */
    function _transferReceived(
        address operator,
        address sender,
        uint256 amount,
        bytes memory data
    ) internal override {}

    /**
     * @dev This method is called after `onApprovalReceived`.
     *  Note: remember that the token contract address is always the message sender.
     * @param sender address The address which called `approveAndCall` function
     * @param amount uint256 The amount of tokens to be spent
     */
    function _approvalReceived(
        address sender,
        uint256 amount,
        bytes memory
    ) internal override {
        _stake(sender, amount);
    }

    function _beforeStake(address staker, uint256 amount) internal virtual {
        require(
            amount >= minStake,
            "StakingRewards: insufficient stake amount"
        );
        require(
            _stakesSummary[staker] + amount <= maxStake,
            "StakingRewards: overflowing stake amount"
        );
        require(
            totalStaking + amount <= totalTokenInPool,
            "StakingRewards: pool is full"
        );
        require(
            block.timestamp <= endDate,
            "StakingRewards: staking time is over"
        );
    }

    /**
     * Do not accept stakes when the staking program is over.
     */
    function _stake(address sender, uint256 amount)
        internal
        isPoolStarted
        whenNotPaused
        nonReentrant
    {
        _beforeStake(sender, amount);

        uint256 allowance = stakingToken().allowance(sender, address(this));

        require(allowance >= amount, "StakingRewards: insufficient allowance");

        if (_stakesSummary[sender] == 0) {
            totalStakers++;
        }
        _stakes[sender].push(Stake(amount, block.timestamp));
        _stakesSummary[sender] += amount;
        totalStaking += amount;

        stakingToken().safeTransferFrom(sender, address(this), amount);
        emit Staked(sender, amount);
    }

    /**
     * Calculate the amount of rewards for a Stake `stake_` since last time the staker withdrawn his rewards.
     */
    function _calculateStakeRewards(Stake memory stake_)
        internal
        view
        virtual
        returns (uint256)
    {
        uint256 timestamp = endDate;
        if (block.timestamp < timestamp) {
            timestamp = block.timestamp;
        }

        uint256 rewards;

        // calculator date
        uint256 numDays = (timestamp - stake_.since) / 1 days;
        if (numDays > 0) {
            rewards =
                (numDays * percentPerDay * stake_.amount) /
                DENOMINATOR_DAY;
        }

        //calculator seconds
        uint256 numSeconds = timestamp - stake_.since - (numDays * 1 days);
        if (numSeconds > 0) {
            rewards +=
                (numSeconds * stake_.amount * percentPerSecond) /
                DENOMINATOR_SECONDS;
        }
        return rewards;
    }

    /**
     * Transfer the current amount of rewards to staker.
     */
    function _withdraw(address staker)
        internal
        virtual
        returns (uint256, uint256)
    {
        uint256 rewardsAmount = 0;
        uint256 stakeAmount = 0;
        require(
            block.timestamp > endDate,
            "StakingRewards: lock time is not over yet"
        );
        for (uint256 i = _lastIndex[staker]; i < _stakes[staker].length; i++) {
            rewardsAmount += _calculateStakeRewards(_stakes[staker][i]);
            stakeAmount += _stakes[staker][i].amount;
            _lastIndex[staker] = i + 1;
        }

        return (rewardsAmount, stakeAmount);
    }

    function withdraw() external isPoolStarted whenNotPaused {
        require(
            _stakesSummary[msg.sender] > 0,
            "StakingRewards: the caller is not a staker"
        );

        (uint256 rewardsAmount, uint256 stakeAmount) = _withdraw(msg.sender);
        require(rewardsAmount > 0, "StakingRewards: already withdraw");
        require(
            stakeAmount > 0,
            "StakingRewards: (sanity check) aldready withdraw"
        );

        assert(_stakesSummary[msg.sender] >= stakeAmount);
        _stakesSummary[msg.sender] -= stakeAmount;
        if (_stakesSummary[msg.sender] == 0) {
            totalStakers--;
        }

        stakingToken().safeTransfer(msg.sender, stakeAmount + rewardsAmount);
        emit Rewarded(msg.sender, stakeAmount, rewardsAmount);
    }

    /**
     * Withdraw all remaining tokens in the pool
     */
    function transferToOwner() external onlyRole(DEFAULT_ADMIN_ROLE) {
        stakingToken().safeTransfer(
            msg.sender,
            stakingToken().balanceOf(address(this))
        );
    }

    /**
     * Return the current amount of stakes of the caller.
     */
    function stakeOf(address staker)
        public
        view
        isPoolStarted
        returns (uint256)
    {
        return _stakesSummary[staker];
    }

    /**
     * Get times stake of staker
     */
    function getTimesStake() external view returns (uint256) {
        return _stakes[msg.sender].length;
    }

    /**
     * Return the current amount of rewards of a staker.
     */
    function earnableOf(address staker)
        external
        view
        isPoolStarted
        returns (uint256)
    {
        uint256 rewardsAmount = 0;
        for (uint256 i = _lastIndex[staker]; i < _stakes[staker].length; i++) {
            rewardsAmount += _calculateStakeRewards(_stakes[staker][i]);
        }

        return rewardsAmount;
    }

    modifier isPoolStarted() {
        require(
            block.timestamp >= startDate,
            "StakingRewards: pool has not been started yet"
        );
        _;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC1363Payable, AccessControl)
        returns (bool)
    {
        return
            interfaceId == type(IERC1363Receiver).interfaceId ||
            interfaceId == type(IERC1363Spender).interfaceId ||
            interfaceId == type(IAccessControl).interfaceId ||
            interfaceId == type(IERC165).interfaceId ||
            super.supportsInterface(interfaceId);
    }
}
