// This file is part of Graylog.
//
// Graylog is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Graylog is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Graylog.  If not, see <http://www.gnu.org/licenses/>.

package graylog

import (
	"github.com/Graylog2/collector-sidecar/common"
)

type RegistrationRequest struct {
	NodeName    string             `json:"node_name"`
	NodeDetails NodeDetailsRequest `json:"node_details"`
}

type NodeDetailsRequest struct {
	OperatingSystem string          `json:"operating_system"`
	IP              string          `json:"ip,omitempty"`
	LogFileList     []common.File   `json:"log_file_list,omitempty"`
	Metrics         *MetricsRequest `json:"metrics,omitempty"`
	Status          *StatusRequest  `json:"status,omitempty"`
}

type StatusRequestBackend struct {
	Id             string `json:"collector_id"`
	Status         int    `json:"status"`
	Message        string `json:"message"`
	VerboseMessage string `json:"verbose_message"`
}

type StatusRequest struct {
	Backends []StatusRequestBackend `json:"collectors"`
	Status   int                    `json:"status"`
	Message  string                 `json:"message"`
}

type MetricsRequest struct {
	Disks75 []string `json:"disks_75"`
	CpuIdle float64  `json:"cpu_idle"`
	Load1   float64  `json:"load_1"`
}
