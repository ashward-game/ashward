const Token = artifacts.require("Token");
const token = require("../token");
const helper = require("../helpers/helper");

module.exports = function (deployer, network, accounts) {
  deployer.then(async function () {

    await deployer.deploy(
      Token,
      token.name,
      token.symbol
    );
    const conToken = await Token.deployed();
    helper.dumpContractAddress("Token", conToken.address);
  });
};
