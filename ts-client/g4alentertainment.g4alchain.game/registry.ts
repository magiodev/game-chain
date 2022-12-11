import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateProject } from "./types/g4alchain/game/tx";
import { MsgCreateProject } from "./types/g4alchain/game/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.game.MsgUpdateProject", MsgUpdateProject],
    ["/g4alentertainment.g4alchain.game.MsgCreateProject", MsgCreateProject],
    
];

export { msgTypes }