# Safrochain Testnet Node Setup Guide

![Safrochain Logo](https://i.ibb.co/99q9HK6/Safrochain-Logo.png)

This guide walks you through setting up a **Safrochain testnet node** in local mode and creating a validator, step by step. It’s designed for beginners and includes copy-paste commands with detailed explanations for **Linux (Ubuntu)**, **macOS**, and **Windows (via WSL2)**. By the end, you’ll have a running node synced with the Safrochain testnet, a validator wallet funded via the faucet, and be ready to stake as a validator.

## 📋 Prerequisites

Before starting, ensure you have the following:

- **System**:
  - Linux (Ubuntu recommended), macOS, or Windows (with WSL2).
  - Minimum: 2GB RAM, 20GB disk space.
  - Windows users must install WSL2 and Ubuntu (see Step 1 for Windows).
- **Internet**: Stable connection for cloning repositories, downloading the genesis file, and syncing the blockchain.
- **Permissions**: Root/admin access (`sudo` on Linux/macOS, admin terminal in WSL2) for installing packages and configuring firewalls.
- **Testnet Configuration**:
  - **Main Node**: `d172b3424a96ebbb806cd19e09d6976db9bb68ea@88.99.211.113:26656`
  - **Genesis File**: `https://safrochain.com/genesis/testnet/genesis.json`
  - **Faucet**: `https://faucet.safrochain.com` (provides 2,500,000,000 `saf` = 2,500 `hela` per request)
- **Testnet Denominations**: Uses `saf` (base unit) and `hela` (1 `hela` = 1,000,000 `saf`) for transactions.
- **Home Directory**: Node configuration and data stored in `$HOME_NODE` (you’ll set this to a custom path, e.g., `$HOME/.safrochain`).

> **Note**: This guide assumes you’re joining an existing testnet and don’t need to create initial accounts or modify the genesis file. All commands are executed manually for clarity and control.

## 🚀 Setup Steps

Follow these steps to set up your Safrochain testnet node and become a validator. Each step includes **commands**, **explanations**, **verification checks**, and **troubleshooting tips** to ensure success.

### Step 1: Install Dependencies

**What it does**: Installs Go 1.23, `git`, `make`, and `jq` (for JSON parsing), and sets up the Go environment required to build and run the Safrochain node.

**Prerequisites**: Administrative privileges (`sudo` on Linux/macOS, admin terminal in WSL2).

**Code**:

#### Linux (Ubuntu)
```bash
sudo apt update
sudo apt install -y git make jq
# Download and install Go 1.23
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
rm go1.23.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
# Set up Go environment
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH
# Verify Go version
if go version | grep -q "go1.23"; then
    echo "Go 1.23 installed successfully."
else
    echo "Error: Go 1.23 not installed. Check installation steps."
    exit 1
fi
```

#### macOS
```bash
# Install git, make, and jq
xcode-select --install || true
brew install jq || true
# Download and install Go 1.23
curl -LO https://go.dev/dl/go1.23.0.darwin-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.0.darwin-amd64.tar.gz
rm go1.23.0.darwin-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
# Set up Go environment
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH
# Make PATH changes permanent
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> $HOME/.zshrc
source $HOME/.zshrc
# Verify Go version
if go version | grep -q "go1.23"; then
    echo "Go 1.23 installed successfully."
else
    echo "Error: Go 1.23 not installed. Check installation steps or run 'sudo rm -rf /usr/local/go' and retry."
    exit 1
fi
```

#### Windows (via WSL2)
```bash
# Install WSL2 and Ubuntu if not already set up
# In a Windows PowerShell (run as Administrator):
# wsl --install
# After Ubuntu setup, run in Ubuntu WSL2 terminal:
sudo apt update
sudo apt install -y git make jq
# Download and install Go 1.23
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
rm go1.23.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
# Set up Go environment
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
mkdir -p $GOPATH
# Verify Go version
if go version | grep -q "go1.23"; then
    echo "Go 1.23 installed successfully."
else
    echo "Error: Go 1.23 not installed. Check installation steps."
    exit 1
fi
```

**Notes**:
- **Linux**: Installs Go 1.23 via official tarball. For other architectures (e.g., ARM), visit `https://go.dev/dl/`.
- **macOS**: Requires Xcode Command Line Tools and Homebrew (`brew install homebrew`). If using Bash, replace `.zshrc` with `.bashrc`.
- **Windows**: Requires WSL2 with Ubuntu. Run `wsl --install` in PowerShell if not set up.
- The `GOPATH` is set to `$HOME/go` for building the Safrochain binary. `jq` is used for JSON validation in later steps.
- **Verification**: Run `go version` to confirm `go1.23.x`. If incorrect, remove existing Go (`sudo rm -rf /usr/local/go`) and retry.

**Verification**:
- Check Go version:
  ```bash
  go version
  ```
  **Expected Output**: `go version go1.23.0 ...`
- Check `jq`:
  ```bash
  jq --version
  ```
  **Expected Output**: `jq-1.6` or similar.

**Troubleshooting**:
- **Wrong Go version**: Run `sudo rm -rf /usr/local/go` and repeat installation.
- **Command not found**: Ensure `sudo` privileges and internet connectivity (`ping google.com`).

---

### Step 2: Clone Repository, Build Binary, and Export Binary Path

**What it does**: Clones the Safrochain node repository, builds the `safrochaind` binary, and adds `~/go/bin` to your PATH for accessibility.

**Prerequisites**: `git` and Go 1.23 installed (Step 1).

**Code**:

#### Linux (Ubuntu) / Windows (WSL2)
```bash
git clone https://github.com/Safrochain-Org/safrochain-node.git
cd safrochain-node
make install
cd ..
# Add ~/go/bin to PATH
export PATH=$PATH:$HOME/go/bin
echo 'export PATH=$PATH:$HOME/go/bin' >> $HOME/.bashrc
source $HOME/.bashrc
# Verify safrochaind
if command -v safrochaind &> /dev/null; then
    echo "safrochaind is accessible."
else
    echo "Error: safrochaind not found. Check if ~/go/bin/safrochaind exists."
    exit 1
fi
```

#### macOS
```bash
git clone https://github.com/Safrochain-Org/safrochain-node.git
cd safrochain-node
make install
cd ..
# Add ~/go/bin to PATH
export PATH=$PATH:$HOME/go/bin
echo 'export PATH=$PATH:$HOME/go/bin' >> $HOME/.zshrc
source $HOME/.zshrc
# Verify safrochaind
if command -v safrochaind &> /dev/null; then
    echo "safrochaind is accessible."
else
    echo "Error: safrochaind not found. Check if ~/go/bin/safrochaind exists."
    exit 1
fi
```

**Notes**:
- The `safrochaind` binary is installed in `~/go/bin`.
- **Linux/WSL2**: PATH changes are saved to `.bashrc`.
- **macOS**: Uses `.zshrc`. For Bash, replace with `.bashrc`.
- For other shells (e.g., Fish), add `export PATH=$PATH:$HOME/go/bin` to the shell’s config (e.g., `~/.config/fish/config.fish`).
- **Verification**: Run `safrochaind version` to confirm the binary is accessible.

**Verification**:
- Check binary:
  ```bash
  ls ~/go/bin/safrochaind
  ```
  **Expected Output**: `/home/<user>/go/bin/safrochaind`
- Check version:
  ```bash
  safrochaind version
  ```
  **Expected Output**: A version number (e.g., `v1.0.0`).

**Troubleshooting**:
- **safrochaind not found**: Verify `~/go/bin` is in PATH (`echo $PATH`) or re-run `source $HOME/.bashrc` (Linux) or `source $HOME/.zshrc` (macOS).
- **make install fails**: Check for errors in `safrochain-node/app.go` or run `go mod tidy` in the repository.

---

### Step 3: Initialize Node

**What it does**: Initializes the node with a unique moniker and creates the configuration directory `$HOME_NODE`.

**Prerequisites**: `safrochaind` binary built (Step 2).

**Code**:
```bash
# Set HOME_NODE (e.g., $HOME/.safrochain)
export HOME_NODE=$HOME/.safrochain
# Make HOME_NODE persistent
echo 'export HOME_NODE=$HOME/.safrochain' >> $HOME/.bashrc
source $HOME/.bashrc
# Initialize node
echo "Enter a moniker for your node (e.g., my-node):"
read MONIKER
safrochaind init "$MONIKER" --chain-id safrochain-testnet --home $HOME_NODE
```

**Notes**:
- The **moniker** is your node’s public name (e.g., `my-node`).
- The `--chain-id safrochain-testnet` matches the testnet’s identifier.
- Creates `$HOME_NODE/config/` with default files (`app.toml`, `config.toml`, `client.toml`, `genesis.json`, etc.).
- `$HOME_NODE` is set to `$HOME/.safrochain` and made persistent for future sessions.

**Verification**:
- Check directory:
  ```bash
  ls $HOME_NODE/config
  ```
  **Expected Output**: `app.toml`, `client.toml`, `config.toml`, `genesis.json`, `node_key.json`, `priv_validator_key.json`
- Verify `$HOME_NODE`:
  ```bash
  echo $HOME_NODE
  ```
  **Expected Output**: `/home/<user>/.safrochain`

**Troubleshooting**:
- **Directory not created**: Ensure `safrochaind` is installed (`which safrochaind`) and you have write permissions (`ls -ld $HOME`).
- **Error during init**: Check disk space (`df -h`) or re-run the command.

---

### Step 4: Configure Genesis File

**What it does**: Downloads the official testnet `genesis.json` from `https://safrochain.com/genesis/testnet/genesis.json` and places it in `$HOME_NODE/config/`.

**Prerequisites**: Internet access and `$HOME_NODE/config/` created (Step 3).

**Code**:
```bash
curl -L -o $HOME_NODE/config/genesis.json https://safrochain.com/genesis/testnet/genesis.json
if [ -f "$HOME_NODE/config/genesis.json" ]; then
    echo "genesis.json downloaded successfully."
else
    echo "Error: Failed to download genesis.json. Check the URL, internet connection, or ensure $HOME_NODE/config/ exists."
    echo "Contact Safrochain’s community (https://github.com/Safrochain-Org, Discord, or Telegram) for the correct genesis file."
    exit 1
fi
```

**Notes**:
- Uses `curl -L` to follow redirects and download the genesis file.
- Overwrites any existing `genesis.json` in `$HOME_NODE/config/`.
- This step avoids custom account creation or address substitution, addressing your previous `substitute_addresses.sh` issue by using the official testnet genesis file.
- To inspect the file:
  ```bash
  cat $HOME_NODE/config/genesis.json
  ```

**Verification**:
- Check file:
  ```bash
  ls $HOME_NODE/config/genesis.json
  ```
  **Expected Output**: `$HOME_NODE/config/genesis.json`
- Validate JSON:
  ```bash
  jq . $HOME_NODE/config/genesis.json
  ```
  **Expected Output**: JSON content without errors.

**Troubleshooting**:
- **Download fails**: Verify the URL in a browser or test internet (`ping google.com`). Try the alternative URL: `https://raw.githubusercontent.com/Safrochain-Org/genesis/refs/heads/main/genesis-testnet.json`.
- **Invalid JSON**: If `jq .` fails, redownload or contact the Safrochain community for the correct file.

---

### Step 5: Configure Node Settings

**What it does**: Sets up `app.toml`, `config.toml`, and `client.toml` with testnet-specific settings, including gas prices, ports, and the main node as a seed peer.

**Prerequisites**: Node initialized (Step 3).

**Code**:
```bash
# Ensure HOME_NODE is set
export HOME_NODE=$HOME/.safrochain
# Create app.toml
cat > "$HOME_NODE/config/app.toml" <<EOL
minimum-gas-prices = "0.001saf"

[api]
enable = true
swagger = true
address = "tcp://0.0.0.0:1317"

[grpc]
enable = true
address = "0.0.0.0:9090"

[grpc-web]
enable = true

[mempool]
max-txs = 5000

[telemetry]
enabled = false

[streaming.abci]
keys = []
plugin = ""
stop-node-on-err = true
EOL
# Create config.toml
echo "Enter your node's external IP (e.g., 192.168.1.100, or press Enter for local):"
read EXTERNAL_IP
if [ -z "$EXTERNAL_IP" ]; then
    EXTERNAL_IP="127.0.0.1"
fi
cat > "$HOME_NODE/config/config.toml" <<EOL
proxy_app = "tcp://127.0.0.1:26658"
moniker = "$MONIKER"
db_backend = "goleveldb"
db_dir = "data"
log_level = "info"
log_format = "plain"
genesis_file = "config/genesis.json"
priv_validator_key_file = "config/priv_validator_key.json"
priv_validator_state_file = "data/priv_validator_state.json"
node_key_file = "config/node_key.json"
abci = "socket"
filter_peers = false

[rpc]
laddr = "tcp://0.0.0.0:26657"
cors_allowed_origins = []
cors_allowed_methods = ["HEAD", "GET", "POST"]
cors_allowed_headers = ["Origin", "Accept", "Content-Type", "X-Requested-With", "X-Server-Time"]
tls_cert_file = ""
tls_key_file = ""
pprof_laddr = "localhost:6060"

[p2p]
laddr = "tcp://0.0.0.0:26656"
external_address = "$EXTERNAL_IP:26656"
seeds = "d172b3424a96ebbb806cd19e09d6976db9bb68ea@88.99.211.113:26656"
persistent_peers = ""
pex = true

[mempool]
type = "flood"
broadcast = true

[statesync]
enable = false

[blocksync]
version = "v0"

[consensus]
timeout_propose = "3s"
timeout_propose_delta = "500ms"
timeout_prevote = "1s"
timeout_prevote_delta = "500ms"
timeout_precommit = "1s"
timeout_precommit_delta = "500ms"
timeout_commit = "5s"
create_empty_blocks = false

[storage]
discard_abci_responses = false

[tx_index]
indexer = "kv"

[instrumentation]
prometheus = true
prometheus_listen_addr = ":26660"
max_open_connections = 3
namespace = "cometbft"
EOL
# Create client.toml
cat > "$HOME_NODE/config/client.toml" <<EOL
chain-id = "safrochain-testnet"
keyring-backend = "os"
output = "json"
node = "tcp://localhost:26657"
EOL
```

**Notes**:
- **Gas prices**: Set to `0.001saf` to match your denomination. If the testnet uses `tSaf` (per the original guide), replace with `0.001tSaf`.
- **Seeds**: Configures the main node `d172b3424a96ebbb806cd19e09d6976db9bb68ea@88.99.211.113:26656` for syncing.
- **External IP**: Defaults to `127.0.0.1` for local operation. For public nodes, enter your server’s IP.
- Enables API (`1317`), gRPC (`9090`), and P2P (`26656`) ports.
- The `$MONIKER` variable uses the value from Step 3.

**Verification**:
- Check files:
  ```bash
  ls $HOME_NODE/config
  ```
  **Expected Output**: `app.toml`, `client.toml`, `config.toml`, `genesis.json`, etc.
- Verify seed peer:
  ```bash
  grep seeds $HOME_NODE/config/config.toml
  ```
  **Expected Output**: `seeds = "d172b3424a96ebbb806cd19e09d6976db9bb68ea@88.99.211.113:26656"`

**Troubleshooting**:
- **Files not created**: Ensure `$HOME_NODE` is set (`echo $HOME_NODE`) and you have write permissions.
- **Incorrect moniker**: Re-run the `config.toml` creation with the correct `$MONIKER`.

---

### Step 6: Open Required Ports

**What it does**: Configures the firewall to allow ports `26656` (P2P), `26657` (RPC), `1317` (API), and `9090` (gRPC) for node communication.

**Prerequisites**: Firewall tools installed (`ufw` for Linux/WSL2, macOS firewall, Windows Firewall).

**Code**:

#### Linux (Ubuntu) / Windows (WSL2)
```bash
sudo ufw allow 26656
sudo ufw allow 26657
sudo ufw allow 1317
sudo ufw allow 9090
sudo ufw deny 26658
sudo ufw deny 6060
sudo ufw enable
```

#### macOS
```bash
# Open ports via terminal (requires sudo)
sudo /sbin/pfctl -f /etc/pf.conf
echo "pass in proto tcp from any to any port {26656, 26657, 1317, 9090}" | sudo pfctl -f -
sudo pfctl -E
# Alternatively, use System Preferences > Security & Privacy > Firewall > Firewall Options
# Add safrochaind and allow ports 26656, 26657, 1317, 9090
```

#### Windows (via WSL2)
```bash
# Run in Ubuntu WSL2 terminal
sudo ufw allow 26656,26657,1317,9090
sudo ufw deny 26658,6060
sudo ufw enable
# For host Windows Firewall, run in PowerShell (as Administrator):
New-NetFirewallRule -DisplayName "Safrochain" -Direction Inbound -Protocol TCP -LocalPort 26656,26657,1317,9090 -Action Allow
```

**Notes**:
- **Linux/WSL2**: `ufw` simplifies firewall management. Verify with `sudo ufw status`.
- **macOS**: `pfctl` changes are temporary; edit `/etc/pf.conf` or use the GUI for persistence.
- **Windows**: Apply rules in both WSL2 and Windows Firewall for external access.
- If using a cloud provider, open these ports in your security group.

**Verification**:
- Check firewall status (Linux/WSL2):
  ```bash
  sudo ufw status
  ```
  **Expected Output**: Shows `26656`, `26657`, `1317`, `9090` allowed, `26658`, `6060` denied.

**Troubleshooting**:
- **Ports not open**: Re-run `ufw` commands or check for conflicting rules (`sudo ufw status`).
- **Access denied**: Ensure cloud provider security groups allow these ports.

---

### Step 7: Start the Node

**What it does**: Starts the node, connects to the main node for syncing, and logs output for monitoring.

**Prerequisites**: Configuration files and genesis file set up (Steps 4–5).

**Code**:
```bash
# Reset node state
safrochaind tendermint unsafe-reset-all --home $HOME_NODE
# Start the node
safrochaind start --home $HOME_NODE > $HOME_NODE/safrochaind.log 2>&1 &
sleep 5
if pgrep safrochaind > /dev/null; then
    echo "Node is running. Logs are in $HOME_NODE/safrochaind.log."
else
    echo "Error: Node failed to start. Check $HOME_NODE/safrochaind.log for details."
    echo "Possible issues and fixes:"
    echo "1. Validator set empty: Verify Step 4 (genesis.json download)."
    echo "2. Redownload genesis.json: curl -L -o $HOME_NODE/config/genesis.json https://safrochain.com/genesis/testnet/genesis.json"
    echo "3. Reset state: safrochaind tendermint unsafe-reset-all --home $HOME_NODE"
    echo "4. Contact Safrochain’s community for an updated genesis file."
    echo "5. OE hash mismatch: Disable OE by adding 'optimistic_execution_enabled = false' under [consensus] in $HOME_NODE/config/config.toml"
    exit 1
fi
```

**Notes**:
- Resets node state to ensure a clean start.
- Logs are saved to `$HOME_NODE/safrochaind.log` for debugging.
- The node connects to the main node (`d172b3424a96ebbb806cd19e09d6976db9bb68ea@88.99.211.113:26656`) specified in `config.toml`.

**Verification**:
- Check logs:
  ```bash
  tail -f $HOME_NODE/safrochaind.log
  ```
  **Expected Output**: Shows node activity, block syncing, etc.
- Check status:
  ```bash
  curl http://localhost:26657/status
  ```
  **Expected Output**: JSON with node info, including `sync_info`.

**Troubleshooting**:
- **Node not starting**: Check logs (`tail -n 20 $HOME_NODE/safrochaind.log`).
- **Validator set empty**:
  - Verify genesis file (`jq . $HOME_NODE/config/genesis.json`).
  - Redownload: `curl -L -o $HOME_NODE/config/genesis.json https://safrochain.com/genesis/testnet/genesis.json`.
  - Reset state: `safrochaind tendermint unsafe-reset-all --home $HOME_NODE`.
- **OE hash mismatch**:
  - Disable OE:
    ```bash
    pkill safrochaind
    echo -e "\n[consensus]\noptimistic_execution_enabled = false" >> $HOME_NODE/config/config.toml
    safrochaind tendermint unsafe-reset-all --home $HOME_NODE
    safrochaind start --home $HOME_NODE
    ```
  - Contact the Safrochain community for guidance.

---

### Step 8: Create Validator Wallet

**What it does**: Creates a wallet for validator transactions and retrieves its address for faucet funding.

**Prerequisites**: Node running (Step 7).

**Code**:
```bash
echo "Enter a name for your validator wallet (e.g., validator):"
read WALLET_NAME
safrochaind keys add "$WALLET_NAME" --keyring-backend os --home $HOME_NODE
WALLET_ADDRESS=$(safrochaind keys show "$WALLET_NAME" -a --home $HOME_NODE)
echo "Your wallet address is: $WALLET_ADDRESS"
echo "Visit https://faucet.safrochain.com, paste your address, and request 2,500,000,000 saf (2,500 hela)."
```

**Notes**:
- The command outputs a **mnemonic phrase**. Save it securely offline (e.g., write it down, don’t store digitally).
- Visit `https://faucet.safrochain.com`, paste `$WALLET_ADDRESS`, and request 2,500,000,000 `saf` (2,500 `hela`).
- Addresses use the `addr_safro` prefix (not `taddr_safro`, as your configuration specifies `saf`).

**Verification**:
- Check wallet:
  ```bash
  safrochaind keys list --home $HOME_NODE --keyring-backend os
  ```
  **Expected Output**: Lists `$WALLET_NAME` with `addr_safro...`.
- Verify balance (after faucet):
  ```bash
  safrochaind query bank balances "$WALLET_ADDRESS" --home $HOME_NODE
  ```
  **Expected Output**: Shows `2500000000saf` if funded.

**Troubleshooting**:
- **Faucet fails**: If `https://faucet.safrochain.com` is down, join Safrochain’s community (GitHub, Discord, Telegram) for tokens.
- **No mnemonic**: Re-run the `keys add` command and save the output.
- **Wrong address prefix**: If the address uses `taddr_safro`, confirm testnet denomination with the community.

---

### Step 9: Create Validator

**What it does**: Submits a transaction to stake tokens and register your node as a validator using a JSON configuration file.

**Prerequisites**:
- Tokens in wallet (from Step 8).
- Node fully synced (`curl http://localhost:26657/status` shows `"catching_up": false`).

**Code**:
```bash
# Check sync status
curl http://localhost:26657/status | jq '.result.sync_info'

# Get validator public key
PUBKEY=$(safrochaind tendermint show-validator --home $HOME_NODE)
# Create validator.json
cat > $HOME/validator.json <<EOL
{
  "pubkey": $PUBKEY,
  "amount": "2500000000saf",
  "moniker": "$MONIKER",
  "identity": "",
  "website": "",
  "security": "",
  "details": "My Safrochain testnet validator",
  "commission-rate": "0.1",
  "commission-max-rate": "0.2",
  "commission-max-change-rate": "0.01",
  "min-self-delegation": "1"
}
EOL
# Verify JSON
jq . $HOME/validator.json
# Submit validator transaction
safrochaind tx staking create-validator $HOME/validator.json \
  --from "$WALLET_NAME" \
  --chain-id safrochain-testnet \
  --fees 5000saf \
  --keyring-backend os \
  --home $HOME_NODE
```

**Notes**:
- **Sync status**: The node must be fully synced before creating a validator. The `watch` command checks every 10 seconds.
- **Amount**: Stakes 2,500,000,000 `saf` (2,500 `hela`), matching the faucet amount. Adjust if you receive more/less tokens.
- **Fees**: Set to `5000saf` (0.005 `hela`). If the testnet uses `tSaf`, replace `saf` with `tSaf` in `amount` and `--fees`.
- The `$MONIKER` variable uses the value from Step 3.
- **Alternative denominations**: If the testnet uses `hela`, set `"amount": "2500hela"` and `--fees 5hela`.

**Verification**:
- Check validator:
  ```bash
  safrochaind query staking validators --home $HOME_NODE | grep -A 5 "$MONIKER"
  ```
  **Expected Output**: Shows your validator details.
- Verify balance:
  ```bash
  safrochaind query bank balances "$WALLET_ADDRESS" --home $HOME_NODE
  ```
  **Expected Output**: Reflects staked amount and remaining balance.

**Troubleshooting**:
- **Node not synced**: Wait for `catching_up: false` or check logs for sync issues.
- **Insufficient funds**: Verify balance (`safrochaind query bank balances`) and request more tokens from the faucet.
- **Invalid JSON**: Check `validator.json` (`jq . $HOME/validator.json`) and ensure `$PUBKEY` is populated.
- **Unknown flag**: Confirm `safrochaind` version supports the `create-validator` command.

---

## 🛠️ Post-Setup

### Monitor the Node
- **Check logs**:
  ```bash
  tail -f $HOME_NODE/safrochaind.log
  ```
  Monitor node activity and block syncing.
- **Check status**:
  ```bash
  curl http://localhost:26657/status
  ```
  View node sync status and details.
- **View node ID**:
  ```bash
  safrochaind tendermint show-node-id --home $HOME_NODE
  ```
  Useful for sharing with other nodes if needed.

### Stop the Node
- Find process ID:
  ```bash
  pgrep safrochaind
  ```
  **Expected Output**: A PID (e.g., `12345`).
- Kill process:
  ```bash
  kill <pid>
  ```
  Replace `<pid>` with the actual PID.

### Troubleshooting
- **Node not starting**:
  - Check logs: `tail -n 20 $HOME_NODE/safrochaind.log`.
  - Verify genesis file: `jq . $HOME_NODE/config/genesis.json`.
- **Empty validator set**:
  - Redownload genesis: `curl -L -o $HOME_NODE/config/genesis.json https://safrochain.com/genesis/testnet/genesis.json`.
  - Reset state: `safrochaind tendermint unsafe-reset-all --home $HOME_NODE`.
  - Contact Safrochain’s community for the correct genesis file.
- **OE hash mismatch**:
  - Ensure the main node seed is correct in `config.toml`.
  - Disable OE:
    ```bash
    pkill safrochaind
    echo -e "\n[consensus]\noptimistic_execution_enabled = false" >> $HOME_NODE/config/config.toml
    safrochaind tendermint unsafe-reset-all --home $HOME_NODE
    safrochaind start --home $HOME_NODE
    ```
- **create-validator errors**:
  - Verify `validator.json`: `jq . $HOME/validator.json`.
  - Check balance: `safrochaind query bank balances "$WALLET_ADDRESS" --home $HOME_NODE`.
  - Confirm sync: `curl http://localhost:26657/status`.
- **No tokens**: Request from `https://faucet.safrochain.com` or community channels.
- **safrochaind not found**: Add `~/go/bin` to PATH: `export PATH=$PATH:$HOME/go/bin`.
- **Firewall issues**: Verify ports: `sudo ufw status`.

## 🌐 Next Steps

- **Join the Community**: Connect with Safrochain’s Discord, Telegram, or forum for support (visit [Safrochain GitHub](https://github.com/Safrochain-Org)).
- **Run a Faucet**: Set up a testnet faucet if needed (contact the community for guidance).
- **Deploy a Block Explorer**: Use tools like [Big Dipper](https://github.com/forbole/big_dipper) for network visualization (ask for setup instructions).

## 📄 License

This guide is based on the Safrochain node setup process, licensed under the [MIT License](https://github.com/Safrochain-Org/safrochain-node/blob/main/LICENSE).

## 🙌 Acknowledgments

Safrochain is built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) and [CometBFT](https://github.com/cometbft/cometbft). Special thanks to the Safrochain community for their contributions!
