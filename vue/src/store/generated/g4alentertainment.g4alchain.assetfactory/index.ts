import { Client, registry, MissingWalletError } from 'G4AL-Entertainment-g4al-chain-client-ts'

import { Class } from "G4AL-Entertainment-g4al-chain-client-ts/g4alentertainment.g4alchain.assetfactory/types"
import { AssetfactoryPacketData } from "G4AL-Entertainment-g4al-chain-client-ts/g4alentertainment.g4alchain.assetfactory/types"
import { NoData } from "G4AL-Entertainment-g4al-chain-client-ts/g4alentertainment.g4alchain.assetfactory/types"
import { Params } from "G4AL-Entertainment-g4al-chain-client-ts/g4alentertainment.g4alchain.assetfactory/types"


export { Class, AssetfactoryPacketData, NoData, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				Class: {},
				ClassAll: {},
				
				_Structure: {
						Class: getStructure(Class.fromPartial({})),
						AssetfactoryPacketData: getStructure(AssetfactoryPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getClass: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Class[JSON.stringify(params)] ?? {}
		},
				getClassAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClassAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: g4alentertainment.g4alchain.assetfactory initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.G4AlentertainmentG4AlchainAssetfactory.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClass({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.G4AlentertainmentG4AlchainAssetfactory.query.queryClass( key.symbol)).data
				
					
				commit('QUERY', { query: 'Class', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClass', payload: { options: { all }, params: {...key},query }})
				return getters['getClass']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClass API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClassAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.G4AlentertainmentG4AlchainAssetfactory.query.queryClassAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.G4AlentertainmentG4AlchainAssetfactory.query.queryClassAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ClassAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClassAll', payload: { options: { all }, params: {...key},query }})
				return getters['getClassAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClassAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateClass({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.G4AlentertainmentG4AlchainAssetfactory.tx.sendMsgUpdateClass({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateClass:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateClass:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgMintNft({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.G4AlentertainmentG4AlchainAssetfactory.tx.sendMsgMintNft({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintNft:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMintNft:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateClass({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.G4AlentertainmentG4AlchainAssetfactory.tx.sendMsgCreateClass({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateClass:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateClass:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateClass({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.G4AlentertainmentG4AlchainAssetfactory.tx.msgUpdateClass({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateClass:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateClass:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgMintNft({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.G4AlentertainmentG4AlchainAssetfactory.tx.msgMintNft({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMintNft:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMintNft:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateClass({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.G4AlentertainmentG4AlchainAssetfactory.tx.msgCreateClass({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateClass:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateClass:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
