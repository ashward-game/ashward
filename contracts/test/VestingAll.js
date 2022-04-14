const SCToken = artifacts.require("Token");
const SCVestingAdvisory = artifacts.require("VestingAdvisory");
const SCVestingIDO = artifacts.require("VestingIDO");
const SCVestingLiquidity = artifacts.require("VestingLiquidity");
const SCVestingMarketing = artifacts.require("VestingMarketing");
const SCVestingPlay2Earn = artifacts.require("VestingPlay2Earn");
const SCVestingPrivate = artifacts.require("VestingPrivate");
const SCVestingReserve = artifacts.require("VestingReserve");
const SCVestingStaking = artifacts.require("VestingStaking");
const SCVestingStrategicPartner = artifacts.require("VestingStrategicPartner");
const SCVestingTeam = artifacts.require("VestingTeam");
const token = require("../token");
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
const util = web3.utils;
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const { assert } = require("chai");
const { ZERO_ADDRESS } = constants;
const vesting = require("../vesting");

contract("VestingAll", async (accounts) => {
    describe("once deployed", async () => {
        var Token;

        beforeEach(async function () {
            Token = await SCToken.deployed();
        });

        context("has correct balance in each vesting pool", async () => {
            it("VestingAdvisory", async () => {
                var contract = await SCVestingAdvisory.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Advisory);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("10000000", "ether")));
            });
            it("VestingIDO", async () => {
                var contract = await SCVestingIDO.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.IDO);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("3000000", "ether")));
            });
            it("VestingLiquidity", async () => {
                var contract = await SCVestingLiquidity.deployed();
                let actual = await Token.balanceOf(contract.address);
                let expected = vesting.Liquidity.mul(new BN(100 - 20)).div(new BN(100));
                actual.should.be.a.bignumber.that.equals(expected);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("68000000", "ether")));
            });
            it("VestingMarketing", async () => {
                var contract = await SCVestingMarketing.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Marketing);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("120000000", "ether")));
            });
            it("VestingPlay2Earn", async () => {
                var contract = await SCVestingPlay2Earn.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Play2Earn);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("300000000", "ether")));
            });
            it("VestingPrivate", async () => {
                var contract = await SCVestingPrivate.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Private);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("100000000", "ether")));
            });
            it("VestingReserve", async () => {
                var contract = await SCVestingReserve.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Reserve);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("80000000", "ether")));
            });
            it("VestingStaking", async () => {
                var contract = await SCVestingStaking.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Staking);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("150000000", "ether")));
            });
            it("VestingStrategicPartner", async () => {
                var contract = await SCVestingStrategicPartner.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.StrategicPartner);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("20000000", "ether")));
            });
            it("VestingTeam", async () => {
                var contract = await SCVestingTeam.deployed();
                let actual = await Token.balanceOf(contract.address);
                actual.should.be.a.bignumber.that.equals(vesting.Team);
                actual.should.be.a.bignumber.that.equals(new BN(util.toWei("120000000", "ether")));
            });
        });
    });
}); 