export default function ({ $axios }, inject) {
  const asset = $axios.create();
  asset.onRequest((config) => {});
  asset.onError((error) => {
    const code = parseInt(error.response && error.response.status);
    if (code === 400) {
      console.log(error);
    }
  });
  inject("asset", asset);
}
