## BPX Execution Client

Official Golang execution layer implementation of the BPX Blockchain.

BPX is a new generation Ethereum-compatible blockchain based upon an innovative consensus algorithm, Proof of Space and Time known from the Chia Network. Proof of Space is a cryptographic technique where farmers prove that they allocate unused hard disk space to the network. Proof of Time increases the overall security of the blockchain.

## Installation

### Windows
1. Download and run the installer EXE file.
2. Follow the wizard instructions.
### Debian / Ubuntu
1. Download DEB package
2. Install it using the following command:
```shell
dpkg -i bpx-execution-client-*.deb
```
### Other linux distributions
1. Download TAR.GZ binaries archive
2. Extract it using the following command:
```shell
tar zxvf bpx-execution-client-*.tar.gz -C /usr/local/bin
```

## Building the source

Building BPX Execution Client requires both a Go (version 1.19 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Usage

> :warning: **You need both [Beacon Client](https://github.com/bpx-chain/bpx-beacon-client) and Execution Client to synchronize with the network.**

### Windows
1. Find the "BPX Execution Client" group in the Start Menu.
2. Use the shortcut "BPX Execution Client" to run a mainnet full node, or "BPX Execution Client (Testnet)" to run a testnet full node
3. Use the shortcut "Attach" to attach to the mainnet node console, or "Attach (Testnet)" to attach to the testnet node console.
### Linux
1. To run mainnet full node use the following command:
```shell
bpx-geth --syncmode snap --http
```
2. To run testnet full node use the following command:
```shell
bpx-geth --testnet --syncmode snap --http
```
3. To attach to the mainnet node console use the following command:
```shell
bpx-geth attach
```
4. To attach to the testnet node console use the following command:
```shell
bpx-geth attach --datadir ~/.bpxchain/execution/testnet
```

The BPX Execution Client comes with several wrappers/executables found in the `cmd`
directory.

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| :--------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`bpx-geth`** | Our main CLI client. It is the entry point into the BPX network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the BPX network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `bpx-geth --help` for command line options. |
|   `clef`   | Stand-alone signing tool, which can be used as a backend signer for `bpx-geth`.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
|  `devp2p`  | Utilities to interact with nodes on the networking layer, without running a full blockchain.                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|  `abigen`  | Source code generator to convert smart contract definitions into easy-to-use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see [Native DApps](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) page for details.                                  |
| `bootnode` | Stripped down version of our execution client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                                                                                                               |
|   `evm`    | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                               |
| `rlpdump`  | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                |
