package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class ClassicSignupAction {
    public static class Req extends JsonSerializable {
    public String value;
    public String type;
    public String password;
    public String firstName;
    public String lastName;
    public String inviteId;
    public String publicJoinKeyId;
    public String workspaceTypeId;
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: Password password
    private MutableLiveData<String> passwordMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPasswordMsg() {
        return passwordMsg;
    }
    public void setPasswordMsg(String v) {
        passwordMsg.setValue(v);
    }
    // upper: FirstName firstName
    private MutableLiveData<String> firstNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFirstNameMsg() {
        return firstNameMsg;
    }
    public void setFirstNameMsg(String v) {
        firstNameMsg.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData<String> lastNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLastNameMsg() {
        return lastNameMsg;
    }
    public void setLastNameMsg(String v) {
        lastNameMsg.setValue(v);
    }
    // upper: InviteId inviteId
    private MutableLiveData<String> inviteIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteIdMsg() {
        return inviteIdMsg;
    }
    public void setInviteIdMsg(String v) {
        inviteIdMsg.setValue(v);
    }
    // upper: PublicJoinKeyId publicJoinKeyId
    private MutableLiveData<String> publicJoinKeyIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPublicJoinKeyIdMsg() {
        return publicJoinKeyIdMsg;
    }
    public void setPublicJoinKeyIdMsg(String v) {
        publicJoinKeyIdMsg.setValue(v);
    }
    // upper: WorkspaceTypeId workspaceTypeId
    private MutableLiveData<String> workspaceTypeIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceTypeIdMsg() {
        return workspaceTypeIdMsg;
    }
    public void setWorkspaceTypeIdMsg(String v) {
        workspaceTypeIdMsg.setValue(v);
    }
    }
    public static class ReqViewModel extends ViewModel {
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    // upper: FirstName firstName
    private MutableLiveData< String > firstName = new MutableLiveData<>();
    public MutableLiveData< String > getFirstName() {
        return firstName;
    }
    public void setFirstName( String  v) {
        firstName.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData< String > lastName = new MutableLiveData<>();
    public MutableLiveData< String > getLastName() {
        return lastName;
    }
    public void setLastName( String  v) {
        lastName.setValue(v);
    }
    // upper: InviteId inviteId
    private MutableLiveData< String > inviteId = new MutableLiveData<>();
    public MutableLiveData< String > getInviteId() {
        return inviteId;
    }
    public void setInviteId( String  v) {
        inviteId.setValue(v);
    }
    // upper: PublicJoinKeyId publicJoinKeyId
    private MutableLiveData< String > publicJoinKeyId = new MutableLiveData<>();
    public MutableLiveData< String > getPublicJoinKeyId() {
        return publicJoinKeyId;
    }
    public void setPublicJoinKeyId( String  v) {
        publicJoinKeyId.setValue(v);
    }
    // upper: WorkspaceTypeId workspaceTypeId
    private MutableLiveData< String > workspaceTypeId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceTypeId() {
        return workspaceTypeId;
    }
    public void setWorkspaceTypeId( String  v) {
        workspaceTypeId.setValue(v);
    }
    // upper: Value value
    private MutableLiveData<String> valueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueMsg() {
        return valueMsg;
    }
    public void setValueMsg(String v) {
        valueMsg.setValue(v);
    }
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
    // upper: Password password
    private MutableLiveData<String> passwordMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPasswordMsg() {
        return passwordMsg;
    }
    public void setPasswordMsg(String v) {
        passwordMsg.setValue(v);
    }
    // upper: FirstName firstName
    private MutableLiveData<String> firstNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFirstNameMsg() {
        return firstNameMsg;
    }
    public void setFirstNameMsg(String v) {
        firstNameMsg.setValue(v);
    }
    // upper: LastName lastName
    private MutableLiveData<String> lastNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLastNameMsg() {
        return lastNameMsg;
    }
    public void setLastNameMsg(String v) {
        lastNameMsg.setValue(v);
    }
    // upper: InviteId inviteId
    private MutableLiveData<String> inviteIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInviteIdMsg() {
        return inviteIdMsg;
    }
    public void setInviteIdMsg(String v) {
        inviteIdMsg.setValue(v);
    }
    // upper: PublicJoinKeyId publicJoinKeyId
    private MutableLiveData<String> publicJoinKeyIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPublicJoinKeyIdMsg() {
        return publicJoinKeyIdMsg;
    }
    public void setPublicJoinKeyIdMsg(String v) {
        publicJoinKeyIdMsg.setValue(v);
    }
    // upper: WorkspaceTypeId workspaceTypeId
    private MutableLiveData<String> workspaceTypeIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceTypeIdMsg() {
        return workspaceTypeIdMsg;
    }
    public void setWorkspaceTypeIdMsg(String v) {
        workspaceTypeIdMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("value")) {
            this.setValueMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("type")) {
            this.setTypeMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("password")) {
            this.setPasswordMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("firstName")) {
            this.setFirstNameMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("lastName")) {
            this.setLastNameMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("inviteId")) {
            this.setInviteIdMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("publicJoinKeyId")) {
            this.setPublicJoinKeyIdMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("workspaceTypeId")) {
            this.setWorkspaceTypeIdMsg(item.messageTranslated);
        }
    });
}
    }
}