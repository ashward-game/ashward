import {
  ABI_VestingPrivate,
  ADDRESS_VestingPrivate,
} from "~/plugins/contracts/VestingPrivate";
import { ethers } from "ethers";
export default {
  getContract(w3provider) {
    let signer = w3provider.getSigner();
    return new ethers.Contract(
      ADDRESS_VestingPrivate,
      ABI_VestingPrivate,
      signer
    );
  },

  async callClaimTGE(w3provider) {
    const contract = this.getContract(w3provider);

    return await contract.claimTGE();
  },

  async callClaim(w3provider) {
    const contract = this.getContract(w3provider);

    return await contract.claim();
  },

  async callVestingOf(w3provider, address) {
    const contract = this.getContract(w3provider);

    return await contract.vestingOf(address);
  },
  async getAmountRemain(w3provider) {
    const contract = this.getContract(w3provider);
    return await contract.amountOfTokensRemaining();
  },

  async callHasClaimedTGE(w3provider, address) {
    const contract = this.getContract(w3provider, address);
    return await contract.hasClaimedTGE(address);
  },
};
