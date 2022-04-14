const SCMock = artifacts.require("OwnerTransferableMock")
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const { assert } = require("chai");
const truffleAssertions = require("truffle-assertions");
const { inLogs } = require("@openzeppelin/test-helpers/src/expectEvent");

contract("OwnerTransferable", async (accounts) => {
    let Mock;
    let owner = accounts[0];
    let client = accounts[1];
    let price = new BN(10);

    beforeEach(async function () {
        Mock = await SCMock.new();
    });

    it("can be called by owner only", async () => {
        try {
            await Mock.transferToOwner({ from: client });
        } catch (error) {
            assert.equal(error.reason, "Ownable: caller is not the owner");
        }
    })

    it("does not transfer if balance is zero", async () => {
        try {
            await Mock.transferToOwner({ from: owner });
        } catch (error) {
            assert.equal(error.reason, "OwnerTransferable: empty balance");
        }
    })

    it("would transfer BNB/ETH back to the owner", async () => {
        let ownerOrigBalance = new BN(await web3.eth.getBalance(owner));

        await Mock.send({ from: client, value: price });

        let transaction = await Mock.transferToOwner({ from: owner });
        const gasUsed = new BN(transaction.receipt.gasUsed);
        const tx = await web3.eth.getTransaction(transaction.tx);
        const gasPrice = new BN(tx.gasPrice);


        // finalOwner = ownerOrigBalance + price 
        let actualOwnerBalance = new BN(await web3.eth.getBalance(owner));
        let gasCost = new BN(gasPrice).mul(new BN(gasUsed));
        actualOwnerBalance.should.be.a.bignumber.that.equal(ownerOrigBalance.add(price).sub(gasCost));
    })
});