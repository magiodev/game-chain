import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateClass } from "./types/g4alchain/assetfactory/tx";
import { MsgDeleteClass } from "./types/g4alchain/assetfactory/tx";
import { MsgUpdateClass } from "./types/g4alchain/assetfactory/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.assetfactory.MsgCreateClass", MsgCreateClass],
    ["/g4alentertainment.g4alchain.assetfactory.MsgDeleteClass", MsgDeleteClass],
    ["/g4alentertainment.g4alchain.assetfactory.MsgUpdateClass", MsgUpdateClass],
    
];

export { msgTypes }