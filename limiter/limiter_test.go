package limiter
import (
	"testing"
	"fmt"
	"runtime"
)

func TestLimiter_MemStat(t *testing.T) {
	info := MemStat()
	fmt.Println(info.Free/1024)
	fmt.Println(runtime.NumCPU())
}

func TestLimiter_DiskStat(t *testing.T) {
	info := DiskStat("/home")
	fmt.Println(info.All/1024)
}

