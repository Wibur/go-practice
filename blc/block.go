package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	// 區塊高度
	Height int64
	// 上一個區塊hash
	PrevBlockHash []byte
	// 交易資料
	Data []byte
	// 時間戳
	TimeStamp int64
	// hash
	Hash []byte
}

func Test() {
	fmt.Println("123")
}

// 建立區塊
func CreateBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	block.SetHash()
	return block
}

func (block Block) SetHash() {
	// height []byte
	heightBytes := IntToHex(block.Height)
	fmt.Println(heightBytes)
	// timestamp => []byte
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	fmt.Println(timeString)
	timeBytes := []byte(timeString)
	fmt.Println(timeBytes)
	// 組合所需屬性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block})
	// 生成hash
	hash := sha256.Sum256(blockBytes)
	fmt.Println(hash)
}
