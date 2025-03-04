package naming

import (
	"fmt"
)

// Name returns the manifest name of a component.
// Example: tempo-simplest-compactor.
func Name(component string, tempoStackName string) string {
	if component == "" || component == "tempo" {
		return DNSName(fmt.Sprintf("tempo-%s", tempoStackName))
	}
	return DNSName(fmt.Sprintf("tempo-%s-%s", tempoStackName, component))
}

// TLSSecretName returns the secret name that stores the TLS cert/key for given component.
func TLSSecretName(component string, tempoStackName string) string {
	return DNSName(fmt.Sprintf("%s-mtls", Name(component, tempoStackName)))
}

// ServiceFqdn returns the fully qualified domain name of a service.
func ServiceFqdn(namespace string, tempoStackName string, component string) string {
	return fmt.Sprintf("%s.%s.svc.cluster.local", Name(component, tempoStackName), namespace)
}

// DefaultServiceAccountName returns the name of the default tempo service account to use.
func DefaultServiceAccountName(name string) string {
	return Name("", name)
}

// SigningCABundleName return CA bundle configmap name.
func SigningCABundleName(name string) string {
	return fmt.Sprintf("tempo-%s-ca-bundle", name)
}

// PrometheusRuleName is the name of the tempo-prometheus-rule.
func PrometheusRuleName(stackName string) string {
	return fmt.Sprintf("%s-prometheus-rule", stackName)
}
