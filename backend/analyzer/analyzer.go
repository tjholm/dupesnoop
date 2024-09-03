package analyzer

import (
	"github.com/shirou/gopsutil/v4/disk"
)

type Analyzer struct {
}

type DiskInfo struct {
	Device string `json:"device"`
	Mount  string `json:"mount"`
	FsType string `json:"fs"`
	Used   uint64 `json:"used"`
	Total  uint64 `json:"total"`
}

func (a *Analyzer) ListDisks() ([]DiskInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	disks := make([]DiskInfo, 0, len(partitions))
	for _, partition := range partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			return nil, err
		}

		disks = append(disks, DiskInfo{
			Device: partition.Device,
			Mount:  partition.Mountpoint,
			FsType: partition.Fstype,
			Used:   usageStat.Used,
			Total:  usageStat.Total,
		})
	}

	return disks, nil
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}
