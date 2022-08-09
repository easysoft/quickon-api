// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/easysoft/quikon-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Backups returns a BackupInformer.
	Backups() BackupInformer
	// Dbs returns a DbInformer.
	Dbs() DbInformer
	// DbBackups returns a DbBackupInformer.
	DbBackups() DbBackupInformer
	// DbRestores returns a DbRestoreInformer.
	DbRestores() DbRestoreInformer
	// DbServices returns a DbServiceInformer.
	DbServices() DbServiceInformer
	// GlobalDBs returns a GlobalDBInformer.
	GlobalDBs() GlobalDBInformer
	// Restores returns a RestoreInformer.
	Restores() RestoreInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Backups returns a BackupInformer.
func (v *version) Backups() BackupInformer {
	return &backupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Dbs returns a DbInformer.
func (v *version) Dbs() DbInformer {
	return &dbInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DbBackups returns a DbBackupInformer.
func (v *version) DbBackups() DbBackupInformer {
	return &dbBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DbRestores returns a DbRestoreInformer.
func (v *version) DbRestores() DbRestoreInformer {
	return &dbRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DbServices returns a DbServiceInformer.
func (v *version) DbServices() DbServiceInformer {
	return &dbServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GlobalDBs returns a GlobalDBInformer.
func (v *version) GlobalDBs() GlobalDBInformer {
	return &globalDBInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Restores returns a RestoreInformer.
func (v *version) Restores() RestoreInformer {
	return &restoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}