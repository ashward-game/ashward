const web3 = require('web3');
const BN = web3.utils.BN;
const Token = artifacts.require("Token");
const VestingLiquidity = artifacts.require("VestingLiquidity");
const vesting = require("../vesting")
const helper = require("../helpers/helper");
const { liquidityAddr } = require('../vesting');

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    if (vesting.liquidityAddr === undefined ||
      vesting.liquidityAddr === "0x00" ||
      vesting.liquidityAddr === "") {
      throw "Liquidity Address is undefined";
    }

    const conToken = await Token.deployed();

    await deployer.deploy(VestingLiquidity, conToken.address);

    const conVestingLiquidity = await VestingLiquidity.deployed();

    // 20 % of liquidity (TGE) is transferred directly to liquidity wallet
    let tgeAmount = vesting.Liquidity.mul(new BN(20)).div(new BN(100));
    await conToken.transfer(vesting.liquidityAddr, tgeAmount);

    let amount = vesting.Liquidity.sub(tgeAmount);
    await conToken.transfer(conVestingLiquidity.address, amount);
    await conToken.addNoTaxAddress(conVestingLiquidity.address);

    // add liquidity address to free-tax list
    await conToken.addNoTaxAddress(liquidityAddr);

    helper.dumpContractAddress("VestingLiquidity", conVestingLiquidity.address);
  });
};
