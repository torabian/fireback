package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class CreateWorkspaceAction {
    public static class Req extends JsonSerializable {
    public String name;
    public WorkspaceEntity workspace;
    public String workspaceId;
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData<String> workspaceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceMsg() {
        return workspaceMsg;
    }
    public void setWorkspaceMsg(String v) {
        workspaceMsg.setValue(v);
    }
    // upper: WorkspaceId workspaceId
    private MutableLiveData<String> workspaceIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceIdMsg() {
        return workspaceIdMsg;
    }
    public void setWorkspaceIdMsg(String v) {
        workspaceIdMsg.setValue(v);
    }
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
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData<String> workspaceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceMsg() {
        return workspaceMsg;
    }
    public void setWorkspaceMsg(String v) {
        workspaceMsg.setValue(v);
    }
    // upper: WorkspaceId workspaceId
    private MutableLiveData<String> workspaceIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWorkspaceIdMsg() {
        return workspaceIdMsg;
    }
    public void setWorkspaceIdMsg(String v) {
        workspaceIdMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("name")) {
            this.setNameMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("workspace")) {
            this.setWorkspaceMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("workspaceId")) {
            this.setWorkspaceIdMsg(item.messageTranslated);
        }
    });
}
    }
}