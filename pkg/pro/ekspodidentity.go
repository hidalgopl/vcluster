package pro

import (
	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	syncertypes "github.com/loft-sh/vcluster/pkg/syncer/types"
	corev1 "k8s.io/api/core/v1"
)

var CreateEKSPodIdentity = func(ctx *synccontext.SyncContext) (syncertypes.EKSPodIdentity, error) {
	return &noOpEKSPodIdentity{}, nil
}

type noOpEKSPodIdentity struct{}

func (n *noOpEKSPodIdentity) Create(ctx *synccontext.SyncContext, sa *corev1.ServiceAccount) error {
	return nil
}

func (n *noOpEKSPodIdentity) Delete(ctx *synccontext.SyncContext, sa *corev1.ServiceAccount) error {
	return nil
}
