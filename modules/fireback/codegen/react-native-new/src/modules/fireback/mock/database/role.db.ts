import { RoleEntity } from "../../sdk/modules/abac/RoleEntity";
import { MemoryEntity } from "./memory-db";

export const MockRoles = new MemoryEntity<RoleEntity>([
  {
    name: "Administrator",
    uniqueId: "administrator",
  },
]);
