package glusterfs

import (
	"github.com/gluster/gluster-csi-driver/pkg/glusterfs/config"

	"github.com/golang/glog"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)

const (
	glusterfsCSIDriverName    = "org.gluster.glusterfs"
	glusterfsCSIDriverVersion = "1.0.0"
)

// GfDriver is the struct embedding information about the connection to gluster
// cluster and configuration of CSI driver.
type GfDriver struct {
	*config.Config
}

// New returns CSI driver
func New(config *config.Config) *GfDriver {
	gfd := &GfDriver{}

	if config == nil {
		glog.Errorf("GlusterFS CSI driver initialization failed")
		return nil
	}

	gfd.Config = config

	glog.V(1).Infof("GlusterFS CSI driver initialized")

	return gfd
}

// NewControllerServer initialize a controller server for GlusterFS CSI driver.
func NewControllerServer() *ControllerServer {
	return &ControllerServer{}
}

// NewNodeServer initialize a node server for GlusterFS CSI driver.
func NewNodeServer(nodeID string) *NodeServer {
	return &NodeServer{
		NodeID: nodeID,
	}
}

// NewIdentityServer initialize an identity server for GlusterFS CSI driver.
func NewIdentityServer() *IdentityServer {
	return &IdentityServer{}
}

// Run start a non-blocking grpc controller,node and identityserver for
// GlusterFS CSI driver which can serve multiple parallel requests
func (g *GfDriver) Run() {
	srv := csicommon.NewNonBlockingGRPCServer()
	srv.Start(g.Endpoint, NewIdentityServer(), nil, NewNodeServer(g.NodeID))
	srv.Wait()
}
