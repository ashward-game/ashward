
const Token = artifacts.require("Token");
const VestingReserve = artifacts.require("VestingReserve");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingReserve, conToken.address);

    const conVestingReserve = await VestingReserve.deployed();

    await conToken.transfer(conVestingReserve.address, vesting.Reserve);
    await conToken.addNoTaxAddress(conVestingReserve.address);

    helper.dumpContractAddress("VestingReserve", conVestingReserve.address);
  });
};
