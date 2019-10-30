package eboost

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xfcd9b7dd
	TestnetMagic wire.BitcoinNet = 0xfbcdccd3
	RegtestMagic wire.BitcoinNet = 0xfabfb5da
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{92}
	MainNetParams.ScriptHashAddrID = []byte{5}
	MainNetParams.Bech32HRPSegwit = "ebst"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{33}
	TestNetParams.ScriptHashAddrID = []byte{55}
	TestNetParams.Bech32HRPSegwit = "tala"
}

// EboostParser handle
type EboostParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewEboostParser returns new EboostParser instance
func NewEboostParser(params *chaincfg.Params, c *btc.Configuration) *EboostParser {
	return &EboostParser{
	BitcoinParser: btc.NewBitcoinParser(params, c),
	baseparser:    &bchain.BaseParser{},
	}
}

// GetChainParams contains network parameters for the main Eboost network,
// and the test Eboost network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to eboost parameters
	// eboost has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *EboostParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *EboostParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
