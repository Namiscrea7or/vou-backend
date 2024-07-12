# Realtime Service

## Debug Boilerplate

```js
const socket = new WebSocket("ws://localhost:8085/ws")
socket.onmessage = console.log
socket.send("hello")
```
