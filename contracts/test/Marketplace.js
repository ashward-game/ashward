const SCNFT = artifacts.require("NFT")
const SCToken = artifacts.require("Token")
const Martketplace = artifacts.require("Marketplace")

const token = require("../token");
const nft = require("../nft");
const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const { ZERO_ADDRESS } = constants;


contract("Marketplace", async (accounts) => {
    let NFT, Token, Market;
    let seller = accounts[0];
    let buyer = accounts[1];
    let keeper = accounts[2];
    let balance = 100;
    let itemPrice = balance;
    let buyerOriginalBalance;
    let sellerOriginalBalance;
    let tokens = [
        {
            metadata: "forsale.json",
            isForSale: true,
            price: 100
        },
        {
            metadata: "notforsale.json",
            isForSale: false,
            price: 0
        },
        {
            metadata: "notowned.json",
            isForSale: false,
            price: 0
        }
    ]

    describe("during deployment", async () => {
        beforeEach(async function () {
            Token = await SCToken.new(token.name, token.symbol);
            NFT = await SCNFT.new(nft.name, nft.symbol, nft.ipfs_url);
        });

        it("requires a non-null ERC1363 token for payment", async () => {
            await expectRevert.unspecified(Martketplace.new.estimateGas(ZERO_ADDRESS, NFT.address));
        });

        it("requires a non-null ERC721 token address", async () => {
            await expectRevert.unspecified(Martketplace.new.estimateGas(Token.address, ZERO_ADDRESS));
        });
    });

    describe("once deployed", async () => {
        async function openOffer(seller, tokenId, price) {
            const priceInWei = helper.wei(price);
            const data = helper.uintToBytes32(priceInWei);
            await NFT.safeTransferFrom(seller, Market.address, tokenId, data);
        }

        async function purchase(buyer, tokenId, price) {
            const priceInWei = helper.wei(price);
            const data = helper.uintToBytes32(tokenId);
            await Token.methods["approveAndCall(address,uint256,bytes)"](Market.address, priceInWei, data, { from: buyer });
        }

        beforeEach(async function () {
            Token = await SCToken.new(token.name, token.symbol);
            NFT = await SCNFT.new(nft.name, nft.symbol, nft.ipfs_url);
            Market = await Martketplace.new(Token.address, NFT.address);
            for (i = 0; i < 3; i++) {
                await NFT.mint(tokens[i].metadata, { from: seller });
            }
            await NFT.safeTransferFrom(seller, keeper, 2, { from: seller });
            await Token.transfer(buyer, helper.wei(balance), { from: seller });
            sellerOriginalBalance = await Token.balanceOf(seller);
            buyerOriginalBalance = await Token.balanceOf(buyer);
        });

        it("allows to trade", async () => {
            await openOffer(seller, 0, itemPrice);
            let actual = await NFT.ownerOf(0);
            actual.should.equal(Market.address);
            actual = await Market.isForSale(0);
            actual.should.equal(true);
            actual = await Market.priceOf(0);
            actual.should.be.a.bignumber.that.equals(helper.wei(itemPrice));

            await purchase(buyer, 0, itemPrice);
            actual = await NFT.ownerOf(0);
            actual.should.equal(buyer);
            actual = await Market.isForSale(0);
            actual.should.equal(false);
            await truffleAssert.fails(Market.priceOf(0));

            buyerBalance = await Token.balanceOf(buyer);
            buyerBalance.should.be.a.bignumber.that.equals(
                buyerOriginalBalance.sub(helper.wei(itemPrice))
            );
            sellerBalance = await Token.balanceOf(seller);
            sellerBalance.should.be.a.bignumber.that.equals(
                sellerOriginalBalance.add(helper.wei(itemPrice))
            );
        });

        it("requires correct format of data", async () => {
            await truffleAssert.fails(NFT.safeTransferFrom(seller, Market.address, 0, [0x01, 0x02]));
            try {
                await NFT.safeTransferFrom(seller, Market.address, 0, [0x01, 0x02]);
            } catch (error) {
                assert.equal(error.reason, "TypeConversion: bytesToUint256 requires input of type bytes with length 32");
            }
        });

        describe("when offering", async () => {
            it("does not allow to give offer for the same token twice", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(openOffer(seller, 0, itemPrice));
                try {
                    await openOffer(seller, 0, itemPrice);
                } catch (error) {
                    assert.equal(error.reason, "ERC721: transfer caller is not owner nor approved");
                }
            });

            it("does not allow to offer others' tokens", async () => {
                await truffleAssert.fails(openOffer(buyer, 0, itemPrice));
                try {
                    await openOffer(buyer, 0, itemPrice);
                } catch (error) {
                    assert.equal(error.reason, "ERC721: transfer of token that is not own");
                }
            })
        });

        describe("when canceling", async () => {
            it("allows to cancel listed offers", async () => {
                await openOffer(seller, 0, itemPrice);

                await Market.cancelOffer(0, { from: seller });
                let actual = await NFT.ownerOf(0);
                actual.should.equal(seller);
                actual = await Market.isForSale(0);
                actual.should.equal(false);
                await truffleAssert.fails(Market.priceOf(0));
            });

            it("does not allow to cancel others' offers", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(Market.cancelOffer(0, { from: buyer }));
                try {
                    await Market.cancelOffer(0, { from: buyer });
                } catch (error) {
                    assert.equal(error.reason, "Marketplace: caller is not the owner");
                }
            });

            it("does not allow to cancel non-listed offers", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(Market.cancelOffer(1, { from: seller }));
                try {
                    await Market.cancelOffer(1, { from: seller });
                } catch (error) {
                    assert.equal(error.reason, "Marketplace: token is not for sale");
                }
            });
        });

        describe("when buying", async () => {
            it("does not allow to buy non-listed offers", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(purchase(buyer, 1, itemPrice));
                try {
                    await purchase(buyer, 1, itemPrice);
                } catch (error) {
                    assert.equal(error.reason, "Marketplace: token is not for sale");
                }
            });

            it("does not allow to buy your own offers", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(purchase(seller, 0, itemPrice));
                try {
                    await purchase(seller, 0, itemPrice);
                } catch (error) {
                    assert.equal(error.reason, "Marketplace: buyer is the owner");
                }
            });

            it("does not allow to buy with a lower price", async () => {
                await openOffer(seller, 0, itemPrice);
                await truffleAssert.fails(purchase(buyer, 0, itemPrice - 1));
                try {
                    await purchase(buyer, 0, itemPrice - 1);
                } catch (error) {
                    assert.equal(error.reason, "Marketplace: must pay exactly the price");
                }
            });
        });
    });
})