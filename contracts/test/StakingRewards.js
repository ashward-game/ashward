var token = require("../token");
const SCToken = artifacts.require("Token");
const StakingRewards = artifacts.require("StakingRewards");
const StakingPool1 = artifacts.require("StakingPool1");
const StakingPool2 = artifacts.require("StakingPool2");

// library test
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
// Enable and inject BN dependency
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const timeMachine = require('ganache-time-traveler');
const helper = require('./helpers/helper');
const staking = require('../staking');
const { assert } = require("chai");
const { ZERO_ADDRESS } = constants;

contract("Staking", async (accounts) => {
    let poolMock, Token, snapshotId;

    let minStake = helper.wei(500);
    let maxStake = helper.wei(1000);
    let secondPerDay = 86400;
    let percent = 2.5;// 2.5%/day
    let percentSeconds = 0.000028935;//0.000028935%/seconds --- ~2.5%/day
    let totalInPool = helper.wei(3000);
    let owner = accounts[0];
    let staker = accounts[1];
    let balance = helper.wei(10000);
    let startDate, endDate;
    let stakeAmount = helper.wei(500);
    let admin_role = "0x0000000000000000000000000000000000000000000000000000000000000000"

    describe("once deployed", async () => {
        beforeEach(async () => {
            let snapshot = await timeMachine.takeSnapshot();
            snapshotId = snapshot['result'];
        })

        afterEach(async () => {
            await timeMachine.revertToSnapshot(snapshotId);
        });

        describe("test input pool", async () => {
            it("test input pool 1", async () => {
                let pool1 = await StakingPool1.deployed();

                let start = await pool1.startDate();
                start.should.be.a.bignumber.that.equal(new BN(1649347200));//2022-04-07 16:00

                let end = await pool1.endDate();
                end.should.be.a.bignumber.that.equal(new BN(1650556800));//2022-04-21 16:00
            })

            it("test input pool 2", async () => {
                let pool2 = await StakingPool2.deployed();
                let start = await pool2.startDate();
                start.should.be.a.bignumber.that.equal(new BN(1649347200));//2022-04-07 16:00

                let end = await pool2.endDate();
                end.should.be.a.bignumber.that.equal(new BN(1651161600));//2022-04-28 16:00
            })
        })

        describe("Pool Mocking(pool 1, pool 2)", async () => {
            beforeEach(async () => {
                Token = await SCToken.new(token.name, token.symbol);
                let lastBlock = await web3.eth.getBlock("latest");
                startDate = lastBlock.timestamp + secondPerDay;
                endDate = startDate + (secondPerDay * 14);

                poolMock = await StakingRewards.new(Token.address, minStake, maxStake, percent * 100, percentSeconds * 10e8, totalInPool, startDate, endDate);

                // enabled transfer
                await Token.addSellingAddress(accounts[9], { from: owner });
            })

            describe("access control", async () => {
                beforeEach(async () => {
                    await truffleAssert.passes(Token.transfer(poolMock.address, balance, { from: owner }));
                })

                it("only owner call pause", async () => {
                    await truffleAssert.fails(poolMock.pause({ from: staker }), truffleAssert.ErrorType.REVERT, "AccessControl: account " + staker.toLowerCase() + " is missing role " + admin_role);
                })

                it("only owner call unpause", async () => {
                    await truffleAssert.fails(poolMock.unpause({ from: staker }), truffleAssert.ErrorType.REVERT, "AccessControl: account " + staker.toLowerCase() + " is missing role " + admin_role);
                })

                it("only owner receive token", async () => {
                    await truffleAssert.fails(poolMock.transferToOwner({ from: staker }), truffleAssert.ErrorType.REVERT, "AccessControl: account " + staker.toLowerCase() + " is missing role " + admin_role);

                    let balancePool = await Token.balanceOf(poolMock.address);
                    balancePool.should.be.a.bignumber.that.equal(balance);
                })

                it("receive all token", async () => {
                    let balancePool = await Token.balanceOf(poolMock.address);
                    balancePool.should.be.a.bignumber.that.equal(balance);

                    let beforeReceive = await Token.balanceOf(owner);
                    await poolMock.transferToOwner({ from: owner });
                    let afteReceive = await Token.balanceOf(owner);

                    let balancePoolFinal = await Token.balanceOf(poolMock.address);
                    balancePoolFinal.should.be.a.bignumber.that.equal(new BN(0));

                    afteReceive.should.be.a.bignumber.that.equal(beforeReceive.add(balance));
                })
            })

            describe("stake to pool", async () => {
                it("cannot stake if pool not started", async () => {
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: pool has not been started yet");
                })

                it("cannot stake if enought balance token", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }), truffleAssert.ErrorType.REVERT, "ERC20: transfer amount exceeds balance");
                })

                it("cannot stake if amount < min", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, 1, { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: insufficient stake amount");
                })

                it("cannot stake if amount > max", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, balance, { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: overflowing stake amount");
                })

                it("cannot stake if total stake > max", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, helper.wei(1000), { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: overflowing stake amount");
                })

                it("cannot stake if end time stake", async () => {
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await timeMachine.advanceBlockAndSetTime(endDate);
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: staking time is over");
                })

                it("stake ok", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await truffleAssert.passes(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }));
                    let total = await poolMock.totalStaking();
                    total.should.be.a.bignumber.that.equal(new BN(stakeAmount))

                    let mystake = await poolMock.stakeOf(staker);
                    mystake.should.be.a.bignumber.that.equal(new BN(stakeAmount));

                    let totalStaker = await poolMock.totalStakers();
                    totalStaker.should.be.a.bignumber.that.equal(new BN(1));

                    let timesStaker = await poolMock.getTimesStake({ from: staker });
                    timesStaker.should.be.a.bignumber.that.equal(new BN(1));
                })

                it("stake many", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await truffleAssert.passes(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }));
                    await truffleAssert.passes(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }));

                    let total = await poolMock.totalStaking();
                    total.should.be.a.bignumber.that.equal(new BN(helper.wei(1000)))

                    let mystake = await poolMock.stakeOf(staker);
                    mystake.should.be.a.bignumber.that.equal(new BN(helper.wei(1000)));

                    let totalStaker = await poolMock.totalStakers();
                    totalStaker.should.be.a.bignumber.that.equal(new BN(1));

                    let timesStaker = await poolMock.getTimesStake({ from: staker });
                    timesStaker.should.be.a.bignumber.that.equal(new BN(2));
                })

                it("cannot stake if total pool is full", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    for (i = 0; i < 3; i++) {
                        let _staker = accounts[i + 5];
                        await truffleAssert.passes(Token.transfer(_staker, helper.wei(1000), { from: owner }));
                        await truffleAssert.passes(Token.approveAndCall(poolMock.address, helper.wei(1000), { from: _staker }));
                    }
                    await truffleAssert.passes(Token.transfer(staker, helper.wei(1000), { from: owner }));

                    let total = await poolMock.totalStaking();
                    total.should.be.a.bignumber.that.equal(new BN(helper.wei(3000)))
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, helper.wei(1000), { from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: pool is full");

                    let totalStaker = await poolMock.totalStakers();
                    totalStaker.should.be.a.bignumber.that.equal(new BN(3));
                })

                it("cannot stake when pool paused", async () => {
                    await timeMachine.advanceTimeAndBlock(secondPerDay);
                    await poolMock.pause({ from: owner });
                    await truffleAssert.passes(Token.transfer(staker, stakeAmount, { from: owner }));
                    await truffleAssert.fails(Token.approveAndCall(poolMock.address, stakeAmount, { from: staker }), truffleAssert.ErrorType.REVERT, "Pausable: paused");
                })
            })

            describe("withdraw pool", async () => {
                beforeEach(async () => {
                    await truffleAssert.passes(Token.transfer(staker, balance, { from: owner }));
                    await truffleAssert.passes(Token.transfer(poolMock.address, balance, { from: owner }));
                })

                it("cannot withdraw if caller not staker", async () => {
                    await timeMachine.advanceBlockAndSetTime(startDate);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });

                    await truffleAssert.fails(poolMock.withdraw({ from: owner }), truffleAssert.ErrorType.REVERT, "StakingRewards: the caller is not a staker");
                })

                it("cannot withdraw if not over time lock(end date)", async () => {
                    await timeMachine.advanceBlockAndSetTime(startDate);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });

                    await truffleAssert.fails(poolMock.withdraw({ from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: lock time is not over yet");
                })

                it("stake 0.5 day => withdraw ok", async () => {
                    let timeStart = endDate - (secondPerDay * 0.5);
                    await timeMachine.advanceBlockAndSetTime(timeStart);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    let since = (await web3.eth.getBlock("latest")).timestamp;

                    await timeMachine.advanceTimeAndBlock(secondPerDay * 2);
                    let actualEarn = await poolMock.earnableOf(staker);
                    let exspectEarn;
                    if (since == timeStart) {
                        exspectEarn = new BN(String((stakeAmount * percentSeconds) * (secondPerDay * 0.5) / 100));
                    } else {
                        exspectEarn = new BN(String((stakeAmount * percentSeconds) * (secondPerDay * 0.5 - (since - timeStart)) / 100));
                    }
                    actualEarn.should.be.a.bignumber.that.equal(exspectEarn);

                    let beforeBalance = await Token.balanceOf(staker);
                    let event = await poolMock.withdraw({ from: staker });
                    truffleAssert.eventEmitted(event, "Rewarded", (evt) => {
                        assert.equal(evt.staker, staker);
                        evt.rewardsAmount.should.be.a.bignumber.that.equal(exspectEarn);
                        evt.stakeAmount.should.be.a.bignumber.that.equal(stakeAmount);
                        return true;
                    })

                    let afterBalance = await Token.balanceOf(staker);
                    afterBalance.should.be.a.bignumber.that.equal(beforeBalance.add(stakeAmount).add(exspectEarn));
                })

                it("stake 1 day => withdraw ok", async () => {
                    let timeStart = endDate - secondPerDay;
                    await timeMachine.advanceBlockAndSetTime(timeStart);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    let since = (await web3.eth.getBlock("latest")).timestamp;

                    await timeMachine.advanceTimeAndBlock(secondPerDay * 2);
                    let actualEarn = await poolMock.earnableOf(staker);
                    let exspectEarn;
                    if (since == timeStart) {
                        exspectEarn = new BN(String(stakeAmount * percent / 100));
                    } else {
                        exspectEarn = new BN(String((stakeAmount * percentSeconds) / 100 * (secondPerDay - (since - timeStart))));
                    }
                    actualEarn.should.be.a.bignumber.that.equal(exspectEarn);

                    let beforeBalance = await Token.balanceOf(staker);
                    let event = await poolMock.withdraw({ from: staker });
                    truffleAssert.eventEmitted(event, "Rewarded", (evt) => {
                        assert.equal(evt.staker, staker);
                        evt.rewardsAmount.should.be.a.bignumber.that.equal(exspectEarn);
                        evt.stakeAmount.should.be.a.bignumber.that.equal(stakeAmount);
                        return true;
                    })

                    let afterBalance = await Token.balanceOf(staker);
                    afterBalance.should.be.a.bignumber.that.equal(beforeBalance.add(stakeAmount).add(exspectEarn));
                })

                it("stake 1.5 day => withdraw ok", async () => {
                    let timeStart = endDate - (secondPerDay * 1.5);
                    await timeMachine.advanceBlockAndSetTime(timeStart);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    let since = (await web3.eth.getBlock("latest")).timestamp;

                    let rewardDay = stakeAmount * percent / 100;
                    await timeMachine.advanceTimeAndBlock(secondPerDay * 2);

                    let actualEarn = await poolMock.earnableOf(staker);
                    let exspectEarn;
                    if (since == timeStart) {
                        exspectEarn = new BN(String(rewardDay + (stakeAmount * percentSeconds / 100 * (secondPerDay * 0.5))));
                    } else {
                        exspectEarn = new BN(String(rewardDay + (stakeAmount * percentSeconds) / 100 * (secondPerDay * 0.5 - (since - timeStart))));
                    }
                    actualEarn.should.be.a.bignumber.that.equal(exspectEarn);

                    let beforeBalance = await Token.balanceOf(staker);
                    let event = await poolMock.withdraw({ from: staker });
                    truffleAssert.eventEmitted(event, "Rewarded", (evt) => {
                        assert.equal(evt.staker, staker);
                        evt.rewardsAmount.should.be.a.bignumber.that.equal(exspectEarn);
                        evt.stakeAmount.should.be.a.bignumber.that.equal(stakeAmount);
                        return true;
                    })

                    let afterBalance = await Token.balanceOf(staker);
                    afterBalance.should.be.a.bignumber.that.equal(beforeBalance.add(stakeAmount).add(exspectEarn));
                })

                it("stake many to withdraw", async () => {
                    let timeStart = endDate - (secondPerDay * 3);
                    await timeMachine.advanceBlockAndSetTime(timeStart);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    let since1 = (await web3.eth.getBlock("latest")).timestamp;

                    await timeMachine.advanceTimeAndBlock(secondPerDay);// add 1 day

                    let actualEarn = await poolMock.earnableOf(staker);
                    let exspectEarn = new BN(String(stakeAmount * percent / 100));
                    actualEarn.should.be.a.bignumber.that.equal(exspectEarn);

                    await Token.transfer(staker, balance, { from: owner });
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });
                    let since2 = (await web3.eth.getBlock("latest")).timestamp;

                    let beforeBalance = await Token.balanceOf(staker);

                    await timeMachine.advanceBlockAndSetTime(endDate + secondPerDay);

                    let reward1;
                    let rewardInSeconds = stakeAmount * percentSeconds / 100;
                    if (since1 == timeStart) {
                        reward1 = exspectEarn.mul(new BN(3));
                    } else {
                        reward1 = exspectEarn.mul(new BN(2)).add(new BN(String(rewardInSeconds * (secondPerDay - (since1 - timeStart)))));
                    }

                    let reward2;
                    if (since2 == timeStart + secondPerDay * 2) {
                        reward2 = exspectEarn.mul(new BN(2));
                    } else {
                        reward2 = exspectEarn.add(new BN(String(rewardInSeconds * (secondPerDay - (since2 - (timeStart + secondPerDay))))));
                    }
                    let exspectEarn2 = reward1.add(reward2);

                    let actualEarn2 = await poolMock.earnableOf(staker);
                    actualEarn2.should.be.a.bignumber.that.equal(exspectEarn2);

                    let event = await poolMock.withdraw({ from: staker });
                    truffleAssert.eventEmitted(event, "Rewarded", (evt) => {
                        assert.equal(evt.staker, staker);
                        evt.rewardsAmount.should.be.a.bignumber.that.equal(exspectEarn2);
                        evt.stakeAmount.should.be.a.bignumber.that.equal(new BN(stakeAmount).mul(new BN(2)));
                        return true;
                    })

                    let afterBalance = await Token.balanceOf(staker);
                    afterBalance.should.be.a.bignumber.that.equal(beforeBalance.add(new BN(stakeAmount).mul(new BN(2))).add(exspectEarn2));
                })

                it("cannot withdraw 2 times", async () => {
                    await timeMachine.advanceBlockAndSetTime(endDate - secondPerDay);
                    await Token.approveAndCall(poolMock.address, stakeAmount, { from: staker });

                    await timeMachine.advanceTimeAndBlock(secondPerDay * 2);

                    await truffleAssert.passes(poolMock.withdraw({ from: staker }));

                    await truffleAssert.fails(poolMock.withdraw({ from: staker }), truffleAssert.ErrorType.REVERT, "StakingRewards: the caller is not a staker");
                })
            })
        })
    })
})