const SCVestingIDO = artifacts.require("VestingIDO");
const SCToken = artifacts.require("Token");

const token = require("../token");
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const vestingConfig = require('../vesting.js');
const { assert } = require("chai");
const timeMachine = require('ganache-time-traveler');

contract("VestingIDO", async (accounts) => {
    let amount = new BN(helper.wei(111));
    let percentOfAmount = new BN(helper.wei(111 * 20 / 100));
    let claimTGE = new BN(helper.wei(111 * 20 / 100));
    let vestingIDO, Token;
    let owner = accounts[0];
    let grantor = accounts[1];
    let beneficiary = accounts[2];
    const TGE_MILESTONE = 1647621000; // 18-03-2022 16:30:00 UTC
    const milestones = [
        1650299400,
        1652891400,
        1655569800,
        1658161800,
        1660840200
    ];

    describe("once deployed", async () => {
        let snapshotId;

        beforeEach(async function () {
            let snapshot = await timeMachine.takeSnapshot();
            snapshotId = snapshot['result'];

            Token = await SCToken.new(token.name, token.symbol);
            vestingIDO = await SCVestingIDO.new(Token.address);
            await Token.addNoTaxAddress(vestingIDO.address);
            await vestingIDO.setGrantor(grantor, { from: owner });
            await vestingIDO.addBeneficiaries([beneficiary], [amount], { from: grantor });
            await Token.transfer(vestingIDO.address, vestingConfig.IDO, { from: owner });
        });

        afterEach(async () => {
            await timeMachine.revertToSnapshot(snapshotId);
        });

        describe("when claiming TGE", async () => {
            it("cannot claim TGE if not yet over 30 minutes passed since TGE time", async () => {
                await truffleAssert.fails(vestingIDO.claimTGE({ from: beneficiary }));
                try {
                    await vestingIDO.claimTGE({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: need to wait for 30 minutes before unlocking TGE tokens");
                }
            });

            it("can be claimed by beneficiaries", async () => {
                let claimed = await vestingIDO.hasClaimedTGE(beneficiary);
                assert.equal(claimed, false);
                await timeMachine.advanceBlockAndSetTime(TGE_MILESTONE);
                let event = await vestingIDO.claimTGE({ from: beneficiary });
                truffleAssert.eventEmitted(event, "TGEClaimed", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    evt.amount.should.be.a.bignumber.that.equal(claimTGE)
                    return true;
                });
                let balanceAfterClaiming = await Token.balanceOf(beneficiary);
                balanceAfterClaiming.should.be.a.bignumber.that.equals(claimTGE);

                claimed = await vestingIDO.hasClaimedTGE(beneficiary);
                assert.equal(claimed, true);
            });

            it("can be claimed at any time", async () => {
                // 18-07-2022 16:00 UTC
                await timeMachine.advanceBlockAndSetTime(milestones[3]);

                let event = await vestingIDO.claimTGE({ from: beneficiary });
                truffleAssert.eventEmitted(event, "TGEClaimed", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    evt.amount.should.be.a.bignumber.that.equal(claimTGE)
                    return true;
                });
                let balanceAfterClaiming = await Token.balanceOf(beneficiary);
                balanceAfterClaiming.should.be.a.bignumber.that.equals(claimTGE);

                let info = await vestingIDO.vestingOf(beneficiary, { from: beneficiary });
                info[0].should.be.a.bignumber.that.equal(amount);
                info[1].should.be.a.bignumber.that.equal(new BN(0));
                let claimed = await vestingIDO.hasClaimedTGE(beneficiary);
                assert.equal(claimed, true);
            });

            it("cannot claim TGE more than once", async () => {
                await timeMachine.advanceBlockAndSetTime(TGE_MILESTONE);
                let event = await vestingIDO.claimTGE({ from: beneficiary });
                truffleAssert.eventEmitted(event, "TGEClaimed", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    evt.amount.should.be.a.bignumber.that.equal(claimTGE)
                    return true;
                });

                await truffleAssert.fails(vestingIDO.claimTGE({ from: beneficiary }));
                try {
                    await vestingIDO.claimTGE({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: already claimed all TGE tokens");
                }
            });
        });

        describe("when a beneficiary claims", async () => {
            beforeEach(async () => {
                await timeMachine.advanceBlockAndSetTime(TGE_MILESTONE);
                await truffleAssert.passes(vestingIDO.claimTGE({ from: beneficiary }));
            });

            it("he cannot claim before the first milestone", async () => {
                await timeMachine.advanceBlockAndSetTime(TGE_MILESTONE + 10); // should be TGE_MILESTONE + 10 seconds now

                await truffleAssert.fails(vestingIDO.claim({ from: beneficiary }));
                try {
                    await vestingIDO.claim({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: cannot unlock tokens for this milestone or already claimed tokens for current milestone");
                }
            });

            it("he receives correct amounts of token for each milestone", async () => {
                let total = new BN(0);
                for (i = 1; i < milestones.length; i++) {
                    await timeMachine.advanceBlockAndSetTime(milestones[i]);
                    let event = await vestingIDO.claim({ from: beneficiary });
                    truffleAssert.eventEmitted(event, "Claimed", (evt) => {
                        assert.equal(evt.beneficiary, beneficiary);
                        evt.amount.should.be.a.bignumber.that.equal(percentOfAmount);
                        total = total.add(evt.amount);
                        return true;
                    })

                    let info = await vestingIDO.vestingOf(beneficiary, { from: beneficiary });
                    info[0].should.be.a.bignumber.that.equal(amount);
                    info[1].should.be.a.bignumber.that.equal(new BN(i));
                }
                amount.should.be.a.bignumber.that.equal(total.add(claimTGE));

                await timeMachine.advanceBlockAndSetTime(milestones[milestones.length - 1] + 1);

                await truffleAssert.fails(vestingIDO.claim({ from: beneficiary }));
                try {
                    await vestingIDO.claim({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: cannot unlock tokens for this milestone or already claimed tokens for current milestone");
                }
            });

            it("he can claim multiple milestones at once", async () => {
                await timeMachine.advanceBlockAndSetTime(milestones[3]);
                let eventClaim = await vestingIDO.claim({ from: beneficiary });
                truffleAssert.eventEmitted(eventClaim, "Claimed", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    evt.amount.should.be.a.bignumber.that.equal(percentOfAmount.mul(new BN(3)));
                    return true;
                })

                let info = await vestingIDO.vestingOf(beneficiary, { from: beneficiary });
                info[0].should.be.a.bignumber.that.equal(amount);
                info[1].should.be.a.bignumber.that.equal(new BN(3));
            })

            it("he can claim all at once", async () => {
                let totalClaim = new BN(0);

                // 14:00 18-10-2022 UTC
                await timeMachine.advanceBlockAndSetTime(1666108800);
                let eventClaim = await vestingIDO.claim({ from: beneficiary });
                truffleAssert.eventEmitted(eventClaim, "Claimed", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    totalClaim = percentOfAmount.mul(new BN(milestones.length - 1));
                    evt.amount.should.be.a.bignumber.that.equal(totalClaim);
                    return true;
                });

                // total claim all = total claim - claimTGE
                totalClaim.should.be.a.bignumber.that.equal(amount.sub(claimTGE));

                let info = await vestingIDO.vestingOf(beneficiary, { from: beneficiary });
                info[0].should.be.a.bignumber.that.equal(amount);
                info[1].should.be.a.bignumber.that.equal(new BN(milestones.length - 1));

                await timeMachine.advanceBlockAndSetTime(1666108801);

                await truffleAssert.fails(vestingIDO.claim({ from: beneficiary }));
                try {
                    await vestingIDO.claim({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: cannot unlock tokens for this milestone or already claimed tokens for current milestone");
                }

                // check balance of pool
                let poolBalance = await Token.balanceOf(vestingIDO.address);
                poolBalance.should.be.a.bignumber.that.equals(vestingConfig.IDO.sub(amount));
            });
        });
    });
});