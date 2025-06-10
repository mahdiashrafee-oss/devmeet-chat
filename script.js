const ws = new WebSocket("ws://"+location.host + "/ws");
const pc = new RTCPeerConnection({
    iceServers: [{urls: "stun:stun.l.google.com:19302"}]
});

const local = document.getElementById("local");
const remote = document.getElementById("remote");

pc.ontrack = ({streams: [stream]}) => {
    remote.srcObject = stream;
};

pc.onicecandidate = ({candidate}) => {
    ws.send(JSON.stringify({candidate}));
};

ws.onmessage = ({data}) => {
    const message = JSON.parse(data);
    if (message.candidate) {
        pc.addIceCandidate(message.candidate);
    } else {
        pc.setRemoteDescription(message.description);
    }
};

navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
    local.srcObject = stream;
    pc.addStream(stream);
});


ws.onopen = async () => {
  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);
  ws.send(JSON.stringify({description: pc.localDescription})); 
}
