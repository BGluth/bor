// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package state

import "github.com/ethereum/go-ethereum/metrics"

var (
	accountUpdatedMeter   = metrics.NewRegisteredMeter("state/update/account", nil)
	storageUpdatedMeter   = metrics.NewRegisteredMeter("state/update/storage", nil)
	accountDeletedMeter   = metrics.NewRegisteredMeter("state/delete/account", nil)
	storageDeletedMeter   = metrics.NewRegisteredMeter("state/delete/storage", nil)
	accountCommittedMeter = metrics.NewRegisteredMeter("state/commit/account", nil)
	storageCommittedMeter = metrics.NewRegisteredMeter("state/commit/storage", nil)

	StorageCacheTotHitMeter   = metrics.NewRegisteredMeter("state/extra_stats/cache_hit_total", nil)
	StorageSnapHitMeter       = metrics.NewRegisteredMeter("state/extra_stats/snap_hit_total", nil)
	StorageTotAccessMeter     = metrics.NewRegisteredMeter("state/extra_stats/tot_accesses", nil)
	StorageTotTrieAccessMeter = metrics.NewRegisteredMeter("state/extra_stats/tot_trie_accesses", nil)
	StorageCacheHitRateMeter  = metrics.NewRegisteredGaugeFloat64("state/extra_stats/cache_hit_rate", nil)

	NodesPerTrieAccess    = metrics.NewRegisteredGaugeFloat64("state/extra_stats/node_accesses_per_storage_trie_access", nil)
	BranchesPerTrieAccess = metrics.NewRegisteredGaugeFloat64("state/extra_stats/node_accesses_per_storage_trie_access", nil)
)
