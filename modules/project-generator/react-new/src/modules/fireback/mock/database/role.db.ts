import { RoleDto } from "../../sdk/abac/RoleDto";
import { MemoryEntity } from "./memory-db";

export const MockRoles = new MemoryEntity<RoleDto>([
  {
    name: "Administrator",
    uniqueId: "administrator",
  },
]);
