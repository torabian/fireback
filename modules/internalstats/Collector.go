// Collector.go gathers every stat InternalStatsSnapshot/InternalStatsStream report -
// around 40 individual measurements across host, CPU, memory, disk, network and this
// process's own Go runtime - into a flat, display-ordered CollectSnapshot() call.
// Both actions (SnapshotActionImplementation.go/StreamActionImplementation.go, via
// Stream.go) and the `internalstats watch` CLI table (Cli.go) call this same function,
// so there is exactly one place that decides what "server health" means here.
//
// gopsutil (github.com/shirou/gopsutil/v3) does the actual OS-level reading - it's
// already cross-platform (linux/darwin/windows), unlike modules/backup/DiskSpace.go's
// hand-rolled per-OS split, which would need to triple for cpu/mem/disk/host/load/net.
package internalstats

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	"github.com/torabian/emi/emigo"
	internalstatsdefs "github.com/torabian/fireback/modules/internalstats/defs"
)

// Severity thresholds applied to every percent-based stat (CPU/memory/swap/disk
// usage). Anything else is reported as SeverityInfo - there's no meaningful
// "too high" reading for e.g. a hostname or a Go version.
const (
	SeverityOk       = "ok"
	SeverityWarn     = "warn"
	SeverityCritical = "critical"
	SeverityInfo     = "info"
)

const (
	warnPercent     = 75.0
	criticalPercent = 90.0
)

// processStart is this process's own start time (package init, i.e. as close to the
// binary's actual start as anything internalstats can observe) - used for the
// "Process Uptime" runtime stat.
var processStart = time.Now()

// CollectSnapshot reads every stat fresh (no caching - each call is a real,
// point-in-time read) and returns them in a stable, category-grouped display order.
func CollectSnapshot() *internalstatsdefs.InternalStatsSnapshotActionRes {
	c := &collector{}

	c.collectHost()
	c.collectCPU()
	c.collectMemory()
	c.collectDisk()
	c.collectNetwork()
	c.collectRuntime()

	hostname := c.hostname
	if hostname == "" {
		hostname, _ = os.Hostname()
	}

	return &internalstatsdefs.InternalStatsSnapshotActionRes{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Hostname:    hostname,
		Items:       emigo.ArrayReplace(c.items),
	}
}

type collector struct {
	items    []internalstatsdefs.InternalStatsSnapshotActionResItems
	hostname string
}

func (c *collector) add(key, label, category, value string, rawValue float64, unit, severity string) {
	c.items = append(c.items, internalstatsdefs.InternalStatsSnapshotActionResItems{
		Key:      key,
		Label:    label,
		Category: category,
		Value:    value,
		RawValue: rawValue,
		Unit:     unit,
		Severity: severity,
	})
}

func (c *collector) addInfo(key, label, category, value string) {
	c.add(key, label, category, value, 0, "", SeverityInfo)
}

func percentSeverity(pct float64) string {
	switch {
	case pct >= criticalPercent:
		return SeverityCritical
	case pct >= warnPercent:
		return SeverityWarn
	default:
		return SeverityOk
	}
}

func (c *collector) addPercent(key, label, category string, pct float64) {
	c.add(key, label, category, fmt.Sprintf("%.1f%%", pct), pct, "percent", percentSeverity(pct))
}

const (
	categoryHost    = "Host"
	categoryCPU     = "CPU"
	categoryMemory  = "Memory"
	categoryDisk    = "Disk"
	categoryNetwork = "Network"
	categoryRuntime = "Runtime"
)

func (c *collector) collectHost() {
	info, err := host.Info()
	if err != nil {
		c.addInfo("host.error", "Host Info", categoryHost, err.Error())
		return
	}

	c.hostname = info.Hostname

	c.addInfo("host.hostname", "Hostname", categoryHost, info.Hostname)
	c.addInfo("host.os", "OS", categoryHost, fmt.Sprintf("%s (%s)", info.Platform, info.OS))
	c.addInfo("host.platformVersion", "OS Version", categoryHost, info.PlatformVersion)
	c.addInfo("host.kernelVersion", "Kernel Version", categoryHost, fmt.Sprintf("%s (%s)", info.KernelVersion, info.KernelArch))
	c.add("host.uptime", "Uptime", categoryHost, formatDuration(time.Duration(info.Uptime)*time.Second), float64(info.Uptime), "seconds", SeverityInfo)
	bootTime := time.Unix(int64(info.BootTime), 0).UTC().Format(time.RFC3339)
	c.addInfo("host.bootTime", "Boot Time", categoryHost, bootTime)
	c.add("host.processes", "Running Processes", categoryHost, fmt.Sprintf("%d", info.Procs), float64(info.Procs), "count", SeverityInfo)
}

