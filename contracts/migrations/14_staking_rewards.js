
const Token = artifacts.require("Token");
const StakingPool1 = artifacts.require("StakingPool1");
const StakingPool2 = artifacts.require("StakingPool2");
const helper = require("../helpers/helper");
const staking = require("../staking")

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    const conToken = await Token.deployed();

    let pool1 = staking.pool1;
    await deployer.deploy(StakingPool1, conToken.address, pool1.minStake, pool1.maxStake, pool1.percent, pool1.percentSeconds, pool1.totalTokenInPool, pool1.startDate, pool1.endDate);
    const conStakingPool1 = await StakingPool1.deployed();

    let pool2 = staking.pool2;
    await deployer.deploy(StakingPool2, conToken.address, pool2.minStake, pool2.maxStake, pool2.percent, pool2.percentSeconds, pool2.totalTokenInPool, pool2.startDate, pool2.endDate);
    const conStakingPool2 = await StakingPool2.deployed();

    await conToken.addNoTaxAddress(conStakingPool1.address);
    await conToken.addNoTaxAddress(conStakingPool2.address);

    await conToken.transfer(conStakingPool1.address, pool1.balance);
    await conToken.transfer(conStakingPool2.address, pool2.balance);

    helper.dumpContractAddress("StakingPool1", conStakingPool1.address);
    helper.dumpContractAddress("StakingPool2", conStakingPool2.address);
  });
};
