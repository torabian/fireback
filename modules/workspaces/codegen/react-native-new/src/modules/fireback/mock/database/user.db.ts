import { UserEntity } from "../../sdk/modules/workspaces/UserEntity";
import { MemoryEntity } from "./memory-db";

export const MockUsers = new MemoryEntity<UserEntity>([
  {
    person: {
      firstName: "Alex",
      lastName: "Mendoz",
    },
    uniqueId: "1",
  },
]);
