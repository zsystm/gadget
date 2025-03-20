# Gadget
> A simple yet powerful CLI tool for various cryptographic and address-related tasks, especially handy in Tendermint(CometBFT) & Ethereum ecosystems.

Gadget provides handy subcommands for working with base64/hex conversions, secp256k1/ed25519 keys, address formats (Bech32, Ethereum hex), and more.

## Features
- Base64 ↔ Hex: Easily convert between base64 and hexadecimal. 
- Ed25519 Key Tools:
  - Extract Tendermint/CometBFT validator addresses from Ed25519 public keys.
  - Generate public keys from Ed25519 private keys.
- Secp256k1 Key Tools:
  - Derive Ethereum accounts from private keys. 
  - Generate Tendermint-style validator (privval) keys. 
  - Convert private keys to public keys.
- Address Conversions:
  - Convert Bech32 to Ethereum hex, or vice versa. 
  - Change Bech32 prefixes.
- Ethereum Utilities:
  - Create new Ethereum accounts. 
  - Check balances via RPC.

## Install
### Option 1) Go Install

```bash
go install github.com/zsystm/gadget/cmd/gadget@latest
```

### Option2) Build from Source
```bash
git clone https://github.com/zsystm/gadget.git
cd gadget
make build
```

## Quick Examples
The following examples demonstrate how to use Gadget for various tasks.
<details> <summary><strong>1) Get Validator Hex Address from Ed25519 Public Key</strong></summary>

```bash
# 1. Check your validator's Ed25519 public key
cometbft show-validator
# Outputs something like:
{
  "type":"tendermint/PubKeyEd25519",
  "value":"xV2T7kMMXB94NOm22wIPrFyaFFGhiodEIliFAaGnODw="
}
```

**Result:**

```bash
address: 9C1950C518E7F2188B054417A9B33CB41B5935B7
```

</details> <details> <summary><strong>2) Generate Secp256k1 Privval Key</strong></summary>

```bash
# Generate a Tendermint-style private validator key from a secp256k1 private key
gadget secp256k1 privval 1afed3c4437316f73d28f69fd5e90ffc551a3862d08c34073e42f89d9dcc7149
```

**Result:**
```json
{
  "Key": {
    "address": "870B2D7410FA447EBB61E03D41F01B746F413137",
    "pub_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "A1Sh0t3iqUEvMEfxBYaLCmRoDfqvNzIw0UjmkYBivSZC"
    },
    "priv_key": {
      "type": "tendermint/PrivKeySecp256k1",
      "value": "FkyINadFe6Dw5TKNHXly4lnxqXnQ63aSwxxfauROJ6M="
    }
  },
  "LastSignState": {
    "height": "0",
    "round": 0,
    "step": 0
  }
}
```

</details> <details> <summary><strong>3) Get Account Info from Secp256k1 Private Key</strong></summary>

```bash
gadget secp256k1 acc story 164c8835a7457ba0f0e5328d1d7972e259f1a979d0eb7692c31c5f6ae44e27a3
```

**Result:**
```json
{
  "ethAddr": "0x881319354734eb9366A26a4f3e640BA55F1a2e0c",
  "accAddr": "story1su9j6aqslfz8awmpuq75ruqmw3h5zvfh7zkaax",
  "valAddr": "storyvaloper1su9j6aqslfz8awmpuq75ruqmw3h5zvfhsdzukd",
  "pubKey": "A1Sh0t3iqUEvMEfxBYaLCmRoDfqvNzIw0UjmkYBivSZC"
}
```

</details> <details> <summary><strong>4) Convert Base64 to Hexadecimal</strong></summary>

```bash
gadget b64-to-hex A1Sh0t3iqUEvMEfxBYaLCmRoDfqvNzIw0UjmkYBivSZC
```

**Result:**
```bash
hex: 0354a1d2dde2a9412f3047f105868b0a64680dfaaf373230d148e6918062bd2642
```

## Commands Overview
The following is a list of available commands in Gadget. Use `gadget help [command]` for detailed help on any command.
- Base64 ↔ Hex 
  - `gadget b64-to-hex`: Convert a base64 string to hexadecimal 
  - `gadget hex-to-b64`: Convert a hexadecimal string to base64

- Address Tools(`addr`)
  - `gadget addr bech-to-eth`: Convert a Bech32 address to a 20-byte Ethereum hexadecimal address
  - `gadget addr eth-to-bech`: Convert a 20-byte Ethereum hexadecimal address to Bech32 
  - `gadget addr change-bech-prefix`: Convert a Bech32 address to another prefix
- Ethereum Tools(`eth`)
  - `gadget eth addr`: Get Ethereum address from a private key
  - `gadget eth get-balance`: Get the balance of an Ethereum address using RPC
  - `gadget eth new-acc`: Generate a random Ethereum account
- Secp256k1 Tools(`secp256k1`)
  - `gadget secp2565k1 acc`: (hex without 0x) secp256k1 private key → account info.
  - `gadget secp256k1 privval`: (hex without 0x) secp256k1 private key → Tendermint/CometBFT privval key
  - `gadget secp256k1 pub`: (hex without 0x) secp256k1 private key → public key
- Ed25519 Tools (`ed25519`)
  - `gadget ed25519 addr-from-pubkey`: Ed25519 public key → Tendermint/CometBFT validator hex address
  - `gadget ed25519 pubkey-from-privkey`: Ed25519 private key → public key

## License

This project is licensed under GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
