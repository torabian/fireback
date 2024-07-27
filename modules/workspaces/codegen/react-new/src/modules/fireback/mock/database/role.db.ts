import { RoleEntity } from "../../sdk/modules/workspaces/RoleEntity";
import { MemoryEntity } from "./memory-db";

export const MockRoles = new MemoryEntity<RoleEntity>([
  {
    name: "Administrator",
    uniqueId: "administrator",
  },
]);
