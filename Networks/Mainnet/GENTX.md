# GENTX & HARDFORK INSTRUCTIONS

### Install & Initialize

-   Install tucd binary

-   Initialize canto node directory

```bash
tucd init <node_name> --chain-id canto_7700-1
```

-   Download the [genesis file](https://github.com/Canto-Network/Canto/raw/genesis/Networks/Mainnet/genesis.json)

```bash
wget https://github.com/Canto-Network/Canto/raw/genesis/Networks/Mainnet/genesis.json -b $HOME/.tucd/config
```

### Create & Submit a GENTX file + genesis.json

A GENTX is a genesis transaction that adds a validator node to the genesis file.

```bash
tucd gentx <key_name> <token-amount>acanto --chain-id=canto_7700-1 --moniker=<your_moniker> --commission-max-change-rate=0.01 --commission-max-rate=0.10 --commission-rate=0.05 --details="<details here>" --security-contact="<email>" --website="<website>"
```

-   Fork [Canto](https://github.com/Canto-Network/Canto)

-   Copy the contents of `${HOME}/.tucd/config/gentx/gentx-XXXXXXXX.json` to `$HOME/Canto/Mainnet/Gentx/<yourvalidatorname>.json`

-   Create a pull request to the genesis branch of the [repository](https://github.com/Canto-Network/Canto/Mainnet/gentx)

### Restarting Your Node

You do not need to reinitialize your Canto Node. Basically a hard fork on Cosmos is starting from block 1 with a new genesis file. All your configuration files can stay the same. Steps to ensure a safe restart

1. Backup your data directory.

-   `mkdir $HOME/canto-backup`

-   `cp $HOME/.tucd/data $HOME/canto-backup/`

2. Remove old genesis

-   `rm $HOME/.tucd/genesis.json`

3. Download new genesis

-   `wget`

4. Remove old data

-   `rm -rf $HOME/.tucd/data`

6. Create a new data directory

-   `mkdir $HOME/.tucd/data`

7. copy the contents of the `priv_validator_state.json` file 

-   `nano $HOME/.tucd/data/priv_validator_state.json`

-   Copy the json string and paste into the file
 {
"height": "0",
 "round": 0,
 "step": 0
 }

If you do not reinitialize then your peer id and ip address will remain the same which will prevent you from needing to update your peers list.

8. Download the new binary

```
cd $HOME/Canto
git checkout <branch>
make install
mv $HOME/go/bin/tucd /usr/bin/
```

9. Restart your node

-   `systemctl restart tucd`

## Emergency Reversion

1. Move your backup data directory into your .tucd directory

-   `mv HOME/canto-backup/data $HOME/.canto/`

2. Download the old genesis file

-   `wget https://github.com/Canto-Network/Canto/raw/main/Mainnet/genesis.json -b $HOME/.tucd/config/`

3. Restart your node

-   `systemctl restart tucd`
