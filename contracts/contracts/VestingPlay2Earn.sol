// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "./utils/Vesting.sol";

contract VestingPlay2Earn is Vesting {
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
            1648225800,
            1650299400,
            1652891400,
            1655569800,
            1658161800,
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
            1695054600,
            1697646600,
            1700325000,
            1702917000,
            1705595400,
            1708273800,
            1710779400,
            1713457800,
            1716049800,
            1718728200,
            1721320200,
            1723998600,
            1726677000,
            1729269000,
            1731947400,
            1734539400,
            1737217800,
            1739896200
        ];
        _claimablePercents[1648225800] = 278; // 2022-03-25 16:30:00UTC
        _claimablePercents[1650299400] = 278; // 2022-04-18 16:30:00UTC
        _claimablePercents[1652891400] = 278; // 2022-05-18 16:30:00UTC
        _claimablePercents[1655569800] = 278; // 2022-06-18 16:30:00UTC
        _claimablePercents[1658161800] = 278; // 2022-07-18 16:30:00UTC
        _claimablePercents[1660840200] = 278; // 2022-08-18 16:30:00UTC
        _claimablePercents[1663518600] = 278; // 2022-09-18 16:30:00UTC
        _claimablePercents[1666110600] = 278; // 2022-10-18 16:30:00UTC
        _claimablePercents[1668789000] = 278; // 2022-11-18 16:30:00UTC
        _claimablePercents[1671381000] = 278; // 2022-12-18 16:30:00UTC
        _claimablePercents[1674059400] = 278; // 2023-01-18 16:30:00UTC
        _claimablePercents[1676737800] = 278; // 2023-02-18 16:30:00UTC
        _claimablePercents[1679157000] = 278; // 2023-03-18 16:30:00UTC
        _claimablePercents[1681835400] = 278; // 2023-04-18 16:30:00UTC
        _claimablePercents[1684427400] = 278; // 2023-05-18 16:30:00UTC
        _claimablePercents[1687105800] = 278; // 2023-06-18 16:30:00UTC
        _claimablePercents[1689697800] = 278; // 2023-07-18 16:30:00UTC
        _claimablePercents[1692376200] = 278; // 2023-08-18 16:30:00UTC
        _claimablePercents[1695054600] = 278; // 2023-09-18 16:30:00UTC
        _claimablePercents[1697646600] = 278; // 2023-10-18 16:30:00UTC
        _claimablePercents[1700325000] = 278; // 2023-11-18 16:30:00UTC
        _claimablePercents[1702917000] = 278; // 2023-12-18 16:30:00UTC
        _claimablePercents[1705595400] = 278; // 2024-01-18 16:30:00UTC
        _claimablePercents[1708273800] = 278; // 2024-02-18 16:30:00UTC
        _claimablePercents[1710779400] = 278; // 2024-03-18 16:30:00UTC
        _claimablePercents[1713457800] = 278; // 2024-04-18 16:30:00UTC
        _claimablePercents[1716049800] = 278; // 2024-05-18 16:30:00UTC
        _claimablePercents[1718728200] = 278; // 2024-06-18 16:30:00UTC
        _claimablePercents[1721320200] = 278; // 2024-07-18 16:30:00UTC
        _claimablePercents[1723998600] = 278; // 2024-08-18 16:30:00UTC
        _claimablePercents[1726677000] = 278; // 2024-09-18 16:30:00UTC
        _claimablePercents[1729269000] = 278; // 2024-10-18 16:30:00UTC
        _claimablePercents[1731947400] = 278; // 2024-11-18 16:30:00UTC
        _claimablePercents[1734539400] = 278; // 2024-12-18 16:30:00UTC
        _claimablePercents[1737217800] = 278; // 2025-01-18 16:30:00UTC
        _claimablePercents[1739896200] = 270; // 2025-02-18 16:30:00UTC
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
