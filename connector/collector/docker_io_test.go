package collector

import (
	"testing"

	api "github.com/fsouza/go-dockerclient"
)

func TestReadIOOpCase(t *testing.T) {
	for _, tc := range []struct{ name, r, w string }{
		{"cgroup v1", "Read", "Write"},
		{"cgroup v2", "read", "write"},
	} {
		var stats api.Stats
		stats.BlkioStats.IOServiceBytesRecursive = []api.BlkioStatsEntry{
			{Major: 259, Minor: 0, Op: tc.r, Value: 4096},
			{Major: 259, Minor: 0, Op: tc.w, Value: 1781760},
			{Major: 259, Minor: 0, Op: "Total", Value: 1785856},
		}
		d := &Docker{}
		d.ReadIO(&stats)
		if d.IOBytesRead != 4096 || d.IOBytesWrite != 1781760 {
			t.Errorf("%s: got read=%d write=%d, want 4096/1781760", tc.name, d.IOBytesRead, d.IOBytesWrite)
		}
	}
}
