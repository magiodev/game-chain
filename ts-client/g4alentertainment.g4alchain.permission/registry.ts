import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateAdministrator } from "./types/g4alchain/permission/tx";
import { MsgDeleteAdministrator } from "./types/g4alchain/permission/tx";
import { MsgUpdateAdministrator } from "./types/g4alchain/permission/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.permission.MsgCreateAdministrator", MsgCreateAdministrator],
    ["/g4alentertainment.g4alchain.permission.MsgDeleteAdministrator", MsgDeleteAdministrator],
    ["/g4alentertainment.g4alchain.permission.MsgUpdateAdministrator", MsgUpdateAdministrator],
    
];

export { msgTypes }