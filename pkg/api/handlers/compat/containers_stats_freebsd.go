//go:build !remote

package compat

import (
	"fmt"
	"time"

	"github.com/containers/podman/v5/libpod"
	"github.com/containers/podman/v5/libpod/define"
	"github.com/docker/docker/api/types/container"
)

func getPreCPUStats(stats *define.ContainerStats) CPUStats {
	return CPUStats{
		CPUUsage: container.CPUUsage{
			TotalUsage: stats.CPUNano,
		},
		CPU:            stats.CPU,
		OnlineCPUs:     0,
		ThrottlingData: container.ThrottlingData{},
	}
}

func statsContainerJSON(ctnr *libpod.Container, stats *define.ContainerStats, preCPUStats CPUStats, onlineCPUs int) (StatsJSON, error) {
	state, err := ctnr.State()
	if err != nil {
		return StatsJSON{}, err
	}
	if state != define.ContainerStateRunning {
		return StatsJSON{}, fmt.Errorf("container %s is not running", ctnr.ID())
	}

	return StatsJSON{
		Stats: Stats{
			Read: time.Now(),
			CPUStats: CPUStats{
				CPUUsage: container.CPUUsage{
					TotalUsage: stats.CPUNano,
				},
				CPU:            stats.CPU,
				OnlineCPUs:     0,
				ThrottlingData: container.ThrottlingData{},
			},
			PreCPUStats: preCPUStats,
			MemoryStats: container.MemoryStats{},
		},
		Name: stats.Name,
		ID:   stats.ContainerID,
	}, nil
}
