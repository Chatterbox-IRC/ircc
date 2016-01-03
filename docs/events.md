## Events

These are events that the IRCC can output.

#### Connected
- Triggered: When a user is fully connected to the IRC server.
- `target`: Sever the user is connected to.
- `msg`: Connection message.
- Format:

```json
{
  "type": "connected",
  "status": "ok",
  "target": "chat.matrix.gf",
  "msg": "Welcome!"
}
```
