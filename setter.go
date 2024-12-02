package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"unsafe"
)

/*
1,2,4,6,8,16,32,64

Changes:

whenever we try to set something, we lazily traverse the ShardManagerKeeper for now, if we find the index in there,cool

if not, we initiate the addition of the next ShardManager with the current size and just wait until we have it

*/

func _setKey(key string, value string) {
	idx := int32(696969696)

	if value, ok := keyManager.Keys.Load(key); ok {
		if intValue, ok := value.(int32); ok {
			idx = int32(intValue)
		} else {
			fmt.Println("NOOOOOOOOOOOOOOOOOOOOOO set-x-x-x-x-x-x-x-x-x-x-xx-x-x-x-x-x-x--x", value, "-->")
			os.Exit(1)
		}
	} else {
		val := atomic.AddInt32(&nextIdx, 1)
		keyManager.Keys.Store(key, val)
		idx = val
	}

	if idx == 696969696 {
		fmt.Println("trying to set non existing shit")
		os.Exit(1)
	}

	shardNumber := idx / ShardSize
	localShardIndex := idx % ShardSize

	fmt.Println("setting key", key, "at", idx, "at shard number", shardNumber, "at local index", localShardIndex)

	fmt.Println("trying to acquire lock to set key")

	newVal := getNewValueData(value)

	ShardManagerKeeper.mutex.Lock()

	// TODO: fix the below shit, it should not be this way
	// fmt.Println("set worker locked acquired")

	if shardNumber >= ShardManagerKeeper.capacity {
		// do soemthing about it, lmao
		ShardManagerKeeper.mutex.Unlock()

		UpgradeShardManagerKeeper(shardNumber)

		for atomic.LoadInt32(&ShardManagerKeeper.capacity) <= shardNumber {
			// just wait it out man
		}

		ShardManagerKeeper.mutex.Lock()
	}

	tmpNow := 0
	for tmpNow 
	shard := shardManager.Shards[shardNumber]

	shardManager.mutex.Unlock()

	fmt.Println("set worker locked released")

	// value is a darn string
	atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&shard.data[localShardIndex])), unsafe.Pointer(newVal))
}
