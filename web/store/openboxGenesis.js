export default {
  namespaced: true,
  state: () => ({
    isLoading: true,
    hash: null,
    signature: null,
    nft: null,
  }),
  mutations: {
    ["SET_IS_LOADING"](state, payload) {
      state.loading = payload;
    },
    ["SET_HASH"](state, payload) {
      state.hash = payload;
    },
    ["SET_SIGNATURE"](state, payload) {
      state.signature = payload;
    },
    ["SET_NFT"](state, payload) {
      state.nft = payload;
    },
  },
  actions: {
    async getCommit({ commit }) {
      return new Promise((resolve, reject) => {
        commit("SET_IS_LOADING", true);
        this.$api
          .get(`/makerand/commit`)
          .then((response) => {
            commit("SET_HASH", response.data.hash);
            commit("SET_SIGNATURE", response.data.signature);
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
    getHash: (state) => {
      if (state) return state.hash;
    },
    getSignature: (state) => {
      if (state) return state.signature;
    },
  },
};
