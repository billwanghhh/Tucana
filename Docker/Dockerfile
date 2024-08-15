FROM ubuntu:latest 

# add your keyname
ARG KEY_NAME <add your keyname>

# add the nickname for your validator on the network
ARG MONIKER <add your moniker>

# add the network id
ARG NETWORK_ID canto_9624-1

# arbitrary validator metadata
ARG DETAILS <details as string>

# add amount you're delegating
ARG TOKEN_DELEGATION <amount>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update &&\ 
	apt-get install -y apt-utils git git-extras software-properties-common vim \
	wget ca-certificates curl build-essential libssl-dev make openssh-server snap

RUN snap install go --classic

# builds out tucd
RUN git clone https://github.com/Canto-Network/Canto-Testnet
WORKDIR /Canto-Testnet/cmd/tucd
RUN go install -tags ledger ./... &&\
    mv $HOME/go/bin/tucd /usr/bin/

# generate keys
RUN tucd keys add $KEY_NAME
    # tucd keys add $KEY_NAME --recover &&\
    # tucd keys add $KEY_NAME --ledger

# initialize node
RUN tucd init $MONIKER --chain-id canto_9624-1

# get genesis file
RUN wget https://github.com/Canto-Network/Canto-Testnet-v2/raw/main/genesis.json -P $HOME/.tucd/config/
RUN sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0001acanto"/g' $HOME/.tucd/config/app.toml
RUN wget https://github.com/Canto-Network/Canto-Testnet-v2/raw/main/Docker/tucd.service -P $HOME/systemd/system/

# run tucd
WORKDIR /Canto-Testnet/cmd/tucd
RUN sudo systemctl daemon-reload &&\
    sudo systemctl enable tucd.service &&\
    systemctl start tucd && journalctl -u tucd -f

ENTRYPOINT ["/bin/bash"]
