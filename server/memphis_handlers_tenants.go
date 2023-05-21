// Copyright 2022-2023 The Memphis.dev Authors
// Licensed under the Memphis Business Source License 1.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// Changed License: [Apache License, Version 2.0 (https://www.apache.org/licenses/LICENSE-2.0), as published by the Apache Foundation.
//
// https://github.com/memphisdev/memphis/blob/master/LICENSE
//
// Additional Use Grant: You may make use of the Licensed Work (i) only as part of your own product or service, provided it is not a message broker or a message queue product or service; and (ii) provided that you do not use, provide, distribute, or make available the Licensed Work as a Service.
// A "Service" is a commercial offering, product, hosted, or managed service, that allows third parties (other than your own employees and contractors acting on your behalf) to access and/or use the Licensed Work or a substantial set of the features or functionality of the Licensed Work to third parties as a software-as-a-service, platform-as-a-service, infrastructure-as-a-service or other similar services that compete with Licensor products or services.
package server

import (
	"memphis/conf"
	"memphis/db"
)

func isGlobalTenantExist() (bool, error) {
	exist, _, err := db.GetGlobalTenant()
	if err != nil {
		return false, err
	} else if !exist {
		return false, nil
	}
	return true, nil
}

func CreateGlobalTenantOnFirstSystemLoad() error {
	exist, _, err := db.GetGlobalTenant()
	if err != nil {
		return err
	}

	if !exist {
		_, err := db.CreateTenant(conf.MEMPHIS_GLOBAL_ACCOUNT_NAME)
		if err != nil {
			return err
		}
	}
	return nil
}

type getTenantMsg struct {
	Acc string `json:"acc"`
	Rtt int    `json:"rtt"`
}

// TODO: remove
func (s *Server) getTenantName(c *client, reply string, msg []byte) {
	var resp getTenantNameResponse
	tenantName, _, err := s.getTenantNameAndMessage(msg)
	if err != nil {
		s.Errorf("getTenantName: " + err.Error())
		return
	}

	resp.TenantName = tenantName
	respondWithResp(conf.MEMPHIS_GLOBAL_ACCOUNT_NAME, s, reply, &resp)
}