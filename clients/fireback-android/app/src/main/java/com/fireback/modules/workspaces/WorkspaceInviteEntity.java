package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class WorkspaceInviteEntity extends JsonSerializable {
    public String coverLetter;
    public String targetUserLocale;
    public String value;
    public WorkspaceEntity workspace;
    public String firstName;
    public String lastName;
    public Boolean used;
    public RoleEntity role;
    public static class VM extends ViewModel {
    // upper: CoverLetter coverLetter
    private MutableLiveData< String > coverLetter = new MutableLiveData<>();
    public MutableLiveData< String > getCoverLetter() {
        return coverLetter;
    }
    public void setCoverLetter( String  v) {
        coverLetter.setValue(v);
    }
    // upper: TargetUserLocale targetUserLocale
    private MutableLiveData< String > targetUserLocale = new MutableLiveData<>();
    public MutableLiveData< String > getTargetUserLocale() {
        return targetUserLocale;
    }
    public void setTargetUserLocale( String  v) {
        targetUserLocale.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
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
    // upper: Used used
    private MutableLiveData< Boolean > used = new MutableLiveData<>();
    public MutableLiveData< Boolean > getUsed() {
        return used;
    }
    public void setUsed( Boolean  v) {
        used.setValue(v);
    }
    // upper: Role role
    private MutableLiveData< RoleEntity > role = new MutableLiveData<>();
    public MutableLiveData< RoleEntity > getRole() {
        return role;
    }
    public void setRole( RoleEntity  v) {
        role.setValue(v);
    }
    }
}