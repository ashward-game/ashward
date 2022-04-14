const IDO = artifacts.require("IDO");
const helper = require("../helpers/helper");

const BUSDAddress = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {
    await deployer.deploy(IDO, BUSDAddress);
    const conIDO = await IDO.deployed();
    helper.dumpContractAddress("IDO", conIDO.address);
    helper.dumpContractAddress("BUSD", BUSDAddress);
  });
};
