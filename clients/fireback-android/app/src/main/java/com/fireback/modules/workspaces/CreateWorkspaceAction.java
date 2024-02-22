package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class CreateWorkspaceAction {
    public static class Req extends JsonSerializable {
    public String name;
    public WorkspaceEntity workspace;
    public String workspaceId;
    }
    public static class ReqViewModel extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
    }
    // upper: WorkspaceId workspaceId
    private MutableLiveData< String > workspaceId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceId() {
        return workspaceId;
    }
    public void setWorkspaceId( String  v) {
        workspaceId.setValue(v);
    }
    }
}