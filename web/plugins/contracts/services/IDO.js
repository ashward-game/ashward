import { ABI_IDO, ADDRESS_IDO } from "~/plugins/contracts/IDO";
import { ethers } from "ethers";
export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(ADDRESS_IDO, ABI_IDO, signer);
  },

  async buy(w3provider, packageType) {
    const contract = this.getContract(w3provider);

    return await contract.buy(packageType);
  },
  async getAmountRemain(w3provider) {
    const contract = this.getContract(w3provider);
    return await contract.amountOfTokensRemaining();
  },

  async getAmountRemain2(w3provider) {
    const contract = new ethers.Contract(ADDRESS_IDO, ABI_IDO, w3provider);
    return await contract.amountOfTokensRemaining();
  },

  async getPaused(w3provider) {
    const contract = this.getContract(w3provider);
    return await contract.paused();
  },
  async getIsPublicSale(w3provider) {
    const contract = this.getContract(w3provider);
    return await contract.isPublicSale();
  },
};
