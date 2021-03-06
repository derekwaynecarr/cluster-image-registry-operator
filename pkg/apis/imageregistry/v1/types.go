package v1

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operatorsv1api "github.com/openshift/api/operator/v1"
)

const (
	OperatorStatusTypeRemoved             = "Removed"
	ImageRegistryPrivateConfiguration     = "image-registry-private-configuration"
	ImageRegistryPrivateConfigurationUser = ImageRegistryPrivateConfiguration + "-user"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ImageRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ImageRegistry `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ImageRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ImageRegistrySpec   `json:"spec"`
	Status            ImageRegistryStatus `json:"status,omitempty"`
}

type ImageRegistryConfigProxy struct {
	HTTP    string `json:"http,omitempty"`
	HTTPS   string `json:"https,omitempty"`
	NoProxy string `json:"noProxy,omitempty"`
}

type ImageRegistryConfigStorageS3 struct {
	Bucket         string `json:"bucket,omitempty"`
	Region         string `json:"region,omitempty"`
	RegionEndpoint string `json:"regionEndpoint,omitempty"`
	Encrypt        bool   `json:"encrypt,omitempty"`
}

type ImageRegistryConfigStorageAzure struct {
	Container string `json:"container,omitempty"`
}

type ImageRegistryConfigStorageGCS struct {
	Bucket string `json:"bucket,omitempty"`
}

type ImageRegistryConfigStorageSwift struct {
	AuthURL   string `json:"authURL,omitempty"`
	Container string `json:"container,omitempty"`
}

type ImageRegistryConfigStorageFilesystem struct {
	VolumeSource corev1.VolumeSource `json:"volumeSource,omitempty"`
}

type ImageRegistryConfigStorage struct {
	Azure      *ImageRegistryConfigStorageAzure      `json:"azure,omitempty"`
	Filesystem *ImageRegistryConfigStorageFilesystem `json:"filesystem,omitempty"`
	GCS        *ImageRegistryConfigStorageGCS        `json:"gcs,omitempty"`
	S3         *ImageRegistryConfigStorageS3         `json:"s3,omitempty"`
	Swift      *ImageRegistryConfigStorageSwift      `json:"swift,omitempty"`
}

type ImageRegistryConfigRequestsLimits struct {
	MaxRunning     int           `json:"maxrunning,omitempty"`
	MaxInQueue     int           `json:"maxinqueue,omitempty"`
	MaxWaitInQueue time.Duration `json:"maxwaitinqueue,omitempty"`
}

type ImageRegistryConfigRequests struct {
	Read  ImageRegistryConfigRequestsLimits `json:"read,omitempty"`
	Write ImageRegistryConfigRequestsLimits `json:"write,omitempty"`
}

type ImageRegistryConfigRoute struct {
	Name       string `json:"name"`
	Hostname   string `json:"hostname"`
	SecretName string `json:"secretName"`
}

type ImageRegistrySpec struct {
	ManagementState operatorsv1api.ManagementState `json:"managementState"`
	HTTPSecret      string                         `json:"httpSecret,omitempty"`
	Proxy           ImageRegistryConfigProxy       `json:"proxy,omitempty"`
	Storage         ImageRegistryConfigStorage     `json:"storage,omitempty"`
	Requests        ImageRegistryConfigRequests    `json:"requests,omitempty"`
	TLS             bool                           `json:"tls,omitempty"`
	CAConfigName    string                         `json:"caConfigName,omitempty"`
	DefaultRoute    bool                           `json:"defaultRoute,omitempty"`
	Routes          []ImageRegistryConfigRoute     `json:"routes,omitempty"`
	Replicas        int32                          `json:"replicas,omitempty"`
	LogLevel        int64                          `json:"logging,omitempty"`
}

type ImageRegistryConfigStorageState struct {
	ImageRegistryConfigStorage `json:",inline"`
}

type ImageRegistryConfigStorageStatus struct {
	// Managed is a boolean which denotes whether or not
	// we created the registry storage medium (such as an
	// S3 bucket
	Managed bool `json:"managed"`

	// State is the current state of the registry
	// storage configuration
	State ImageRegistryConfigStorageState `json:"state"`
}

const (
	// StorageExists denotes whether or not the registry storage medium exists
	StorageExists = "StorageExists"

	// StorageTagged denotes whether or not the registry storage medium
	// that we created was tagged correctly
	StorageTagged = "StorageTagged"

	// StorageEncrypted denotes whether or not the registry storage medium
	// that we created has encryption enabled
	StorageEncrypted = "StorageEncrypted"

	// StorageIncompleteUploadCleanupEnabled denotes whethere or not the registry storage
	// medium is configured to automatically cleanup incomplete uploads
	StorageIncompleteUploadCleanupEnabled = "StorageIncompleteUploadCleanupEnabled"
)

type ImageRegistryStatus struct {
	operatorsv1api.OperatorStatus `json:",inline"`

	InternalRegistryHostname string `json:"internalRegistryHostname"`

	Storage ImageRegistryConfigStorageStatus `json:"storage"`
}
