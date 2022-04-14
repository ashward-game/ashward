const SCVestingTeam = artifacts.require("VestingTeam");
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

contract("VestingTeam", async (accounts) => {
    let VestingTeam, Token;
    let owner = accounts[0];
    let grantor = accounts[1];
    let beneficiary = accounts[2];
    let amount = new BN(helper.wei(111));
    const TGE_MILESTONE = 1647621000; // 18-03-2022 16:30:00 UTC

    describe("once deployed", async () => {
        beforeEach(async function () {
            Token = await SCToken.new(token.name, token.symbol);
            VestingTeam = await SCVestingTeam.new(Token.address);
            await Token.transfer(VestingTeam.address, vestingConfig.Team, { from: owner });
            await VestingTeam.setGrantor(grantor, { from: owner });
            await VestingTeam.addBeneficiaries([beneficiary], [amount], { from: grantor });

        });

        it("cannot claim TGE", async () => {
            await timeMachine.advanceBlockAndSetTime(TGE_MILESTONE);

            await truffleAssert.fails(VestingTeam.claimTGE({ from: beneficiary }));
            try {
                await VestingTeam.claimTGE({ from: beneficiary });
            } catch (error) {
                assert.equal(error.reason, "Vesting: no TGE tokens");
            }
        });
    });
});