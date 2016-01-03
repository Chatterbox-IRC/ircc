## Actions

These are actions you can perform using IRCC.

#### Nickname
- Description: Set or change a user's nickname.
- `nick`: Desired nickname.
- Format:

```json
{
  "type": "nick",
  "nick": "jharkness"
}
```

#### User
- Description: Set a user's user name and real name. Can only be run at start of connection.
- `user`: Desired username.
- `name`: Desired real name.
- Format:

```json
{
  "type": "user",
  "user": "msmith",
  "name": "Mickey Smith"
}
```

#### Quit
- Description: Close IRC connection.
- `msg`: Message to display while quitting (optional).
- Format:

```json
{
  "type": "quit",
  "msg": "I'm going to lunch"
}

```

#### Join
- Description: Join a channel.
- `channel`: Desired channel.
- `password`: Channel password (optional).
- Format:

```json
{
  "type": "join",
  "channel": "#satellite5",
  "password": "badwolf"
}
```

#### Part
- Description: Leave a channel.
- `channel`: Desired channel.
- Format:

```json
{
  "type": "part",
  "channel": "#doomsday",
}
```

#### Message
- Description: Send a message.
- `target`: channel or person to receive the message.
- `notice`: should the message be sent as a notice (bool)?
- `msg`: Message contents.
- Format:

```json
{
  "type": "msg",
  "target": "doctor",
  "notice": false,
  "msg": "I can't believe you just said that."
}
```
