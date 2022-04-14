
const Token = artifacts.require("Token");
const VestingPrivate = artifacts.require("VestingPrivate");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingPrivate, conToken.address);

    const conVestingPrivate = await VestingPrivate.deployed();

    await conToken.transfer(conVestingPrivate.address, vesting.Private);
    await conToken.addNoTaxAddress(conVestingPrivate.address);

    helper.dumpContractAddress("VestingPrivate", conVestingPrivate.address);
  });
};
