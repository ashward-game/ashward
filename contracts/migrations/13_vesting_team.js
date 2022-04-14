
const Token = artifacts.require("Token");
const VestingTeam = artifacts.require("VestingTeam");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingTeam, conToken.address);

    const conVestingTeam = await VestingTeam.deployed();

    await conToken.transfer(conVestingTeam.address, vesting.Team);
    await conToken.addNoTaxAddress(conVestingTeam.address);

    helper.dumpContractAddress("VestingTeam", conVestingTeam.address);
  });
};
