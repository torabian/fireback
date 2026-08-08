import type { FileEntity } from "../../sdk/legacy-types/FileEntity";
import { MemoryEntity } from "./memory-db";

export const MockFiles = new MemoryEntity<FileEntity>([]);
