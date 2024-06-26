package packetcaptures

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCaptureSettings struct {
	FileCount                 *int64 `json:"fileCount,omitempty"`
	FileSizeInBytes           *int64 `json:"fileSizeInBytes,omitempty"`
	SessionTimeLimitInSeconds *int64 `json:"sessionTimeLimitInSeconds,omitempty"`
}
