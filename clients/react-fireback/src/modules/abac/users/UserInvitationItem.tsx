import { PendingWorkspaceInviteEntity } from "src/sdk/fireback";

export function UserInvitationItem({
  invite,
}: {
  invite: PendingWorkspaceInviteEntity;
}) {
  return (
    <div className="mb-5">
      <h2>{invite.role?.name}</h2>
      <p>For the {invite.workspaceName}</p>
      <button className="btn btn-sm btn-primary">Accept</button>{" "}
      <button className="btn btn-sm btn-danger ">Reject</button>
    </div>
  );
}
