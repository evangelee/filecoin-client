package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/types"
)

// 简单的充值逻辑
func main() {
	// client := filecoin.New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")
	client := filecoin.NewClient("http://113.142.2.194:41234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.L2rJyd_Q77XW8MD02arfPbaR1GuBn8cFAfBAIhMHmmM")

	addr, err := client.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		// t.Error(err)
		panic(err)
	}
	// t.Log(addr)
	fmt.Println(addr)

	ki, err := client.WalletExport(context.Background(), addr)
	if err != nil {
		// t.Error(err)
		panic(err)
	}

	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	// t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
	fmt.Println(ki.Type, hex.EncodeToString(ki.PrivateKey))

	// job := &RechargeFilJob{}

	// // 处理区块652243
	// if err := job.mapHeight(client, 1956904); err != nil {
	// 	panic(err)
	// }

	// todo
}

// 充值处理任务
type RechargeFilJob struct {
}

func (job *RechargeFilJob) mapHeight(c *filecoin.Client, height int64) error {
	ts, err := c.ChainGetTipSetByHeight(context.Background(), height, nil)
	if err != nil {
		return err
	}
	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			return err
		}

		// BlsMessages
		for _, msg := range bm.BlsMessages {
			err := job.handleMessage(msg.Cid(), msg)
			if err != nil {
				return err
			}
		}

		// SecpkMessages
		for _, msg := range bm.SecpkMessages {
			err := job.handleSecpkMessage(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (job *RechargeFilJob) handleSecpkMessage(msg *types.SignedMessage) error {
	return job.handleMessage(msg.Cid(), msg.Message)
}

func (job *RechargeFilJob) handleMessage(msgCid cid.Cid, msg *types.Message) error {
	// 发送交易并且值大于0
	if msg.Method != 0 || msg.Value.IsZero() {
		return nil
	}

	value := filecoin.ToFil(msg.Value)

	// 有可能重复
	// 请根据msgCid自行去重复
	println(msgCid.String(), value.String())

	// todo

	return nil
}

// 结果如下：
/*
bafy2bzacedv5q3nam6icb4qwuftxw3nazn33iu3vd6zvwkcr6ljgfbkjmow62 5.002
bafy2bzacecb52i5423svvxgyfz7dvfvl4wssi4vraiyilkckomytplvbtxbgo 1.676583814823873898
bafy2bzaceabqawv4iwnjn4xusseqgxenprq36w6xxyveryflr3e7dpfmroprw 9.999853889870997446
bafy2bzacedvdkwm7js4qjgzyiqdzcoy3kgf7fdooowzbchj4rnzlw6wuk4bp6 3.978744820695744384
bafy2bzaced7tmalo62e74xsveb2ahmb6metcb6kwo5jq35neic3xlzwphm7am 3100
bafy2bzacecdykvlsswqlrkgohqsvr4cwv7rph4hzabbfg2p4iqixcm3iv62ly 220
bafy2bzaceakfahxbzfvson5ylx7nsa2velqfy37xmdz5bycyxds4rsyefittu 220
bafy2bzacec3rlx3yoezxzpnzyzehu5kenyjoso2gogkkpiw5wwbjv7hcjmy5y 9.99
bafy2bzacebrnc5tactfdeddxmpiyy5wppfc4gyc45zscwymn4r2pm4uwmasx4 17.90844787
*/
