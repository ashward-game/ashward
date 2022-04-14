import { ethers } from "ethers";
import StakingRewards from "~/plugins/contracts/services/StakingRewards";
import moment from "moment";

export default {
  data() {
    return {
      pool: {
        totalStaked: 0,
        totalStakers: 0,
        myStakes: 0,
        amount: 0,
        totalRewarded: 0,
        earned: 0,
        isLoading: false,
        typePool: undefined,
        startDate: "",
        endDate: "",
      },
    };
  },
  mounted() {},
  computed: {
    provider() {
      if (this.$store.state.ether?.isConnectedWeb3) {
        return this.$web3Provider();
      } else {
        return new ethers.providers.JsonRpcProvider(process.env.RPC_PROVIDER);
      }
    },
    account() {
      if (this.$store.state.ether?.web3Account)
        return this.$store.state.ether?.web3Account;
    },
    checkEndTime() {
      if (this.pool.endDate) {
        const now = moment(moment.utc());

        let distance = moment.duration(
          moment.unix(this.pool.endDate).diff(now)
        );
        if (distance < 0) return false;
      }
      return true;
    },
  },
  watch: {
    account(oldValue, newValue) {
      if (this.pool.typePool && oldValue) {
        this.getMyStakes(this.pool.typePool);
        this.getEarned(this.pool.typePool);
      }
    },
  },
  methods: {
    initPool(pool) {
      this.pool.typePool = pool;
      var wsProvider = new ethers.providers.WebSocketProvider(
        process.env.RPC_URL
      );
        this.getDataContract(pool);
        if (this.account) {
          this.getMyStakes(pool);
          this.getEarned(pool);
        }
      wsProvider.on("block", async (blockNumber) => {
        console.log("New Block: " + blockNumber);
          this.getDataContract(pool);
          if (this.account) {
            this.getMyStakes(pool);
            this.getEarned(pool);
          }
      });
    },
    setPoolSelected() {
      let self = this;
      this.$store.dispatch("staking/setPoolSelected", {
        ...self.pool,
      });
    },
    async getEarned(pool) {
      try {
        if (this.account) {
          this.pool.earned = await StakingRewards.earned(
            this.provider,
            this.account,
            pool
          );
        }
      } catch (e) {
        console.error(this.catchError(e));
      }
    },
    async getMyStakes(pool) {
      try {
        if (this.account) {
          this.pool.myStakes = await StakingRewards.myStakes(
            this.provider,
            this.account,
            pool
          );
        }
      } catch (e) {
        console.error(this.catchError(e));
      }
    },
    async getDataContract(pool) {
      try {
        this.pool.totalStakers = await StakingRewards.totalStakers(
          this.provider,
          pool
        );

        this.pool.totalStaked = await StakingRewards.totalStaking(
          this.provider,
          pool
        );

        this.pool.startDate = await StakingRewards.getStartDate(
          this.provider,
          pool
        );
        this.pool.endDate = await StakingRewards.getEndDate(
          this.provider,
          pool
        );
      } catch (e) {
        console.error(this.catchError(e));
      }
    },
    async withdraw(pool) {
      try {
        if (!this.account) {
          this.$toast.add({
            severity: "warn",
            summary: "You must connect your wallet",
            life: 3000,
          });
          return;
        }
        this.isLoading = true;
        const receipt = await StakingRewards.withdraw(this.provider, pool);
        const tx = await receipt.wait();
        if (tx) {
          this.isLoading = false;
          this.$toast.add({
            severity: "success",
            summary: "Exit successful",
            life: 3000,
          });
        }
      } catch (e) {
        console.error(this.catchError(e));
        this.$toast.add({
          severity: "error",
          summary: "Create transaction error",
          life: 3000,
        });
      }
      this.isLoading = false;
    },
    async getRewards(pool) {
      try {
        if (!this.account) {
          this.$toast.add({
            severity: "warn",
            summary: "You must connect your wallet",
            life: 3000,
          });
          return;
        }
        this.isLoading = true;
        const receipt = await StakingRewards.getRewards(this.provider, pool);
        const tx = await receipt.wait();
        if (tx) {
          this.isLoading = false;
          this.$toast.add({
            severity: "success",
            summary: "Get reward successful",
            life: 3000,
          });
        }
      } catch (e) {
        console.error(this.catchError(e));
        this.$toast.add({
          severity: "error",
          summary: "Create transaction error",
          life: 3000,
        });
      }
      this.isLoading = false;
    },
  },
};
