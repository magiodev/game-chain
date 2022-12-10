// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import G4AlentertainmentG4AlchainG4Alchain from './g4alentertainment.g4alchain.g4alchain'


export default { 
  G4AlentertainmentG4AlchainG4Alchain: load(G4AlentertainmentG4AlchainG4Alchain, 'g4alentertainment.g4alchain.g4alchain'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}