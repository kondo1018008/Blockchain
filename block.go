package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte // 前のブロックのハッシュ値
	Hash []byte
}

func (b *Block) SetHash(){
	// 実際のハッシュ値の計算はノードがマイニングする形で計算されるが、そのシステムはのちに実装
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // 時間を[]byteにキャストする。（メモリ確保注意）
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // 第一引数の要素を空のスライスで結合
	hash := sha256.Sum256(headers) // チェックサムを行う

	b.Hash = hash[:] // hashの全要素をb.Hashに代入
}

func NewBlock(data string, prevBlockHash []byte) *Block{// ブロックの生成を行う
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}