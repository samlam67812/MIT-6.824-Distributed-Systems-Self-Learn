package lock

import (
	"time"

	"6.5840/kvtest1"
)

type Lock struct {
	// IKVClerk is a go interface for k/v clerks: the interface hides
	// the specific Clerk type of ck but promises that ck supports
	// Put and Get.  The tester passes the clerk in when calling
	// MakeLock().
	ck kvtest.IKVClerk
	lockStr string
	id string

}

// The tester calls MakeLock() and passes in a k/v clerk; your code can
// perform a Put or Get by calling lk.ck.Put() or lk.ck.Get().
//
// Use l as the key to store the "lock state" (you would have to decide
// precisely what the lock state is).
func MakeLock(ck kvtest.IKVClerk, l string) *Lock {
	lk := &Lock{
		ck: ck,
		lockStr: l,
		id: kvtest.RandValue(8), // unique identifier for each lock client
	}

	return lk
}

func (lk *Lock) Acquire() {
	for {
		id, version, err := lk.ck.Get(lk.lockStr)
		if id == lk.id {
			return
		}
		if err == rpc.ErrNoKey || id == "" {
			err := lk.ck.Put(lk.lockStr, lk.id, version)
			if err == rpc.OK {
				return
			}
		}
		time.Sleep(10 * time.Millisecond)
	}

}

func (lk *Lock) Release() {
	id, version, err := lk.ck.Get(lk.lockStr)
	if err == rpc.OK && id == lk.id {
		lk.ck.Put(lk.lockStr, "", version)
	}

}
