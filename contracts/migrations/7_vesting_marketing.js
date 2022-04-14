
const Token = artifacts.require("Token");
const VestingMarketing = artifacts.require("VestingMarketing");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingMarketing, conToken.address);

    const conVestingMarketing = await VestingMarketing.deployed();

    await conToken.transfer(conVestingMarketing.address, vesting.Marketing);
    await conToken.addNoTaxAddress(conVestingMarketing.address);

    helper.dumpContractAddress("VestingMarketing", conVestingMarketing.address);
  });
};
