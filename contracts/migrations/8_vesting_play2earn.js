
const Token = artifacts.require("Token");
const VestingPlay2Earn = artifacts.require("VestingPlay2Earn");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingPlay2Earn, conToken.address);

    const conVestingPlay2Earn = await VestingPlay2Earn.deployed();

    await conToken.transfer(conVestingPlay2Earn.address, vesting.Play2Earn);
    await conToken.addNoTaxAddress(conVestingPlay2Earn.address);

    helper.dumpContractAddress("VestingPlay2Earn", conVestingPlay2Earn.address);
  });
};
