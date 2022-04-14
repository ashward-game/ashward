import {
  ABI_STAKINGPOOL1,
  ADDRESS_STAKINGPOOL1,
} from "~/plugins/contracts/StakingPool1";
import {
  ABI_STAKINGPOOL2,
  ADDRESS_STAKINGPOOL2,
} from "~/plugins/contracts/StakingPool2";
import {
  ABI_STAKINGPOOL3,
  ADDRESS_STAKINGPOOL3,
} from "~/plugins/contracts/StakingPool3";
import { ethers } from "ethers";

export default {
  getContract(w3provider, pool) {
    if (!this.checkPool(pool)) return null;

    let signer = w3provider.getSigner();
    let contractStaking = null;
    if (pool === "pool1") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL1,
        ABI_STAKINGPOOL1,
        signer
      );
    } else if (pool === "pool2") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL2,
        ABI_STAKINGPOOL2,
        signer
      );
    } else if (pool === "pool3") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL3,
        ABI_STAKINGPOOL3,
        signer
      );
    }

    return contractStaking;
  },
  getContractWeb3(w3provider, pool) {
    if (!this.checkPool(pool)) return null;

    let contractStaking = null;
    if (pool === "pool1") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL1,
        ABI_STAKINGPOOL1,
        w3provider
      );
    } else if (pool === "pool2") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL2,
        ABI_STAKINGPOOL2,
        w3provider
      );
    } else if (pool === "pool3") {
      contractStaking = new ethers.Contract(
        ADDRESS_STAKINGPOOL3,
        ABI_STAKINGPOOL3,
        w3provider
      );
    }

    return contractStaking;
  },
  checkPool(pool) {
    return !(pool !== "pool1" && pool !== "pool2" && pool !== "pool3");
  },

  async totalStakers(w3provider, pool) {
    const contractStaking = this.getContractWeb3(w3provider, pool);

    const totalStakers = await contractStaking.totalStakers();
    return totalStakers.toString();
  },

  async totalStaking(w3provider, pool) {
    const contractStaking = this.getContractWeb3(w3provider, pool);

    const totalStakedWei = await contractStaking.totalStaking();
    return ethers.utils.formatUnits(totalStakedWei, 18).toString();
  },

  async myStakes(w3provider, address, pool) {
    let signer = w3provider.getSigner();
    const contractStaking = this.getContract(w3provider, pool);

    let signed = await contractStaking.connect(signer);
    const myStakesBN = await signed.stakeOf(address);
    return ethers.utils.formatUnits(myStakesBN, 18).toString();
  },

  async earned(w3provider, address, pool) {
    let signer = w3provider.getSigner();

    const contractStaking = this.getContract(w3provider, pool);

    let signed = await contractStaking.connect(signer);
    const earnedBN = await signed.earnableOf(address);
    return ethers.utils.formatUnits(earnedBN, 18).toString();
  },

  async withdraw(w3provider, pool) {
    const contractStaking = this.getContract(w3provider, pool);
    return contractStaking.withdraw();
  },

  async getStartDate(w3provider, pool) {
    const contractStaking = this.getContractWeb3(w3provider, pool);
    const startDate = await contractStaking.startDate();
    return startDate.toString();
  },
  async getEndDate(w3provider, pool) {
    const contractStaking = this.getContractWeb3(w3provider, pool);
    const endDate = await contractStaking.endDate();
    return endDate.toString();
  },
};
