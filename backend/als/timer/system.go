package timer

import (
	"runtime"
	"strconv"
	"time"

	"github.com/X-Zero-L/als/als/client"
)

func UpdateSystemResource() {
	var m runtime.MemStats
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		runtime.ReadMemStats(&m)
		client.BroadCastMessage("MemoryUsage", strconv.Itoa(int(m.Sys)))
	}

}
