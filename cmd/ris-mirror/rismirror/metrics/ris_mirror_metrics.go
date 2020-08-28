package metrics

import (
	"net"

	vrf_metrics "github.com/bio-routing/bio-rd/routingtable/vrf/metrics"
)

// RISMirrorMetrics contains per router BMP metrics
type RISMirrorMetrics struct {
	Routers []*RISMirrorRouterMetrics
}

// RISMirrorRouterMetrics contains a routers RIS mirror metrics
type RISMirrorRouterMetrics struct {
	// Routers IP Address
	Address net.IP

	// SysName of the monitored router
	SysName string

	// VRFMetrics represent per VRF metrics
	VRFMetrics []*vrf_metrics.VRFMetrics
}
