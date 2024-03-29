// Copyright (c) 2022-2022 北京渠成软件有限公司(Beijing Qucheng Software Co., Ltd. www.qucheng.com) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Z PUBLIC LICENSE 1.2 (ZPL 1.2)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// license that can be found in the LICENSE file.

package v1beta1

const (
	BackupNameLabel  = "easycorp.io/backup_name"
	RestoreNameLabel = "easycorp.io/restore_name"

	BackupSetNameLabel  = "easycorp.io/backup_set_name"
	RestoreSetNameLabel = "easycorp.io/restore_set_name"

	ApplicationNameLabel    = "easycorp.io/app_release_name"
	StorageProfileRoleLabel = "easycorp.io/storage_role"

	PvcNameLabel = "easycorp.io/pvc_name"
)

const (
	StorageProfileRoleVolume = "volume"
	StorageProfileRoleDB     = "db"
	StorageProfileRoleRemote = "remote"
)

const (
	SelectorReleaseKey = "release"
)

const (
	PvcBackupExcludeAnnotation = "easycorp.io/pvc_backup_exclude"
	BslRemoteDumpWeight        = "easycorp.io/remote_dump_weight"
)
