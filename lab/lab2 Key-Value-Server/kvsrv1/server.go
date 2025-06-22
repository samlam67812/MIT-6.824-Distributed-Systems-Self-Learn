package kvsrv

import (
	"log"
	"sync"

	"6.5840/kvsrv1/rpc"
	"6.5840/labrpc"
	"6.5840/tester1"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}


type KVServer struct {
	mu sync.Mutex

	// Your definitions here.
	pairs map[string]*ValueHandle
}
type ValueHandle struct {
	Value string
	Version rpc.Tversion
}


func MakeKVServer() *KVServer {
	kv := &KVServer{
		pairs: make(map[string]*ValueHandle),
	}
	return kv
}

// Get returns the value and version for args.Key, if args.Key
// exists. Otherwise, Get returns ErrNoKey.
func (kv *KVServer) Get(args *rpc.GetArgs, reply *rpc.GetReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	
	v, ok := kv.pairs[args.Key]; if ok {
		reply.Err = rpc.OK
		reply.Value = v.Value
		reply.Version = v.Version
	} else {
		reply.Err = rpc.ErrNoKey
		reply.Value = ""
		reply.Version = 0
		return
	}

}

// Update the value for a key if args.Version matches the version of
// the key on the server. If versions don't match, return ErrVersion.
// If the key doesn't exist, Put installs the value if the
// args.Version is 0, and returns ErrNoKey otherwise.
func (kv *KVServer) Put(args *rpc.PutArgs, reply *rpc.PutReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	key := args.Key
	version := args.Version
	v, ok := kv.pairs[key];
	if !ok {
		if version == 0 {
			reply.Err = rpc.OK
			kv.pairs[key] = &ValueHandle{args.Value, version + 1}
		} else {
			reply.Err = rpc.ErrNoKey
			return
		}
	} else {
		if v.Version != args.Version {
			reply.Err = rpc.ErrVersion
			return
		}
		reply.Err = rpc.OK
		kv.pairs[key] = &ValueHandle{args.Value, version + 1}
	}

}

// You can ignore Kill() for this lab
func (kv *KVServer) Kill() {
}


// You can ignore all arguments; they are for replicated KVservers
func StartKVServer(ends []*labrpc.ClientEnd, gid tester.Tgid, srv int, persister *tester.Persister) []tester.IService {
	kv := MakeKVServer()
	return []tester.IService{kv}
}
