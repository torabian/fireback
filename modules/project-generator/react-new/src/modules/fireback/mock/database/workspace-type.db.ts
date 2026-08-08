import { WorkspaceTypeDto } from "../../sdk/abac/WorkspaceTypeDto";
import { MemoryEntity } from "./memory-db";

export const MockWorkspaceType = new MemoryEntity<WorkspaceTypeDto>([
  {
    title: "Student workspace type",
    uniqueId: "1",
    slug: "/student",
  },
]);
