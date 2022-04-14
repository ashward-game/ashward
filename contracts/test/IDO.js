const SCBusdMock = artifacts.require("BusdMock");
const SCIDO = artifacts.require("IDO");
const SCIDOMock = artifacts.require("IDOMock")

const chai = require("chai");
var should = require("chai").should();
const {
  BN,
  constants,
  expectRevert,
  expectEvent,
} = require("@openzeppelin/test-helpers");
chai.use(require("chai-bn")(BN));
const truffleAssert = require("truffle-assertions");
const helper = require("./helpers/helper");
const { assert } = require("chai");
const { ZERO_ADDRESS } = constants;

contract("IDO", async (accounts) => {
  let BusdMock;
  let IDO;
  let owner = accounts[0];
  let subscriber = accounts[1];
  let non_subscriber = accounts[2];
  let admin_role =
    "0x0000000000000000000000000000000000000000000000000000000000000000";
  let subscriber_role =
    "0x3f483399a73bbfbc7e47cea702709b2441abfc4e8152100709ca14556e321303";

  const package100 = 0;
  const package200 = 1;
  const package100Price = 100;
  const package200Price = 200;

  const Package100TokenAmount = "3333340000000000000000";
  const Package200TokenAmount = "6666670000000000000000";

  async function buyPackage100() {
    await truffleAssert.passes(BusdMock.approve(
      IDO.address,
      helper.wei(package100Price).toString(),
      { from: subscriber }
    ));

    let event = await IDO.buy(package100, { from: subscriber });
    truffleAssert.eventEmitted(event, "Buy", (evt) => {
      assert.equal(evt.buyer, subscriber);
      assert.equal(evt.amount, Package100TokenAmount);
      return true;
    });
  }

  async function buySuccessPackage200() {
    // check transfer token ERC20
    let subscriberBalanceBeforeTransfer = new BN(
      await BusdMock.balanceOf(subscriber)
    );
    subscriberBalanceBeforeTransfer.should.be.a.bignumber.that.equals("0");
    await BusdMock.transfer(subscriber, helper.wei(package200Price));
    let subscriberBalanceAfterTransfer = new BN(
      await BusdMock.balanceOf(subscriber)
    );

    subscriberBalanceAfterTransfer.should.be.a.bignumber.that.equals(
      new BN(helper.wei(package200Price))
    );

    //add subscriber
    await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));

    //check allowance before
    let allowanceBefore = await BusdMock.allowance(subscriber, IDO.address);
    allowanceBefore.should.be.a.bignumber.that.equals("0");

    //approve contract
    await BusdMock.approve(
      IDO.address,
      helper.wei(package200Price).toString(),
      { from: subscriber }
    );

    //check allowance after approve
    let allowanceAfter = await BusdMock.allowance(subscriber, IDO.address);
    allowanceAfter.should.be.a.bignumber.that.equals(
      helper.wei(package200Price)
    );

    const balanceIDOBeforeBuy = new BN(await BusdMock.balanceOf(IDO.address));

    //buy
    let event = await IDO.buy(package200, { from: subscriber });
    truffleAssert.eventEmitted(event, "Buy", (evt) => {
      assert.equal(evt.buyer, subscriber);
      assert.equal(evt.amount, Package200TokenAmount);
      return true;
    });

    //check balance IDO and subscriber after buy
    let subscriberBalanceAfterBuy = new BN(
      await BusdMock.balanceOf(subscriber)
    );
    subscriberBalanceAfterBuy.should.be.a.bignumber.that.equals("0");

    const balanceIDOAfterBuy = new BN(await BusdMock.balanceOf(IDO.address));
    balanceIDOAfterBuy.should.be.a.bignumber.that.equals(
      balanceIDOBeforeBuy.add(helper.wei(package200Price))
    );
  }

  describe("once deployed", async () => {
    beforeEach(async function () {
      BusdMock = await SCBusdMock.new();
      IDO = await SCIDO.new(BusdMock.address);
    });

    describe("has correct information", async () => {
      describe("on the packages", async () => {
        it("100 BUSD for 3333.34 token", async () => {
          let actualBUSD = await IDO.Package100BUSDAmount();
          actualBUSD.should.be.a.bignumber.that.equals(helper.wei(100));
          let actualToken = await IDO.Package100TokenAmount();
          actualToken.should.be.a.bignumber.that.equals(helper.wei(3333.34));
        });
        it("200 BUSD for 6666.67 token", async () => {
          let actualBUSD = await IDO.Package200BUSDAmount();
          actualBUSD.should.be.a.bignumber.that.equals(helper.wei(200));
          let actualToken = await IDO.Package200TokenAmount();
          actualToken.should.be.a.bignumber.that.equals(helper.wei(6666.67))
        });
      });

      it("on the total amount of token available for IDO (3.000.000 tokens)", async () => {
        let actual = await IDO.amountOfTokensRemaining();
        actual.should.be.a.bignumber.that.equals(helper.wei(3000000))
      });
    });

    describe("managing subscribers/whitelist", async () => {
      let subscribers = [accounts[1], accounts[2]];
      describe("can be done by ADMIN", async () => {
        it("for single subscriber", async () => {
          await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));
        });

        it("for multiple subscribers", async () => {
          await truffleAssert.passes(IDO.addSubscribers(subscribers, {
            from: owner,
          }));
        });
      });

      describe("cannot be done by NON-ADMIN", async () => {
        it("for single subscriber", async () => {
          await truffleAssert.fails(IDO.addSubscriber(owner, { from: subscriber }));
          try {
            await IDO.addSubscriber(owner, { from: subscriber });
          } catch (error) {
            assert.equal(
              error.reason,
              "AccessControl: account " +
              subscriber.toLowerCase() +
              " is missing role " +
              admin_role
            );
          }
        });

        it("for multiple subscribers", async () => {
          await truffleAssert.fails(IDO.addSubscribers(subscribers, { from: subscriber }));
          try {
            await IDO.addSubscribers(subscribers, { from: subscriber });
          } catch (error) {
            assert.equal(
              error.reason,
              "AccessControl: account " +
              subscriber.toLowerCase() +
              " is missing role " +
              admin_role
            );
          }
        });
      });
    });

    describe("managing public sale", async () => {
      it("can be done by ADMIN", async () => {
        await truffleAssert.passes(IDO.publicSale({ from: owner }));
      });

      it("cannot be done by NON-ADMIN", async () => {
        await truffleAssert.fails(IDO.publicSale({ from: subscriber }));
        try {
          await IDO.publicSale({ from: subscriber });
        } catch (error) {
          assert.equal(
            error.reason,
            "AccessControl: account " +
            subscriber.toLowerCase() +
            " is missing role " +
            admin_role
          );
        }
      });
    });

    describe("managing stopping sale", async () => {
      it("can be done by ADMIN", async () => {
        await truffleAssert.passes(IDO.stop({ from: owner }));
      });

      it("cannot be done by NON-ADMIN", async () => {
        await truffleAssert.fails(IDO.stop({ from: subscriber }));
        try {
          await IDO.stop({ from: subscriber });
        } catch (error) {
          assert.equal(
            error.reason,
            "AccessControl: account " +
            subscriber.toLowerCase() +
            " is missing role " +
            admin_role
          );
        }
      });
    });

    describe("when public sale", async () => {
      beforeEach(async () => {
        await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));

        await truffleAssert.passes(BusdMock.transfer(subscriber, helper.wei(package100Price)));
        await truffleAssert.passes(BusdMock.transfer(non_subscriber, helper.wei(package100Price)));
      });

      describe("is not yet openned", async () => {
        it("whitelist-ers can buy IDO", async () => {
          await truffleAssert.passes(BusdMock.approve(
            IDO.address,
            helper.wei(package100Price).toString(),
            { from: subscriber }
          ));

          await truffleAssert.passes(IDO.buy(
            package100,
            { from: subscriber }
          ));
        });

        it("non-whitelist-ers cannot buy IDO", async () => {
          await truffleAssert.fails(IDO.buy(package100, { from: non_subscriber }));
          try {
            await IDO.buy(package100, { from: non_subscriber });
          } catch (error) {
            assert.equal(
              error.reason,
              "IDO: either caller is not in the whitelist or public sale is not ready"
            );
          }
        })
      });

      describe("is openned", async () => {
        beforeEach(async () => {
          await truffleAssert.passes(IDO.publicSale({ from: owner }));
        });

        it("whitelist-ers can buy IDO", async () => {
          await truffleAssert.passes(BusdMock.approve(
            IDO.address,
            helper.wei(package100Price).toString(),
            { from: subscriber }
          ));

          await truffleAssert.passes(IDO.buy(
            package100,
            { from: subscriber }
          ));

        });

        it("non-whitelist-ers can buy IDO", async () => {
          await truffleAssert.passes(BusdMock.approve(
            IDO.address,
            helper.wei(package100Price).toString(),
            { from: non_subscriber }
          ));

          await truffleAssert.passes(IDO.buy(
            package100,
            { from: non_subscriber }
          ));
        })
      });
    });

    describe("when stopping sale", async () => {
      beforeEach(async () => {
        await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));
        await truffleAssert.passes(IDO.publicSale({ from: owner }));
        await truffleAssert.passes(IDO.stop({ from: owner }));
      });

      it("whitelist-ers cannot buy IDO", async () => {
        await truffleAssert.fails(IDO.buy(package100, { from: subscriber }));
        try {
          await IDO.buy(package100, { from: subscriber });
        } catch (error) {
          assert.equal(
            error.reason,
            "Pausable: paused"
          );
        }
      });

      it("non-whitelist-ers cannot buy IDO", async () => {
        await truffleAssert.fails(IDO.buy(package100, { from: non_subscriber }));
        try {
          await IDO.buy(package100, { from: non_subscriber });
        } catch (error) {
          assert.equal(
            error.reason,
            "Pausable: paused"
          );
        }
      })
    });

    describe("when on sale", async () => {
      let buyerBalanceBeforeBuy = package100Price + package200Price;
      beforeEach(async () => {
        await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));
        await truffleAssert.passes(BusdMock.transfer(subscriber, helper.wei(buyerBalanceBeforeBuy)));
      });

      it("can buy pack 100 BUSD", async () => {
        await buyPackage100();
        let subscriberBalanceAfterBuy = new BN(
          await BusdMock.balanceOf(subscriber)
        );
        subscriberBalanceAfterBuy.should.be.a.bignumber.that.equals(helper.wei(buyerBalanceBeforeBuy - package100Price));

        const balanceIDOAfterBuy = new BN(await BusdMock.balanceOf(IDO.address));
        balanceIDOAfterBuy.should.be.a.bignumber.that.equals(
          (helper.wei(package100Price))
        );
      });

      it("can buy pack 200 BUSD", async () => {
        await truffleAssert.passes(BusdMock.approve(
          IDO.address,
          helper.wei(package200Price).toString(),
          { from: subscriber }
        ));

        let event = await IDO.buy(package200, { from: subscriber });
        truffleAssert.eventEmitted(event, "Buy", (evt) => {
          assert.equal(evt.buyer, subscriber);
          assert.equal(evt.amount, Package200TokenAmount);
          return true;
        });

        let subscriberBalanceAfterBuy = new BN(
          await BusdMock.balanceOf(subscriber)
        );
        subscriberBalanceAfterBuy.should.be.a.bignumber.that.equals(helper.wei(buyerBalanceBeforeBuy - package200Price));

        const balanceIDOAfterBuy = new BN(await BusdMock.balanceOf(IDO.address));
        balanceIDOAfterBuy.should.be.a.bignumber.that.equals(
          (helper.wei(package200Price))
        );
      });

      it("cannot sale more than amount of available tokens for IDO", async () => {
        // this test uses a mock contract which is identical to the real contract, 
        // except that the amount of available tokens is much smaller.
        let IDOMock = await SCIDOMock.new(BusdMock.address);
        await truffleAssert.passes(IDOMock.publicSale({ from: owner }));

        let tokenAmount = await IDOMock.amountOfTokensRemaining();
        let package200TokenAmount = await IDOMock.Package200TokenAmount();

        let canBuyTimes = tokenAmount.div(package200TokenAmount);

        // shoule be able to buy 499 times pack 200
        for (i = 1; i <= canBuyTimes; i++) {
          await truffleAssert.passes(BusdMock.transfer(accounts[i], helper.wei(package200Price)));

          await truffleAssert.passes(BusdMock.approve(IDOMock.address, helper.wei(package200Price).toString(), { from: accounts[i] }));
          await truffleAssert.passes(IDOMock.buy(package200, { from: accounts[i] }));
        }

        // and should fail next
        await truffleAssert.passes(BusdMock.transfer(accounts[500], helper.wei(package200Price)));
        await truffleAssert.passes(BusdMock.approve(IDOMock.address,
          helper.wei(package200Price).toString(), { from: accounts[500] }));

        await truffleAssert.fails(IDOMock.buy(package200, { from: accounts[500] }));
        try {
          await IDOMock.buy(package200, { from: accounts[500] });
        } catch (error) {
          assert.equal(
            error.reason,
            "IDO: amount of tokens available for IDO is running out"
          );
        }

      });

      it("allows each address can buy at most 1 time", async () => {
        await buyPackage100();

        await truffleAssert.fails(IDO.buy(package100, { from: subscriber }));
        try {
          await IDO.buy(package100, { from: subscriber });
        } catch (error) {
          assert.equal(error.reason, "IDO: caller has reached her quota");
        }
      });
    });

    describe("when collecting BUSD", async () => {
      beforeEach(async () => {
        await truffleAssert.passes(IDO.addSubscriber(subscriber, { from: owner }));
        await truffleAssert.passes(BusdMock.transfer(subscriber, helper.wei(package100Price)));

      });
      it("is only doable by ADMIN", async () => {
        await buyPackage100();
        const ownerBalanceBefore = new BN(await BusdMock.balanceOf(owner));

        const balanceERC20IDO = new BN(await BusdMock.balanceOf(IDO.address));

        balanceERC20IDO.should.be.a.bignumber.that.equal(
          helper.wei(package100Price)
        );

        await truffleAssert.passes(IDO.collectBUSD({ from: owner }));

        const ownerBalanceAfter = await BusdMock.balanceOf(owner);

        ownerBalanceAfter.should.be.a.bignumber.that.equal(
          ownerBalanceBefore.add(helper.wei(package100Price))
        );
      });

      it("is not doable by NON-ADMIN ", async () => {
        await buyPackage100();

        await truffleAssert.fails(IDO.collectBUSD({ from: subscriber }));
        try {
          await IDO.collectBUSD({ from: subscriber });
        } catch (error) {
          assert.equal(
            error.reason,
            "AccessControl: account " +
            subscriber.toLowerCase() +
            " is missing role " +
            admin_role
          );
        }
      });
    });
  });
});
