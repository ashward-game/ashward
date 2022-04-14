/**
 * This module contains all constants related to the openbox genesis.
 */
 const web3 = require('web3');

 // Note: decimals are not supported in BN library.
 // https://github.com/indutny/bn.js/#usage
 const rareBoxPriceinWei = web3.utils.toWei(process.env.RARE_BOX_PRICE);
 const legendBoxPriceinWei = web3.utils.toWei(process.env.LEGEND_BOX_PRICE);
 const mythBoxPriceinWei = web3.utils.toWei(process.env.MYTH_BOX_PRICE);
 
 
 module.exports = Object.freeze({
     numRareBoxes: process.env.NUM_RARE_BOXES,
     rareBoxPrice: rareBoxPriceinWei,
     numLegendBoxes: process.env.NUM_LEGEND_BOXES,
     legendBoxPrice: legendBoxPriceinWei,
     numMythBoxes: process.env.NUM_MYTH_BOXES,
     mythBoxPrice: mythBoxPriceinWei,
 });