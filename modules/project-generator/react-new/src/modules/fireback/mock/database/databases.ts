import { EmailProviderDto } from "../../sdk/messaging/EmailProviderDto";
import { EmailSenderDto } from "../../sdk/messaging/EmailSenderDto";
import { PublicJoinKeyDto } from "../../sdk/abac/PublicJoinKeyDto";
import { WorkspaceDto } from "../../sdk/abac/WorkspaceDto";
import { WorkspaceInviteDto } from "../../sdk/abac/WorkspaceInviteDto";
import { MemoryEntity } from "./memory-db";

export const mdb = {
  emailProvider: new MemoryEntity<EmailProviderDto>([]),
  emailSender: new MemoryEntity<EmailSenderDto>([]),
  workspaceInvite: new MemoryEntity<WorkspaceInviteDto>([]),
  publicJoinKey: new MemoryEntity<PublicJoinKeyDto>([]),
  workspaces: new MemoryEntity<WorkspaceDto>([]),
};
