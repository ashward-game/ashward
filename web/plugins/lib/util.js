import moment from "moment";
import whiteListIDO from "@/static/data/IDOWhitelist";

export default {
  checkInWhiteListIDO(address) {
    for (let i = 0; i < whiteListIDO.length; i++) {
      if (whiteListIDO[i].toLowerCase() === address.toLowerCase()) return true;
    }
    return false;
  },
  textAddress(text) {
    return text.slice(0, 5) + "..." + text.slice(-5);
  },
  startTimer(idText, eventTime, textExpire) {
    let countDownDate = moment(eventTime, "YYYY-MM-DD HH:mm:ss");

    // Update the count down every 1 second
    let x = setInterval(function () {
      // Get today's date and time
      let now = moment(moment.utc().format("YYYY-MM-DD HH:mm:ss"));

      // Find the distance between now and the count down date
      let distance = moment.duration(countDownDate.diff(now));

      // Time calculations for days, hours, minutes and seconds
      let days = Math.floor(distance / (1000 * 60 * 60 * 24)).toString();
      let hours = Math.floor(
        (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
      ).toString();
      let minutes = Math.floor(
        (distance % (1000 * 60 * 60)) / (1000 * 60)
      ).toString();
      let seconds = Math.floor((distance % (1000 * 60)) / 1000).toString();

      hours = hours.length === 1 ? "0" + hours : hours;
      minutes = minutes.length === 1 ? "0" + minutes : minutes;
      seconds = seconds.length === 1 ? "0" + seconds : seconds;

      // Display the result in the element with id="demo"
      document.getElementById(idText).innerHTML =
        hours + ":" + minutes + ":" + seconds;

      // If the count down is finished, write some text
      if (distance < 0) {
        document.getElementById(idText).innerHTML = textExpire;
        clearInterval(x);
      }
    }, 1000);
  },
  numberWithCommas(x) {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") ?? 0;
  },
};
