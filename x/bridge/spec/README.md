# `bridge`

## Abstract

The `x/bridge` module handles cross-chain transfers of ERC-20 tokens between
external blockchains (Ethereum) and Fury's EVM, and an internal bridge to
convert between ERC20 tokens on Fury EVM and Fury Cosmos `sdk.Coin`.

This allows the following:

* Transfer between Ethereum ERC20 <-> Fury EVM ERC20. E.g. transferring WETH or
  USDC to Fury EVM.
* Conversion between Fury EVM ERC20 <-> Fury Cosmos `sdk.Coin`. E.g. converting
  ERC20 WETH on Fury EVM to Cosmos `sdk.Coin` to be able to use WETH in native
  Fury DeFi modules such as Mint, Lend, Swap, etc.

There is no direct path from Ethereum ERC20 to Fury Cosmos `sdk.Coin`. To get an
ERC20 token on Ethereum into the Fury Cosmos ecosystem, tokens must be bridged
to the Fury EVM first, then converted to an `sdk.Coin`.

<p align="center">
  <img src="./images/high-level-overview.drawio.png">
</p>

## Contents

1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
4. **[Events](04_events.md)**
5. **[Params](05_params.md)**
