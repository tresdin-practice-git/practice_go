let appendMessage = message => {
    let container = document.querySelector("message__container");
    let div = document.createElement('div');
    div.textContent = message.content;
    container.appendChild(div);
};
let ws = new WebSocket("ws://localhost:3000");
ws.onopen = event => {
    console.log("WebSocket open");
};
ws.onerror = event => {
    console.log('Error occurs');
};
ws.onmessage = event => {
    let data = JSON.parse(data);
    appendMessage(data);
};
