// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package fftypes

import (
	"context"
	"crypto/sha256"
	"database/sql/driver"
	"encoding/json"

	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/log"
)

// JSONObject is a holder of a hash, that can be used to correlate onchain data with off-chain data.
type JSONObject map[string]interface{}

// Scan implements sql.Scanner
func (jd *JSONObject) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string, []byte:
		if src == "" {
			return nil
		}
		return json.Unmarshal(src.([]byte), &jd)

	default:
		return i18n.NewError(context.Background(), i18n.MsgScanFailed, src, jd)
	}

}

func (jd JSONObject) GetString(ctx context.Context, key string) string {
	s, _ := jd.GetStringOk(ctx, key)
	return s
}

func (jd JSONObject) GetStringOk(ctx context.Context, key string) (string, bool) {
	vInterace, ok := jd[key]
	if ok && vInterace != nil {
		if vString, ok := vInterace.(string); ok {
			return vString, true
		}
	}
	log.L(ctx).Errorf("Invalid string value '%+v' for key '%s'", vInterace, key)
	return "", false
}

func (jd JSONObject) GetObject(ctx context.Context, key string) JSONObject {
	ob, _ := jd.GetObjectOk(ctx, key)
	return ob
}

func (jd JSONObject) GetObjectOk(ctx context.Context, key string) (JSONObject, bool) {
	vInterace, ok := jd[key]
	if ok && vInterace != nil {
		if vMap, ok := vInterace.(map[string]interface{}); ok {
			return JSONObject(vMap), true
		}
	}
	log.L(ctx).Errorf("Invalid object value '%+v' for key '%s'", vInterace, key)
	return JSONObject{}, false // Ensures a non-nil return
}

func ToJSONObjectArray(unknown interface{}) (JSONObjectArray, bool) {
	vMap, ok := unknown.([]interface{})
	joa := make(JSONObjectArray, len(vMap))
	if !ok {
		joa, ok = unknown.(JSONObjectArray) // Case that we're passed a JSONObjectArray directly
	}
	for i, joi := range vMap {
		jo, childOK := joi.(map[string]interface{})
		if childOK {
			joa[i] = JSONObject(jo)
		}
		ok = ok && childOK
	}
	return joa, ok
}

func (jd JSONObject) GetObjectArray(ctx context.Context, key string) JSONObjectArray {
	oa, _ := jd.GetObjectArrayOk(ctx, key)
	return oa
}

func (jd JSONObject) GetObjectArrayOk(ctx context.Context, key string) (JSONObjectArray, bool) {
	vInterace, ok := jd[key]
	if ok && vInterace != nil {
		return ToJSONObjectArray(vInterace)
	}
	log.L(ctx).Errorf("Invalid object value '%+v' for key '%s'", vInterace, key)
	return JSONObjectArray{}, false // Ensures a non-nil return
}

// Value implements sql.Valuer
func (jd JSONObject) Value() (driver.Value, error) {
	return json.Marshal(&jd)
}

func (jd JSONObject) String() string {
	b, _ := json.Marshal(&jd)
	return string(b)
}

func (jd JSONObject) Hash(ctx context.Context, jsonDesc string) (*Bytes32, error) {
	b, err := json.Marshal(&jd)
	if err != nil {
		return nil, i18n.NewError(ctx, i18n.MsgJSONObjectParseFailed, jsonDesc)
	}
	var b32 Bytes32 = sha256.Sum256(b)
	return &b32, nil
}

// JSONObjectArray is an array of JSONObject
type JSONObjectArray []JSONObject

// Scan implements sql.Scanner
func (jd *JSONObjectArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string, []byte:
		if src == "" {
			return nil
		}
		return json.Unmarshal(src.([]byte), &jd)

	default:
		return i18n.NewError(context.Background(), i18n.MsgScanFailed, src, jd)
	}

}

func (jd JSONObjectArray) Value() (driver.Value, error) {
	return json.Marshal(&jd)
}

func (jd JSONObjectArray) String() string {
	b, _ := json.Marshal(&jd)
	return string(b)
}

func (jd JSONObjectArray) Hash(ctx context.Context, jsonDesc string) (*Bytes32, error) {
	b, err := json.Marshal(&jd)
	if err != nil {
		return nil, i18n.NewError(ctx, i18n.MsgJSONObjectParseFailed, jsonDesc)
	}
	var b32 Bytes32 = sha256.Sum256(b)
	return &b32, nil
}