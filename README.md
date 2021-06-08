# Go-CLI

Official node for running stakers in Golang.

## Installation

### Prerequisites
* Golang 1.15 or later must be installed.
* Latest stable version of node is required.
* Silicon chip based Mac users must go for node 15.3.0+

### Building the source
1. Run `npm install` to install the node dependencies.
2. Run `npm run build` to build the binary. While building the binary, supply the provider RPC url and the gas multiplier.
3. The binary will be generated at `build/bin`.

## Commands

Go to the `build/bin` directory where the razor binary is generated.

`cd build/bin`

### Create Account
Create an account using the `create` command
```
$ ./razor create <password>
```

Example:

```
$ ./razor create deadbeef
```

### Stake

If you have a minimum of 1000 razors in your account, you can stake those using the stake command.
```
$ ./razor stake --address <address> --password <password> --amount <amount>
```

Example:
```
$ ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef --amount 1000
```

### Vote
You can start voting once you've staked some razors
```
$ ./razor vote --address <address> --password <password>
```

Example:
```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef
```

### Unstake
If you wish to withdraw your funds, you can run the `unstake` command followed by the `withdraw` command.
```
$ ./razor unstake --address <address> --password <password>
```

Example:
```
$ ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef
```

### Withdraw
Once `unstake` has been called, you can withdraw your funds using the `withdraw` command

```
$ ./razor withdraw --address <address> --password <password>
```

Example:

```
$ ./razor withdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef
```
### Transfer
Transfers razor to other accounts.

```
$ ./razor transfer --amount <amount> --to <transfer_to_address> --from <transfer_from_address> --password <password>
```

Example:
```
$ ./razor transfer --amount 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef
```

### Create Job
You can create new jobs using `creteJob` command.

```
$ ./razor createJob --url <URL> --selector <selector_comma_seperated> --name <name> --fee <fee_to_lock> --address <address> --password <password>
```

Example:
```
$ ./razor createJob --url https://www.alphavantage.co/query\?function\=GLOBAL_QUOTE\&symbol\=MSFT\&apikey\=demo --selector "Global Quote,05. price" --fee 100 --name msft --repeat false --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --password deadbeef 
```

### Set Config
The config is set while the build is generated, but if you need to change your provider or the gas multiplier, you can use the `setconfig` command.

```
$ ./razor setconfig --provider <rpc_provider> --gasmultiplier <multiplier_value>
```

Example:
```
$ ./razor setconfig --provider https://infura/v3/matic --gasmultiplier 1.5
```