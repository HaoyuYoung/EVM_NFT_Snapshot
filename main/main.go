package main

import EVM_NFT_Snapshot "ElfinSnapshot"

const rpc = "YOUR RPC PROVIDER"
const addr = "NFT CONTRACT ADDRESS"

func main() {
	EVM_NFT_Snapshot.GetAllID(rpc, addr)
	ownList := EVM_NFT_Snapshot.GetAllOwners(rpc, addr, "./allID.csv")
	EVM_NFT_Snapshot.GetBalanceAndLastAction(rpc, addr, ownList)
	EVM_NFT_Snapshot.Check("./result.csv")
}
