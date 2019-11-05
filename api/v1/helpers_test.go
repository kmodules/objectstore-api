package v1

import (
	"fmt"
	"testing"

	core "k8s.io/api/core/v1"
)

type test struct {
	name                  string
	backend               Backend
	expectedContainer     string
	expectedPrefix        string
	expectedMaxConnection int
	expectedLocation      string
	expectedProvider      string
	expectedEndpoint      string
	expected              bool
}

var testCases = func() []test {
	return []test{
		{
			name: ProviderB2,
			backend: Backend{
				B2: &B2Spec{
					Bucket:         "stash-backup",
					Prefix:         "/source/data",
					MaxConnections: 2,
				},
				StorageSecretName: "b2-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderB2, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderB2,
			expectedMaxConnection: 2,
			expectedEndpoint:      "",
		},
		{
			name: ProviderLocal,
			backend: Backend{
				Local: &LocalSpec{
					MountPath: "/safe/data",
					VolumeSource: core.VolumeSource{
						HostPath: &core.HostPathVolumeSource{
							Path: "/data/stash-test",
						},
					},
					SubPath: "/stash/backup",
				},
				StorageSecretName: "local-secret",
			},
			expectedContainer:     "/safe/data",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderLocal, "/safe/data"),
			expectedPrefix:        "",
			expectedProvider:      ProviderLocal,
			expectedMaxConnection: 0,
			expectedEndpoint:      "",
		},
		{
			name: ProviderS3,
			backend: Backend{
				S3: &S3Spec{
					Bucket:   "stash-backup",
					Prefix:   "/source/data",
					Endpoint: "s3.amazonaws.com",
				},
				StorageSecretName: "s3-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderS3, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderS3,
			expectedMaxConnection: 0,
			expectedEndpoint:      "s3.amazonaws.com",
		},
		{
			name: ProviderGCS,
			backend: Backend{
				GCS: &GCSSpec{
					Bucket:         "stash-backup",
					Prefix:         "/source/data",
					MaxConnections: 2,
				},
				StorageSecretName: "gcs-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderGCS, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderGCS,
			expectedMaxConnection: 2,
			expectedEndpoint:      "",
		},
		{
			name: ProviderAzure,
			backend: Backend{
				Azure: &AzureSpec{
					Container:      "stash-backup",
					Prefix:         "/source/data",
					MaxConnections: 2,
				},
				StorageSecretName: "azure-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderAzure, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderAzure,
			expectedMaxConnection: 2,
			expectedEndpoint:      "",
		},
		{
			name: ProviderSwift,
			backend: Backend{
				Swift: &SwiftSpec{
					Container: "stash-backup",
					Prefix:    "/source/data",
				},
				StorageSecretName: "swift-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderSwift, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderSwift,
			expectedMaxConnection: 0,
			expectedEndpoint:      "",
		},
		{
			name: ProviderRest,
			backend: Backend{
				Rest: &RestServerSpec{
					URL: "http://rest-server.demo.svc:8000/stash-backup",
				},
				StorageSecretName: "rest-secret",
			},
			expectedContainer:     "rest-server.demo.svc:8000",
			expectedLocation:      "",
			expectedPrefix:        "/stash-backup",
			expectedProvider:      ProviderRest,
			expectedMaxConnection: 0,
			expectedEndpoint:      "http://rest-server.demo.svc:8000/stash-backup",
		},
	}

}

func TestBackend_Container(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got, err := backend.Container()
			if (err != nil) != tt.expected {
				t.Errorf("Backend.Container() error = %v, expectedErr %v", err, tt.expected)
				return
			}
			if got != tt.expectedContainer {
				t.Errorf("Backend.Container() = %v, expected result %v", got, tt.expectedContainer)
			}
		})
	}
}

func TestBackend_Location(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got, err := backend.Location()
			if (err != nil) != tt.expected && backend.Rest == nil {
				t.Errorf("Backend.Location() error = %v, expectedErr %v", err, tt.expected)
				return
			}

			if got != tt.expectedLocation {
				t.Errorf("Backend.Location() = %v, expected result %v", got, tt.expectedLocation)
			}
		})
	}
}

func TestBackend_Prefix(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got, err := backend.Prefix()
			if (err != nil) != tt.expected {
				t.Errorf("Backend.Prefix() error = %v, expectedErr %v", err, tt.expected)
				return
			}
			if got != tt.expectedPrefix {
				t.Errorf("Backend.Prefix() = %v, expected result %v", got, tt.expectedPrefix)
			}
		})
	}
}

func TestBackend_Provider(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got, err := backend.Provider()
			if (err != nil) != tt.expected {
				t.Errorf("Backend.Provider() error = %v, expectedErr %v", err, tt.expected)
				return
			}
			if got != tt.expectedProvider {
				t.Errorf("Backend.Provider() = %v, expected  result %v", got, tt.expectedProvider)
			}
		})
	}
}

func TestBackend_MaxConnections(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got := backend.MaxConnections()
			if got != tt.expectedMaxConnection {
				t.Errorf("Backend.MaxConnection() = %v, expected result %v", got, tt.expectedMaxConnection)
			}
		})
	}
}

func TestBackend_Endpoint(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			backend := Backend{
				StorageSecretName: tt.backend.StorageSecretName,
				Local:             tt.backend.Local,
				S3:                tt.backend.S3,
				GCS:               tt.backend.GCS,
				Azure:             tt.backend.Azure,
				Swift:             tt.backend.Swift,
				B2:                tt.backend.B2,
				Rest:              tt.backend.Rest,
			}
			got, ok := backend.Endpoint()
			if !ok != tt.expected && (backend.S3 != nil || backend.Rest != nil) {
				t.Errorf("Backend.Endpoint() got = %v, expected %v", ok, tt.expected)
				return
			}
			if got != tt.expectedEndpoint {
				t.Errorf("Backend.MaxConnection() = %v, expected result %v", got, tt.expectedEndpoint)
				return
			}
		})
	}
}
