import { FileEntity } from "../../sdk/modules/workspaces/FileEntity";
import { MemoryEntity } from "./memory-db";

export const MockFiles = new MemoryEntity<FileEntity>([]);
