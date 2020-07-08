package dlkaddhtlog

import (
	"github.com/libp2p/go-libp2p-kad-dht/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("lkaddhtp2p")
}
