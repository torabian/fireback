package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserWorkspaceEntity extends JsonSerializable {
    public UserEntity user;
    public WorkspaceEntity workspace;
    public static class VM extends ViewModel {
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
    }
    }
}