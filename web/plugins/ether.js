import Web3Modal from "web3modal";
import { ethers } from "ethers";
import WalletConnectProvider from "@walletconnect/web3-provider";
var web3Provider, web3Account, web3Modal;
async function connectWeb3(context) {
  try {
    const providerOptions = {
      walletconnect: {
        package: WalletConnectProvider, // required
        options: {
          rpc: {
            [process.env.CHAIN_ID]: process.env.RPC_URL,
          },
        },
      },
    };

    web3Modal = new Web3Modal({
      cacheProvider: true,
      providerOptions,
    });

    const web3ModalProvider = await web3Modal.connect();
    web3Provider = new ethers.providers.Web3Provider(web3ModalProvider);

    // FIXME remove this when going to production
    const { chainId } = await web3Provider.getNetwork();

    if (chainId.toString() !== process.env.CHAIN_ID) {
      if (process.env.CHAIN_ID === "97")
        alert("Currently we only support BSC testnet");
      else if (process.env.CHAIN_ID === "56")
        alert("Currently we only support BSC mainnet");
      else alert("Currently we only support ganache");

      return;
    }

    localStorage.setItem("web3ModalConnected", true);

    web3Account = web3ModalProvider.selectedAddress;

    web3ModalProvider.on("connect", async (chainId) => {
      this.$emit("Web3Connect", web3Account);
      window.location.reload();
    });

    web3ModalProvider.on("accountsChanged", async (accounts) => {
      if (accounts.length > 0) web3Account = accounts[0];
      else web3Account = undefined;
    });

    // see https://docs.ethers.io/v5/concepts/best-practices/#best-practices--network-changes
    web3ModalProvider.on("chainChanged", async (chainId) => {
      window.location.reload();
    });

    web3ModalProvider.on("disconnect", async () => {
      web3Account = undefined;
      localStorage.removeItem("web3ModalConnected");
    });
    const storeModule = {
      state: () => ({
        isConnectedWeb3: false,
        web3Account: web3Account,
      }),
      mutations: {
        ["SET_IS_CONNECTED_WEB3"](state, payload) {
          state.isConnectedWeb3 = payload;
        },
        ["SET_WEB3_ACCOUNT"](state, payload) {
          state.web3Account = payload;
        },
      },
    };
    context.store.registerModule("ether", storeModule);
    context.store.commit("SET_IS_CONNECTED_WEB3", true);
  } catch (e) {
    console.log(e);
  }
}

async function disconnectWallet(context, web3Modal) {
  await web3Modal.clearCachedProvider();
  localStorage.removeItem("web3ModalConnected");
  context.store.commit("SET_IS_CONNECTED_WEB3", false);
  context.store.commit("SET_WEB3_ACCOUNT", null);
  web3Account = undefined;
  web3Provider = undefined;
}

export default ({ app }, inject) => {
  inject("ether", () => connectWeb3(app.context));
  inject("web3Provider", () => web3Provider);
  inject("disconnectWallet", () => disconnectWallet(app.context, web3Modal));
};
