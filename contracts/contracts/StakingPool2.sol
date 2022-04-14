// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "./utils/StakingRewards.sol";

contract StakingPool2 is StakingRewards {
    constructor(
        IERC1363 acceptedToken_,
        uint256 minStake_,
        uint256 maxStake_,
        uint256 percentPerDay_,
        uint256 percentPerSecond_,
        uint256 totalTokenInPool_,
        uint256 startDate_,
        uint256 endDate_
    )
        StakingRewards(
            acceptedToken_,
            minStake_,
            maxStake_,
            percentPerDay_,
            percentPerSecond_,
            totalTokenInPool_,
            startDate_,
            endDate_
        )
    {}
}
