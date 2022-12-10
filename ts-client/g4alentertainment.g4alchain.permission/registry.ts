import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateDeveloper } from "./types/g4alchain/permission/tx";
import { MsgCreateDeveloper } from "./types/g4alchain/permission/tx";
import { MsgUpdateAdministrator } from "./types/g4alchain/permission/tx";
import { MsgCreateAdministrator } from "./types/g4alchain/permission/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.permission.MsgUpdateDeveloper", MsgUpdateDeveloper],
    ["/g4alentertainment.g4alchain.permission.MsgCreateDeveloper", MsgCreateDeveloper],
    ["/g4alentertainment.g4alchain.permission.MsgUpdateAdministrator", MsgUpdateAdministrator],
    ["/g4alentertainment.g4alchain.permission.MsgCreateAdministrator", MsgCreateAdministrator],

];

export { msgTypes }