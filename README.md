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
  backup                     start a backup of the local data
  buy_bitcoin                generate a URL to buy bitcoin
  check_message              verify a signed message
  clear                      clear the screen
  close_lsp_channels         close all LSP channels
  connect                    initialize an SDK instance
  connect_lsp                connect to another LSP
  disconnect                 stop the SDK instance
  execute_dev_command        execute a low level node command (used for debugging)
  exit                       exit the shell
  fetch_onchain_fees         fetch the current fees for a reverse swap
  help                       use 'help [command]' for command help
  in_progress_reverse_swaps  fetch the in-progress reverse swaps
  in_progress_swap           fetch the in-progress swap
  list_lsps                  list available LSPs
  list_payments              list and filter payments
  list_refundables           list refundable swap addresses
  lnurl_auth                 authenticate using LNURL-auth
  lnurl_pay                  send payment using LNURL-pay
  lnurl_withdraw             receive payment using LNURL-withdraw
  lsp_info                   get the latest LSP information
  max_reverse_swap_amount    calculate the maximum amount for a reverse swap
  node_credentials           get the node credentials
  node_info                  get the latest node state
  open_channel_fee           calculate the open channel fee
  parse                      parse text to get its type and relevant metadata
  payment_by_hash            fetch a payment by its hash
  prepare_refund             prepare a refund transaction for an incomplete swap
  prepare_sweep              calculate the fee for a potential transaction
  receive_onchain            receive on-chain using a swap
  receive_payment            generate a bolt11 invoice
  recommended_fees           list recommended fees based on the mempool
  refund                     broadcast a refund transaction for an incomplete swap
  report_payment_failure     send a payment failure report
  send_onchain               send on-chain using a reverse swap
  send_payment               send a lightning payment
  send_spontaneous_payment   send a spontaneous (keysend) payment
  service_health_check       fetch the latest service health check
  set_api_key                configure the Breez API key
  sign_message               sign a message
  static_backup              fetch the static backup data
  sweep                      send on-chain funds to an external address
  sync                       sync local data with remote node

sdk>  
```