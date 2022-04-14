import { ABI_BUSD, ADDRESS_BUSD } from "~/plugins/contracts/BUSD";
import { ADDRESS_IDO } from "~/plugins/contracts/IDO";
import { ethers } from "ethers";
export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(ADDRESS_BUSD, ABI_BUSD, signer);
  },

  async approveBUSD(w3provider, amountInWei) {
    const contract = this.getContract(w3provider);

    return await contract.approve(ADDRESS_IDO, amountInWei);
  },

  async getAllowance(w3provider, owner, spender) {
    const contract = this.getContract(w3provider);

    return await contract.allowance(owner, spender);
  },
};
