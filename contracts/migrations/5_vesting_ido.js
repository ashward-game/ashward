
const Token = artifacts.require("Token");
const VestingIDO = artifacts.require("VestingIDO");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingIDO, conToken.address);

    const conVestingIDO = await VestingIDO.deployed();

    await conToken.transfer(conVestingIDO.address, vesting.IDO);
    await conToken.addNoTaxAddress(conVestingIDO.address);

    helper.dumpContractAddress("VestingIDO", conVestingIDO.address);
  });
};
