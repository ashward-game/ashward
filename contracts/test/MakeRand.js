
const MockMakeRand = artifacts.require("MakeRandMock")

const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const { assert } = require("chai");

contract("MakeRand", async (accounts) => {
    let MakeRand, publicKey, privateKey;

    beforeEach(async function () {
        let key = helper.generateKey();
        privateKey = key.privateKey;
        publicKey = key.address;

        MakeRand = await MockMakeRand.new(publicKey);
    });
    describe("when committing", async () => {
        it("accepts fresh commitments", async () => {
            let random = helper.random();
            let sign = helper.sign(privateKey, random);
            let hash = sign.messageHash;
            assert.equal(hash, helper.hash(random));

            let signature = sign.signature;
            let pRandom = helper.random();

            let event = await MakeRand.commit(hash, signature, pRandom, { from: accounts[0] });
            truffleAssert.eventEmitted(event, "Committed", (ev) => {
                assert.equal(ev.serverHash, hash);
                assert.equal(ev.clientRandom, pRandom);
                return true;
            });
        });

        it("does not accept submitted commitments", async () => {
            let random = helper.random();
            let sign = helper.sign(privateKey, random);
            let hash = sign.messageHash;
            let signature = sign.signature;
            let pRandom = helper.random();

            let event = await MakeRand.commit(hash, signature, pRandom, { from: accounts[0] });
            truffleAssert.eventEmitted(event, "Committed", (ev) => {
                assert.equal(ev.serverHash, hash);
                assert.equal(ev.clientRandom, pRandom);
                return true;
            });

            try {
                await MakeRand.commit(hash, signature, pRandom, { from: accounts[0] });
            } catch (error) {
                assert.equal(error.reason, "MakeRand: hash value already exists");
            }
        });

        it("does not accepts commitments with invalid signatures", async () => {
            let pk = helper.generateKey().privateKey;
            let random = helper.random();
            let sign = helper.sign(pk, random);
            let hash = sign.messageHash;
            let signature = sign.signature;
            let pRandom = helper.random();

            try {
                await MakeRand.commit(hash, signature, pRandom, { from: accounts[0] });
            } catch (error) {
                assert.equal(error.reason, "MakeRand: signature is invalid");
            }
        });

        it("returns correct value of client's randomness", async () => {
            let random = helper.random();
            let sign = helper.sign(privateKey, random);
            let hash = sign.messageHash;
            let signature = sign.signature;
            let pRandom = helper.random();
            await MakeRand.commit(hash, signature, pRandom, { from: accounts[0] });
            let cRand = await MakeRand.getClientRandom(hash, { from: accounts[0] });
            assert.equal(cRand, pRandom);
        });

        it("client random must be not equal zero", async () => {
            let random = helper.random();
            let pRandom = web3.utils.toHex(0);
            assert.equal(pRandom, "0x0");
            let sign = helper.sign(privateKey, random);
            let hash = sign.messageHash;
            let signature = sign.signature;


            await truffleAssert.fails(MakeRand.commit(hash, signature, pRandom, { from: accounts[0] }), truffleAssert.ErrorType.REVERT, "MakeRand: client random must be not equal 0");
        });
    });
});