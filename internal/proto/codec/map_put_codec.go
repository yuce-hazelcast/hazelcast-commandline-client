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
package codec

import (
	iserialization "github.com/hazelcast/hazelcast-go-client"
	proto "github.com/hazelcast/hazelcast-go-client"
)

const (
	// hex: 0x010100
	MapPutCodecRequestMessageType = int32(65792)
	// hex: 0x010101
	MapPutCodecResponseMessageType = int32(65793)

	MapPutCodecRequestThreadIdOffset   = proto.PartitionIDOffset + proto.IntSizeInBytes
	MapPutCodecRequestTtlOffset        = MapPutCodecRequestThreadIdOffset + proto.LongSizeInBytes
	MapPutCodecRequestInitialFrameSize = MapPutCodecRequestTtlOffset + proto.LongSizeInBytes
)

// Puts an entry into this map with a given ttl (time to live) value.Entry will expire and get evicted after the ttl
// If ttl is 0, then the entry lives forever.This method returns a clone of the previous value, not the original
// (identically equal) value previously put into the map.Time resolution for TTL is seconds. The given TTL value is
// rounded to the next closest second value.

func EncodeMapPutRequest(name string, key iserialization.Data, value iserialization.Data, threadId int64, ttl int64) *proto.ClientMessage {
	clientMessage := proto.NewClientMessageForEncode()
	clientMessage.SetRetryable(false)

	initialFrame := proto.NewFrameWith(make([]byte, MapPutCodecRequestInitialFrameSize), proto.UnfragmentedMessage)
	EncodeLong(initialFrame.Content, MapPutCodecRequestThreadIdOffset, threadId)
	EncodeLong(initialFrame.Content, MapPutCodecRequestTtlOffset, ttl)
	clientMessage.AddFrame(initialFrame)
	clientMessage.SetMessageType(MapPutCodecRequestMessageType)
	clientMessage.SetPartitionId(-1)

	EncodeString(clientMessage, name)
	EncodeData(clientMessage, key)
	EncodeData(clientMessage, value)

	return clientMessage
}

func DecodeMapPutResponse(clientMessage *proto.ClientMessage) iserialization.Data {
	frameIterator := clientMessage.FrameIterator()
	// empty initial frame
	frameIterator.Next()

	return DecodeNullableForData(frameIterator)
}
