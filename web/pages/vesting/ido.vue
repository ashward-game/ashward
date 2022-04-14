<template>
  <div class="wrap__page litepaper">
    <div class="litepaper__nav mb-4">
      <ul class="nav nav-custom justify-content-center">
        <li class="nav-item">
          <n-link class="nav-link" to="/launchpad">Info</n-link>
        </li>
        <li class="nav-item">
          <n-link class="nav-link active" to="/vesting/ido">Vesting</n-link>
        </li>
      </ul>
    </div>

    <div class="container container-790">
      <div class="el__box mb-4">
        <p class="font18 mb-2">Investor’s zone</p>
        <!--        <div class="el__box__time text-center mb-3">-->
        <!--          <p class="mb-2">Sale starts in</p>-->
        <!--          <p class="-time">-->
        <!--            &lt;!&ndash; time = 2 day &ndash;&gt;-->
        <!--            <Countdown />-->
        <!--          </p>-->
        <!--        </div>-->
        <p class="mb-3" v-if="inVestingPool">
          You claimed: {{ amountClaimed }} ASC
        </p>

        <div class="progress__wrap mb-4" v-if="inVestingPool">
          <div class="progress">
            <div
              class="progress-bar"
              role="progressbar"
              :style="{
                width: progressStatus + '%',
              }"
              aria-valuenow="50"
              aria-valuemin="0"
              aria-valuemax="100"
            />
          </div>
          <div class="d-flex justify-content-between mt-3">
            <p>{{ parseFloat(progressStatus).toFixed(2) }}%</p>
            <p>{{ amountClaimed }} / {{ dataVestingAccount[0] }} ASC</p>
          </div>
        </div>

        <div class="el__wrap maxwidth-380">
          <ul class="el__list -investor -light mb-3">
            <li>
              <div class="li__label">Wallet:</div>
              <div class="li__val ms-auto">
                {{ account }}
              </div>
            </li>
            <li>
              <div class="li__label">Status:</div>
              <div class="li__val ms-auto -ongoing" v-if="inVestingPool">
                Ongoing
              </div>
              <div class="li__val ms-auto" v-else>Not in vesting pool</div>
            </li>
            <li>
              <div class="li__label">Sale</div>
              <div class="li__val ms-auto">
                18-Mar-2022<br />13:00 - 15:00 (UTC)
              </div>
            </li>
            <li>
              <div class="li__label">Vesting starts:</div>
              <div class="li__val ms-auto">
                18-Mar-2022<br />16:30 (UTC)<br />30 mins after public sale
              </div>
            </li>
            <li>
              <div class="li__label">Vesting period:</div>
              <div class="li__val ms-auto">
                TGE 20%, 1 month cliff,<br />
                vest monthly in 5 months
              </div>
            </li>
          </ul>
        </div>
      </div>

      <div class="el__box" v-if="inVestingPool">
        <table class="table -table-6col align-middle">
          <thead>
            <tr>
              <th scope="col" />
              <th scope="col">Round</th>
              <th scope="col">Time</th>
              <th scope="col">Token unlocked</th>
              <th scope="col">Token amount</th>
              <th scope="col" class="td__stt" />
            </tr>
          </thead>
          <tbody>
            <tr
              v-if="isLoadingDataBC"
              class="js-tr-toggle"
              v-for="(item1, key1) in claimableMilestones"
              :key="key1"
            >
              <td scope="row">1</td>
              <td>IDO</td>
              <td>
                {{ formatTimestampToString(item1)[0] }}<br />{{
                  formatTimestampToString(item1)[1]
                }}
                (UTC)
              </td>
              <td class="td-toggle d-md-none"><i class="ic-arrow-down" /></td>
              <td>Loading</td>
              <td>Loading</td>
              <td class="td__stt">Loading</td>
            </tr>
            <tr
              class="js-tr-toggle"
              v-for="(item, key) in claimableMilestones"
              :key="key"
              v-if="!isLoadingDataBC"
            >
              <td scope="row">{{ key + 1 }}</td>
              <td>IDO</td>
              <td>
                {{ formatTimestampToString(item)[0] }}<br />{{
                  formatTimestampToString(item)[1]
                }}
                (UTC)
              </td>
              <td class="td-toggle d-md-none"><i class="ic-arrow-down" /></td>
              <td>
                {{
                  parseFloat(
                    getPercentMilestone(claimablePercents[key]) * 100
                  ).toFixed(2)
                }}
                %
              </td>
              <td>
                {{
                  parseFloat(
                    dataVestingAccount[0] *
                      getPercentMilestone(claimablePercents[key])
                  ).toFixed(2)
                }}
              </td>
              <td class="td__stt -claimed">
                <div v-if="key === 0" class="flex justify-center">
                  <div v-if="isClaimedTGE">Claimed</div>
                  <div v-else>
                    <button
                      class="btnz mb-2 me-2 js-attr-button"
                      style="min-width: 100px !important"
                      @click="claim('tge')"
                    >
                      <vue-element-loading
                        :active="isLoadingCall"
                        spinner="bar-fade-scale"
                        color="#FF6700"
                      />
                      Claim
                    </button>
                  </div>
                </div>
                <div v-else class="flex justify-center">
                  <div :class="{ td__stt: statusClaimData[key] === 'Pending' }">
                    {{ statusClaimData[key] }}
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <p class="text-center table-note" v-if="isClaimedTGE">
          <button
            @click="claim('normal')"
            class="btnz w-4 mb-2 me-2"
            style="min-width: 100px !important"
            :disabled="!isClaimable"
          >
            <vue-element-loading
              :active="isLoadingCall"
              spinner="bar-fade-scale"
              color="#FF6700"
            />
            Claim
          </button>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import $ from "jquery";
