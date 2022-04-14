export default function ({ $axios }, inject) {
  let configUrl;
  const api = $axios.create();
  api.onRequest((config) => {
    configUrl = config;
    config.url = "/api/v1" + config.url;
  });
  api.onError((error) => {
    const code = parseInt(error.response && error.response.status);
    if (code === 400) {
      console.log(error);
    }
  });
  inject("api", api);
}
