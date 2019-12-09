package rpc

import (
	"github.com/chislab/go-fiscobcos/common/hexutil"
	"github.com/pborman/uuid"
	"strings"
)

type ChannelPack int

const (
	TYPE_RPC ChannelPack = 0x12
	TYPE_HEATBEAT ChannelPack = 0x13
	TYPE_AMOP_REQ ChannelPack = 0x30
	TYPE_AMOP_RESP ChannelPack = 0x31
	TYPE_TOPIC_REPORT ChannelPack = 0x32
	TYPE_TOPIC_MULTICAST ChannelPack = 0x35
	TYPE_TX_COMMITTED ChannelPack = 0x1000
	TYPE_TX_BLOCKNUM ChannelPack = 0x1001
)

func GenMsgSeq() ([]byte, error) {
	uid := uuid.New()
	splited := strings.Split(uid, "-")
	uid = "0X"
	for _, v := range splited {
		uid += v
	}
	return hexutil.Decode(strings.ToUpper(uid))
}

func GenZeroSeq() ([]byte, error) {
	return hexutil.Decode("0x00000000000000000000000000000000")
}

func SockReq() {
	//tls.LoadX509KeyPair()
}