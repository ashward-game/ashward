import Vue from "vue";
Vue.mixin({
  data() {
    return {};
  },
  methods: {
    catchError(error) {
      // try to extract error message, otherwise return raw error
      let formatted_error;

      if (error.message.startsWith("invalid ENS name")) {
        formatted_error = "Missing or invalid parameter.";
      } else if (error.message.startsWith("invalid BigNumber string")) {
        formatted_error = "Invalid number parameter.";
      } else {
        try {
          let errors = JSON.stringify(error).match(EXTRACT_ERROR_MESSAGE);
          formatted_error = errors[errors.length - 1];
        } catch (e) {
          formatted_error = error.message;
        }
      }

      return formatted_error;
    },
    camelCase(text) {
      return text
        .replace(/(?:^\w|[A-Z]|\b\w)/g, function (word, index) {
          return index === 0 ? word.toLowerCase() : word.toUpperCase();
        })
        .replace(/\s+/g, "");
    },
  },
});
