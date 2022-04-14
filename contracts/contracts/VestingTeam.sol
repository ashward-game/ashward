// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "./utils/Vesting.sol";

contract VestingTeam is Vesting {
    uint256 private immutable _tge_percent;
    // claimable milestones
    // each milestone is defined by the timestamp of the claimable datetime
    // index 0 is the timestamp of the starting date
    uint256[] private _claimableMilestones;

    // claimable percent of each milestone
    // claimable percent of the first milestone is always 0
    mapping(uint256 => uint256) private _claimablePercents;

    constructor(address token) Vesting(token) {
        _tge_percent = 0; // 0.00%
        _claimableMilestones = [
            1647621000,
            1681835400,
            1684427400,
            1687105800,
            1689697800,
            1692376200,
            1695054600,
            1697646600,
            1700325000,
            1702917000,
            1705595400,
            1708273800,
            1710779400
        ];
        _claimablePercents[1681835400] = 833; // 2023-04-18 16:30:00UTC
        _claimablePercents[1684427400] = 833; // 2023-05-18 16:30:00UTC
        _claimablePercents[1687105800] = 833; // 2023-06-18 16:30:00UTC
        _claimablePercents[1689697800] = 833; // 2023-07-18 16:30:00UTC
        _claimablePercents[1692376200] = 833; // 2023-08-18 16:30:00UTC
        _claimablePercents[1695054600] = 833; // 2023-09-18 16:30:00UTC
        _claimablePercents[1697646600] = 833; // 2023-10-18 16:30:00UTC
        _claimablePercents[1700325000] = 833; // 2023-11-18 16:30:00UTC
        _claimablePercents[1702917000] = 833; // 2023-12-18 16:30:00UTC
        _claimablePercents[1705595400] = 833; // 2024-01-18 16:30:00UTC
        _claimablePercents[1708273800] = 833; // 2024-02-18 16:30:00UTC
        _claimablePercents[1710779400] = 837; // 2024-03-18 16:30:00UTC
    }

    function _tgePercent() internal view override returns (uint256) {
        return _tge_percent;
    }

    function _percent(uint256 lastIndex)
        internal
        view
        override
        returns (uint256, uint256)
    {
        uint256 percent = 0;
        uint256 currentMilestone = 0;
        for (uint256 i = lastIndex + 1; i < _claimableMilestones.length; i++) {
            if (block.timestamp >= _claimableMilestones[i]) {
                percent += _claimablePercents[_claimableMilestones[i]];
                currentMilestone = i;
            } else {
                break;
            }
        }
        return (percent, currentMilestone);
    }
}
