package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
class SendEmailActionReqDto extends JsonSerializable {
    public String ToAddress;
    public String Body;
}
class SendEmailActionResDto extends JsonSerializable {
    public String QueueId;
}
class SendEmailWithProviderActionReqDto extends JsonSerializable {
    public EmailProviderEntity EmailProvider;
    public String ToAddress;
    public String Body;
}
class SendEmailWithProviderActionResDto extends JsonSerializable {
    public String QueueId;
}
class GsmSendSmsActionReqDto extends JsonSerializable {
    public String ToNumber;
    public String Body;
}
class GsmSendSmsActionResDto extends JsonSerializable {
    public String QueueId;
}
class GsmSendSmsWithProviderActionReqDto extends JsonSerializable {
    public GsmProviderEntity GsmProvider;
    public String ToNumber;
    public String Body;
}
class GsmSendSmsWithProviderActionResDto extends JsonSerializable {
    public String QueueId;
}
class ClassicSigninActionReqDto extends JsonSerializable {
    public String Value;
    public String Password;
}
class ClassicSignupActionReqDto extends JsonSerializable {
    public String Value;
    public String Type;
    public String Password;
    public String FirstName;
    public String LastName;
    public String InviteId;
    public String PublicJoinKeyId;
    public String WorkspaceTypeId;
}
class CreateWorkspaceActionReqDto extends JsonSerializable {
    public String Name;
    public WorkspaceEntity Workspace;
    public String WorkspaceId;
}
class CheckClassicPassportActionReqDto extends JsonSerializable {
    public String Value;
}
class CheckClassicPassportActionResDto extends JsonSerializable {
    public Boolean Exists;
}
class ClassicPassportOtpActionReqDto extends JsonSerializable {
    public String Value;
    public String Otp;
}
class ClassicPassportOtpActionResDto extends JsonSerializable {
    public int SuspendUntil;
    public UserSessionDto Session;
    public int ValidUntil;
    public int BlockedUntil;
    public int SecondsToUnblock;
}