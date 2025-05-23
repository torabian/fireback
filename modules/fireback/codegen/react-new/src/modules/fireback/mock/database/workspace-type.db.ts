import { WorkspaceTypeEntity } from "../../sdk/modules/abac/WorkspaceTypeEntity";
import { MemoryEntity } from "./memory-db";

export const MockWorkspaceType = new MemoryEntity<WorkspaceTypeEntity>([
  {
    title: "Student workspace type",
    uniqueId: "1",
    slug: "/student",
  },
]);
