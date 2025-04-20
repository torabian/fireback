import { FileEntity } from "../../sdk/modules/abac/FileEntity";
import { MemoryEntity } from "./memory-db";

export const MockFiles = new MemoryEntity<FileEntity>([]);
