// Copyright 2014 The go-fiscobcos Authors
// This file is part of the go-fiscobcos library.
//
// The go-fiscobcos library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-fiscobcos library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-fiscobcos library. If not, see <http://www.gnu.org/licenses/>.

// Package types contains data types related to FiscoBcos consensus.
package types

import (
	"encoding/binary"
	"math/big"
	"reflect"

	"github.com/chislab/go-fiscobcos/common"
	"github.com/chislab/go-fiscobcos/common/hexutil"
	"github.com/chislab/go-fiscobcos/rlp"
	"golang.org/x/crypto/sha3"
)

var (
	EmptyRootHash  = DeriveSha(Transactions{})
	EmptyUncleHash = rlpHash([]*Header(nil))
)

// A BlockNonce is a 64-bit hash which proves (combined with the
// mix-hash) that a sufficient amount of computation has been carried
// out on a block.
type BlockNonce [8]byte

// EncodeNonce converts the given integer to a block nonce.
func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// Uint64 returns the integer value of a block nonce.
func (n BlockNonce) Uint64() uint64 {
	return binary.BigEndian.Uint64(n[:])
}

// MarshalText encodes n as a hex string with 0x prefix.
func (n BlockNonce) MarshalText() ([]byte, error) {
	return hexutil.Bytes(n[:]).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *BlockNonce) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BlockNonce", input, n[:])
}

//go:generate gencodec -type Header -field-override headerMarshaling -out gen_header_json.go

// Header represents a block header in the FiscoBcos blockchain.
type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"            gencodec:"required"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	RandomId    BlockNonce     `json:"randomid"`
}

// field type overrides for gencodec
type headerMarshaling struct {
	Difficulty *hexutil.Big
	Number     *hexutil.Big
	GasLimit   hexutil.Uint64
	GasUsed    hexutil.Uint64
	Time       hexutil.Uint64
	Extra      hexutil.Bytes
	Hash       common.Hash `json:"hash"` // adds call to Hash() in MarshalJSON
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}

var headerSize = common.StorageSize(reflect.TypeOf(Header{}).Size())

// Size returns the approximate memory used by all internal contents. It is used
// to approximate and limit the memory consumption of various caches.
func (h *Header) Size() common.StorageSize {
	return headerSize + common.StorageSize(len(h.Extra)+(h.Difficulty.BitLen()+h.Number.BitLen())/8)
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Body is a simple (mutable, non-safe) data container for storing and moving
// a block's data contents (transactions and uncles) together.
type Body struct {
	Transactions []*Transaction
	Uncles       []*Header
}

type ClientVersion struct {
	Build_Time         string `json:"Build Time"`
	Build_Type         string `json:"Build Type"`
	Chain_Id           string `json:"Chain Id"`
	FISCO_BCOS_Version string `json:"FISCO-BCOS Version"`
	Git_Branch         string `json:"Git Branch"`
	Git_Commit_Hash    string `json:"Git Commit Hash"`
	Supported_Version  string `json:"Supported Version"`
}

type SyncStatus struct {
	BlockNumber        int    `json:"blockNumber"`
	GenesisHash        string `json:"genesisHash"`
	IsSyncing          bool   `json:"isSyncing"`
	KnownHighestNumber int    `json:"knownHighestNumber"`
	KnownLatestHash    string `json:"knownLatestHash"`
	LatestHash         string `json:"latestHash"`
	NodeID             string `json:"nodeId"`
	Peers              []struct {
		BlockNumber int    `json:"blockNumber"`
		GenesisHash string `json:"genesisHash"`
		LatestHash  string `json:"latestHash"`
		NodeID      string `json:"nodeId"`
	} `json:"peers"`
	ProtocolID int    `json:"protocolId"`
	TxPoolSize string `json:"txPoolSize"`
}

type Block struct {
	DbHash       common.Hash        `json:"dbHash"`
	ExtraData    []interface{} `json:"extraData"`
	GasLimit     string        `json:"gasLimit"`
	GasUsed      string        `json:"gasUsed"`
	Hash         common.Hash   `json:"hash"`
	LogsBloom    string        `json:"logsBloom"`
	Number       string        `json:"number"`
	ParentHash   common.Hash        `json:"parentHash"`
	ReceiptsRoot string        `json:"receiptsRoot"`
	Sealer       string        `json:"sealer"`
	SealerList   []string      `json:"sealerList"`
	StateRoot    string        `json:"stateRoot"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Receipt `json:"transactions"`
	TransactionsRoot string `json:"transactionsRoot"`
}

type TotalTransactionCount struct {
	BlockNumber string `json:"blockNumber"`
	TxSum       string `json:"txSum"`
}

type TransactionByHash struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
}

