import { ABI_TOKEN, ADDRESS_TOKEN } from "~/plugins/contracts/Token";
import { ADDRESS_STAKINGREWARDS } from "~/plugins/contracts/StakingRewards";
import {
  ABI_STAKINGPOOL1,
  ADDRESS_STAKINGPOOL1,
} from "~/plugins/contracts/StakingPool1";

import { ethers } from "ethers";
import {
  ABI_STAKINGPOOL2,
  ADDRESS_STAKINGPOOL2,
} from "~/plugins/contracts/StakingPool2";
import {
  ABI_STAKINGPOOL3,
  ADDRESS_STAKINGPOOL3,
} from "~/plugins/contracts/StakingPool3";
export default {
  async balanceOf(web3provider, account) {
    console.log("okok");
    let contract = new ethers.Contract(ADDRESS_TOKEN, ABI_TOKEN, web3provider);
    let balanceWei = await contract.balanceOf(account);
    return ethers.utils.formatUnits(balanceWei, 18).toString();
  },

  async stake(web3provider, amount, pool) {
    let signer = web3provider.getSigner();
    let contract = new ethers.Contract(ADDRESS_TOKEN, ABI_TOKEN, signer);
    let signed = await contract.connect(signer);
    // call overloading function in ethers
    // see https://stackoverflow.com/q/68289806
    // https://github.com/ethers-io/ethers.js/issues/407
    let addressPool = null;
    if (pool === "pool1") {
      addressPool = ADDRESS_STAKINGPOOL1;
    } else if (pool === "pool2") {
      addressPool = ADDRESS_STAKINGPOOL2;
    } else if (pool === "pool3") {
      addressPool = ADDRESS_STAKINGPOOL3;
    }
    console.log(pool);
    return signed["approveAndCall(address,uint256)"](addressPool, amount);
  },
};
