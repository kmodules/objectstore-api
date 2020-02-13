/*
Copyright The Kmodules Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	expectedMaxConnection int64
	expectedLocation      string
	expectedProvider      string
	expectedEndpoint      string
	expectedRegion        string
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
			expectedRegion:        "",
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
			expectedRegion:        "",
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
			expectedRegion:        "",
		},
		{
			name: ProviderS3 + "_with_region",
			backend: Backend{
				S3: &S3Spec{
					Bucket:   "stash-backup",
					Prefix:   "/source/data",
					Endpoint: "s3.amazonaws.com",
					Region:   "my.custom.region",
				},
				StorageSecretName: "s3-secret",
			},
			expectedContainer:     "stash-backup",
			expectedLocation:      fmt.Sprintf("%s:%s", ProviderS3, "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderS3,
			expectedMaxConnection: 0,
			expectedEndpoint:      "s3.amazonaws.com",
			expectedRegion:        "my.custom.region",
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
			expectedLocation:      fmt.Sprintf("%s:%s", "gs", "stash-backup"),
			expectedPrefix:        "/source/data",
			expectedProvider:      ProviderGCS,
			expectedMaxConnection: 2,
			expectedEndpoint:      "",
			expectedRegion:        "",
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
			expectedRegion:        "",
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
			expectedRegion:        "",
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
			expectedRegion:        "",
		},
	}

}

func TestBackend_Container(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			container, err := tt.backend.Container()
			if err != nil {
				t.Errorf("fail to get container, reason: %v", err)
				return
			}
			if container != tt.expectedContainer {
				t.Errorf("expected Container: %v, found: %v", tt.expectedContainer, container)
			}
		})
	}
}

func TestBackend_Location(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			location, err := tt.backend.Location()
			if err != nil && tt.backend.Rest == nil {
				t.Errorf("fail to get location, reason: %v", err)
				return
			}
			if location != tt.expectedLocation {
				t.Errorf("expected Location: %v, found: %v", tt.expectedLocation, location)
			}
		})
	}
}

func TestBackend_Prefix(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			prefix, err := tt.backend.Prefix()
			if err != nil {
				t.Errorf("fail to get prefix, reason: %v", err)
				return
			}
			if prefix != tt.expectedPrefix {
				t.Errorf("expected prefix: %v, found: %v", tt.expectedPrefix, prefix)
			}
		})
	}
}

func TestBackend_Provider(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := tt.backend.Provider()
			if err != nil {
				t.Errorf("fail to get provider, reason: %v", err)
				return
			}
			if provider != tt.expectedProvider {
				t.Errorf("expected provider: %v, found: %v", tt.expectedProvider, provider)
			}
		})
	}
}

func TestBackend_MaxConnections(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			maxConnection := tt.backend.MaxConnections()
			if maxConnection != tt.expectedMaxConnection {
				t.Errorf("expected maxconnection: %v, found: %v", tt.expectedMaxConnection, maxConnection)
			}
		})
	}
}

func TestBackend_Endpoint(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			endpoint, ok := tt.backend.Endpoint()
			if !ok && (tt.backend.S3 != nil || tt.backend.Rest != nil) {
				t.Errorf("fail to get endpoint")
				return
			}
			if endpoint != tt.expectedEndpoint {
				t.Errorf("expected endpoint: %v, found: %v", tt.expectedEndpoint, endpoint)
				return
			}
		})
	}
}

func TestBackend_Region(t *testing.T) {
	for _, tt := range testCases() {
		t.Run(tt.name, func(t *testing.T) {
			region, ok := tt.backend.Region()
			if !ok && tt.backend.S3 != nil {
				t.Errorf("fail to get region")
				return
			}
			if region != tt.expectedRegion {
				t.Errorf("expected region: %v, found: %v", tt.expectedRegion, region)
				return
			}
		})
	}
}
