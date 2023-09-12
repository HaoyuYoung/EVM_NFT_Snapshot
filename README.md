# EVM_NFT_Snapshot
A NFT Snapshot tool for All EVM NFTs, using Golang

## About
It's a simple tool for Go users to run a snapshot for EVM NFTs, all you need is a reliable RPC provider.

### Functions
This simple tool can get all owners of a specific NFT collection, and record their last active time.
All results saved as CSV files.

## How to Use:
1. Copy the abi to abi/NFT.abi
2. run .\abigen.exe --abi NFT.abi --pkg abi --type NFT --out NFT.go
3. Replace all const to yours
4. Build and Run
