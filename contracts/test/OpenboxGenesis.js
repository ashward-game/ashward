const SCNFT = artifacts.require("NFT")
const SCOpenbox = artifacts.require("OpenboxGenesis")

const chai = require("chai");
var should = require("chai").should();
const { BN, constants, expectRevert, expectEvent } = require('@openzeppelin/test-helpers');
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require('./helpers/helper');
const { assert } = require("chai");
const { ZERO_ADDRESS } = constants;
const nft = require("../nft");
const openboxGenesis = require("../openboxGenesis");
const { web3 } = require("@openzeppelin/test-helpers/src/setup");

contract("OpenboxGenesis", async (accounts) => {
    let Openbox, publicKey, NFT, privateKey;
    let owner = accounts[0];
    let subscriber = accounts[1];
    let subscriberNotInWhitelist = accounts[3];
    let opener = accounts[2];
    let admin_role = "0x0000000000000000000000000000000000000000000000000000000000000000";
    let subscribe_role = "0x3f483399a73bbfbc7e47cea702709b2441abfc4e8152100709ca14556e321303";
    let open_role = "0x4c419879f6b579c6ec1e6a6a12f4bd58e14d8daf5d4cff9fcd2dffb91775d339";
    let priceRare = new BN(openboxGenesis.rareBoxPrice);
    let priceLegend = new BN(openboxGenesis.legendBoxPrice);
    let priceMyth = new BN(openboxGenesis.mythBoxPrice);
    let metadata = "metadata1.json";
    let rareBox = 0;
    let legendBox = 1;
    let mythBox = 2;
    openboxGenesis.numLegendBoxes = 1;
    let priceDecimals = helper.wei(0.5);

    function commit() {
        let random = helper.random();
        let sign = helper.sign(privateKey, random);
        let hash = sign.messageHash;
        assert.equal(hash, helper.hash(random));

        let signature = sign.signature;
        let cRandom = helper.random();
        return {
            random,
            cRandom,
            hash,
            signature
        };
    }

    describe("once deployed", async () => {
        beforeEach(async function () {
            let key = helper.generateKey();
            privateKey = key.privateKey;
            publicKey = key.address;
            NFT = await SCNFT.new(nft.name, nft.symbol, nft.ipfs_url);
            Openbox = await SCOpenbox.new(publicKey,
                NFT.address,
                openboxGenesis.numRareBoxes,
                openboxGenesis.rareBoxPrice,
                openboxGenesis.numLegendBoxes,
                openboxGenesis.legendBoxPrice,
                openboxGenesis.numMythBoxes,
                openboxGenesis.mythBoxPrice);
            await NFT.setupMinter(Openbox.address, { from: owner });
        });

        it("has correct information on number of boxes for each type and their prices", async () => {
            // amount box
            let numRare = await Openbox.numRareBoxes();
            numRare.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numRareBoxes));

            let numLegend = await Openbox.numLegendBoxes();
            numLegend.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numLegendBoxes));

            let numMyth = await Openbox.numMythBoxes();
            numMyth.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numMythBoxes));

            // price box
            let priceRare = await Openbox.rareBoxPrice();
            priceRare.should.be.a.bignumber.that.equals(openboxGenesis.rareBoxPrice);

            let priceLegend = await Openbox.legendBoxPrice();
            priceLegend.should.be.a.bignumber.that.equals(openboxGenesis.legendBoxPrice);

            let priceMyth = await Openbox.mythBoxPrice();
            priceMyth.should.be.a.bignumber.that.equals(openboxGenesis.mythBoxPrice);
        });

        describe("adding subscribers", async () => {
            describe("can be done by admin", async () => {
                it("for single subscriber", async () => {
                    let event = await Openbox.addSubscriber(subscriber, { from: owner });
                    truffleAssert.eventEmitted(event, "SubscriberRegistered", (evt) => {
                        assert.equal(evt.subscriber, subscriber);
                        return true;
                    });
                });

                it("for multiple subscribers", async () => {
                    let event = await Openbox.addSubscribers([subscriber, owner], { from: owner });
                    truffleAssert.eventEmitted(event, "SubscriberRegistered", (evt, i) => {
                        if (i == 0) {
                            assert.equal(evt.subscriber, subscriber);
                        } else {
                            assert.equal(evt.subscriber, owner);
                        }
                        return true;
                    });
                });
            });

            describe("cannot be done by non-admin", async () => {
                it("for single subscriber", async () => {
                    try {
                        await Openbox.addSubscriber(owner, { from: subscriber });
                    } catch (error) {
                        assert.equal(error.reason, "AccessControl: account " + subscriber.toLowerCase() + " is missing role " + admin_role);
                    }
                });

                it("for multiple subscribers", async () => {
                    try {
                        await Openbox.addSubscribers([owner, subscriber], { from: subscriber });
                    } catch (error) {
                        assert.equal(error.reason, "AccessControl: account " + subscriber.toLowerCase() + " is missing role " + admin_role);
                    }
                });
            });
        });

        describe("when buying boxes", async () => {
            it("is OK for rare boxes", async () => {
                let { random, cRandom, hash, signature } = commit();
                let subscriberBalance = new BN(await web3.eth.getBalance(subscriber));
                let openboxBalance = new BN(await web3.eth.getBalance(Openbox.address));
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));

                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let numRare = await Openbox.numRareBoxes();
                numRare.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numRareBoxes - 1));

                const gasUsed = new BN(event.receipt.gasUsed);
                const tx = await web3.eth.getTransaction(event.tx);
                const gasPrice = new BN(tx.gasPrice);
                const final = new BN(await web3.eth.getBalance(subscriber));

                // subscriberBalance = final + (gasPrice * gasUsed) + price
                subscriberBalance.should.be.a.bignumber.that.equals(
                    final.add(gasPrice.mul(gasUsed)).add(priceRare)
                );
                // finalOB = openboxBalance + price
                const finalOB = new BN(await web3.eth.getBalance(Openbox.address));
                finalOB.should.be.a.bignumber.that.equals(
                    openboxBalance.add(priceRare)
                );
            });

            it("is OK for legend boxes", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));

                let subscriberBalance = new BN(await web3.eth.getBalance(subscriber));
                let openboxBalance = new BN(await web3.eth.getBalance(Openbox.address));

                let event = await Openbox.buyBox(legendBox, hash, signature, cRandom, { from: subscriber, value: priceLegend });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, legendBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let numLegend = await Openbox.numLegendBoxes();
                numLegend.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numLegendBoxes - 1));

                const gasUsed = new BN(event.receipt.gasUsed);
                const tx = await web3.eth.getTransaction(event.tx);
                const gasPrice = new BN(tx.gasPrice);
                const final = new BN(await web3.eth.getBalance(subscriber));

                subscriberBalance.should.be.a.bignumber.that.equals(
                    final.add(gasPrice.mul(gasUsed)).add(priceLegend)
                );

                const finalOB = new BN(await web3.eth.getBalance(Openbox.address));
                finalOB.should.be.a.bignumber.that.equals(
                    openboxBalance.add(priceLegend)
                );
            });

            it("is OK for mythical boxes", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));

                let subscriberBalance = new BN(await web3.eth.getBalance(subscriber));
                let openboxBalance = new BN(await web3.eth.getBalance(Openbox.address));

                let event = await Openbox.buyBox(mythBox, hash, signature, cRandom, { from: subscriber, value: priceMyth });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, mythBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let numMyth = await Openbox.numMythBoxes();
                numMyth.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numMythBoxes - 1));

                const gasUsed = new BN(event.receipt.gasUsed);
                const tx = await web3.eth.getTransaction(event.tx);
                const gasPrice = new BN(tx.gasPrice);
                const final = new BN(await web3.eth.getBalance(subscriber));

                subscriberBalance.should.be.a.bignumber.that.equals(
                    final.add(gasPrice.mul(gasUsed)).add(priceMyth)
                );

                const finalOB = new BN(await web3.eth.getBalance(Openbox.address));
                finalOB.should.be.a.bignumber.that.equals(
                    openboxBalance.add(priceMyth)
                );
            });

            it("does not allow to buy when the buyer is not a subscriber", async () => {
                let { random, cRandom, hash, signature } = commit();
                try {
                    await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: owner, value: priceRare });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: either caller is not in the whitelist or public sell is not ready");
                }
            });

            it("does not allow to buy rare boxes if the buyer pays less than the price", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                try {
                    await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: 0 });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: must pay exactly the price");
                }
            });

            it("does not allow to buy legend boxes if the buyer pays less than the price", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                try {
                    await Openbox.buyBox(legendBox, hash, signature, cRandom, { from: subscriber, value: 0 });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: must pay exactly the price");
                }
            });

            it("does not allow to buy mythical boxes if the buyer pays less than the price", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                try {
                    await Openbox.buyBox(mythBox, hash, signature, cRandom, { from: subscriber, value: 0 });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: must pay exactly the price");
                }
            });

            it("does not allow to buy boxes with invalid type", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                await truffleAssert.fails(Openbox.buyBox(-1, hash, signature, cRandom, { from: subscriber, value: priceRare }));
                await truffleAssert.fails(Openbox.buyBox(3, hash, signature, cRandom, { from: subscriber, value: priceRare }));
            });

            it("does not allow to reuse the same server hash for more than 1 transaction", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                try {
                    await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                } catch (error) {
                    assert.equal(error.reason, "MakeRand: hash value already exists");
                }
            });

            it("does not allow a subscriber to buy more than his quota", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let commit2 = commit();
                await truffleAssert.passes(Openbox.buyBox(rareBox, commit2.hash, commit2.signature, commit2.cRandom, { from: subscriber, value: priceRare }));

                let commit3 = commit();
                try {
                    await Openbox.buyBox(rareBox, commit3.hash, commit3.signature, commit3.cRandom, { from: subscriber, value: priceRare });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: can only buy at most 2 boxes");
                }
            });

            it("does not allow a subscriber to buy more than his quota if addSubscriber 2 times", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let commit2 = commit();
                await truffleAssert.passes(Openbox.buyBox(rareBox, commit2.hash, commit2.signature, commit2.cRandom, { from: subscriber, value: priceRare }));

                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));

                let commit3 = commit();
                try {
                    await Openbox.buyBox(rareBox, commit3.hash, commit3.signature, commit3.cRandom, { from: subscriber, value: priceRare });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: can only buy at most 2 boxes");
                }
            });

            it("buy box with cRandom is zero", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));

                // client random equal 0
                cRandom = web3.utils.toHex(0);
                assert.equal(cRandom,"0x0");

                await truffleAssert.fails(Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare }),truffleAssert.ErrorType.REVERT,"MakeRand: client random must be not equal 0");
            });
        });

        describe("when opening boxes", async () => {
            it("allows only admin to open boxes", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                try {
                    await Openbox.openBox(hash, false, metadata, { from: subscriber });
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + subscriber.toLowerCase() + " is missing role " + open_role);
                }
            });

            it("does not open if the box cannot be opened", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                await Openbox.setupOpener(opener, { from: owner });

                let event2 = await Openbox.openBox(hash, true, metadata, { from: opener });
                truffleAssert.eventEmitted(event2, "BoxOpened", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.isEmpty, true);
                    assert.equal(evt.tokenID, 0);
                    return true;
                });

                await truffleAssert.fails(NFT.ownerOf(0, { from: subscriber }));
            });

            it("opens if the box can be opened", async () => {
                let { random, cRandom, hash, signature } = commit();
                await truffleAssert.passes(Openbox.addSubscriber(subscriber, { from: owner }));
                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriber, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                await Openbox.setupOpener(opener, { from: owner });

                let event2 = await Openbox.openBox(hash, false, metadata, { from: opener });
                truffleAssert.eventEmitted(event2, "BoxOpened", (evt) => {
                    assert.equal(evt.buyer, subscriber);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.isEmpty, false);
                    assert.equal(evt.tokenID, 0);
                    return true;
                });

                let ownerNft = await NFT.ownerOf(0, { from: subscriber });
                assert.equal(ownerNft, subscriber);


                let meta = await NFT.tokenURI(0);
                assert.equal(meta, nft.ipfs_url + metadata);
            });
        });

        describe("price decimals", async () => {
            it("price is decimals", async () => {
                let key = helper.generateKey();
                privateKey = key.privateKey;
                publicKey = key.address;
                NFT = await SCNFT.new(nft.name, nft.symbol, nft.ipfs_url);

                Openbox = await SCOpenbox.new(publicKey,
                    NFT.address,
                    openboxGenesis.numRareBoxes,
                    priceDecimals,
                    openboxGenesis.numLegendBoxes,
                    openboxGenesis.legendBoxPrice,
                    openboxGenesis.numMythBoxes,
                    openboxGenesis.mythBoxPrice);

                let priceBoxRare = await Openbox.rareBoxPrice();
                assert.equal(priceBoxRare.toString(), priceDecimals.toString());
            });
        });

        describe("Change status public sell", async () => {
            it("not allows user to call publicSale", async () => {
                try {
                    await Openbox.publicSale({from: subscriber});
                } catch (error) {
                    assert.equal(error.reason, "AccessControl: account " + subscriber.toLowerCase() + " is missing role " + admin_role);
                }
            });

            it("allows only admin to call publicSale", async () => {
                const event = await Openbox.publicSale({from: owner});
                assert.equal(await Openbox.isPublicSale(), true);

                truffleAssert.eventEmitted(event, "PublicSaleOpened", () => {
                    return true;
                });
            });

            it("public sale is already opened", async () => {
               const event = await Openbox.publicSale({from: owner});
                assert.equal(await Openbox.isPublicSale(), true);

              truffleAssert.eventEmitted(event, "PublicSaleOpened", () => {
                return true;
              });

              try {
                  await Openbox.publicSale({from: owner});
              } catch (error) {
                  assert.equal(error.reason, "OpenboxGenesis: public sale is already opened");
              }
            });

        });

        describe("buy box when user not in whitelist after allow public sell", async () => {
            it("buy box not in whitelist", async () => {

                let {random, cRandom, hash, signature} = commit();
                //before admin change status => publicSell = false
                try {
                    await Openbox.buyBox(rareBox, hash, signature, cRandom, {from: subscriberNotInWhitelist, value: priceRare});
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: either caller is not in the whitelist or public sell is not ready");
                }

                await Openbox.publicSale({from: owner});

                //after admin change status => publicSell = true
                let subscriberNotInWhitelistBalance = new BN(await web3.eth.getBalance(subscriberNotInWhitelist));
                let openboxBalance = new BN(await web3.eth.getBalance(Openbox.address));

                const hasRole = await Openbox.hasRole(subscribe_role, subscriberNotInWhitelist);
                //check user not in whitelist
                assert.equal(hasRole, false);

                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, {
                    from: subscriberNotInWhitelist,
                    value: priceRare
                });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriberNotInWhitelist);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let numRare = await Openbox.numRareBoxes();
                numRare.should.be.a.bignumber.that.equals(new BN(openboxGenesis.numRareBoxes - 1));

                const gasUsed = new BN(event.receipt.gasUsed);
                const tx = await web3.eth.getTransaction(event.tx);
                const gasPrice = new BN(tx.gasPrice);
                const final = new BN(await web3.eth.getBalance(subscriberNotInWhitelist));

                // subscriberNotInWhitelistBalance = final + (gasPrice * gasUsed) + price
                subscriberNotInWhitelistBalance.should.be.a.bignumber.that.equals(
                    final.add(gasPrice.mul(gasUsed)).add(priceRare)
                );
                // finalOB = openboxBalance + price
                const finalOB = new BN(await web3.eth.getBalance(Openbox.address));
                finalOB.should.be.a.bignumber.that.equals(
                    openboxBalance.add(priceRare)
                );
            });

            it("user not in whitelist buy only 2 boxes", async () => {
                let { random, cRandom, hash, signature } = commit();

                await truffleAssert.fails(Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriberNotInWhitelist, value: priceRare }));

                await truffleAssert.passes(Openbox.publicSale({from:owner}));

                let event = await Openbox.buyBox(rareBox, hash, signature, cRandom, { from: subscriberNotInWhitelist, value: priceRare });
                truffleAssert.eventEmitted(event, "BoxBought", (evt) => {
                    assert.equal(evt.buyer, subscriberNotInWhitelist);
                    assert.equal(evt.boxGrade, rareBox);
                    assert.equal(evt.serverHash, hash);
                    assert.equal(evt.clientRandom, cRandom);
                    return true;
                });

                let commit2 = commit();
                await truffleAssert.passes(Openbox.buyBox(rareBox, commit2.hash, commit2.signature, commit2.cRandom, { from: subscriberNotInWhitelist, value: priceRare }));

                let commit3 = commit();
                try {
                    await Openbox.buyBox(rareBox, commit3.hash, commit3.signature, commit3.cRandom, { from: subscriberNotInWhitelist, value: priceRare });
                } catch (error) {
                    assert.equal(error.reason, "OpenboxGenesis: can only buy at most 2 boxes");
                }
            });
        });
    });
});
