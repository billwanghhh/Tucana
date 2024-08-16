KEY1="key1"
KEY2="key2"
KEY3="key3"
KEY4="key4"
CHAINID="canto_7900-1"
MONIKER1="validator1"
MONIKER2="validator2"
MONIKER3="validator3"
MONIKER4="validator4"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"


# if $KEY exists it should be deleted
cantod keys add $KEY1 --keyring-backend $KEYRING --algo $KEYALGO
cantod keys add $KEY2 --keyring-backend $KEYRING --algo $KEYALGO
cantod keys add $KEY3 --keyring-backend $KEYRING --algo $KEYALGO
cantod keys add $KEY4 --keyring-backend $KEYRING --algo $KEYALGO


# Set moniker and chain-id for Canto (Moniker can be anything, chain-id must be an integer)
cantod init $MONIKER1 --chain-id $CHAINID

# Change parameter token denominations to acanto
cat $HOME/.cantod/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["staking"]["params"]["max_validators"]=7' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
# cat $HOME/.cantod/config/genesis.json | jq '.app_state["distribution"]["params"]["base_proposer_reward"]="0.010000000000000000"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
# cat $HOME/.cantod/config/genesis.json | jq '.app_state["distribution"]["params"]["bonus_proposer_reward"]="0.040000000000000000"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["gov"]["params"]["expedited_min_deposit"][0]["denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
# cat $HOME/.cantod/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["coinswap"]["params"]["pool_creation_fee"]["denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json
cat $HOME/.cantod/config/genesis.json | jq '.app_state["coinswap"]["standard_denom"]="acanto"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json

# change downtime slash to longer
cat $HOME/.cantod/config/genesis.json | jq '.app_state["slashing"]["params"]["signed_blocks_window"]=30000' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json

# change max gas
cat $HOME/.cantod/config/genesis.json | jq '.consensus["params"]["block"]["max_gas"]="30000000"' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json



#sed -i '/\[json-rpc\]/a enable-websocket = true' $HOME/.cantod/config/app.toml
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $HOME/.cantod/config/config.toml
sed -i 's/prometheus = false/prometheus = true/' $HOME/.cantod/config/config.toml
sed -i 's/minimum-gas-prices = "0acanto"/minimum-gas-prices = "0.0001acanto"/' $HOME/.cantod/config/app.toml
sed -i 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3"/' $HOME/.cantod/config/app.toml
sed -i "s/chain-id = \"ethermint_9000-1\"/chain-id = \"$CHAINID\"/" $HOME/.cantod/config/client.toml

# Allocate genesis accounts (cosmos formatted addresses)
cantod add-genesis-account $KEY1 1050000000000000000000000000acanto --keyring-backend $KEYRING
cantod add-genesis-account $KEY2 1000000000000000000000000000acanto --keyring-backend $KEYRING
cantod add-genesis-account $KEY3 1000000000000000000000000000acanto --keyring-backend $KEYRING
cantod add-genesis-account $KEY4 1000000000000000000000000000acanto --keyring-backend $KEYRING

# Update total supply with claim values
#validators_supply=$(cat $HOME/.cantod/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=4050000000000000000000000000
cat $HOME/.cantod/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.cantod/config/tmp_genesis.json && mv $HOME/.cantod/config/tmp_genesis.json $HOME/.cantod/config/genesis.json


cantod init $MONIKER2 --chain-id $CHAINID --home ~/.cantod2
cp -r ~/.cantod/keyring-test $HOME/.cantod2
cp $HOME/.cantod/config/genesis.json $HOME/.cantod2/config/genesis.json

cantod init $MONIKER3 --chain-id $CHAINID --home ~/.cantod3
cp -r ~/.cantod/keyring-test $HOME/.cantod3
cp $HOME/.cantod/config/genesis.json $HOME/.cantod3/config/genesis.json

cantod init $MONIKER4 --chain-id $CHAINID --home ~/.cantod4
cp -r ~/.cantod/keyring-test $HOME/.cantod4
cp $HOME/.cantod/config/genesis.json $HOME/.cantod4/config/genesis.json


echo $KEYRING
echo $KEY1
# Sign genesis transaction
mkdir $HOME/.cantod/config/gentx
cantod gentx $KEY1 900000000000000000000000acanto --commission-rate 0.1 --commission-max-rate 1.0 --commission-max-change-rate 1.0 --keyring-backend $KEYRING --chain-id $CHAINID --output-document $HOME/.cantod/config/gentx/gentx-1.json
cantod gentx $KEY2 900000000000000000000000acanto --commission-rate 0.1 --commission-max-rate 1.0 --commission-max-change-rate 1.0 --keyring-backend $KEYRING --chain-id $CHAINID --output-document $HOME/.cantod/config/gentx/gentx-2.json --home ~/.cantod2
cantod gentx $KEY3 900000000000000000000000acanto --commission-rate 0.1 --commission-max-rate 1.0 --commission-max-change-rate 1.0 --keyring-backend $KEYRING --chain-id $CHAINID --output-document $HOME/.cantod/config/gentx/gentx-3.json --home ~/.cantod3
cantod gentx $KEY4 900000000000000000000000acanto --commission-rate 0.1 --commission-max-rate 1.0 --commission-max-change-rate 1.0 --keyring-backend $KEYRING --chain-id $CHAINID --output-document $HOME/.cantod/config/gentx/gentx-4.json --home ~/.cantod4


# Collect genesis tx
cantod collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
cantod validate-genesis

cp $HOME/.cantod/config/genesis.json $HOME/.cantod2/config/genesis.json
cp $HOME/.cantod/config/genesis.json $HOME/.cantod3/config/genesis.json
cp $HOME/.cantod/config/genesis.json $HOME/.cantod4/config/genesis.json

# edit each config.toml and app.toml to change ports, moniker and add peers, if run 4 node under the same ip, set `allow_duplicate_ip = true` in config.toml, then start 4 nodes
# cantod start --chain-id canto_7900-1
# cantod start --chain-id canto_7900-1 --home ~/.cantod2
# cantod start --chain-id canto_7900-1 --home ~/.cantod3
# cantod start --chain-id canto_7900-1 --home ~/.cantod4


