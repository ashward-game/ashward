import COOKIES_DEVICE from './cookies-device'

export const state = () => ({
  device: 'mobile',
  isLoading: false,
  isWalletConnect: false
})

export const mutations = {
  loading (state, payload) {
    state.isLoading = payload
  },
  walletConnect (state, payload) {
    state.isWalletConnect = payload
  },
  mutate (state, payload) {
    state[payload.property] = typeof payload.with === 'object' && payload.with !== null ? { ...state[payload.property], ...payload.with } : payload.with
  }
}
export const actions = { ...COOKIES_DEVICE }
