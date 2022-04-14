import { ethers } from "ethers";

export default {
  namespaced: true,
  state: () => ({
    isLoading: true,
    filter: null,
    loading: true,
    nfts: [],
    total: 0,
    queryObject: {
      type: "character",
      rarity: null,
      star: null,
      level: null,
      order_by_price: null,
      page: 1,
      limit: 20,
    },
    nft: null,
  }),
  mutations: {
    ["SET_LOADING"](state, payload) {
      state.loading = payload;
    },
    ["SET_MESSAGE"](state, payload) {
      state.message = payload;
    },
    ["SET_FILTER"](state, payload) {
      state.filter = payload;
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
    ["SET_PAGE"](state, payload) {
      state.page = payload;
    },
    ["SET_IS_LOADING"](state, payload) {
      state.isLoading = payload;
    },
    ["SET_NFT"](state, payload) {
      state.nft = payload;
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
    async fetchNfts({ commit, state }) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get("/marketplace", { params: state.queryObject })
          .then((response) => {
            const dataConverse = response.data.data.map((item) => {
              return { ...item, price: ethers.utils.formatEther(item.price) };
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
    async resetFilter({ commit, state, dispatch }) {
      commit("SET_QUERY_OBJECT", {
        ...state.queryObject,
        rarity: null,
        star: null,
        level: null,
        order_by_price: null,
      });
      dispatch("fetchNfts");
    },
    async fetchNftDetail({ commit, state, dispatch }, id) {
      commit("SET_IS_LOADING", true);
      return new Promise((resolve, reject) => {
        this.$api
          .get(`/marketplace/nft/${id}`)
          .then(async (response) => {
            if (response.data && response.data.metadata_uri) {
              const historyConverse = response.data.marketplaces.map((item) => {
                return { ...item, price: ethers.utils.formatEther(item.price) };
              });
              const dataConverse = {
                ...response.data,
                price:
                  response.data.price &&
                  ethers.utils.formatEther(response.data.price),
                marketplaces: historyConverse,
              };
              commit("SET_NFT", {
                ...dataConverse,
              });
            }
            resolve(response);
          })
          .catch((error) => {
            console.log(error);
            reject(error);
          });
      });
    },
  },
  getters: {
    listNfts: (state) => {
      if (state) return state.nfts;
    },
  },
};
