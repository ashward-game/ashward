import {
  ABI_MARKETPLACE,
  ADDRESS_MARKETPLACE,
} from "~/plugins/contracts/Marketplace";
const BN = require("bn.js");
import { ethers } from "ethers";
import { ABI_TOKEN, ADDRESS_TOKEN } from "../Token";
import { ABI_NFT, ADDRESS_NFT } from "../NFT";
export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(ADDRESS_MARKETPLACE, ABI_MARKETPLACE, signer);
  },
  async openOffer(w3provider, seller, tokenId, priceInWei) {
    let signer = w3provider.getSigner();
    let contract = new ethers.Contract(ADDRESS_NFT, ABI_NFT, signer);
    let signed = await contract.connect(signer);
    let data = new BN(priceInWei.toString()).toArray("be", 32);
    return await signed["safeTransferFrom(address,address,uint256,bytes)"](
      seller,
      ADDRESS_MARKETPLACE,
      tokenId,
      data
    );
  },
  async cancel(w3provider, tokenId) {
    const marketplaceContract = this.getContract(w3provider);
    return await marketplaceContract.cancelOffer(tokenId);
  },

  async purchase(w3provider, buyer, tokenId, priceInWei) {
    let signer = w3provider.getSigner();
    let contractToken = new ethers.Contract(ADDRESS_TOKEN, ABI_TOKEN, signer);
    const data = new BN(parseInt(tokenId)).toArray("be", 32);
    let signed = await contractToken.connect(signer);
    return signed["approveAndCall(address,uint256,bytes)"](
      ADDRESS_MARKETPLACE,
      priceInWei,
      data,
      { from: buyer }
    );
  },
};
