package glusterfs

import (
	"github.com/gluster/gluster-csi-driver/pkg/glusterfs/config"
	"os"
	"testing"

	"github.com/kubernetes-csi/csi-test/pkg/sanity"
	"k8s.io/kubernetes/pkg/util/mount"
)

type volume struct {
	Size     uint64
	snapList []string
}

func TestDriverSuite(t *testing.T) {
	glusterMounter = &mount.FakeMounter{}
	socket := "/tmp/csi.sock"
	endpoint := "unix://" + socket

	//cleanup socket file if already present
	os.Remove(socket)

	_, err := os.Create(socket)
	if err != nil {
		t.Fatal("Failed to create a socket file")
	}
	defer os.Remove(socket)

	d := GfDriver{
		Config: &config.Config{
			Endpoint: endpoint,
			NodeID:   "testing",
		},
	}

	go d.Run()

	mntStageDir := "/tmp/mntStageDir"
	mntDir := "/tmp/mntDir"
	defer os.RemoveAll(mntStageDir)
	defer os.RemoveAll(mntDir)

	cfg := &sanity.Config{
		StagingPath: mntStageDir,
		TargetPath:  mntDir,
		Address:     endpoint,
	}

	sanity.Test(t, cfg)
}