func (c *collector) collectCPU() {
	if info, err := cpu.Info(); err == nil && len(info) > 0 {
		c.addInfo("cpu.model", "CPU Model", categoryCPU, info[0].ModelName)
		// gopsutil reads Mhz from sysctl on Apple Silicon, which doesn't expose a
		// real clock speed there and returns a tiny bogus value (e.g. 4) instead of
		// erroring - filtered out here rather than showing a misleading number.
		if info[0].Mhz >= 100 {
			c.add("cpu.mhz", "CPU Frequency", categoryCPU, fmt.Sprintf("%.0f MHz", info[0].Mhz), info[0].Mhz, "mhz", SeverityInfo)
		} else {
			c.addInfo("cpu.mhz", "CPU Frequency", categoryCPU, "n/a")
		}
	} else {
		c.addInfo("cpu.model", "CPU Model", categoryCPU, "unknown")
	}

	if physical, err := cpu.Counts(false); err == nil {
		c.add("cpu.physicalCores", "Physical Cores", categoryCPU, fmt.Sprintf("%d", physical), float64(physical), "count", SeverityInfo)
	}
	logical := runtime.NumCPU()
	c.add("cpu.logicalCores", "Logical Cores", categoryCPU, fmt.Sprintf("%d", logical), float64(logical), "count", SeverityInfo)

	// A 200ms sample is enough for a live meaningful reading without stalling
	// CollectSnapshot (and therefore every caller: the HTTP snapshot, the reactive
	// stream's per-tick push, and the CLI watch loop) noticeably.
	if pct, err := cpu.Percent(200*time.Millisecond, false); err == nil && len(pct) > 0 {
		c.addPercent("cpu.usedPercent", "CPU Usage", categoryCPU, pct[0])
	} else {
		c.addInfo("cpu.usedPercent", "CPU Usage", categoryCPU, "n/a")
	}

	// load.Avg is unix-only in gopsutil (returns an error on Windows) - reported as
	// n/a there instead of failing the whole snapshot.
	if avg, err := load.Avg(); err == nil {
		c.add("cpu.load1", "Load Avg (1m)", categoryCPU, fmt.Sprintf("%.2f", avg.Load1), avg.Load1, "load", SeverityInfo)
		c.add("cpu.load5", "Load Avg (5m)", categoryCPU, fmt.Sprintf("%.2f", avg.Load5), avg.Load5, "load", SeverityInfo)
		c.add("cpu.load15", "Load Avg (15m)", categoryCPU, fmt.Sprintf("%.2f", avg.Load15), avg.Load15, "load", SeverityInfo)
	} else {
		c.addInfo("cpu.load1", "Load Avg (1m)", categoryCPU, "n/a")
		c.addInfo("cpu.load5", "Load Avg (5m)", categoryCPU, "n/a")
		c.addInfo("cpu.load15", "Load Avg (15m)", categoryCPU, "n/a")
	}
}

func (c *collector) collectMemory() {
	if vm, err := mem.VirtualMemory(); err == nil {
		c.add("mem.total", "Memory Total", categoryMemory, formatBytes(vm.Total), float64(vm.Total), "bytes", SeverityInfo)
		c.add("mem.used", "Memory Used", categoryMemory, formatBytes(vm.Used), float64(vm.Used), "bytes", SeverityInfo)
		c.add("mem.free", "Memory Free", categoryMemory, formatBytes(vm.Free), float64(vm.Free), "bytes", SeverityInfo)
		c.add("mem.available", "Memory Available", categoryMemory, formatBytes(vm.Available), float64(vm.Available), "bytes", SeverityInfo)
		c.addPercent("mem.usedPercent", "Memory Used", categoryMemory, vm.UsedPercent)
		c.add("mem.cached", "Memory Cached", categoryMemory, formatBytes(vm.Cached), float64(vm.Cached), "bytes", SeverityInfo)
	} else {
		c.addInfo("mem.error", "Memory", categoryMemory, err.Error())
	}

	if sm, err := mem.SwapMemory(); err == nil {
		c.add("swap.total", "Swap Total", categoryMemory, formatBytes(sm.Total), float64(sm.Total), "bytes", SeverityInfo)
		c.add("swap.used", "Swap Used", categoryMemory, formatBytes(sm.Used), float64(sm.Used), "bytes", SeverityInfo)
		c.addPercent("swap.usedPercent", "Swap Used", categoryMemory, sm.UsedPercent)
	}
}

