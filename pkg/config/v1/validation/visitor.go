// Copyright 2023 The frp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"errors"
	"fmt"
	"net"
	"slices"

	v1 "github.com/imneov/kube-frp/pkg/config/v1"
)

func ValidateVisitorConfigurer(c v1.VisitorConfigurer) error {
	base := c.GetBaseConfig()
	if err := validateVisitorBaseConfig(base); err != nil {
		return err
	}

	switch v := c.(type) {
	case *v1.STCPVisitorConfig:
	case *v1.SUDPVisitorConfig:
	case *v1.XTCPVisitorConfig:
		return validateXTCPVisitorConfig(v)
	default:
		return errors.New("unknown visitor config type")
	}
	return nil
}

func validateVisitorBaseConfig(c *v1.VisitorBaseConfig) error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.ServerName == "" {
		return errors.New("server name is required")
	}

	if c.BindPort == 0 && !c.AutoAssignPort {
		return errors.New("bind port is required when auto assign port is disabled")
	}

	if c.IPRange != "" {
		// Validate IP range format (e.g., "172.1.1.0/24")
		_, _, err := net.ParseCIDR(c.IPRange)
		if err != nil {
			return fmt.Errorf("invalid IP range format: %v", err)
		}
	}

	return nil
}

func validateXTCPVisitorConfig(c *v1.XTCPVisitorConfig) error {
	if !slices.Contains([]string{"kcp", "quic"}, c.Protocol) {
		return fmt.Errorf("protocol should be kcp or quic")
	}
	return nil
}
