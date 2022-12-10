import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateProject } from "./types/g4alchain/game/tx";
import { MsgUpdateProject } from "./types/g4alchain/game/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/g4alentertainment.g4alchain.game.MsgCreateProject", MsgCreateProject],
    ["/g4alentertainment.g4alchain.game.MsgUpdateProject", MsgUpdateProject],
    
];

export { msgTypes }