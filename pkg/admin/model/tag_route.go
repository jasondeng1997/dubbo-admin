/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

type TagRouteDto struct {
	Base

	Tags []Tag `json:"tags" binding:"required"`

	Priority      int    `json:"priority"`
	Enabled       bool   `json:"enabled" binding:"required"`
	Force         bool   `json:"force"`
	Runtime       bool   `json:"runtime"`
	ConfigVersion string `json:"configVersion" binding:"required"`
}

type TagRoute struct {
	Priority int    `json:"priority"`
	Enabled  bool   `json:"enabled"`
	Force    bool   `json:"force"`
	Runtime  bool   `json:"runtime"`
	Key      string `json:"key"`
	Tags     []Tag  `json:"tags"`
}

type Tag struct {
	Name      string       `json:"name" binding:"required"`
	Match     []ParamMatch `json:"match" binding:"required"`
	Addresses []string     `json:"addresses"`
}
