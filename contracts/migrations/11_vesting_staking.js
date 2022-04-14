
const Token = artifacts.require("Token");
const VestingStaking = artifacts.require("VestingStaking");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingStaking, conToken.address);

    const conVestingStaking = await VestingStaking.deployed();

    await conToken.transfer(conVestingStaking.address, vesting.Staking);
    await conToken.addNoTaxAddress(conVestingStaking.address);

    helper.dumpContractAddress("VestingStaking", conVestingStaking.address);
  });
};
