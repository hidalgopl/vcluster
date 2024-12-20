package types

import (
	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	corev1 "k8s.io/api/core/v1"
)

type EKSPodIdentity interface {
	Create(ctx *synccontext.SyncContext, sa *corev1.ServiceAccount) error
	Delete(ctx *synccontext.SyncContext, sa *corev1.ServiceAccount) error
}
