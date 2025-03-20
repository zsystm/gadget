# Gadget
A simple utility tool for various cryptographic and address-related tasks.

## Building
To build the gadget tool, use:

```bash
make build
```

## Usage Examples
install binary first: `go install github.com/zsystm/gadget/cmd/gadget@latest`

### Get Validator Hex Address from Ed25519 Public Key
```bash
# 1. Get the validator public key
cometbft show-validator
{"type":"tendermint/PubKeyEd25519","value":"xV2T7kMMXB94NOm22wIPrFyaFFGhiodEIliFAaGnODw="}
# 2. Get the validator hex address
gadget ed25519 addr-from-pubkey xV2T7kMMXB94NOm22wIPrFyaFFGhiodEIliFAaGnODw=
```
**Result:**

```bash
address: 9C1950C518E7F2188B054417A9B33CB41B5935B7
```

### Secp256k1 Private Validator Key
Generate a Secp256k1 Privval Key from hexadecimal secp256k1 private key:

```bash
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

### Secp256k1 Account Information
Get account information from private key:

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

### Base64 to Hexadecimal Conversion
Convert a base64 string to a hexadecimal representation:

```bash
gadget b64-to-hex A1Sh0t3iqUEvMEfxBYaLCmRoDfqvNzIw0UjmkYBivSZC
```
**Result:**
```
hex: 0354a1d2dde2a9412f3047f105868b0a64680dfaaf373230d148e6918062bd2642
```

## Available Commands

- `gadget b64-to-hex`: Convert a base64 string to hexadecimal
- `gadget hex-to-b64`: Convert a hexadecimal string to base64
- `addr`: Tools for address conversion.
  - `gadget addr bech-to-eth`: Convert a Bech32 address to a 20-byte Ethereum hexadecimal address
  - `gadget addr eth-to-bech`: Convert a 20-byte Ethereum hexadecimal address to Bech32 
  - `gadget addr change-bech-prefix`: Convert a Bech32 address to another prefix
- `eth`: Ethereum-related commands
  - `gadget eth addr`: Get Ethereum address from a private key
  - `gadget eth get-balance`: Get the balance of an Ethereum address using RPC
  - `gadget eth new-acc`: Generate a random Ethereum account
- `secp256k1`: Secp256k1-related commands
  - `gadget secp2565k1 acc`: Convert a secp256k1 private key(hexadecimal wihtout 0x) to an account info
  - `gadget secp256k1 privval`: Generate a privval key from a secp256k1 private key(hexadecimal without 0x)
  - `gadget secp256k1 pub`: Generate a public key from a secp256k1 private key(hexadecimal without 0x)
- `ed25519`: Ed25519-related commands
  - `gadget ed25519 addr-from-pubkey`: Get the hex address from an Ed25519 public key
  - `gadget ed25519 pubkey-from-privkey`: Generate a public key from an Ed25519 private key

For detailed help on any command, use:

```bash
gadget help [command]
```
