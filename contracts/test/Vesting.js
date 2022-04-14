const SCMockVesting = artifacts.require("MockVesting");
const SCToken = artifacts.require("Token");

const token = require("../token");
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const { assert } = require("chai");
const { ZERO_ADDRESS } = constants;

contract("Vesting", async (accounts) => {
    let MockVesting, Token;
    let owner = accounts[0];
    let grantor = accounts[1];
    let beneficiary = accounts[2];
    let amount100 = new BN(helper.wei(100));
    let amount200 = new BN(helper.wei(200));
    let amount1000 = new BN(helper.wei(1000));
    let admin_role = "0x0000000000000000000000000000000000000000000000000000000000000000";
    let grantor_role = "0xd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf";
    let beneficiary_role = "0xc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e";

    describe("once deployed", async () => {
        context("admin control", async () => {
            beforeEach(async function () {
                Token = await SCToken.new(token.name, token.symbol);
                MockVesting = await SCMockVesting.new(Token.address);
            });

            it("only admin set grantor", async () => {
                await truffleAssert.fails(MockVesting.setGrantor(grantor, { from: grantor }));
                try {
                    await MockVesting.setGrantor(grantor, { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + grantor.toLowerCase() + " is missing role " + admin_role);
                }
            })

            it("only admin remove grantor", async () => {
                let eventSetGrantor = await MockVesting.setGrantor(grantor, { from: owner });
                truffleAssert.eventEmitted(eventSetGrantor, "RoleGranted", (evt) => {
                    assert.equal(evt.role, grantor_role);
                    assert.equal(evt.account, grantor);
                    assert.equal(evt.sender, owner);
                    return true;
                })

                await truffleAssert.fails(MockVesting.removeGrantor(grantor, { from: grantor }));
                try {
                    await MockVesting.removeGrantor(grantor, { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + grantor.toLowerCase() + " is missing role " + admin_role);
                }
            })
        })

        context("grantor control", async () => {
            beforeEach(async function () {
                Token = await SCToken.new(token.name, token.symbol);
                MockVesting = await SCMockVesting.new(Token.address);
                await Token.addNoTaxAddress(MockVesting.address);
                await MockVesting.setGrantor(grantor, { from: owner });
            });

            it("only add beneficiary by grantor", async () => {
                await truffleAssert.fails(MockVesting.addBeneficiaries([beneficiary], [amount100], { from: beneficiary }));
                try {
                    await MockVesting.addBeneficiaries([beneficiary], [amount100], { from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + beneficiary.toLowerCase() + " is missing role " + grantor_role);
                }
            })

            it("beneficiaries and amounts' length should be equal", async () => {
                await truffleAssert.fails(MockVesting.addBeneficiaries([beneficiary], [], { from: grantor }));
                try {
                    await MockVesting.addBeneficiaries([beneficiary], [], { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: beneficiaries and amounts' length should be equal");
                }
            })

            it("address beneficiary not is address 0", async () => {
                await truffleAssert.fails(MockVesting.addBeneficiaries([ZERO_ADDRESS], [amount100], { from: grantor }));
                try {
                    await MockVesting.addBeneficiaries([ZERO_ADDRESS], [amount100], { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: beneficiary address must not be 0");
                }
            })

            it("amount beneficiary not equal 0", async () => {
                await truffleAssert.fails(MockVesting.addBeneficiaries([beneficiary], [0], { from: grantor }));
                try {
                    await MockVesting.addBeneficiaries([beneficiary], [0], { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: total amount must be greater than 0");
                }
            })

            it("add beneficiary ok", async () => {
                let beneficiaries = [beneficiary, accounts[3]];
                let amounts = [amount100, amount200];

                let event = await MockVesting.addBeneficiaries(beneficiaries, amounts, { from: grantor });
                truffleAssert.eventEmitted(event, "BeneficiaryRegistered", (evt, i) => {
                    assert.equal(evt.beneficiary, beneficiaries[i]);
                    evt.amount.should.be.a.bignumber.that.equal(amounts[i])
                    return true;
                });

                let info = await MockVesting.vestingOf(beneficiary, { from: beneficiary });
                info[0].should.be.a.bignumber.that.equal(amount100);
                info[1].should.be.a.bignumber.that.equal(new BN(0));
            })

            it("not add beneficiary 2 times", async () => {
                let beneficiaries = [beneficiary, accounts[3], beneficiary];
                let amounts = [amount100, amount200, amount100];

                await truffleAssert.fails(MockVesting.addBeneficiaries(beneficiaries, amounts, { from: grantor }));
                try {
                    await MockVesting.addBeneficiaries(beneficiaries, amounts, { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: beneficiary is already in pool")
                }
            })

            it("not suspend if beneficiary not already", async () => {
                await truffleAssert.fails(MockVesting.suspendBeneficiary(beneficiary, { from: grantor }));
                try {
                    await MockVesting.suspendBeneficiary(beneficiary, { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: beneficiary must be employed")
                }
            })

            it("suspend beneficiary ok", async () => {
                await truffleAssert.passes(MockVesting.addBeneficiaries([beneficiary], [amount100], { from: grantor }));

                let event = await MockVesting.suspendBeneficiary(beneficiary, { from: grantor });
                truffleAssert.eventEmitted(event, "BeneficiarySuspended", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    return true;
                })
                await truffleAssert.fails(MockVesting.vestingOf(beneficiary, { from: beneficiary }));
                try {
                    await MockVesting.vestingOf(beneficiary, { from: beneficiary });
                } catch (error) {
                    let reason = error.message.replace("Returned error: VM Exception while processing transaction: revert ", "")
                    assert.equal(reason, "Vesting: beneficiary is not in pool");
                }
            })

            it("suspend beneficiary and add beneficiary", async () => {
                await truffleAssert.passes(MockVesting.addBeneficiaries([beneficiary], [amount100], { from: grantor }));

                let event = await MockVesting.suspendBeneficiary(beneficiary, { from: grantor });
                truffleAssert.eventEmitted(event, "BeneficiarySuspended", (evt) => {
                    assert.equal(evt.beneficiary, beneficiary);
                    return true;
                })

                await truffleAssert.fails(MockVesting.addBeneficiaries([beneficiary], [amount100], { from: grantor }));
                try {
                    await MockVesting.addBeneficiaries([beneficiary], [amount100], { from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: beneficiary is already in pool");
                }
            })

            it("only grantor collect token", async () => {
                await truffleAssert.fails(MockVesting.collectToken({ from: beneficiary }));
                try {
                    await MockVesting.collectToken({ from: beneficiary });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + beneficiary.toLowerCase() + " is missing role " + grantor_role);
                }
            })

            it("not collect token if balance equal 0", async () => {
                await truffleAssert.fails(MockVesting.collectToken({ from: owner }));
                try {
                    await MockVesting.collectToken({ from: owner });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: current balance is zero");
                }
            })

            it("collect token ok", async () => {
                await Token.transfer(MockVesting.address, amount1000);

                let beforeClaim = await Token.balanceOf(owner, { from: owner });
                let balanceOfVesting = await Token.balanceOf(MockVesting.address, { from: owner });
                await MockVesting.collectToken({ from: owner });
                let afterClaim = await Token.balanceOf(owner, { from: owner });
                afterClaim.should.be.a.bignumber.that.equal(beforeClaim.add(balanceOfVesting));
            })

            it("not collect token 2 times", async () => {
                await Token.transfer(MockVesting.address, amount1000);

                let beforeClaim = await Token.balanceOf(owner, { from: owner });
                let balanceOfVesting = await Token.balanceOf(MockVesting.address, { from: owner });
                await MockVesting.collectToken({ from: owner });
                let afterClaim = await Token.balanceOf(owner, { from: owner });
                afterClaim.should.be.a.bignumber.that.equal(beforeClaim.add(balanceOfVesting));

                await truffleAssert.fails(MockVesting.collectToken({ from: owner }));
                try {
                    await MockVesting.collectToken({ from: owner });
                } catch (error) {
                    assert.equal(error.reason, "Vesting: current balance is zero")
                }
            })

        })

        context("beneficiary control", async () => {
            beforeEach(async () => {
                beforeEach(async function () {
                    Token = await SCToken.new(token.name, token.symbol);
                    MockVesting = await SCMockVesting.new(Token.address);
                    await Token.addNoTaxAddress(MockVesting.address);
                    await MockVesting.setGrantor(grantor, { from: owner });
                    await MockVesting.addBeneficiaries([beneficiary], [amount100], { from: grantor });

                    Token.transfer(MockVesting.address, amount1000, { from: owner });
                });
            })

            it("only beneficiary claim", async () => {
                await truffleAssert.fails(MockVesting.claim({ from: grantor }));
                try {
                    await MockVesting.claim({ from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + grantor.toLowerCase() + " is missing role " + beneficiary_role);
                }
            })

            it("only beneficiary claim TGE", async () => {
                await truffleAssert.fails(MockVesting.claimTGE({ from: grantor }));
                try {
                    await MockVesting.claimTGE({ from: grantor });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + grantor.toLowerCase() + " is missing role " + beneficiary_role);
                }
            })
        })
    })

})