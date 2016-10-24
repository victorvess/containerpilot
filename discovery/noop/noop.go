package noop

import (
    "fmt"
    "os"
	"github.com/joyent/containerpilot/discovery"
)

func init() {
	discovery.RegisterBackend("noop", ConfigHook)
}

// Noop is a bogus service discovery backend
type Noop struct {
}


// ConfigHook is the hook to register with the Noop backend
func ConfigHook(raw interface{}) (discovery.ServiceBackend, error) {
    //fmt.Println("In ConfigHook (Noop)")
    //os.Setenv("MOO", "MOO_CP")
    return new(Noop), nil
}


// Deregister removes this instance from the registry
func (c *Noop) Deregister(service *discovery.ServiceDefinition) {
}

// MarkForMaintenance removes this instance from the registry
func (c *Noop) MarkForMaintenance(service *discovery.ServiceDefinition) {
}

// SendHeartbeat refreshes the TTL of this associated etcd node
func (c *Noop) SendHeartbeat(service *discovery.ServiceDefinition) {
}

// CheckForUpstreamChanges checks another etcd node for changes
func (c *Noop) CheckForUpstreamChanges(backendName, backendTag string) bool {
    return false
}