// rootPath is the filesystem disk.Usage is read for - the volume this process itself
// runs from, same intent as modules/backup/DiskSpace.go's AvailableDiskSpace but read
// via gopsutil (which already knows the right per-OS root) instead of a hand-rolled
// unix/windows split.
func rootPath() string {
	if runtime.GOOS == "windows" {
		return `C:\`
	}
	return "/"
}

func (c *collector) collectDisk() {
	usage, err := disk.Usage(rootPath())
	if err != nil {
		c.addInfo("disk.error", "Disk", categoryDisk, err.Error())
		return
	}

	c.add("disk.total", "Disk Total", categoryDisk, formatBytes(usage.Total), float64(usage.Total), "bytes", SeverityInfo)
	c.add("disk.used", "Disk Used", categoryDisk, formatBytes(usage.Used), float64(usage.Used), "bytes", SeverityInfo)
	c.add("disk.free", "Disk Free", categoryDisk, formatBytes(usage.Free), float64(usage.Free), "bytes", SeverityInfo)
	c.addPercent("disk.usedPercent", "Disk Used", categoryDisk, usage.UsedPercent)
}

func (c *collector) collectNetwork() {
	counters, err := net.IOCounters(false)
	if err != nil || len(counters) == 0 {
		c.addInfo("net.error", "Network", categoryNetwork, "n/a")
		return
	}

	total := counters[0]
	c.add("net.bytesSent", "Bytes Sent", categoryNetwork, formatBytes(total.BytesSent), float64(total.BytesSent), "bytes", SeverityInfo)
	c.add("net.bytesRecv", "Bytes Received", categoryNetwork, formatBytes(total.BytesRecv), float64(total.BytesRecv), "bytes", SeverityInfo)
	c.add("net.packetsSent", "Packets Sent", categoryNetwork, fmt.Sprintf("%d", total.PacketsSent), float64(total.PacketsSent), "count", SeverityInfo)
	c.add("net.packetsRecv", "Packets Received", categoryNetwork, fmt.Sprintf("%d", total.PacketsRecv), float64(total.PacketsRecv), "count", SeverityInfo)
}

func (c *collector) collectRuntime() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.addInfo("runtime.goVersion", "Go Version", categoryRuntime, runtime.Version())
	c.add("runtime.goroutines", "Goroutines", categoryRuntime, fmt.Sprintf("%d", runtime.NumGoroutine()), float64(runtime.NumGoroutine()), "count", SeverityInfo)
	c.add("runtime.heapAlloc", "Heap Alloc", categoryRuntime, formatBytes(m.HeapAlloc), float64(m.HeapAlloc), "bytes", SeverityInfo)
	c.add("runtime.heapSys", "Heap Sys", categoryRuntime, formatBytes(m.HeapSys), float64(m.HeapSys), "bytes", SeverityInfo)
	c.add("runtime.numGC", "GC Runs", categoryRuntime, fmt.Sprintf("%d", m.NumGC), float64(m.NumGC), "count", SeverityInfo)
	uptime := time.Since(processStart)
	c.add("runtime.processUptime", "Process Uptime", categoryRuntime, formatDuration(uptime), uptime.Seconds(), "seconds", SeverityInfo)
}

// formatBytes renders a byte count the way an operator reads it - binary units
// (1024-based, like `free`/`df`), one decimal place, capped at TB since nothing this
// module reports is expected to exceed that.
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit && exp < 4; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGT"[exp])
}

// formatDuration renders a duration as "3d 4h 12m" (dropping smaller units once a
// bigger one is present, dropping seconds entirely once minutes are), rather than Go's
// default "76h32m32.427s" - readable at a glance in the watch table.
func formatDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	totalSeconds := int64(d.Seconds())
	days := totalSeconds / 86400
	hours := (totalSeconds % 86400) / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	switch {
	case days > 0:
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	case hours > 0:
		return fmt.Sprintf("%dh %dm", hours, minutes)
	case minutes > 0:
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	default:
		return fmt.Sprintf("%ds", seconds)
	}
}
