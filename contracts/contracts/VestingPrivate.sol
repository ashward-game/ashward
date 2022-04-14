// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "./utils/Vesting.sol";

contract VestingPrivate is Vesting {
    uint256 private immutable _tge_percent;
    // claimable milestones
    // each milestone is defined by the timestamp of the claimable datetime
    // index 0 is the timestamp of the starting date
    uint256[] private _claimableMilestones;

    // claimable percent of each milestone
    // claimable percent of the first milestone is always 0
    mapping(uint256 => uint256) private _claimablePercents;

    constructor(address token) Vesting(token) {
        _tge_percent = 800; // 8.00%
        _claimableMilestones = [
            1647621000,
            1655569800,
            1660840200,
            1663518600,
            1666110600,
            1668789000,
            1671381000,
            1674059400,
            1676737800,
            1679157000,
            1681835400,
            1684427400,
            1687105800,
            1689697800,
            1692376200,
            1695054600
        ];
        _claimablePercents[1655569800] = 613; // 2022-06-18 16:30:00UTC
        _claimablePercents[1660840200] = 613; // 2022-08-18 16:30:00UTC
        _claimablePercents[1663518600] = 613; // 2022-09-18 16:30:00UTC
        _claimablePercents[1666110600] = 613; // 2022-10-18 16:30:00UTC
        _claimablePercents[1668789000] = 613; // 2022-11-18 16:30:00UTC
        _claimablePercents[1671381000] = 613; // 2022-12-18 16:30:00UTC
        _claimablePercents[1674059400] = 613; // 2023-01-18 16:30:00UTC
        _claimablePercents[1676737800] = 613; // 2023-02-18 16:30:00UTC
        _claimablePercents[1679157000] = 613; // 2023-03-18 16:30:00UTC
        _claimablePercents[1681835400] = 613; // 2023-04-18 16:30:00UTC
        _claimablePercents[1684427400] = 613; // 2023-05-18 16:30:00UTC
        _claimablePercents[1687105800] = 613; // 2023-06-18 16:30:00UTC
        _claimablePercents[1689697800] = 613; // 2023-07-18 16:30:00UTC
        _claimablePercents[1692376200] = 613; // 2023-08-18 16:30:00UTC
        _claimablePercents[1695054600] = 618; // 2023-09-18 16:30:00UTC
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
