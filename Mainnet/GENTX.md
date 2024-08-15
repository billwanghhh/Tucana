# GENTX & HARDFORK INSTRUCTIONS

### Install & Initialize 

* Install tucd binary

* Initialize canto node directory 
```bash
tucd init <node_name> --chain-id <chain_id>
```
* Download the [genesis file](https://github.com/Canto-Network/Canto/raw/main/Mainnet/genesis.json)
```bash
wget https://github.com/Canto-Network/Canto/raw/main/Mainnet/genesis.json -b $HOME/.tucd/config
```

### Add a Genesis Account
A genesis account is required to create a GENTX

```bash
tucd add-genesis-account <address-or-key-name> acanto --chain-id <chain-id>
```
### Create & Submit a GENTX file + genesis.json
A GENTX is a genesis transaction that adds a validator node to the genesis file.
```bash
tucd gentx <key_name> <token-amount>acanto --chain-id=<chain_id> --moniker=<your_moniker> --commission-max-change-rate=0.01 --commission-max-rate=0.10 --commission-rate=0.05 --details="<details here>" --security-contact="<email>" --website="<website>"
```
* Fork [Canto](https://github.com/Canto-Network/Canto)

* Copy the contents of `${HOME}/.tucd/config/gentx/gentx-XXXXXXXX.json` to `$HOME/Canto/Mainnet/gentx/<yourvalidatorname>.json`

* Copy the genesis.json file `${HOME}/.tucd/config/genesis.json` to `$HOME/Canto/Mainnet/Genesis-Files/`

* Create a pull request to the main branch of the [repository](https://github.com/Canto-Network/Canto/Mainnet/gentx)

### Restarting Your Node

You do not need to reinitialize your Canto Node. Basically a hard fork on Cosmos is starting from block 1 with a new genesis file. All your configuration files can stay the same. Steps to ensure a safe restart

1) Backup your data directory. 
* `mkdir $HOME/canto-backup` 

* `cp $HOME/.tucd/data $HOME/canto-backup/`

2) Remove old genesis 

* `rm $HOME/.tucd/genesis.json`

3) Download new genesis

* `wget`

4) Remove old data

* `rm -rf $HOME/.tucd/data`

5) Create a new data directory

* `mkdir $HOME/.tucd/data`

If you do not reinitialize then your peer id and ip address will remain the same which will prevent you from needing to update your peers list.

7) Download the new binary
```
cd $HOME/Canto
git checkout <branch>
make install
mv $HOME/go/bin/tucd /usr/bin/
```


6) Restart your node

* `systemctl restart tucd`

## Emergency Reversion

1) Move your backup data directory into your .tucd directory 

* `mv HOME/canto-backup/data $HOME/.canto/`

2) Download the old genesis file

* `wget https://github.com/Canto-Network/Canto/raw/main/Mainnet/genesis.json -b $HOME/.tucd/config/`

3) Restart your node

* `systemctl restart tucd`