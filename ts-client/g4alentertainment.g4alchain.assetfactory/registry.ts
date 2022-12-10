import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateClass } from "./types/g4alchain/assetfactory/tx";
import { MsgMintNft } from "./types/g4alchain/assetfactory/tx";
import { MsgCreateClass } from "./types/g4alchain/assetfactory/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.assetfactory.MsgUpdateClass", MsgUpdateClass],
    ["/g4alentertainment.g4alchain.assetfactory.MsgMintNft", MsgMintNft],
    ["/g4alentertainment.g4alchain.assetfactory.MsgCreateClass", MsgCreateClass],
    
];

export { msgTypes }