const WebSocket = require("ws");

const ws = new WebSocket("wss://cekirdektenyetisenler.kartaca.com/ws");

function returnKey(text) {
  const arr1 = [
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
    "g",
    "h",
    "i",
    "j",
    "k",
    "l",
    "m",
  ];
  const arr2 = [
    "z",
    "y",
    "x",
    "w",
    "v",
    "u",
    "t",
    "s",
    "r",
    "q",
    "p",
    "o",
    "n",
  ];

  let text1 = String(text);
  let text2 = "";

  for (var i = 0; i < text1.length; i++) {
    let ch = text1.charAt(i);

    if (arr1.includes(ch)) {
      let index = arr1.indexOf(ch);
      text2 = text2.concat(arr2[index]);
    } else if (arr2.includes(ch)) {
      let index = arr2.indexOf(ch);
      text2 = text2.concat(arr1[index]);
    } else {
      text2 = text2.concat(ch);
    }
  }

  const regex = /registrationKey\s*:\s*([\w\d]+)/;

  const match = regex.exec(text2);

  if (match && match[1]) {
    var registrationKey = match[1];
  } else {
    console.log("Registration key not found.");
    var registrationKey = undefined;
  }

  return registrationKey;
}

ws.on("open", function () {
  console.log("WebSocket client connected");
});

ws.on("message", function (message) {
  console.log(`Received message: ${message}`);
  const key = returnKey(message);
  const sendM = {
    type: "REGISTER",
    name: "Talha",
    surname: "Ünal",
    email: "talhaunal7@gmail.com",
    registrationKey: key,
  };
  if (key) {
    ws.send(JSON.stringify(sendM));
  }
});

ws.on("close", function () {
  console.log("WebSocket client disconnected");
});
