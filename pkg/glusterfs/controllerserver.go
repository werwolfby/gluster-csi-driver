package glusterfs

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	volumeOwnerAnn = "VolumeOwner"
)

// ControllerServer struct of GlusterFS CSI driver with supported methods of CSI
// controller server spec.
type ControllerServer struct {
}

// CsiDrvParam stores csi driver specific request parameters.
// This struct will be used to gather specific fields of CSI driver:
// For eg. csiDrvName, csiDrvVersion..etc and also gather
// parameters passed from SC which not part of gluster volcreate api.
// GlusterCluster - The resturl of gluster cluster
// GlusterUser - The gluster username who got access to the APIs.
// GlusterUserToken - The password/token of glusterUser to connect to
// glusterCluster.
// GlusterVersion - Says the version of the glustercluster
// running in glusterCluster endpoint.
// All of these fields are optional and can be used if needed.
type CsiDrvParam struct {
	GlusterCluster   string
	GlusterUser      string
	GlusterUserToken string
	GlusterVersion   string
	CsiDrvName       string
	CsiDrvVersion    string
}

// ProvisionerConfig is the combined configuration of gluster cli vol create
// request and CSI driver specific input
type ProvisionerConfig struct {
	// csiConf  *CsiDrvParam
}

// CreateVolume creates and starts the volume
func (cs *ControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// DeleteVolume deletes the given volume.
func (cs *ControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerPublishVolume return Unimplemented error
func (cs *ControllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerUnpublishVolume return Unimplemented error
func (cs *ControllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ValidateVolumeCapabilities checks whether the volume capabilities requested
// are supported.
func (cs *ControllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ValidateVolumeCapabilities() - request is nil")
	}

	volumeID := req.GetVolumeId()
	if volumeID == "" {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities() - VolumeId is nil")
	}

	reqCaps := req.GetVolumeCapabilities()
	if reqCaps == nil {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities() - VolumeCapabilities is nil")
	}

	var vcaps []*csi.VolumeCapability_AccessMode
	for _, mode := range []csi.VolumeCapability_AccessMode_Mode{
		csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	} {
		vcaps = append(vcaps, &csi.VolumeCapability_AccessMode{Mode: mode})
	}
	capSupport := false

	for _, cap := range reqCaps {
		for _, m := range vcaps {
			if m.Mode == cap.AccessMode.Mode {
				capSupport = true
			}
		}
	}

	if !capSupport {
		return nil, status.Errorf(codes.NotFound, "%v not supported", req.GetVolumeCapabilities())
	}

	resp := &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeCapabilities: req.VolumeCapabilities,
		},
	}

	glog.V(1).Infof("GlusterFS CSI driver volume capabilities: %+v", resp)
	return resp, nil
}

// ListVolumes returns a list of volumes
func (cs *ControllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// GetCapacity returns the capacity of the storage pool
func (cs *ControllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerGetCapabilities returns the capabilities of the controller service.
func (cs *ControllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	resp := &csi.ControllerGetCapabilitiesResponse{}

	return resp, nil
}

// CreateSnapshot create snapshot of an existing PV
func (cs *ControllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// DeleteSnapshot delete provided snapshot of a PV
func (cs *ControllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ListSnapshots list the snapshots of a PV
func (cs *ControllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
