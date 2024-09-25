# Gadget
A simple utility tool for various cryptographic and address-related tasks.

## Building
To build the gadget tool, use:

```bash
go build -o gadget
```

## Usage Examples
install binary first: `go install github.com/zsystm/gadget@latest`

### Secp256k1 Private Validator Key
Generate a Secp256k1 private validator key:

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

- `b64-to-hex`  — Convert a base64 string to hexadecimal
- `bech-to-eth` — Convert a Bech32 address to a 20-byte Ethereum hexadecimal address
- `eth`         — Ethereum-related commands
- `eth-to-bech` — Convert a 20-byte Ethereum hexadecimal address to Bech32
- `hex-to-b64`  — Convert a hexadecimal string to base64
- `other-prefix`— Convert a Bech32 address to another prefix
- `secp256k1`   — Secp256k1-related commands

For detailed help on any command, use:

```bash
gadget help [command]
```
