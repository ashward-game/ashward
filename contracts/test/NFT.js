const SCNFT = artifacts.require("NFT")
const truffleAssert = require("truffle-assertions");
const chai = require("chai");
const nft = require("../nft");
const { assert } = require("chai");
const { ZERO_ADDRESS } = require("@openzeppelin/test-helpers/src/constants");
const BN = web3.utils.BN;
var should = chai.should();
chai.use(require("chai-bn")(BN));

contract("NFT", async (accounts) => {
    let NFT;
    let owner = accounts[0];
    let client = accounts[1];
    let minter = accounts[2];
    let adminRole = "0x0000000000000000000000000000000000000000000000000000000000000000";
    let minterRole = "0x9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6";

    let metadata = [
        "1.json",
        "2.json",
        "3.json"
    ];

    beforeEach(async function () {
        NFT = await SCNFT.new(nft.name, nft.symbol, nft.ipfs_url);
    });

    it("Cannot be minted by non-minter", async () => {
        try {
            await NFT.mint(metadata[0], { from: owner });
        } catch (error) {
            assert.equal(error.reason, "AccessControl: account " + owner.toLowerCase() + " is missing role " + minterRole);
        }
    })

    it("Cannot set minter if account not admin", async () => {
        try {
            await NFT.setupMinter(minter, { from: client });
        } catch (error) {
            assert.equal(error.reason, "AccessControl: account " + client.toLowerCase() + " is missing role " + adminRole);
        }
    })

    beforeEach(async function () {
        await NFT.setupMinter(minter, { from: owner });
    });

    it("Cannot be minted if TokenURI is empty", async () => {
        try {
            await NFT.mint("", { from: minter });
        } catch (error) {
            assert.equal(error.reason, "NFT: token's URI must not be empty");
        }
    })

    it("Cannot be minted if owner is address 0", async () => {
        try {
            await NFT.mintAndTransfer(metadata[0], ZERO_ADDRESS, { from: minter });
        } catch (error) {
            assert.equal(error.reason, "NFT: owner must not be 0");
        }
    })

    it("Can be minted by minter", async () => {
        let eventMinted = await NFT.mint(metadata[0], { from: minter });
        // Token 0 belongs to Owner
        truffleAssert.eventEmitted(
            eventMinted,
            "Transfer",
            (ev) => {
                assert.equal(ev.from, "0x0000000000000000000000000000000000000000");
                assert.equal(ev.to, minter);
                assert.equal(ev.tokenId, 0);
                return true;
            },
            "Minted event should have triggered");
    })

    it("Cannot be transferred by non-owner", async () => {
        // mint
        await NFT.mint(metadata[0], { from: minter });

        // not owner
        await NFT.approve(client, 0, { from: minter });
        try {
            // Token 0 belongs to Minter
            await NFT.safeTransferFrom(client, minter, 0, { from: client });
        } catch (error) {
            assert.equal(error.reason, "ERC721: transfer of token that is not own");
        }
    })

    it("Can be transferred by owner", async () => {
        // mint
        await NFT.mint(metadata[0], { from: minter });

        await NFT.approve(client, 0, { from: minter });

        let eventTransfer = await NFT.safeTransferFrom(minter, client, 0, { from: minter });
        // Token 0 belongs to client
        truffleAssert.eventEmitted(
            eventTransfer,
            "Transfer",
            (ev) => {
                assert.equal(ev.from, minter);
                assert.equal(ev.to, client);
                assert.equal(ev.tokenId, 0);
                return true;
            },
            "Transfer event should have triggered");
    })

    it("Mint multiple NFT by owner", async () => {
        for (let i = 0; i < 3; i++) {
            // mint
            let eventTransfer = await NFT.mint(metadata[i], { from: minter });

            truffleAssert.eventEmitted(
                eventTransfer,
                "Transfer",
                (ev) => {
                    assert.equal(ev.from, "0x0000000000000000000000000000000000000000");
                    assert.equal(ev.to, minter);
                    assert.equal(ev.tokenId, i);
                    return true;
                },
                "Transfer event should have triggered");
        }
    })

    it("Mint NFT and transfer by owner", async () => {
        // mint
        let eventTransfer = await NFT.mintAndTransfer(metadata[1], client, { from: minter });

        truffleAssert.eventEmitted(
            eventTransfer,
            "Transfer",
            (ev) => {
                assert.equal(ev.from, "0x0000000000000000000000000000000000000000");
                assert.equal(ev.to, client);
                assert.equal(ev.tokenId, 0);
                return true;
            },
            "Transfer event should have triggered");
    })
})
