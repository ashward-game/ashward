/**
 * This module contains all constants related to Staking.
 */
const web3 = require('web3');
const util = web3.utils;

let pool1 = {
    minStake: util.toWei("500", "ether"),
    maxStake: util.toWei("1000", "ether"),
    totalTokenInPool: util.toWei("300000", "ether"),
    balance: util.toWei("105000", "ether"),
    percent: 250,// 2.5%/day
    percentSeconds: 28935,// 0.000028935%/seconds
    startDate: (new Date("2022-04-07 16:00")).getTime() / 1000,
    endDate: (new Date("2022-04-21 16:00")).getTime() / 1000,
}

let pool2 = {
    minStake: util.toWei("1000", "ether"),
    maxStake: util.toWei("2000", "ether"),
    totalTokenInPool: util.toWei("200000", "ether"),
    balance: util.toWei("105000", "ether"),
    percent: 250,// 2.5%/day
    percentSeconds: 28935,// 0.000028935%/seconds
    startDate: (new Date("2022-04-07 16:00")).getTime() / 1000,
    endDate: (new Date("2022-04-28 16:00")).getTime() / 1000,
}

module.exports = Object.freeze({
    pool1,
    pool2,
});

