import { EmailProviderEntity } from "../../sdk/modules/workspaces/EmailProviderEntity";
import { EmailSenderEntity } from "../../sdk/modules/workspaces/EmailSenderEntity";
import { PublicJoinKeyEntity } from "../../sdk/modules/workspaces/PublicJoinKeyEntity";
import { WorkspaceEntity } from "../../sdk/modules/workspaces/WorkspaceEntity";
import { WorkspaceInviteEntity } from "../../sdk/modules/workspaces/WorkspaceInviteEntity";
import { MemoryEntity } from "./memory-db";

export const mdb = {
  emailProvider: new MemoryEntity<EmailProviderEntity>([]),
  emailSender: new MemoryEntity<EmailSenderEntity>([]),
  workspaceInvite: new MemoryEntity<WorkspaceInviteEntity>([]),
  publicJoinKey: new MemoryEntity<PublicJoinKeyEntity>([]),
  workspaces: new MemoryEntity<WorkspaceEntity>([]),
};
