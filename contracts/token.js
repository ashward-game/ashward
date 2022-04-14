/**
 * This module contains all constants related to ERC20/ERC1363 tokens.
 */
const web3 = require('web3');
const BN = web3.utils.BN;
const decimals = 18;
const TokenTotalSupplyInEther = process.env.TOKEN_TOTAL_SUPPLY;
const TokenTotalSupplyInWei = new BN(web3.utils.toWei(TokenTotalSupplyInEther, "ether"));

module.exports = Object.freeze({
  name: process.env.TOKEN_NAME,
  symbol: process.env.TOKEN_SYMBOL,
  totalSupply: TokenTotalSupplyInWei,
  decimals: decimals
});

