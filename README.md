# Little Ethereum transaction info getter 

## How to build

- glide install
- go build

## Example

`./ether-get-tx-info --server=https://mainnet.infura.io 
--txhash=0xa237249a37164af117481ab85bb63088a269e162997068ff431d66294693b703`

```
(*types.Transaction)(0xc4200d6000)({
 data: (types.txdata) {
  AccountNonce: (uint64) 0,
  Price: (*big.Int)(0xc42031c160)(41000000000),
  GasLimit: (uint64) 37887,
  Recipient: (*common.Address)(0xc4204aa080)((len=20 cap=20) 0xe3881F8B2f1f8B3eE69090CA5e3a422ea4b75fc7),
  Amount: (*big.Int)(0xc42031c2a0)(0),
  Payload: ([]uint8) (len=68 cap=68) {
   00000000  a9 05 9c bb 00 00 00 00  00 00 00 00 00 00 00 00  |................|
   00000010  50 cb ac 09 55 e3 9d 39  6b df 3f 09 90 dd 4a 01  |P...U..9k.?...J.|
   00000020  d9 0d 2d c7 00 00 00 00  00 00 00 00 00 00 00 00  |..-.............|
   00000030  00 00 00 00 00 00 00 00  00 00 00 0b 0f 11 97 29  |...............)|
   00000040  63 b0 00 00                                       |c...|
  },
  V: (*big.Int)(0xc42031c2c0)(37),
  R: (*big.Int)(0xc42031c300)(39692809628277133672067410501376237568860675272681850344637427981536657501370),
  S: (*big.Int)(0xc42031c320)(23438084937191390759889755981969716394886023311604423297938951629806955216956),
  Hash: (*common.Hash)(0xc4204aa040)((len=32 cap=32) 0xa237249a37164af117481ab85bb63088a269e162997068ff431d66294693b703)
 },
 hash: (atomic.Value) {
  v: (interface {}) <nil>
 },
 size: (atomic.Value) {
  v: (interface {}) <nil>
 },
 from: (atomic.Value) {
  v: (types.sigCache) {
   signer: (*ethclient.senderFromServer)(0xc4203b4080)({
    addr: (common.Address) (len=20 cap=20) 0x0260Cc43809ad303228d3bCfaa1Fd808cb75fbAF,
    blockhash: (common.Hash) (len=32 cap=32) 0x6cf187bacef46c4a8f7cba21d368dc55aa803b7aa7b076f6098ff777b6f72798
   }),
   from: (common.Address) (len=20 cap=20) 0x0260Cc43809ad303228d3bCfaa1Fd808cb75fbAF
  }
 }
})

```
