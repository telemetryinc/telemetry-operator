package controller

import (
	telemetryv1 "github.com/telemetryinc/telemetry-operator/api/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	sccNonroot    = "nonroot"
	sccPrivileged = "privileged"
)

func (r *TelemetryReconciler) openshiftSCCRole(cr *telemetryv1.Telemetry, scc string) *rbacv1.Role {
	return &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-" + scc,
			Namespace: cr.Namespace,
			Labels:    Labels(cr, "roles"),
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{"security.openshift.io"},
				Resources:     []string{"securitycontextconstraints"},
				ResourceNames: []string{scc},
				Verbs:         []string{"use"},
			},
		},
	}
}

func (r *TelemetryReconciler) openshiftSCCRoleBinding(cr *telemetryv1.Telemetry, component, scc string) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-" + component + "-" + scc,
			Namespace: cr.Namespace,
			Labels:    Labels(cr, component),
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      cr.Name + "-" + component,
				Namespace: cr.Namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "Role",
			Name:     cr.Name + "-" + scc,
		},
	}
}
