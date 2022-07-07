package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

func main() {
	// 设置网络类型
	address.CurrentNetwork = address.Mainnet
	// client := filecoin.New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")
	client := filecoin.NewClient("http://113.142.2.194:41234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.L2rJyd_Q77XW8MD02arfPbaR1GuBn8cFAfBAIhMHmmM")

	// addr1, _ := address.NewFromString("f1ntod647g54mv7pqbkniqnyov6k7thr2uxdec42i")
	// b, err := client.WalletBalance(context.Background(), addr1)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(filecoin.ToFil(b))
	// aa := []byte("7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22732b6138364378324842426e35686a5837647945484d7a456977613245305a4b61444b3173524b635a69493d227d")

	// 生产新的地址
	// 新地址有转入fil才激活，不然没法用
	// ki, addr, err := local.WalletNew(types.KTSecp256k1)
	// if err != nil {
	// 	panic(err)
	// }
	// kb, _ := json.Marshal(ki)
	// println(hex.EncodeToString(kb))
	// println(hex.EncodeToString(ki.PrivateKey))
	// println(addr.String())
	// return
	// 50a5e6234f5fdfc026bd889347409e11b6ef5b6034a7b0572d7e24ed1e9ba0e4
	// f1dynqskhlixt5eswpff3a72ksprqmeompv3pbesy

	// cd2980bc01d648d09a3698e275bc503229f45a4bd3d4050e5dd289286c969186
	// t1624zejtbagscwaswsojkye4dojegpqts7robb2a

	// 48157655123d812fff893ab99952ab52f6329654be40aa168efc1322bfc9eaa1
	// f17appra7cxzbrwiwuanlx3mu3qffl6qq75xfrsqy

	fromAddress, _ := address.NewFromString("f1g6ckkx7pbj2g6wnmns55ubz5jlexvq6eksw7apa")
	toAddress, _ := address.NewFromString("f1fpuzqvfdtmxpo6doexgxg6sapxabzkugmsy5yxq")
	// 转移0.001FIL到f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy
	msg := &types.Message{
		Version:    0,
		To:         toAddress,
		From:       fromAddress,
		Nonce:      19,
		Value:      filecoin.FromFil(decimal.NewFromFloat(0.001)),
		GasLimit:   2624630,
		GasFeeCap:  abi.NewTokenAmount(10000),
		GasPremium: abi.NewTokenAmount(10000),
		Method:     0,
		Params:     nil,
	}

	// client := filecoin.New("https://1v4969j3e9DkiSSXX4APjD2UKTX:742bb6ab7bd0e495d8f232657c2ae82b@filecoin.infura.io")

	// 最大手续费0.0001 FIL
	//maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))

	// 估算GasLimit
	//msg, err = client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	//if err != nil {
	//	panic(err)
	//}
	// pk, err := hex.DecodeString("7b2254797065223a22736563703235366b31222c22507269766174654b6579223a224a42474c4244365445336d73655a757a3468413836703157746e5531564266537479727435796a71566a4d3d227d")
	// if err != nil {
	// 	panic(err)
	// }
	// pk := []byte("cd2980bc01d648d09a3698e275bc503229f45a4bd3d4050e5dd289286c969186")

	aa, err := hex.DecodeString("7b2254797065223a22736563703235366b31222c22507269766174654b6579223a224a42474c4244365445336d73655a757a3468413836703157746e5531564266537479727435796a71566a4d3d227d")
	if err != nil {
		panic(err)
	}
	keyInfo := new(types.KeyInfo)
	json.Unmarshal(aa, &keyInfo)
	fmt.Println(keyInfo.Type)
	fmt.Println(keyInfo.PrivateKey)
	fmt.Println(hex.EncodeToString(keyInfo.PrivateKey))
	// return

	// 离线签名
	s, err := local.WalletSignMessage(types.KTSecp256k1, keyInfo.PrivateKey, msg)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(s.Signature.Data))
	// 47bcbb167fd9040bd02dba02789bc7bc0463c290db1be9b07065c12a64fb84dc546bef7aedfba789d0d7ce2c4532f8fa0d2dd998985ad3ec1a8b064c26e4625a01

	// 验证签名
	if err := local.WalletVerifyMessage(s); err != nil {
		panic(err)
	}

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		panic(err)
	}

	println(mid.String())
}
