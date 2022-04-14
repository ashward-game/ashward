var token = require('../token')
const SCToken = artifacts.require("Token");

const chai = require('chai');
var should = require('chai').should()
const BN = web3.utils.BN;
const util = web3.utils;
// Enable and inject BN dependency
chai.use(require('chai-bn')(BN));
const truffleAssert = require("truffle-assertions");
const timeMachine = require('ganache-time-traveler');

contract("Token", accounts => {
    let TGE_PLUS_1_HOUR = 1647622800;

    describe("for deployed instance", async () => {
        var Token;

        beforeEach(async function () {
            Token = await SCToken.deployed();
        });

        it("Has a name", async () => {
            (await Token.name()).should.equal(token.name);
        });

        it("Has a symbol", async () => {
            (await Token.symbol()).should.equal(token.symbol);
        });

        it("Has 18 decimals", async function () {
            (await Token.decimals()).should.be.a.bignumber.that.equals(new BN(token.decimals));
        });

        it("Has correct initial supply", async () => {
            let actual = await Token.totalSupply();
            actual.should.be.a.bignumber.that.equals(token.totalSupply);
        });

        it("has enough balance for 3rd parties", async () => {
            let actual = await Token.balanceOf(accounts[0]);
            actual.should.be.a.bignumber.that.equals(new BN(util.toWei("12000000", "ether")));
        });

        it("Only transfers in wei", async () => {
            let sender = accounts[0];
            let receiver = accounts[1];
            let amount = new BN(10).pow(new BN(token.decimals)).mul(new BN(1));
            await Token.transfer(receiver, amount, { from: sender });
            let actual = await Token.balanceOf(receiver);
            actual.should.be.a.bignumber.that.equals(amount);
        });
    });

    describe("once deployed", async () => {
        let Token;
        let admin = accounts[0];
        let governor = accounts[0];
        let freetaxer = accounts[1];
        let nonfreetaxer = accounts[2];
        let sellingaddress = accounts[3];
        let nonadmin = accounts[4];

        beforeEach(async function () {
            Token = await SCToken.new(
                token.name,
                token.symbol);
        });

        describe("access control", async () => {
            it("pause/unpause", async () => {
                await truffleAssert.passes(Token.pause({ from: admin }));
                await truffleAssert.fails(Token.pause({ from: nonadmin }));

                await truffleAssert.passes(Token.unpause({ from: admin }));
                await truffleAssert.fails(Token.unpause({ from: nonadmin }));
            });

            it("set governor", async () => {
                await truffleAssert.passes(Token.setGovernor(governor, {
                    from: admin,
                }));

                await truffleAssert.fails(Token.setGovernor(governor, {
                    from: nonadmin,
                }));
            });

            it("set tax", async () => {
                await truffleAssert.passes(Token.setTax(1, {
                    from: admin,
                }));

                await truffleAssert.fails(Token.setTax(1, {
                    from: nonadmin,
                }));
            });

            it("add/remove selling address", async () => {
                await truffleAssert.passes(Token.addSellingAddress(sellingaddress, { from: admin }));
                await truffleAssert.fails(Token.addSellingAddress(sellingaddress, { from: nonadmin }));

                await truffleAssert.passes(Token.removeSellingAddress(sellingaddress, { from: admin }));
                await truffleAssert.fails(Token.removeSellingAddress(sellingaddress, { from: nonadmin }));
            });

            it("add/remove no tax address", async () => {
                await truffleAssert.passes(Token.addNoTaxAddress(freetaxer, { from: admin }));
                await truffleAssert.fails(Token.addNoTaxAddress(freetaxer, { from: nonadmin }));

                await truffleAssert.passes(Token.removeNoTaxAddress(freetaxer, { from: admin }));
                await truffleAssert.fails(Token.removeNoTaxAddress(freetaxer, { from: nonadmin }));
            });
        });

        describe("selling tax", async () => {
            let snapshotId;
            beforeEach(async () => {
                let snapshot = await timeMachine.takeSnapshot();
                snapshotId = snapshot['result'];

                await truffleAssert.passes(Token.addNoTaxAddress(freetaxer, { from: admin }));
                await truffleAssert.passes(Token.transfer(freetaxer, 200, { from: admin }));
                await truffleAssert.passes(Token.transfer(nonfreetaxer, 200, { from: admin }));
            });

            afterEach(async () => {
                await timeMachine.revertToSnapshot(snapshotId);
            });

            describe("when transferring before adding selling address", async () => {
                it("free-taxers can still do it", async () => {
                    await truffleAssert.passes(Token.transfer(nonadmin, 100, { from: freetaxer }));
                    await truffleAssert.passes(Token.transfer(nonadmin, 100, { from: admin }));
                });

                it("others cannot do it", async () => {
                    await truffleAssert.fails(Token.transfer(nonadmin, 100, { from: nonfreetaxer }));
                });
            });

            describe("before TGE + 1 hour", async () => {
                beforeEach(async () => {
                    await truffleAssert.passes(Token.addSellingAddress(sellingaddress, { from: admin }));
                });

                describe("using transfer", async () => {
                    it("no tax for normal transfer", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(nonadmin);
                        await truffleAssert.passes(Token.transfer(nonadmin, 100, { from: nonfreetaxer }));
                        let balanceAfter = await Token.balanceOf(nonadmin);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("no tax for free-taxers when transfering to selling address", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(sellingaddress);
                        await truffleAssert.passes(Token.transfer(sellingaddress, 100, { from: freetaxer }));
                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("tax for non-free-taxers when transfering to selling address", async () => {
                        // tax is 8%
                        let balanceFromBefore = await Token.balanceOf(nonfreetaxer);
                        let balanceToBefore = await Token.balanceOf(sellingaddress);
                        let balanceGovernorBefore = await Token.balanceOf(governor);
                        await truffleAssert.passes(Token.transfer(sellingaddress, 100, { from: nonfreetaxer }));

                        let balanceFromAfter = await Token.balanceOf(nonfreetaxer);
                        balanceFromAfter.should.be.a.bignumber.that.equals(balanceFromBefore.sub(new BN(100)));

                        let balanceToAfter = await Token.balanceOf(sellingaddress);
                        balanceToAfter.should.be.a.bignumber.that.equals(balanceToBefore.add(new BN(92)));
                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore.add(new BN(8)));
                    });
                });

                describe("using transferFrom", async () => {
                    beforeEach(async () => {
                        await truffleAssert.passes(Token.approve(admin, 200, { from: freetaxer }));
                        await truffleAssert.passes(Token.approve(admin, 200, { from: nonfreetaxer }));
                    });
                    it("no tax for normal transfer", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(nonadmin);
                        await truffleAssert.passes(Token.transferFrom(nonfreetaxer, nonadmin, 100, { from: admin }));
                        let balanceAfter = await Token.balanceOf(nonadmin);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("no tax for free-taxers when transfering to selling address", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(sellingaddress);
                        await truffleAssert.passes(Token.transferFrom(freetaxer, sellingaddress, 100, { from: admin }));
                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("tax for non-free-taxers when transfering to selling address", async () => {
                        // default tax is 8%
                        let balanceFromBefore = await Token.balanceOf(nonfreetaxer);
                        let balanceToBefore = await Token.balanceOf(sellingaddress);
                        let balanceGovernorBefore = await Token.balanceOf(governor);
                        await truffleAssert.passes(Token.transferFrom(nonfreetaxer, sellingaddress, 100, { from: admin }));

                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceToBefore.add(new BN(92)));
                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore.add(new BN(8)));
                    });
                });
            });

            describe("after TGE + 1 hour", async () => {
                beforeEach(async () => {
                    await truffleAssert.passes(Token.addSellingAddress(sellingaddress, { from: admin }));
                    await timeMachine.advanceBlockAndSetTime(TGE_PLUS_1_HOUR + 1);
                });

                describe("using transfer", async () => {
                    it("no tax for normal transfer", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(nonadmin);
                        await truffleAssert.passes(Token.transfer(nonadmin, 100, { from: nonfreetaxer }));
                        let balanceAfter = await Token.balanceOf(nonadmin);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("no tax for free-taxers when transfering to selling address", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(sellingaddress);
                        await truffleAssert.passes(Token.transfer(sellingaddress, 100, { from: freetaxer }));
                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("tax for non-free-taxers when transfering to selling address", async () => {
                        // default tax is 2%
                        let balanceFromBefore = await Token.balanceOf(nonfreetaxer);
                        let balanceToBefore = await Token.balanceOf(sellingaddress);
                        let balanceGovernorBefore = await Token.balanceOf(governor);
                        await truffleAssert.passes(Token.transfer(sellingaddress, 100, { from: nonfreetaxer }));

                        let balanceFromAfter = await Token.balanceOf(nonfreetaxer);
                        balanceFromAfter.should.be.a.bignumber.that.equals(balanceFromBefore.sub(new BN(100)));

                        let balanceToAfter = await Token.balanceOf(sellingaddress);
                        balanceToAfter.should.be.a.bignumber.that.equals(balanceToBefore.add(new BN(98)));
                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore.add(new BN(2)));
                    });
                });

                describe("using transferFrom", async () => {
                    beforeEach(async () => {
                        await truffleAssert.passes(Token.approve(admin, 200, { from: freetaxer }));
                        await truffleAssert.passes(Token.approve(admin, 200, { from: nonfreetaxer }));
                    });
                    it("no tax for normal transfer", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(nonadmin);
                        await truffleAssert.passes(Token.transferFrom(nonfreetaxer, nonadmin, 100, { from: admin }));
                        let balanceAfter = await Token.balanceOf(nonadmin);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("no tax for free-taxers when transfering to selling address", async () => {
                        let balanceGovernorBefore = await Token.balanceOf(governor);

                        let balanceBefore = await Token.balanceOf(sellingaddress);
                        await truffleAssert.passes(Token.transferFrom(freetaxer, sellingaddress, 100, { from: admin }));
                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(100)));

                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore);
                    });

                    it("tax for non-free-taxers when transfering to selling address", async () => {
                        // default tax is 2%
                        let balanceBefore = await Token.balanceOf(sellingaddress);
                        let balanceGovernorBefore = await Token.balanceOf(governor);
                        await truffleAssert.passes(Token.transferFrom(nonfreetaxer, sellingaddress, 100, { from: admin }));

                        let balanceAfter = await Token.balanceOf(sellingaddress);
                        balanceAfter.should.be.a.bignumber.that.equals(balanceBefore.add(new BN(98)));
                        let balanceGovernorAfter = await Token.balanceOf(governor);
                        balanceGovernorAfter.should.be.a.bignumber.that.equals(balanceGovernorBefore.add(new BN(2)));
                    });
                });
            });
        });
    });
})
