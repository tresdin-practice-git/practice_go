

let appendMessage = message => {
    let container = document.querySelector("message__container");
    let div = document.createElement('div');
    div.textContent = message.content;
    container.appendChild(div);
};
let ws = new WebSocket("ws://localhost:3000/ws");
ws.onopen = event => {
    console.log("WebSocket open");
};
ws.onerror = event => {
    console.log('Error occurs', event);
};
ws.onmessage = event => {
    let data = JSON.parse(data);
    appendMessage(data);
};
let inputText = document.querySelector('.message__input_text');
inputText.onkeydown = event => {
    if(event.key !== 'Enter'){
        return;
    }
    ws.send(JSON.stringify({content: inputText.textContent}));
    inputText.textContent = '';
};
