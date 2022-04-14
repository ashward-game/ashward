import { ethers } from "ethers";

export default {
  namespaced: true,
  state: () => ({
    isLoading: true,
    poolSelected: undefined,
    rewards: {
      giftboxes: [],
      giftcodes: [],
      refcode: {},
    },
  }),
  mutations: {
    ["SET_IS_LOADING"](state, payload) {
      state.isLoading = payload;
    },
    ["SET_POOL_SELECTED"](state, payload) {
      state.poolSelected = payload;
    },
    ["SET_REWARDS"](state, payload) {
      state.rewards = payload;
    },
  },
  actions: {
    setPoolSelected({ commit, state, dispatch }, pool) {
      commit("SET_POOL_SELECTED", {
        ...state.poolSelected,
        ...pool,
      });
    },
    async fetchRewards({ commit, state }, address) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get(`/stakingrewards/${address}`)
          .then((response) => {
            commit("SET_REWARDS", response.data);
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
  },
  getters: {},
};
