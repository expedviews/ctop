package connector

import (
	"testing"
	"time"

	api "github.com/fsouza/go-dockerclient"
)

func TestCalcUptime(t *testing.T) {
	now := time.Now()
	mk := func(started, finished time.Time, running, restarting bool) *api.Container {
		c := &api.Container{}
		c.State.StartedAt = started
		c.State.FinishedAt = finished
		c.State.Running = running
		c.State.Restarting = restarting
		return c
	}
	for _, tc := range []struct {
		name string
		c    *api.Container
		want string
	}{
		{"running", mk(now.Add(-6*time.Minute), time.Time{}, true, false), "6 minutes"},
		{"never started", mk(time.Time{}, time.Time{}, false, false), ""},
		{"never started, running flag", mk(time.Time{}, time.Time{}, true, false), ""},
		{"stopped", mk(now.Add(-3*time.Hour), now.Add(-1*time.Hour), false, false), "2 hours"},
		{"restarted: old FinishedAt", mk(now.Add(-6*time.Minute), now.Add(-2*time.Hour), false, true), "6 minutes"},
	} {
		if got := calcUptime(tc.c); got != tc.want {
			t.Errorf("%s: calcUptime = %q, want %q", tc.name, got, tc.want)
		}
	}
}
