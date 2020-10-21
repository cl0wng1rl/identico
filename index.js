const message = document.querySelector("#message");
const complexity = document.querySelector("#complexity");
const url = document.querySelector("#url");

const identicon = document.querySelector("#identicon");

const mainURL = "https://identico.herokuapp.com/";

function debounce(callback, wait) {
  let timeout = null;

  return function () {
    const callNow = !timeout;
    const next = () => callback.apply(this, arguments);

    clearTimeout(timeout);
    timeout = setTimeout(next, wait);

    if (callNow) next();
  };
}

const makeURL = (messageValue, complexityValue) => {
  return `${mainURL}?message=${messageValue}&quadrant=${complexityValue}`;
};

const setURL = () => {
  const messageValue = message.value;
  const complexityValue = complexity.value;
  const urlValue = makeURL(messageValue, complexityValue);
  url.value = urlValue;
  identicon.src = urlValue;
};

const setImageSource = (urlValue) => (identicon.src = urlValue);
const setURLTextboxValue = (urlValue) => (url.value = urlValue);

const handleChange = () => {
  const urlValue = makeURL(message.value, complexity.value);
  setImageSource(urlValue);
  setURLTextboxValue(urlValue);
};

const handleURLChange = () => {
  setImageSource(url.value);
};

message.addEventListener("keyup", debounce(handleChange, 300));
complexity.addEventListener("change", handleChange);
url.addEventListener("keyup", debounce(handleURLChange, 300));
