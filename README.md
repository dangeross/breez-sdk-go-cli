```
$ go run . -d data/test

__________                                  _________________   ____  __.
\______   \_______   ____   ____ ________  /   _____/\______ \ |    |/ _|
 |    |  _/\_  __ \_/ __ \_/ __ \\___   /  \_____  \  |    |  \|      <  
 |    |   \ |  | \/\  ___/\  ___/ /    /   /        \ |    `   \    |  \ 
 |______  / |__|    \___  >\___  >_____ \ /_______  //_______  /____|__ \
        \/              \/     \/      \/         \/         \/        \/


A Breez SDK Go CLI

Commands:
=========
  clear                     clear the screen
  connect                   initialize an SDK instance
  exit                      exit the shell
  help                      use 'help [command]' for command help
  lnurl_auth                authenticate using LNURL-auth
  lnurl_pay                 send payment using LNURL-pay
  lnurl_withdraw            receive payment using LNURL-withdraw
  node_info                 get the latest node state
  receive_onchain           receive on-chain using a swap
  receive_payment           generate a bolt11 invoice
  send_onchain              send on-chain using a reverse swap
  send_payment              send a lightning payment
  send_spontaneous_payment  send a spontaneous (keysend) payment
  set_api_key               configure the Breez API key

sdk>  
```