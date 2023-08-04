// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/easysoft/quickon-api/client/clientset/versioned/typed/qucheng/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeQuchengV1beta1 struct {
	*testing.Fake
}

func (c *FakeQuchengV1beta1) Backups(namespace string) v1beta1.BackupInterface {
	return &FakeBackups{c, namespace}
}

func (c *FakeQuchengV1beta1) BackupSets(namespace string) v1beta1.BackupSetInterface {
	return &FakeBackupSets{c, namespace}
}

func (c *FakeQuchengV1beta1) Dbs(namespace string) v1beta1.DbInterface {
	return &FakeDbs{c, namespace}
}

func (c *FakeQuchengV1beta1) DbBackups(namespace string) v1beta1.DbBackupInterface {
	return &FakeDbBackups{c, namespace}
}

func (c *FakeQuchengV1beta1) DbDumpReplicas(namespace string) v1beta1.DbDumpReplicaInterface {
	return &FakeDbDumpReplicas{c, namespace}
}

func (c *FakeQuchengV1beta1) DbRestores(namespace string) v1beta1.DbRestoreInterface {
	return &FakeDbRestores{c, namespace}
}

func (c *FakeQuchengV1beta1) DbServices(namespace string) v1beta1.DbServiceInterface {
	return &FakeDbServices{c, namespace}
}

func (c *FakeQuchengV1beta1) DeleteBackupRequests(namespace string) v1beta1.DeleteBackupRequestInterface {
	return &FakeDeleteBackupRequests{c, namespace}
}

func (c *FakeQuchengV1beta1) ResticRepositories(namespace string) v1beta1.ResticRepositoryInterface {
	return &FakeResticRepositories{c, namespace}
}

func (c *FakeQuchengV1beta1) ResticSnapshotDumps(namespace string) v1beta1.ResticSnapshotDumpInterface {
	return &FakeResticSnapshotDumps{c, namespace}
}

func (c *FakeQuchengV1beta1) Restores(namespace string) v1beta1.RestoreInterface {
	return &FakeRestores{c, namespace}
}

func (c *FakeQuchengV1beta1) RestoreSets(namespace string) v1beta1.RestoreSetInterface {
	return &FakeRestoreSets{c, namespace}
}

func (c *FakeQuchengV1beta1) StorageProfiles(namespace string) v1beta1.StorageProfileInterface {
	return &FakeStorageProfiles{c, namespace}
}

func (c *FakeQuchengV1beta1) VolumeBackups(namespace string) v1beta1.VolumeBackupInterface {
	return &FakeVolumeBackups{c, namespace}
}

func (c *FakeQuchengV1beta1) VolumeRestores(namespace string) v1beta1.VolumeRestoreInterface {
	return &FakeVolumeRestores{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeQuchengV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
