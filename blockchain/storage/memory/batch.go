package memory

import (
	"github.com/PrivixAI-labs/Privix-node/blockchain/storage"
	"github.com/PrivixAI-labs/Privix-node/helper/hex"
)

var _ storage.Batch = (*batchMemory)(nil)

type batchMemory struct {
	db           map[string][]byte
	keysToDelete [][]byte
	valuesToPut  [][2][]byte
}

func NewBatchMemory(db map[string][]byte) *batchMemory {
	return &batchMemory{db: db}
}

func (b *batchMemory) Delete(key []byte) {
	b.keysToDelete = append(b.keysToDelete, key)
}

func (b *batchMemory) Put(k []byte, v []byte) {
	b.valuesToPut = append(b.valuesToPut, [2][]byte{k, v})
}

func (b *batchMemory) Write() error {
	for _, x := range b.keysToDelete {
		delete(b.db, hex.EncodeToHex(x))
	}

	for _, x := range b.valuesToPut {
		b.db[hex.EncodeToHex(x[0])] = x[1]
	}

	return nil
}
