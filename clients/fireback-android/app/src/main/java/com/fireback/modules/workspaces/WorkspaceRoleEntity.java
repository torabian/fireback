package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class WorkspaceRoleEntity extends JsonSerializable {
    public UserWorkspaceEntity userWorkspace;
    public RoleEntity role;
    public static class VM extends ViewModel {
    // upper: UserWorkspace userWorkspace
    private MutableLiveData< UserWorkspaceEntity > userWorkspace = new MutableLiveData<>();
    public MutableLiveData< UserWorkspaceEntity > getUserWorkspace() {
        return userWorkspace;
    }
    public void setUserWorkspace( UserWorkspaceEntity  v) {
        userWorkspace.setValue(v);
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