
const Token = artifacts.require("Token");
const VestingAdvisory = artifacts.require("VestingAdvisory");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingAdvisory, conToken.address);

    const conVestingAdvisory = await VestingAdvisory.deployed();

    await conToken.transfer(conVestingAdvisory.address, vesting.Advisory);
    await conToken.addNoTaxAddress(conVestingAdvisory.address);

    helper.dumpContractAddress("VestingAdvisory", conVestingAdvisory.address);
  });
};
