package worker

import (
	"log"

	"github.com/c9s/goprocinfo/linux"
)

// Stats aggregates system resource statistics for memory, disk, CPU, load average, and task count.
type Stats struct {
	// Memory statistics from /proc/meminfo
	MemStats *linux.MemInfo

	// Disk usage statistics for the root filesystem
	DiskStats *linux.Disk

	// CPU usage statistics from /proc/stat
	CpuStats *linux.CPUStat

	// Load average statistics from /proc/loadavg
	LoadStats *linux.LoadAvg
	//
	// Number of tasks currently managed by the system
	TaskCount int
}

// MemUsedKb returns the amount of memory used in kilobytes.
func (s *Stats) MemUsedKb() uint64 {
	return s.MemStats.MemTotal - s.MemStats.MemAvailable
}

// MemUsedPercent returns the percentage of memory currently in use.
func (s *Stats) MemUsedPercent() uint64 {
	if s.MemStats.MemTotal == 0 {
		return 0
	}
	return s.MemStats.MemAvailable / s.MemStats.MemTotal
}

// MemAvailableKb returns the amount of memory available in kilobytes.
func (s *Stats) MemAvailableKb() uint64 {
	return s.MemStats.MemAvailable
}

// MemTotalKb returns the total amount of memory in kilobytes.
func (s *Stats) MemTotalKb() uint64 {
	return s.MemStats.MemTotal
}

// DiskTotal returns the total disk space in bytes.
func (s *Stats) DiskTotal() uint64 {
	return s.DiskStats.All
}

// DiskFree returns the amount of free disk space in bytes.
func (s *Stats) DiskFree() uint64 {
	return s.DiskStats.Free
}

// DiskUsed returns the amount of used disk space in bytes.
func (s *Stats) DiskUsed() uint64 {
	return s.DiskStats.Used
}

// CpuUsage returns the CPU usage as a percentage (0.0 to 1.0).
func (s *Stats) CpuUsage() float64 {
	idle := s.CpuStats.Idle + s.CpuStats.IOWait
	nonIdle := s.CpuStats.User + s.CpuStats.Nice + s.CpuStats.System + s.CpuStats.IRQ + s.CpuStats.SoftIRQ + s.CpuStats.Steal
	total := idle + nonIdle

	if total == 0 {
		return 0.00
	}

	return (float64(total) - float64(idle)) / float64(total)
}

// GetStats retrieves and aggregates system resource statistics.
func GetStats() *Stats {
	return &Stats{
		MemStats:  GetMemoryInfo(),
		DiskStats: GetDiskInfo(),
		CpuStats:  GetCpuStats(),
		LoadStats: GetLoadAvg(),
	}
}

// GetMemoryInfo retrieves memory statistics from /proc/meminfo.
func GetMemoryInfo() *linux.MemInfo {
	memstats, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Printf("Error reading from /proc/meminfo: %v", err)
		return &linux.MemInfo{}
	}
	return memstats
}

// GetDiskInfo retrieves disk usage statistics for the root filesystem.
func GetDiskInfo() *linux.Disk {
	diskstats, err := linux.ReadDisk("/")
	if err != nil {
		log.Printf("Error reading disk stats from /: %v", err)
		return &linux.Disk{}
	}
	return diskstats
}

// GetCpuStats retrieves CPU usage statistics from /proc/stat.
func GetCpuStats() *linux.CPUStat {
	stats, err := linux.ReadStat("/proc/stat")
	if err != nil {
		log.Printf("Error reading from /proc/stat: %v", err)
		return &linux.CPUStat{}
	}
	return &stats.CPUStatAll
}

// GetLoadAvg retrieves system load average statistics from /proc/loadavg.
func GetLoadAvg() *linux.LoadAvg {
	loadavg, err := linux.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		log.Printf("Error reading from /proc/loadavg: %v", err)
		return &linux.LoadAvg{}
	}
	return loadavg
}

