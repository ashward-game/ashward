import { ethers } from "ethers";

export default {
  namespaced: true,
  state: () => ({
    isLoading: true,
    nfts: [],
    total: 0,
    totalHistory: 0,
    queryObject: {
      address: null,
      type: null,
      rarity: null,
      star: null,
      level: null,
      order_by_price: null,
      page: 1,
      limit: 20,
      search: null,
    },
    queryHistory: {
      page: 1,
      limit: 5,
      address: null,
    },
    nft: null,
    histories: [],
    info: null,
  }),
  mutations: {
    ["SET_LOADING"](state, payload) {
      state.loading = payload;
    },
    ["SET_QUERY_OBJECT"](state, payload) {
      state.queryObject = payload;
    },
    ["SET_NFTS"](state, payload) {
      state.nfts = payload;
    },
    ["SET_TOTAL"](state, payload) {
      state.total = payload;
    },
    ["SET_IS_LOADING"](state, payload) {
      state.isLoading = payload;
    },
    ["SET_NFT"](state, payload) {
      state.nft = payload;
    },
    ["SET_HISTORIES"](state, payload) {
      state.histories = payload;
    },
    ["SET_TOTAL_HISTORY"](state, payload) {
      state.totalHistory = payload;
    },
    ["SET_QUERY_HISTORY"](state, payload) {
      state.queryHistory = payload;
    },
    ["SET_INFO"](state, payload) {
      state.info = payload;
    },
  },
  actions: {
    setQuery({ commit, state, dispatch }, type) {
      commit("SET_QUERY_OBJECT", {
        ...state.queryObject,
        ...type,
      });
      dispatch("fetchNfts");
    },
    setQueryHistory({ commit, state, dispatch }, address) {
      commit("SET_QUERY_HISTORY", {
        ...state.queryHistory,
        ...address,
      });
      dispatch("fetchHistory");
    },
    async fetchNfts({ commit, state }) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get(`/nft/${state.queryObject.address}/lists`, {
            params: state.queryObject,
          })
          .then((response) => {
            const dataConverse = response.data.data.map((item) => {
              return {
                ...item,
                price: item.price && ethers.utils.formatEther(item.price),
              };
            });
            commit("SET_NFTS", dataConverse);
            commit("SET_TOTAL", response.data.total);
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    async fetchHistory({ commit, state }) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get(`/marketplace/${state.queryHistory.address}/history`, {
            params: state.queryHistory,
          })
          .then((response) => {
            const dataConverse = response.data.data.map((item) => {
              return { ...item, price: ethers.utils.formatEther(item.price) };
            });
            commit("SET_TOTAL_HISTORY", response.data.total);
            commit("SET_HISTORIES", dataConverse);
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    async fetchInfo({ commit, state }, address) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get(`/nft/${address}`)
          .then((response) => {
            commit("SET_INFO", response.data);
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
  },
  getters: {
    listNfts: (state) => {
      if (state) return state.nfts;
    },
    listHistories: (state) => {
      if (state) return state.histories;
    },
  },
};