import { ethers } from "ethers";
import VestingIDO from "~/plugins/contracts/services/VestingIDO";
import moment from "moment";
const BN = require("bn.js");

export default {
  data() {
    return {
      claimableMilestones: [
        1647621000, 1652891400, 1655569800, 1658161800, 1660840200,
      ],
      claimablePercents: [2000, 2000, 2000, 2000, 2000],
      denominator: 10000,
      statusClaimData: [],
      dataVestingAccount: [],
      isLoadingCall: false,
      isClaimedTGE: undefined,
      isLoadingDataBC: true,
      inVestingPool: false,
    };
  },
  async mounted() {
    if (this.$device.isMobile) {
      $(".js-tr-toggle").click(function (e) {
        $(this).next().toggle();
      });
    }

    const web3ModalConnected = localStorage.getItem("web3ModalConnected");

    if (web3ModalConnected) {
      await this.$ether();
    }

    if (this.account) {
      await this.initData();
    }
  },
  computed: {
    provider() {
      if (this.$store.state.ether?.isConnectedWeb3) {
        return this.$web3Provider();
      } else {
        return new ethers.providers.JsonRpcProvider(process.env.RPC_PROVIDER);
      }
    },
    account() {
      if (
        this.$store.state.ether?.web3Account &&
        this.$store.state.ether?.isConnectedWeb3
      )
        return this.$store.state.ether?.web3Account;
    },
    amountClaim() {
      if (this.dataVestingAccount[0])
        return parseFloat(
          new Intl.NumberFormat().format(this.dataVestingAccount[0])
        ).toFixed(2);
    },
    isClaimable() {
      if (this.account) {
        return this.statusClaimData.includes("Claimable");
      }
    },
    amountClaimed() {
      if (
        this.account &&
        this.dataVestingAccount.length > 0 &&
        this.dataVestingAccount[0] &&
        this.statusClaimData.length > 0
      ) {
        let amount = 0;
        for (let i = 0; i < this.statusClaimData.length; i++) {
          if (this.statusClaimData[i] === "Claimed")
            amount +=
              parseFloat(this.dataVestingAccount[0]) *
              this.getPercentMilestone(this.claimablePercents[i]);
        }
        return parseFloat(amount).toFixed(2);
      }
    },
    progressStatus() {
      if (this.amountClaimed && this.dataVestingAccount[0]) {
        return (this.amountClaimed / this.dataVestingAccount[0]) * 100;
      }
      return 0;
    },
  },
  methods: {
    async initData() {
      try {
        const dataVestingAccountBN = await VestingIDO.callVestingOf(
          this.provider,
          this.account
        );
        this.dataVestingAccount = dataVestingAccountBN.map((item, key) => {
          if (key === 0) {
            return parseFloat(ethers.utils.formatEther(item)).toFixed(2);
          } else {
            return ethers.BigNumber.from(item).toNumber();
          }
        });
        this.inVestingPool = true;
        this.isClaimedTGE = await VestingIDO.callHasClaimedTGE(
          this.provider,
          this.account
        );

        this.setStatusClaim();
      } catch (error) {
        if (error.data) {
          let reason = "";
          if (error.data.message) reason = error.data.message;
          else reason = error.data.data.stack;
          if (reason.includes("Vesting: beneficiary is not in pool")) {
            this.inVestingPool = false;
          }
        }
      }
      this.isLoadingDataBC = false;
    },
    async claim(type) {
      if (this.account) {
        try {
          this.isLoadingCall = true;
          let receipt;

          if (type === "tge") {
            receipt = await VestingIDO.callClaimTGE(this.provider);
          } else receipt = await VestingIDO.callClaim(this.provider);

          const tx = await receipt.wait();
          if (tx) {
            if (type === "tge") this.isClaimedTGE = true;
            else await this.initData();

            this.$toast.add({
              severity: "success",
              summary: "Claim successful",
              life: 3000,
            });
          }
        } catch (error) {
          console.log(error);
          if (error.data) {
            let reason = "";
            let message = "";
            if (error.data.message) reason = error.data.message;
            else reason = error.data.data.stack;
            if (
              reason.includes(
                "Vesting: cannot unlock tokens for this milestone or already claimed tokens for current milestone"
              )
            ) {
              message = "Cannot unlock tokens for this milestone";
            }

            if (reason.includes("Vesting: already claimed all TGE tokens")) {
              message = "You claimed all TGE tokens";
            }

            if (reason.includes("Vesting: no TGE tokens")) {
              message = "No TGE tokens";
            }

            if (
              reason.includes(
                "Vesting: need to wait for 30 minutes before unlocking TGE tokens"
              )
            ) {
              message =
                "Need to wait for 30 minutes before unlocking TGE tokens";
            }

            if (message !== "") {
              this.$toast.add({
                severity: "error",
                summary: message,
                life: 3000,
              });
            }
          }
        }
        this.isLoadingCall = false;
      }
    },
    formatTimestampToString(timestamp) {
      return moment
        .unix(timestamp)
        .utc()
        .format("DD-MMM-YYYY HH:mm:ss")
        .split(" ");
    },

    getDataClaimableTime() {
      this.statusClaimData = [];
      for (let i = 0; i < this.claimableMilestones.length; i++) {
        const time = moment.unix(this.claimableMilestones[i]);

        let now = moment(moment.utc());
        let distance = moment.duration(time.diff(now));
        if (distance < 0) {
          if (!this.isClaimedTGE && i === 0) {
            this.statusClaimData[i] = "Claim";
          } else if (this.dataVestingAccount[1] >= i) {
            this.statusClaimData[i] = "Claimed";
          } else if (this.dataVestingAccount[1] < i) {
            this.statusClaimData[i] = "Claimable";
          }
        } else {
          this.statusClaimData[i] = "Pending";
        }
      }
    },
    setStatusClaim() {
      let self = this;
      let x = setInterval(function () {
        self.getDataClaimableTime();
      }, 1000);
    },
    getPercentMilestone(percentMilestone) {
      return percentMilestone / this.denominator;
    },
  },
};
</script>
