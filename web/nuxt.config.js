require("dotenv").config();
require("dotenv").config({ path: `.env.${process.env.ASHWARD_ENV}` });
import webpack from "webpack";
export default {
  env: {
    RPC_URL: process.env.RPC_URL || "",
    CHAIN_ID: process.env.CHAIN_ID || 0,
    API_URL: process.env.API_URL || "",
    ASSET_URL: process.env.ASSET_URL || "",
    RPC_PROVIDER: process.env.RPC_PROVIDER || "",
    EVENT_TIME_IDO_WHITELIST: process.env.EVENT_TIME_IDO_WHITELIST || "",
    EVENT_TIME_IDO_FCFS: process.env.EVENT_TIME_IDO_FCFS || "",
    PORT: process.env.PORT || 3005,
  },
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title:
      "Ashward | A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic",
    htmlAttrs: {
      lang: "en",
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { name: "format-detection", content: "telephone=no" },
      // { name: 'google-site-verification', content: 'iE6fr0K52OKJQ2UvDvUZ9YFnnIaEiA2lz4Z3OxZ51as' },

      // base meta
      {
        name: "keywords",
        content:
          "nft game 2022, nft android, nft ios, nft games android, nft games for android, android nft games, nft games ios, upcoming nft games 2022,game nft android, free nft games android, ios nft games, game nft crypto android, nft games for ios, best crypto games, play to earn crypto, game coin crypto, nft game crypto, best crypto games 2021, crypto games android, crypto gaming coins, crypto game, play to earn crypto games 2021, play to earn crypto games 2022, bomb crypto game, free crypto games, earn crypto playing games, upland crypto, games to earn crypto, play to earn games crypto, new crypto games, hot nft game, Ashward , Ashward nft, Ashward game, Ashward ios, Ashward android, Ashward PC, Ashward gameplay, ASC",
      },
      {
        name: "description",
        content:
          "A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },
      {
        name: "subject",
        content: "An NFT game sets in medieval theme and pixel graphics",
      },
      { name: "copyright", content: "Copyright © 2022 Ashward" },
      { name: "language", content: "EN" },
      // { name: 'revised', content: '' },
      { name: "author", content: "Ashward Game" },
      // { name: 'designer', content: '' },
      { name: "url", content: "https://ashward.io" },
      { name: "identifier-URL", content: "https://ashward.io" },
      {
        name: "pagename",
        content:
          "Ashward | A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },
      { name: "coverage", content: "Worldwide" },
      { name: "distribution", content: "Global" },
      { name: "rating", content: "General" },
      { name: "revisit-after", content: "1 days" },
      { name: "date", content: "Jan. 23, 2022" },
      { name: "search_date", content: "2022-01-23" },
      { "http-equiv": "Pragma", content: "no-cache" },
      { "http-equiv": "Cache-Control", content: "no-cache" },
      { "http-equiv": "imagetoolbar", content: "no" },
      { "http-equiv": "x-dns-prefetch-control", content: "off" },
      // opengraph
      {
        property: "og:title",
        name: "og:title",
        content:
          "Ashward | A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },
      // https://ogp.me/#types
      { property: "og:type", name: "og:type", content: "website" },
      { property: "og:url", name: "og:url", content: "https://ashward.io" },
      {
        property: "og:image",
        name: "og:image",
        content: "https://ashward.io/assets/Gamecover.png",
      },
      // { property: 'og:image:width', name: 'og:image:width', content: '' },
      // { property: 'og:image:height', name: 'og:image:height', content: '' },
      {
        property: "og:image:alt",
        name: "og:image:alt",
        content:
          "Ashward | A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },
      {
        property: "og:site_name",
        name: "og:site_name",
        content: "https://ashward.io",
      },
      {
        property: "og:description",
        name: "og:description",
        content:
          "A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },

      // { name: 'fb:page_id', content: '' },
      // { name: 'application-name', content: '' },
      // { name: 'og:email', content: '' },
      // { name: 'og:latitude', content: '' },
      // { name: 'og:longitude', content: '' },
      // { name: 'og:street-address', content: '' },
      // { name: 'og:region', content: 'Ha noi' },
      // { name: 'og:country-name', content: 'Viet Nam' },
      // IOS
      {
        name: "apple-mobile-web-app-title",
        content:
          "Ashward | A virtual world that ultilizes the play-to-earn game model with medieval setting and pixel graphic.",
      },
      { name: "apple-mobile-web-app-capable", content: "yes" },
      { name: "apple-touch-fullscreen", content: "yes" },
      //
      {
        name: "news_keywords",
        content:
          "nft game 2022, nft android, nft ios, nft games android, nft games for android, android nft games, nft games ios, upcoming nft games 2022,game nft android, free nft games android, ios nft games, game nft crypto android, nft games for ios, best crypto games, play to earn crypto, game coin crypto, nft game crypto, best crypto games 2021, crypto games android, crypto gaming coins, crypto game, play to earn crypto games 2021, play to earn crypto games 2022, bomb crypto game, free crypto games, earn crypto playing games, upland crypto, games to earn crypto, play to earn games crypto, new crypto games, hot nft game, Ashward , Ashward nft, Ashward game, Ashward ios, Ashward android, Ashward PC, Ashward gameplay, ASC",
      },
    ],
    link: [
      { rel: "icon", type: "image/png", href: "/tokenfinal.png" },
      {
        rel: "apple-touch-icon",
        type: "image/png",
        href: "/tokenfinal.png",
        sizes: "72x72",
      },
      {
        rel: "apple-touch-startup-image",
        type: "image/png",
        href: "/tokenfinal.png",
        sizes: "72x72",
      },
      { rel: "canonical", href: "https://ashward.io" },
    ],
    script: [
      {
        src: "https://cdn.jsdelivr.net/npm/particles.js@2.0.0/particles.min.js",
      },
    ],
    script: [
      {
        src: '/bootstrap.bundle.min.js'
      }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["static/assets/scss/main.scss"],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    "~/plugins/axios/api.js",
    "~/plugins/axios/asset.js",
    "~/plugins/mixins/index",
    { src: "~/plugins/loading.js", mode: "client" },
    { src: "~/plugins/ether.js", mode: "client" },
    { src: "~/plugins/theme/core.js", ssr: false },

    "~/plugins/global",
    {
      src: "~/plugins/vue-countdown.js",
      ssr: false,
    },
    {
      src: "~/plugins/vue-flickity.js",
      ssr: false,
    },
    {
      src: "~/plugins/vue-lazyload-video.js",
      ssr: false,
    },
    {
      src: "~/plugins/vuex-persistedstate.js",
      ssr: false,
    },
    {
      src: "~/plugins/vue-spine.js",
      ssr: false,
    },
    {
      src: "~/plugins/wow-animation.js",
      ssr: false,
    },
    { src: "~/plugins/bootstrap.client.js" },
    {
      src: "~/plugins/ion-rangeslider.js",
      ssr: false,
    },
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    "@nuxtjs/eslint-module",
    // https://go.nuxtjs.dev/tailwindcss
    "@nuxtjs/tailwindcss",
    "@nuxt/image",
    "nuxt-font-loader",
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    "@nuxtjs/axios",

    // https://go.nuxtjs.dev/pwa
    "@nuxtjs/pwa",

    // https://github.com/nuxt-community/google-gtag-module
    "@nuxtjs/google-gtag",
    "primevue/nuxt",
    "vue-scrollto/nuxt",
    "@nuxtjs/device",
    "@nuxtjs/style-resources",
    "nuxt-user-agent",
    "nuxt-clipboard2",
    "@nuxtjs/redirect-module",
    "@nuxtjs/component-cache",
    "@/modules/axCache",
    "nuxt-webfontloader",
    ["cookie-universal-nuxt", { parseJSON: false }],
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    proxy: true,
  },
  proxy: {
    "/api/": {
      target: process.env.API_URL,
      changeOrigin: true,
    },
    //    "/assets/": {
    //      target: process.env.ASSET_URL,
    //      changeOrigin: true,
    //    },
  },

  server: {
    port: process.env.PORT,
  },

  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: "en",
    },
  },

  primevue: {
    theme: "saga-blue",
    ripple: true,
    components: ["Toast"],
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    plugins: [
      new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "window.jQuery": "jquery",
      }),
    ],
    publicPath: process.env.PUBLIC_URL || "/_nuxt/",
    parallel: true,
    extractCSS: {
      // ignoreOrder: true
    },
    optimizeCSS: false,
    babel: {
      compact: true,
    },
    optimization: {
      minimize: false,
    },
    splitChunks: {
      chunks: "all",
      automaticNameDelimiter: ".",
      name: true,
      cacheGroups: {},
    },
    transpile: ["primevue"],
    loaders: {
      cssModules: {
        modules: {
          localIdentName: "[hash:base64:4]",
        },
      },
    },
    extend(config, { isDev, isClient, loaders: { vue } }) {
      if (isClient) {
        vue.transformAssetUrls.img = ["data-src", "src"];
        vue.transformAssetUrls.source = ["data-srcset", "srcset"];
      }
    },
  },

  "google-gtag": {
    id: "GTM-W7TJP3J",
    config: {
      anonymize_ip: true, // anonymize IP
      send_page_view: false, // might be necessary to avoid duplicated page track on page reload
      linker: {
        domains: ["ashward.io"],
      },
    },
    debug: false, // enable to track in dev mode
    disableAutoPageTrack: false, // disable if you don't want to track each page route with router.afterEach(...).
    additionalAccounts: [
      {
        id: "GTM-W7TJP3J", // required if you are adding additional accounts
        config: {
          send_page_view: false, // optional configurations
        },
      },
    ],
  },
};
