import { ABI_NFT, ADDRESS_NFT } from "~/plugins/contracts/NFT";
import { ethers } from "ethers";

export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(ADDRESS_NFT, ABI_NFT, signer);
  },
  async getUriString(w3provider, tokenId) {
    const nftContract = this.getContract(w3provider);
    return await nftContract.tokenURI(parseInt(tokenId));
  },
};
