/**
 * This module contains all constants related to Vesting.
 */
const web3 = require('web3');
const util = web3.utils;
const BN = web3.utils.BN;

const StrategicPartner = new BN(util.toWei("20000000", "ether"))
const Private = new BN(util.toWei("100000000", "ether"))
const IDO = new BN(util.toWei("3000000", "ether"))
const Advisory = new BN(util.toWei("10000000", "ether"))
const Team = new BN(util.toWei("120000000", "ether"))
const Marketing = new BN(util.toWei("120000000", "ether"))
const Play2Earn = new BN(util.toWei("300000000", "ether"))
const Staking = new BN(util.toWei("150000000", "ether"))
const Liquidity = new BN(util.toWei("85000000", "ether"))
const Reserve = new BN(util.toWei("80000000", "ether"))

const liquidityAddr = process.env.LIQUIDITY_ADDRESS;


module.exports = Object.freeze({
    StrategicPartner,
    Private,
    IDO,
    Advisory,
    Team,
    Marketing,
    Play2Earn,
    Staking,
    Liquidity,
    Reserve,
    liquidityAddr
});

