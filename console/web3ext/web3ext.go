// Copyright 2015 The go-ethereum Authors
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

// package web3ext contains geth specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"admin":      Admin_JS,
	"chequebook": Chequebook_JS,
	"clique":     Clique_JS,
	"ethash":     Ethash_JS,
	"debug":      Debug_JS,
	"eth":        Eth_JS,
	"miner":      Miner_JS,
	"net":        Net_JS,
	"personal":   Personal_JS,
	"rpc":        RPC_JS,
	"shh":        Shh_JS,
	"swarmfs":    SWARMFS_JS,
	"txpool":     TxPool_JS,
	"lk":         LK_JS,
}

const Chequebook_JS = `
web3._extend({
	property: 'chequebook',
	methods: [
		new web3._extend.Method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const Clique_JS = `
web3._extend({
	property: 'clique',
	methods: [
		new web3._extend.Method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.Method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const Ethash_JS = `
web3._extend({
	property: 'ethash',
	methods: [
		new web3._extend.Method({
			name: 'getWork',
			call: 'ethash_getWork',
			params: 0
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'ethash_getHashrate',
			params: 0
		}), 
		new web3._extend.Method({
			name: 'submitWork',
			call: 'ethash_submitWork',
			params: 3,
		}),
		new web3._extend.Method({
			name: 'submitHashRate',
			call: 'ethash_submitHashRate',
			params: 2,
		}),
	]
});
`

const Admin_JS = `
web3._extend({
	property: 'admin',
	methods: [
		new web3._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'addTrustedPeer',
			call: 'admin_addTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'removeTrustedPeer',
			call: 'admin_removeTrustedPeer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.Method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const Debug_JS = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.Method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockSer',
			call: 'debug_getBlockSer',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getBlockBinary',
			call: 'debug_getBlockBinary',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter]
		}),
	],
	properties: []
});
`

const Eth_JS = `
web3._extend({
	property: 'eth',
	methods: [
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'eth_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'submitTransaction',
			call: 'eth_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Method({
			name: 'getRawTransaction',
			call: 'eth_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.Method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'eth_getRawTransactionByBlockHashAndIndex' : 'eth_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
		new web3._extend.Method({
			name: 'getStorageRoot',
			call: 'eth_getStorageRoot',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, web3._extend.formatters.inputBlockNumberFormatter]
		}),
		new web3._extend.Method({
			name: 'sendRawUTXOTransaction',
			call: 'eth_sendRawUTXOTransaction',
			params: 1
		}),
		new web3._extend.Method({
			name: 'sendRawTx',
			call: 'eth_sendRawTx',
			params: 2
		}),
	]
});
`

const Miner_JS = `
web3._extend({
	property: 'miner',
	methods: [
		new web3._extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'stop',
			call: 'miner_stop'
		}),
		new web3._extend.Method({
			name: 'setEtherbase',
			call: 'miner_setEtherbase',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new web3._extend.Method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'setRecommitInterval',
			call: 'miner_setRecommitInterval',
			params: 1,
		}),
		new web3._extend.Method({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
	],
	properties: []
});
`

const Net_JS = `
web3._extend({
	property: 'net',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const Personal_JS = `
web3._extend({
	property: 'personal',
	methods: [
		new web3._extend.Method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.Method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
		new web3._extend.Method({
			name: 'signTransaction',
			call: 'personal_signTransaction',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, null]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RPC_JS = `
web3._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const Shh_JS = `
web3._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SWARMFS_JS = `
web3._extend({
	property: 'swarmfs',
	methods:
	[
		new web3._extend.Method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.Method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.Method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxPool_JS = `
web3._extend({
	property: 'txpool',
	methods: [
		new web3._extend.Method({
			name: 'getStatus',
			call: 'txpool_getStatus'
			params: 1
		}),
	],
	properties:
	[
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.futureTxs = web3._extend.utils.toDecimal(status.futureTxs);
				status.goodTxs = web3._extend.utils.toDecimal(status.goodTxs);

				return status;
			}
		}),
	]
});
`

const LK_JS = `
web3._extend({
	property: 'lk',
	methods: [
		new web3._extend.Method({
			name: 'getLogs',
			call: 'lk_getLogs',
			params: 1
		}),
	],
});
`
