./neth  server --chain testnet.json --libp2p 0.0.0.0:10001 --nat 0.0.0.0 --jsonrpc 0.0.0.0:7545 --seal  --data-dir=node/node1 --grpc-address 0.0.0.0:20001


./neth server --chain testnet.json --libp2p 0.0.0.0:10002 --nat 0.0.0.0 --jsonrpc 0.0.0.0:7545 --seal  --data-dir=node/node2 --grpc-address 0.0.0.0:20002

 ./neth server --chain testnet.json --libp2p 0.0.0.0:10003 --nat 0.0.0.0 --jsonrpc 0.0.0.0:7545 --seal  --data-dir=node/node3 --grpc-address 0.0.0.0:20003

./neth server --chain testnet.json --libp2p 0.0.0.0:10004 --nat 0.0.0.0 --jsonrpc 0.0.0.0:7545 --seal  --data-dir=node/node4 --grpc-address 0.0.0.0:20004


./neth genesis --pos --ibft-validator-type ecdsa --bootnode /ip4/localhost/tcp/10008/p2p/16Uiu2HAmAN3M2CKfjFENKoTVApk1AYzU2F6dcTpn8MFkLqbMgGq2 --ibft-validator 0xE37e24c159943E5204b53e5e3898ea5Ee39Bf941 --ibft-validator 0x5c59dd8B90Ad01d1865d84DdaA01CdD5c7316a92 --ibft-validator 0xE67d7Db0B36A36f6C3cFF364EA0BfD9a1Ff21AdF --ibft-validator 0xa657F44b509ffb6F1Eb3BA0C5D08974D5cf3245E --premine 0xC6bF9487F53deE994210b7844636685b9C75bBA7:1000000000000000000000000 --dir test.json

./neth genesis --pos --ibft-validator-type ecdsa  