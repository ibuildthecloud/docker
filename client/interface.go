package client // import "github.com/docker/docker/client"

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	volumetypes "github.com/docker/docker/api/types/volume"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

// CommonAPIClient is the common methods between stable and experimental versions of APIClient.
type CommonAPIClient interface {
	ContainerAPIClient
	DistributionAPIClient
	ImageAPIClient
	SystemAPIClient
	VolumeAPIClient
	ClientVersion() string
	DaemonHost() string
	HTTPClient() *http.Client
	ServerVersion(ctx context.Context) (types.Version, error)
	NegotiateAPIVersion(ctx context.Context)
	NegotiateAPIVersionPing(types.Ping)
	DialHijack(ctx context.Context, url, proto string, meta map[string][]string) (net.Conn, error)
	Dialer() func(context.Context) (net.Conn, error)
	Close() error
}

// ContainerAPIClient defines API client methods for the containers
type ContainerAPIClient interface {
	ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error)
	ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error)
	ContainerCreate(ctx context.Context, config *containertypes.Config, hostConfig *containertypes.HostConfig, networkingConfig *networktypes.NetworkingConfig, platform *specs.Platform, containerName string) (containertypes.ContainerCreateCreatedBody, error)
	ContainerDiff(ctx context.Context, container string) ([]containertypes.ContainerChangeResponseItem, error)
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
	ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error
	ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error
	ContainerExport(ctx context.Context, container string) (io.ReadCloser, error)
	ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error)
	ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error)
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)
	ContainerPause(ctx context.Context, container string) error
	ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error
	ContainerRename(ctx context.Context, container, newContainerName string) error
	ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error
	ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error
	ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error)
	ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error)
	ContainerStatsOneShot(ctx context.Context, container string) (types.ContainerStats, error)
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerStop(ctx context.Context, container string, timeout *time.Duration) error
	ContainerTop(ctx context.Context, container string, arguments []string) (containertypes.ContainerTopOKBody, error)
	ContainerUnpause(ctx context.Context, container string) error
	ContainerUpdate(ctx context.Context, container string, updateConfig containertypes.UpdateConfig) (containertypes.ContainerUpdateOKBody, error)
	ContainerWait(ctx context.Context, container string, condition containertypes.WaitCondition) (<-chan containertypes.ContainerWaitOKBody, <-chan error)
	CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error)
	CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error
	ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
}

// DistributionAPIClient defines API client methods for the registry
type DistributionAPIClient interface {
	DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registry.DistributionInspect, error)
}

// ImageAPIClient defines API client methods for the images
type ImageAPIClient interface {
	ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
	BuildCachePrune(ctx context.Context, opts types.BuildCachePruneOptions) (*types.BuildCachePruneReport, error)
	BuildCancel(ctx context.Context, id string) error
	ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error)
	ImageHistory(ctx context.Context, image string) ([]image.HistoryResponseItem, error)
	ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error)
	ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error)
	ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
	ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error)
	ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
	ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error)
	ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
	ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error)
	ImageSave(ctx context.Context, images []string) (io.ReadCloser, error)
	ImageTag(ctx context.Context, image, ref string) error
	ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error)
}

// SystemAPIClient defines API client methods for the system
type SystemAPIClient interface {
	Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
	Info(ctx context.Context) (types.Info, error)
	RegistryLogin(ctx context.Context, auth types.AuthConfig) (registry.AuthenticateOKBody, error)
	DiskUsage(ctx context.Context) (types.DiskUsage, error)
	Ping(ctx context.Context) (types.Ping, error)
}

// VolumeAPIClient defines API client methods for the volumes
type VolumeAPIClient interface {
	VolumeCreate(ctx context.Context, options volumetypes.VolumeCreateBody) (types.Volume, error)
	VolumeInspect(ctx context.Context, volumeID string) (types.Volume, error)
	VolumeInspectWithRaw(ctx context.Context, volumeID string) (types.Volume, []byte, error)
	VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumeListOKBody, error)
	VolumeRemove(ctx context.Context, volumeID string, force bool) error
	VolumesPrune(ctx context.Context, pruneFilter filters.Args) (types.VolumesPruneReport, error)
}
