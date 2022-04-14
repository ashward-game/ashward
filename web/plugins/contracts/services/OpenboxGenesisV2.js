import {
  ABI_OPENBOXGENESISV2,
  ADDRESS_OPENBOXGENESISV2,
} from "~/plugins/contracts/OpenboxGenesisV2";
var crypto = require("crypto");
const BN = require("bn.js");
import { ethers } from "ethers";
export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(
      ADDRESS_OPENBOXGENESISV2,
      ABI_OPENBOXGENESISV2,
      signer
    );
  },
  async getAmountLeftBox(w3provider) {
    const openBoxContract = this.getContract(w3provider);
    const numRareBoxes = await openBoxContract.numRareBoxes();
    const numLegendBoxes = await openBoxContract.numLegendBoxes();
    const numMythBoxes = await openBoxContract.numMythBoxes();
    return { Rare: numRareBoxes, Legend: numLegendBoxes, Myth: numMythBoxes };
  },
  async getPriceBox(w3provider) {
    const openBoxContract = this.getContract(w3provider);
    const rareBoxPrice = await openBoxContract.rareBoxPrice();
    const legendBoxPrice = await openBoxContract.legendBoxPrice();
    const mythBoxPrice = await openBoxContract.mythBoxPrice();
    return { Rare: rareBoxPrice, Legend: legendBoxPrice, Myth: mythBoxPrice };
  },
  async buyBox(w3provider, boxType, hash, signature, priceInWei) {
    var CRandom = "0x" + crypto.randomBytes(32).toString("hex");
    const openBoxContract = this.getContract(w3provider);
    return await openBoxContract.buyBox(boxType, hash, signature, CRandom, {
      value: priceInWei,
    });
  },

  async inWhiteList(w3provider, address) {
    const openBoxContract = this.getContract(w3provider);
    return await openBoxContract.hasRole(
      "0x3f483399a73bbfbc7e47cea702709b2441abfc4e8152100709ca14556e321303",
      address
    );
  },
};
