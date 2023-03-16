/*
Copyright © 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package service

import (
	"context"
	"time"

	"github.com/frabits/frabit/common/log"
	pxb "github.com/frabits/frabit/pkg/xtrabackup"

	"go.uber.org/zap"
)

type BackupType string

const (
	Logical  BackupType = "logical"
	Physical BackupType = "physical"
)

// BackupService implement DB backup task
type BackupService struct {
	StartDatetime time.Time
	EndDatetime   time.Time
	PXB           pxb.Xtrabackup
	Type          BackupType
	Logger        *zap.Logger
}

func newBackupService() *BackupService {
	log.Info("create BackupService")
	return &BackupService{
		Logger: log.Logger,
	}
}

// Start a backup job
func (bak *BackupService) Start(ctx context.Context, bakType BackupType) error {
	bak.Logger.Info("create BackupService")
	bak.Type = bakType
	return nil
}

func (bak *BackupService) Stop() error {

	return nil
}

func (bak *BackupService) Cancel() error {

	return nil
}

func init() {
	newBackupService()
}
