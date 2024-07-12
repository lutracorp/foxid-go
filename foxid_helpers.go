package foxid

import (
	"math/rand"
	"os"
	"strconv"
)

func datacenterOrRand() uint16 {
	if datacenter, ok := os.LookupEnv("FOXID_DATACENTER"); ok {
		if u, err := strconv.ParseUint(datacenter, 10, 16); err == nil {
			return uint16(u)
		}
	}

	return uint16(rand.Uint32())
}

func workerOrPid() uint16 {
	if worker, ok := os.LookupEnv("FOXID_WORKER"); ok {
		if u, err := strconv.ParseUint(worker, 10, 16); err == nil {
			return uint16(u)
		}
	}

	return uint16(os.Getpid())
}
