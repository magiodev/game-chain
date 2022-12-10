import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateNft } from "./types/g4alchain/assetfactory/tx";
import { MsgBurnNft } from "./types/g4alchain/assetfactory/tx";
import { MsgUpdateClass } from "./types/g4alchain/assetfactory/tx";
import { MsgCreateClass } from "./types/g4alchain/assetfactory/tx";
import { MsgMintNft } from "./types/g4alchain/assetfactory/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.assetfactory.MsgUpdateNft", MsgUpdateNft],
    ["/g4alentertainment.g4alchain.assetfactory.MsgBurnNft", MsgBurnNft],
    ["/g4alentertainment.g4alchain.assetfactory.MsgUpdateClass", MsgUpdateClass],
    ["/g4alentertainment.g4alchain.assetfactory.MsgCreateClass", MsgCreateClass],
    ["/g4alentertainment.g4alchain.assetfactory.MsgMintNft", MsgMintNft],
    
];

export { msgTypes }