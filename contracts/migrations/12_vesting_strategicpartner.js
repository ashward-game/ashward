
const Token = artifacts.require("Token");
const VestingStrategicPartner = artifacts.require("VestingStrategicPartner");
const vesting = require("../vesting")
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    await deployer.deploy(VestingStrategicPartner, conToken.address);

    const conVestingStrategicPartner = await VestingStrategicPartner.deployed();

    await conToken.transfer(conVestingStrategicPartner.address, vesting.StrategicPartner);
    await conToken.addNoTaxAddress(conVestingStrategicPartner.address);

    helper.dumpContractAddress("VestingStrategicPartner", conVestingStrategicPartner.address);
  });
};
