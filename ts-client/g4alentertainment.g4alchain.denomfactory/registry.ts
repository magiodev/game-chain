import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteDenom } from "./types/g4alchain/denomfactory/tx";
import { MsgCreateDenom } from "./types/g4alchain/denomfactory/tx";
import { MsgUpdateDenom } from "./types/g4alchain/denomfactory/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.denomfactory.MsgDeleteDenom", MsgDeleteDenom],
    ["/g4alentertainment.g4alchain.denomfactory.MsgCreateDenom", MsgCreateDenom],
    ["/g4alentertainment.g4alchain.denomfactory.MsgUpdateDenom", MsgUpdateDenom],
    
];

export { msgTypes }