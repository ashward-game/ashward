const { BN } = require("@openzeppelin/test-helpers/src/setup");

const util = web3.utils;
const eth = web3.eth;

async function getLastBlockTimestamp() {
  return (await eth.getBlock("latest")).timestamp;
}

function wei(amount) {
  return new util.BN(util.toWei(String(amount), "ether"));
}

function uintToBytes32(number) {
  const numberBN = new BN(number);
  return numberBN.toArray("be", 32);
}

function generateKey() {
  let key = eth.accounts.create();
  return key;
}

function sign(privateKey, random) {
  return eth.accounts.sign(random, privateKey);
}

function random() {
  return util.randomHex(32);
}

function hash(random) {
  return eth.accounts.hashMessage(random);
}

function hexToNumber(random) {
  return util.toBN(random);
}
module.exports = {
  getLastBlockTimestamp,
  wei,
  uintToBytes32,
  generateKey,
  sign,
  random,
  hash,
  hexToNumber,
};
