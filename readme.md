# Safrochain

![Safrochain Logo](https://i.ibb.co/99q9HK6D/Safrochain-Logo.png)

**Safrochain** is a Cosmos SDK-based blockchain designed for secure, scalable, and interoperable decentralized applications. Powered by its native token `SAF` and featuring full IBC (Inter-Blockchain Communication) support, Safrochain empowers developers to build innovative dApps that integrate seamlessly with the Cosmos ecosystem.

This guide provides comprehensive instructions to set up a Safrochain node, sync with the network, create accounts, and interact with the chain.

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Node Setup](#node-setup)
  - [Clone the Repository](#clone-the-repository)
  - [Build Safrochain](#build-safrochain)
  - [Initialize the Node](#initialize-the-node)
  - [Synchronize with the Main Network](#synchronize-with-the-main-network)
- [Account Creation](#account-creation)
- [Running the Node](#running-the-node)
- [Faucet](#faucet)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

---

## Features

- **Cosmos SDK-Based**: Built with the Cosmos SDK for performance and modularity.
- **Native Token**: `SAF` (base denom: `saf`, display denom: `SAF`, where `1 SAF = 10^6 microsaf`).
- **IBC Enabled**: Interoperable with other Cosmos chains.
- **Custom Module**: Includes a dedicated `safrochain` module ([Module Docs](#)).
- **Public Endpoints**:
  - RPC: `https://rpcsafro.cardanotask.com`
  - REST: `https://safro.cardanotask.com`
- **Chain ID**: `safrochain`

---

## Prerequisites

Ensure you have the following installed:

- **Go (>=1.18)**: [Install Go](https://go.dev/doc/install)
- **Git**: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **Make**: Usually included in Linux distributions
- **jq**: Optional, for JSON parsing
- **Other dependencies** from Cosmos SDK

To install prerequisites on Ubuntu/Debian:

```bash
sudo apt update && sudo apt install -y git make jq build-essential

# Install Go
wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

---

## Installation

### Clone the Repository

**Option 1: Fork and Clone**
```bash
git clone https://github.com/<your-username>/safrochain.git
cd safrochain
```

**Option 2: Direct Clone**
```bash
git clone https://github.com/safrochain/safrochain.git
cd safrochain
```

---

### Build Safrochain

```bash
make install
```
This will install `safrochaind` to `$GOPATH/bin`.

Ensure `$GOPATH/bin` is in your environment:
```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
safrochaind version
```

---

### Initialize the Node

```bash
safrochaind init <your-moniker> --chain-id safrochain
```
Replace `<your-moniker>` with a unique identifier for your node.

---

### Synchronize with the Main Network

Download the latest genesis file:
```bash
curl http://180.149.197.176:26657/genesis | jq '.result.genesis' > ~/.safrochaind/config/genesis.json
```

Configure the persistent peer:
```bash
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$(curl -s http://180.149.197.176:26657/status | jq -r '.result.node_info.id')@180.149.197.176:26656\"/" ~/.safrochaind/config/config.toml
```

Set minimum gas price:
```bash
sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"0.0001saf\"/" ~/.safrochaind/config/app.toml
```

Download address book (optional):
```bash
wget -O ~/.safrochaind/config/addrbook.json http://180.149.197.176:26657/addrbook.json
```

---

## Account Creation

Create a new key:
```bash
safrochaind keys add <your-key-name>
```
Backup the mnemonic securely.

List all keys:
```bash
safrochaind keys list
```

---

## Running the Node

Start the node:
```bash
safrochaind start --minimum-gas-prices=0.0001saf
```

To run in background:
```bash
nohup safrochaind start --minimum-gas-prices=0.0001saf > safrochaind.log 2>&1 &
```

Check node status:
```bash
curl http://localhost:26657/status
```

---

## Faucet

To receive test `SAF` tokens:

1. Visit: https://faucet.cardanotask.com
2. Enter your `addr_safro...` address

Check balance:
```bash
safrochaind query bank balances <your-address> --node https://rpcsafro.cardanotask.com
```

---

## Contributing

We welcome contributions!

- Fork the repo
- Create your branch: `git checkout -b feature/my-feature`
- Commit: `git commit -m "Add my feature"`
- Push: `git push origin feature/my-feature`
- Open a Pull Request

Please follow the Code of Conduct and project guidelines.

---

## License

Safrochain is licensed under the **Apache License 2.0**. See the [LICENSE](./LICENSE) file for details.

---

## Contact

- GitHub: [github.com/safrochain/safrochain](https://github.com/safrochain/safrochain)
- Faucet: [faucet.cardanotask.com](https://faucet.cardanotask.com)
- Node IP: `180.149.197.176`
- Community: *Discord/Telegram links coming soon*
- Issues: [GitHub Issues](https://github.com/safrochain/safrochain/issues)

For technical help, reach out via GitHub or community channels.

---

Happy building on **Safrochain**! 🚀

