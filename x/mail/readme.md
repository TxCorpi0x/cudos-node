# Mail module

Mail module is responsible for store and retrieval of messages, messages are simple texts that canbe forwarded as well.

## Params

```bash
./build/cudos-noded q mail params
```

Returns tha parameters of main module, maximum character count is the only parameter of this module.

```yaml
params:
  max_body_characters: 32
```

## Send Message

```bash
./build/cudos-noded tx mail send-message cudos1t9ce5kdq0a4u3ct3gjjzpv36jskyc0quf7ykz2 "Test Message 1" --from test0 --keyring-backend test --chain-id cudos-node -y
```

```bash
./build/cudos-noded tx mail send-message cudos1drwvvar643pg3g9azhrlc7z2vse3dzk5c7ummm "Test Message 2 with forwarded 1" 1 --from test0 --keyring-backend test --chain-id cudos-node -y
```

```bash
./build/cudos-noded tx mail send-message cudos17se9a4p6mr3ljlzgx9n6terlj7uy43rllfks9c "Test Message 3 with forwarded 2" 2 --from test0 --keyring-backend test --chain-id cudos-node -y
```

```bash
./build/cudos-noded tx mail send-message $(./build/cudos-noded keys show test0 --keyring-backend test --address) "Test Message 4 with forwarded 3" 3 --from test1 --keyring-backend test --chain-id cudos-node -y
```

## Query

### All Messages

```bash
./build/cudos-noded q mail list-message 
```

List of all messages:

```yaml
Message:
- body: Test Message 1
  forwardedId: "0"
  id: "1"
  receiver: cudos1t9ce5kdq0a4u3ct3gjjzpv36jskyc0quf7ykz2
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881573"
- body: Test Message 2 with forwarded 1
  forwardedId: "1"
  id: "2"
  receiver: cudos1drwvvar643pg3g9azhrlc7z2vse3dzk5c7ummm
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881578"
- body: Test Message 3 with forwarded 2
  forwardedId: "2"
  id: "3"
  receiver: cudos17se9a4p6mr3ljlzgx9n6terlj7uy43rllfks9c
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881583"
- body: Test Message 4 with forwarded 3
  forwardedId: "3"
  id: "4"
  receiver: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  sender: cudos1uu6vn6zgp0mu0mglakaunm7k8e5uy6u05gfztd
  ts: "1718881588"
pagination:
  next_key: null
  total: "0"
```

### Messages sent by test0

``` bash
./build/cudos-noded q mail list-message-sent $(./build/cudos-noded keys show test0 --keyring-backend test --address)
```

List of messages sent by test0:

```yml
Message:
- body: Test Message 1
  forwardedId: "0"
  id: "1"
  receiver: cudos1t9ce5kdq0a4u3ct3gjjzpv36jskyc0quf7ykz2
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881573"
- body: Test Message 2 with forwarded 1
  forwardedId: "1"
  id: "2"
  receiver: cudos1drwvvar643pg3g9azhrlc7z2vse3dzk5c7ummm
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881578"
- body: Test Message 3 with forwarded 2
  forwardedId: "2"
  id: "3"
  receiver: cudos17se9a4p6mr3ljlzgx9n6terlj7uy43rllfks9c
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881583"
pagination:
  next_key: null
  total: "0"
```

### Messages received by test0

``` bash
./build/cudos-noded q mail list-message-received $(./build/cudos-noded keys show test0 --keyring-backend test --address)
```

List of messages received by test0:

```yml
Message:
- body: Test Message 4 with forwarded 3
  forwardedId: "3"
  id: "4"
  receiver: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  sender: cudos1uu6vn6zgp0mu0mglakaunm7k8e5uy6u05gfztd
  ts: "1718881588"
pagination:
  next_key: null
  total: "0"
```

### Message Details

``` bash
./build/cudos-noded q mail show-message 4
```

Message details with the history of all forwards.

```yml
Message:
  body: Test Message 4 with forwarded 3
  forwardedId: "3"
  id: "4"
  receiver: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  sender: cudos1uu6vn6zgp0mu0mglakaunm7k8e5uy6u05gfztd
  ts: "1718881588"
forwards:
- body: Test Message 3 with forwarded 2
  forwardedId: "2"
  id: "3"
  receiver: cudos17se9a4p6mr3ljlzgx9n6terlj7uy43rllfks9c
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881583"
- body: Test Message 2 with forwarded 1
  forwardedId: "1"
  id: "2"
  receiver: cudos1drwvvar643pg3g9azhrlc7z2vse3dzk5c7ummm
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881578"
- body: Test Message 1
  forwardedId: "0"
  id: "1"
  receiver: cudos1t9ce5kdq0a4u3ct3gjjzpv36jskyc0quf7ykz2
  sender: cudos106kt3wzxwhv8k3jrvc2ac0a63hp38p889fpgen
  ts: "1718881573"
```
