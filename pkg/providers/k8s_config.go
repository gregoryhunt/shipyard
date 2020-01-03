package providers

import (
	"github.com/shipyard-run/shipyard/pkg/clients"
	"github.com/shipyard-run/shipyard/pkg/config"
	"golang.org/x/xerrors"
)

type K8sConfig struct {
	config *config.K8sConfig
	client clients.Kubernetes
}

// NewK8sConfig creates a provider which can create and destroy kubernetes configuration
func NewK8sConfig(c *config.K8sConfig, kc clients.Kubernetes) *K8sConfig {
	return &K8sConfig{c, kc}
}

// Create the Kubernetes resources defined by the config
func (c *K8sConfig) Create() error {
	err := c.setup()
	if err != nil {
		return err
	}

	return c.client.Apply(c.config.Paths, c.config.WaitUntilReady)
}

// Destroy the Kubernetes resources defined by the config
func (c *K8sConfig) Destroy() error {
	err := c.setup()
	if err != nil {
		return err
	}

	return c.client.Delete(c.config.Paths)
}

// Lookup the Kubernetes resources defined by the config
func (c *K8sConfig) Lookup() (string, error) {
	return "", nil
}

func (c *K8sConfig) setup() error {
	_, destPath, _ := CreateKubeConfigPath(c.config.ClusterRef.Name)
	err := c.client.SetConfig(destPath)
	if err != nil {
		return xerrors.Errorf("unable to create Kubernetes client: %w", err)
	}

	return nil
}
