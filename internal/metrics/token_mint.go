// Copyright © 2025 Kaleido, Inc.
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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var MintSubmittedCounter *prometheus.CounterVec
var MintConfirmedCounter *prometheus.CounterVec
var MintRejectedCounter *prometheus.CounterVec
var MintHistogram *prometheus.HistogramVec

// MintSubmittedCounterName is the prometheus metric for tracking the total number of mints submitted
var MintSubmittedCounterName = "ff_mint_submitted_total"

// MintConfirmedCounterName is the prometheus metric for tracking the total number of mints confirmed
var MintConfirmedCounterName = "ff_mint_confirmed_total"

// MintRejectedCounterName is the prometheus metric for tracking the total number of mints rejected
var MintRejectedCounterName = "ff_mint_rejected_total"

// MintHistogramName is the prometheus metric for tracking the total number of mints - histogram
var MintHistogramName = "ff_mint_histogram"

func InitTokenMintMetrics() {
	MintSubmittedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: MintSubmittedCounterName,
		Help: "Number of submitted mints",
	}, namespaceLabels)
	MintConfirmedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: MintConfirmedCounterName,
		Help: "Number of confirmed mints",
	}, namespaceLabels)
	MintRejectedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: MintRejectedCounterName,
		Help: "Number of rejected mints",
	}, namespaceLabels)
	MintHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: MintHistogramName,
		Help: "Histogram of mints, bucketed by time to finished",
	}, namespaceLabels)
}

func RegisterTokenMintMetrics() {
	registry.MustRegister(MintSubmittedCounter)
	registry.MustRegister(MintConfirmedCounter)
	registry.MustRegister(MintRejectedCounter)
	registry.MustRegister(MintHistogram)
}
