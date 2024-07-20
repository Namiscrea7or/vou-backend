# Realtime Service

## Quiz Game Scenario

Convention: X (can be VOU Admin, or a Brand Admin, based on biz)

1. X creates a game session,
2. Multiple players connect to the game,
3. X starts the game,
4. X manually controls the game session or the game will automatically continue with predefined configuration,
5. The game session ends, pass the game result to Gifting/Voucher service.

## Debug Boilerplate

```js
const socket = new WebSocket("ws://localhost:8085/ws")
socket.onmessage = console.log
socket.send("hello")
```
