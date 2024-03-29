package business

import (
	"github.com/kiali/kiali/kubernetes"
)

// ValidProxyLogLevels are the application log levels supported by the envoy admin interface.
var ValidProxyLogLevels = []string{"off", "trace", "debug", "info", "warning", "error", "critical"}

// IsValidLogLevel determines if the provided string is a valid proxy log level.
// This can be called before calling SetLogLevel.
func IsValidProxyLogLevel(level string) bool {
	for _, l := range ValidProxyLogLevels {
		if level == l {
			return true
		}
	}
	return false
}

// ProxyLoggingService is a thin layer over the kube interface for proxy logging functions.
type ProxyLoggingService struct {
	k8s         kubernetes.ClientInterface
	proxyStatus *ProxyStatusService
}

// SetLogLevel sets the pod's proxy log level.
func (in *ProxyLoggingService) SetLogLevel(namespace, pod, level string) error {
	// Ensure pod exists
	if _, err := in.k8s.GetPod(namespace, pod); err != nil {
		return err
	}
	return in.k8s.SetProxyLogLevel(namespace, pod, level)
}
