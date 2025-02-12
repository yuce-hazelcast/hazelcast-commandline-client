/*
 * Copyright (c) 2008-2021, Hazelcast, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package it

import (
	"sync/atomic"
)

// from hazelcast-go-client/internal/proxy/reference_id_generator.go
type ReferenceIDGenerator struct {
	nextID int64
}

func NewReferenceIDGenerator(nextID int64) *ReferenceIDGenerator {
	return &ReferenceIDGenerator{nextID: nextID}
}

func (gen *ReferenceIDGenerator) NextID() int64 {
	return atomic.AddInt64(&gen.nextID, 1)
}
