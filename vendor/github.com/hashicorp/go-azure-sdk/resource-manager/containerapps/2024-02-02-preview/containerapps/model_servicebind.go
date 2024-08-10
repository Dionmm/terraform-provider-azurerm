package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceBind struct {
	ClientType     *string            `json:"clientType,omitempty"`
	CustomizedKeys *map[string]string `json:"customizedKeys,omitempty"`
	Name           *string            `json:"name,omitempty"`
	ServiceId      *string            `json:"serviceId,omitempty"`
}
