package limiter

import (
	"runtime"
	"syscall"
)

func MemStat() MemStatus {
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc

	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err == nil {
		mem.All = sysInfo.Totalram
		mem.Free = sysInfo.Freeram
		mem.Used = mem.All - mem.Free
	}
	return mem
}

func CpuStat() {

}

func DiskStat(path string) DiskStatus {
	disk := DiskStatus{}
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err == nil {
		disk.All = fs.Blocks * uint64(fs.Bsize)
		disk.Free = fs.Blocks * uint64(fs.Bsize)
		disk.Used = disk.All - disk.Free
	}
	return disk
}