import { EmailProviderEntity } from "../../sdk/modules/abac/EmailProviderEntity";
import { EmailSenderEntity } from "../../sdk/modules/abac/EmailSenderEntity";
import { PublicJoinKeyEntity } from "../../sdk/modules/abac/PublicJoinKeyEntity";
import { WorkspaceEntity } from "../../sdk/modules/abac/WorkspaceEntity";
import { WorkspaceInviteEntity } from "../../sdk/modules/abac/WorkspaceInviteEntity";
import { MemoryEntity } from "./memory-db";

export const mdb = {
  emailProvider: new MemoryEntity<EmailProviderEntity>([]),
  emailSender: new MemoryEntity<EmailSenderEntity>([]),
  workspaceInvite: new MemoryEntity<WorkspaceInviteEntity>([]),
  publicJoinKey: new MemoryEntity<PublicJoinKeyEntity>([]),
  workspaces: new MemoryEntity<WorkspaceEntity>([]),
};
