// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	"github.com/easysoft/quickon-api/client/clientset/versioned/scheme"
	v1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	rest "k8s.io/client-go/rest"
)

type QuchengV1beta1Interface interface {
	RESTClient() rest.Interface
	BackupsGetter
	DbsGetter
	DbBackupsGetter
	DbRestoresGetter
	DbServicesGetter
	DeleteBackupRequestsGetter
	GlobalDBsGetter
	RestoresGetter
}

// QuchengV1beta1Client is used to interact with features provided by the qucheng group.
type QuchengV1beta1Client struct {
	restClient rest.Interface
}

func (c *QuchengV1beta1Client) Backups(namespace string) BackupInterface {
	return newBackups(c, namespace)
}

func (c *QuchengV1beta1Client) Dbs(namespace string) DbInterface {
	return newDbs(c, namespace)
}

func (c *QuchengV1beta1Client) DbBackups(namespace string) DbBackupInterface {
	return newDbBackups(c, namespace)
}

func (c *QuchengV1beta1Client) DbRestores(namespace string) DbRestoreInterface {
	return newDbRestores(c, namespace)
}

func (c *QuchengV1beta1Client) DbServices(namespace string) DbServiceInterface {
	return newDbServices(c, namespace)
}

func (c *QuchengV1beta1Client) DeleteBackupRequests(namespace string) DeleteBackupRequestInterface {
	return newDeleteBackupRequests(c, namespace)
}

func (c *QuchengV1beta1Client) GlobalDBs(namespace string) GlobalDBInterface {
	return newGlobalDBs(c, namespace)
}

func (c *QuchengV1beta1Client) Restores(namespace string) RestoreInterface {
	return newRestores(c, namespace)
}

// NewForConfig creates a new QuchengV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*QuchengV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new QuchengV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*QuchengV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &QuchengV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new QuchengV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *QuchengV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new QuchengV1beta1Client for the given RESTClient.
func New(c rest.Interface) *QuchengV1beta1Client {
	return &QuchengV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *QuchengV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
