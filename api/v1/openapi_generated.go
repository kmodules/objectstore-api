// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kmodules.xyz/objectstore-api/api/v1.AzureSpec":      schema_kmodulesxyz_objectstore_api_api_v1_AzureSpec(ref),
		"kmodules.xyz/objectstore-api/api/v1.B2Spec":         schema_kmodulesxyz_objectstore_api_api_v1_B2Spec(ref),
		"kmodules.xyz/objectstore-api/api/v1.Backend":        schema_kmodulesxyz_objectstore_api_api_v1_Backend(ref),
		"kmodules.xyz/objectstore-api/api/v1.GCSSpec":        schema_kmodulesxyz_objectstore_api_api_v1_GCSSpec(ref),
		"kmodules.xyz/objectstore-api/api/v1.LocalSpec":      schema_kmodulesxyz_objectstore_api_api_v1_LocalSpec(ref),
		"kmodules.xyz/objectstore-api/api/v1.RestServerSpec": schema_kmodulesxyz_objectstore_api_api_v1_RestServerSpec(ref),
		"kmodules.xyz/objectstore-api/api/v1.S3Spec":         schema_kmodulesxyz_objectstore_api_api_v1_S3Spec(ref),
		"kmodules.xyz/objectstore-api/api/v1.SwiftSpec":      schema_kmodulesxyz_objectstore_api_api_v1_SwiftSpec(ref),
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_AzureSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"container": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maxConnections": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_B2Spec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maxConnections": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_Backend(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"storageSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"local": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.LocalSpec"),
						},
					},
					"s3": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.S3Spec"),
						},
					},
					"gcs": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.GCSSpec"),
						},
					},
					"azure": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.AzureSpec"),
						},
					},
					"swift": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.SwiftSpec"),
						},
					},
					"b2": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.B2Spec"),
						},
					},
					"rest": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kmodules.xyz/objectstore-api/api/v1.RestServerSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"kmodules.xyz/objectstore-api/api/v1.AzureSpec", "kmodules.xyz/objectstore-api/api/v1.B2Spec", "kmodules.xyz/objectstore-api/api/v1.GCSSpec", "kmodules.xyz/objectstore-api/api/v1.LocalSpec", "kmodules.xyz/objectstore-api/api/v1.RestServerSpec", "kmodules.xyz/objectstore-api/api/v1.S3Spec", "kmodules.xyz/objectstore-api/api/v1.SwiftSpec"},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_GCSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maxConnections": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_LocalSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"hostPath": {
						SchemaProps: spec.SchemaProps{
							Description: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
							Ref:         ref("k8s.io/api/core/v1.HostPathVolumeSource"),
						},
					},
					"emptyDir": {
						SchemaProps: spec.SchemaProps{
							Description: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
							Ref:         ref("k8s.io/api/core/v1.EmptyDirVolumeSource"),
						},
					},
					"gcePersistentDisk": {
						SchemaProps: spec.SchemaProps{
							Description: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
							Ref:         ref("k8s.io/api/core/v1.GCEPersistentDiskVolumeSource"),
						},
					},
					"awsElasticBlockStore": {
						SchemaProps: spec.SchemaProps{
							Description: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
							Ref:         ref("k8s.io/api/core/v1.AWSElasticBlockStoreVolumeSource"),
						},
					},
					"gitRepo": {
						SchemaProps: spec.SchemaProps{
							Description: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
							Ref:         ref("k8s.io/api/core/v1.GitRepoVolumeSource"),
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
							Ref:         ref("k8s.io/api/core/v1.SecretVolumeSource"),
						},
					},
					"nfs": {
						SchemaProps: spec.SchemaProps{
							Description: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
							Ref:         ref("k8s.io/api/core/v1.NFSVolumeSource"),
						},
					},
					"iscsi": {
						SchemaProps: spec.SchemaProps{
							Description: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
							Ref:         ref("k8s.io/api/core/v1.ISCSIVolumeSource"),
						},
					},
					"glusterfs": {
						SchemaProps: spec.SchemaProps{
							Description: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
							Ref:         ref("k8s.io/api/core/v1.GlusterfsVolumeSource"),
						},
					},
					"persistentVolumeClaim": {
						SchemaProps: spec.SchemaProps{
							Description: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
							Ref:         ref("k8s.io/api/core/v1.PersistentVolumeClaimVolumeSource"),
						},
					},
					"rbd": {
						SchemaProps: spec.SchemaProps{
							Description: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
							Ref:         ref("k8s.io/api/core/v1.RBDVolumeSource"),
						},
					},
					"flexVolume": {
						SchemaProps: spec.SchemaProps{
							Description: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
							Ref:         ref("k8s.io/api/core/v1.FlexVolumeSource"),
						},
					},
					"cinder": {
						SchemaProps: spec.SchemaProps{
							Description: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
							Ref:         ref("k8s.io/api/core/v1.CinderVolumeSource"),
						},
					},
					"cephfs": {
						SchemaProps: spec.SchemaProps{
							Description: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
							Ref:         ref("k8s.io/api/core/v1.CephFSVolumeSource"),
						},
					},
					"flocker": {
						SchemaProps: spec.SchemaProps{
							Description: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
							Ref:         ref("k8s.io/api/core/v1.FlockerVolumeSource"),
						},
					},
					"downwardAPI": {
						SchemaProps: spec.SchemaProps{
							Description: "DownwardAPI represents downward API about the pod that should populate this volume",
							Ref:         ref("k8s.io/api/core/v1.DownwardAPIVolumeSource"),
						},
					},
					"fc": {
						SchemaProps: spec.SchemaProps{
							Description: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
							Ref:         ref("k8s.io/api/core/v1.FCVolumeSource"),
						},
					},
					"azureFile": {
						SchemaProps: spec.SchemaProps{
							Description: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
							Ref:         ref("k8s.io/api/core/v1.AzureFileVolumeSource"),
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMap represents a configMap that should populate this volume",
							Ref:         ref("k8s.io/api/core/v1.ConfigMapVolumeSource"),
						},
					},
					"vsphereVolume": {
						SchemaProps: spec.SchemaProps{
							Description: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
							Ref:         ref("k8s.io/api/core/v1.VsphereVirtualDiskVolumeSource"),
						},
					},
					"quobyte": {
						SchemaProps: spec.SchemaProps{
							Description: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
							Ref:         ref("k8s.io/api/core/v1.QuobyteVolumeSource"),
						},
					},
					"azureDisk": {
						SchemaProps: spec.SchemaProps{
							Description: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
							Ref:         ref("k8s.io/api/core/v1.AzureDiskVolumeSource"),
						},
					},
					"photonPersistentDisk": {
						SchemaProps: spec.SchemaProps{
							Description: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
							Ref:         ref("k8s.io/api/core/v1.PhotonPersistentDiskVolumeSource"),
						},
					},
					"projected": {
						SchemaProps: spec.SchemaProps{
							Description: "Items for all in one resources secrets, configmaps, and downward API",
							Ref:         ref("k8s.io/api/core/v1.ProjectedVolumeSource"),
						},
					},
					"portworxVolume": {
						SchemaProps: spec.SchemaProps{
							Description: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
							Ref:         ref("k8s.io/api/core/v1.PortworxVolumeSource"),
						},
					},
					"scaleIO": {
						SchemaProps: spec.SchemaProps{
							Description: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
							Ref:         ref("k8s.io/api/core/v1.ScaleIOVolumeSource"),
						},
					},
					"storageos": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
							Ref:         ref("k8s.io/api/core/v1.StorageOSVolumeSource"),
						},
					},
					"csi": {
						SchemaProps: spec.SchemaProps{
							Description: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
							Ref:         ref("k8s.io/api/core/v1.CSIVolumeSource"),
						},
					},
					"ephemeral": {
						SchemaProps: spec.SchemaProps{
							Description: "Ephemeral represents a volume that is handled by a cluster storage driver (Alpha feature). The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.\n\nUse this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity\n   tracking are needed,\nc) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through\n   a PersistentVolumeClaim (see EphemeralVolumeSource for more\n   information on the connection between this volume type\n   and PersistentVolumeClaim).\n\nUse PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.\n\nUse CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.\n\nA pod can use both types of ephemeral volumes and persistent volumes at the same time.",
							Ref:         ref("k8s.io/api/core/v1.EphemeralVolumeSource"),
						},
					},
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"subPath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.AWSElasticBlockStoreVolumeSource", "k8s.io/api/core/v1.AzureDiskVolumeSource", "k8s.io/api/core/v1.AzureFileVolumeSource", "k8s.io/api/core/v1.CSIVolumeSource", "k8s.io/api/core/v1.CephFSVolumeSource", "k8s.io/api/core/v1.CinderVolumeSource", "k8s.io/api/core/v1.ConfigMapVolumeSource", "k8s.io/api/core/v1.DownwardAPIVolumeSource", "k8s.io/api/core/v1.EmptyDirVolumeSource", "k8s.io/api/core/v1.EphemeralVolumeSource", "k8s.io/api/core/v1.FCVolumeSource", "k8s.io/api/core/v1.FlexVolumeSource", "k8s.io/api/core/v1.FlockerVolumeSource", "k8s.io/api/core/v1.GCEPersistentDiskVolumeSource", "k8s.io/api/core/v1.GitRepoVolumeSource", "k8s.io/api/core/v1.GlusterfsVolumeSource", "k8s.io/api/core/v1.HostPathVolumeSource", "k8s.io/api/core/v1.ISCSIVolumeSource", "k8s.io/api/core/v1.NFSVolumeSource", "k8s.io/api/core/v1.PersistentVolumeClaimVolumeSource", "k8s.io/api/core/v1.PhotonPersistentDiskVolumeSource", "k8s.io/api/core/v1.PortworxVolumeSource", "k8s.io/api/core/v1.ProjectedVolumeSource", "k8s.io/api/core/v1.QuobyteVolumeSource", "k8s.io/api/core/v1.RBDVolumeSource", "k8s.io/api/core/v1.ScaleIOVolumeSource", "k8s.io/api/core/v1.SecretVolumeSource", "k8s.io/api/core/v1.StorageOSVolumeSource", "k8s.io/api/core/v1.VsphereVirtualDiskVolumeSource"},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_RestServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_S3Spec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_objectstore_api_api_v1_SwiftSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"container": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}
