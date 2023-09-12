package EVM_NFT_Snapshot

import (
	"ElfinSnapshot/abi"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"io"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
)

func GetAllID(rpcURL, contractAddr string) {
	//Typically, TokenID is from 0 to total supply, but some NFT has a discrete ID, so we can use this function to get all ID ourselves
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress(contractAddr)

	NFT, err := abi.NewNFT(address, rpcClient)
	if err != nil {
		fmt.Println(err)
	}

	total, err := NFT.TotalSupply(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	t := total.Int64()
	fmt.Println(t)
	var ids [][]string
	var wg sync.WaitGroup
	wg.Add(1)
	for i := int64(0); i <= t; i++ {
		time.Sleep(50 * time.Millisecond)
		i := i
		go func() {
			of, err := NFT.TokenByIndex(&bind.CallOpts{
				Pending:     false,
				From:        common.Address{},
				BlockNumber: nil,
				Context:     nil,
			}, big.NewInt(i))
			if err != nil {
				fmt.Println(i, err)
				return
			} else {
				d := [][]string{
					{of.String(), strconv.FormatInt(i, 10)},
				}
				fmt.Println(i, d)
				ids = append(ids, d...)
			}

		}()
	}
	f, err := os.Create("allID.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(ids)
	w.Flush()
	fmt.Println("Total", len(ids))

}

func GetAllOwners(rpcURL, contractAddr, filePath string) []common.Address {
	//To get all owners And eliminate the contract address.
	f, err := os.Open(filePath) // load the file
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	rpcClient, err := ethclient.Dial(rpcURL)
	address := common.HexToAddress(contractAddr)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	var list []common.Address
	var addListStr [][]string
	NFT, err := abi.NewNFT(address, rpcClient)
	if err != nil {
		fmt.Println(err)
	}

	count := int64(0)
	var wg sync.WaitGroup
	wg.Add(1)
	for _, i := range csvData {
		time.Sleep(50 * time.Millisecond)
		i, _ := strconv.Atoi(i[0])
		go func() {
			of, err := NFT.OwnerOf(&bind.CallOpts{
				Pending:     false,
				From:        common.Address{},
				BlockNumber: nil,
				Context:     nil,
			}, big.NewInt(int64(i)))
			if err != nil {
				fmt.Println(i, err)
				return
			}
			at, err := rpcClient.CodeAt(ctx, of, nil)

			if err != nil {
				fmt.Println(i, err)
				return
			}
			if len(at) == 0 {
				fmt.Println(of)
				list = append(list, of)
				d := [][]string{
					{strconv.FormatInt(int64(i), 10), of.Hex()},
				}
				addListStr = append(addListStr, d...)
				count += 1
				fmt.Println(i, "commonUser")
				fmt.Println("total", count)
			} else {
				count += 1

				fmt.Println(i, "Contracts!")
				fmt.Println("total", count)

			}

		}()

	}

	ff, err := os.Create("allTokenAndOwner.csv")
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	w := csv.NewWriter(ff)
	w.WriteAll(addListStr)
	w.Flush()

	f, err = os.Open("./allTokenAndOwner.csv") // 读取文件
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	reader = csv.NewReader(f)
	csvData, err = reader.ReadAll() // 读取全部数据
	if err != nil {
		fmt.Println(err)
	}

	var result []common.Address
	var adds []string
	temp := map[string]struct{}{}
	for _, line := range csvData {
		add := line[1]
		fmt.Println(1)

		if _, ok := temp[add]; !ok {
			temp[add] = struct{}{}
			adds = append(adds, add)
			result = append(result, common.HexToAddress(add))
		}

	}
	newf, err := os.Create("allHolders.csv")
	if err != nil {
		panic(err)
	}
	defer newf.Close()
	w = csv.NewWriter(newf) //创建一个新的写入文件流
	w.Write(adds)           //写入数据
	w.Flush()
	return result

}

func GetBalanceAndLastAction(rpcURL, contractAddr string, list []common.Address) {
	rpcClient, err := ethclient.Dial(rpcURL)
	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 30 * time.Second,
		Key:     "Your-API of ChainScan, e.g. BSCSCAN",
		BaseURL: "API URL",
		Verbose: false,
	})
	var start, end int
	//block height, start to end
	start = 1
	end = 99999999
	address := common.HexToAddress(contractAddr)
	if err != nil {
		panic(err)
	}

	NFT, err := abi.NewNFT(address, rpcClient)
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	var m [][]string
	fmt.Println(len(list))
	for i, item := range list {
		fmt.Println(i, item)
		item := item
		time.Sleep(220 * time.Millisecond)
		go func() {
			balance, err := NFT.BalanceOf(&bind.CallOpts{
				Pending:     false,
				From:        common.Address{},
				BlockNumber: nil,
				Context:     nil,
			}, item)
			if err != nil {
				fmt.Println(err)
				return
			}
			if balance != nil {
				address, err := client.NormalTxByAddress(item.Hex(), &start, &end, 1, 1, true)
				if err != nil {
					fmt.Println(err)
					d := [][]string{
						{item.Hex(), balance.String()},
					}
					m = append(m, d...)
				} else {
					d := [][]string{
						{item.Hex(), balance.String(), address[0].TimeStamp.Time().String()},
					}
					fmt.Println(i, d)
					m = append(m, d...)
				}

			} else {
				d := [][]string{
					{item.Hex(), "0"},
				}
				m = append(m, d...)

			}

		}()

	}
	f, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f) //创建一个新的写入文件流

	w.WriteAll(m) //写入数据
	w.Flush()

}

func Check(path string) {
	f, err := os.Create("checked.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)    //创建一个新的写入文件流
	ff, err := os.Open(path) // 读取文件
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	defer ff.Close()
	var m, n [][]string
	reader := csv.NewReader(ff)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		n = append(n, rec)
	}

	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 30 * time.Second,
		Key:     "Your-API of ChainScan, e.g. BSCSCAN",
		BaseURL: "API URL",
		Verbose: false,
	})
	var start, end int
	//block height, start to end
	start = 1
	end = 99999999
	for _, item := range n {

		if len(item) == 2 {
			address, err := client.NormalTxByAddress(item[0], &start, &end, 1, 1, true)
			if err != nil {
				fmt.Println(err)
				m = append(m, item)
				fmt.Println("Checked for missing, not filled", item)
			} else {
				m = append(m, append(item, address[0].TimeStamp.Time().String()))
				fmt.Println("Checked for missing, filled in", append(item, address[0].TimeStamp.Time().String()))
			}
			time.Sleep(200 * time.Millisecond)
		} else if len(item) == 3 {
			m = append(m, item)
			fmt.Println("Not missing, continue", item)
		}
	}
	w.WriteAll(m) //写入数据
	w.Flush()
}
