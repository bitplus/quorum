package pubsubconn

import (
	quorumpb "github.com/rumsystem/quorum/internal/pkg/pb"
)

type Chain interface {
	HandleTrx(trx *quorumpb.Trx) error
	HandleBlock(block *quorumpb.Block) error
}

type PubSubConn interface {
	JoinChannelAsExchange(cId string) error
	JoinChannel(cId string, chain Chain) error
	Publish(data []byte) error
}
