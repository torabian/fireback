package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PendingWorkspaceInviteEntity extends JsonSerializable {
    public String value;
    public String type;
    public String coverLetter;
    public String workspaceName;
    public RoleEntity role;
    public static class VM extends ViewModel {
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
    // upper: CoverLetter coverLetter
    private MutableLiveData< String > coverLetter = new MutableLiveData<>();
    public MutableLiveData< String > getCoverLetter() {
        return coverLetter;
    }
    public void setCoverLetter( String  v) {
        coverLetter.setValue(v);
    }
    // upper: WorkspaceName workspaceName
    private MutableLiveData< String > workspaceName = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceName() {
        return workspaceName;
    }
    public void setWorkspaceName( String  v) {
        workspaceName.setValue(v);
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