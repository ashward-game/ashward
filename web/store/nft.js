import { ethers } from "ethers";

export default {
  namespaced: true,
  state: () => ({
    isLoading: true,
    nft: null,
  }),
  mutations: {
    ["SET_IS_LOADING"](state, payload) {
      state.isLoading = payload;
    },
    ["SET_NFT"](state, payload) {
      state.nft = payload;
    },
  },
  actions: {
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
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            console.log(error);
            reject(error);
          });
      });
    },
    async fetchMetadata({ commit, state, dispatch }, tokenURI) {
      commit("SET_IS_LOADING", true);
      return new Promise((resolve, reject) => {
        this.$asset
          .get(tokenURI.replace(process.env.ASSET_URL, "/ipfs/"), {
            proxy: {
              host: process.env.ASSET_URL,
            },
          })
          .then((response) => {
            commit("SET_NFT", {
              ...state.nft,
              ...response.data,
            });
            commit("SET_IS_LOADING", false);
            resolve(response);
          })
          .catch((error) => {
            console.log(error);
            reject(error);
          });
      });
    },
    async fetchAll({ commit, state, dispatch }, id) {
      await dispatch("fetchNftDetail", id);
      await dispatch("fetchMetadata", state.nft.metadata_uri);
    },
  },
  getters: {},
};
