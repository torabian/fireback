package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class WorkspaceConfigEntity extends JsonSerializable {
    public int disablePublicWorkspaceCreation;
    public WorkspaceEntity workspace;
    public String zoomClientId;
    public String zoomClientSecret;
    public Boolean allowPublicToJoinTheWorkspace;
    public static class VM extends ViewModel {
    // upper: DisablePublicWorkspaceCreation disablePublicWorkspaceCreation
    private MutableLiveData< Integer > disablePublicWorkspaceCreation = new MutableLiveData<>();
    public MutableLiveData< Integer > getDisablePublicWorkspaceCreation() {
        return disablePublicWorkspaceCreation;
    }
    public void setDisablePublicWorkspaceCreation( Integer  v) {
        disablePublicWorkspaceCreation.setValue(v);
    }
    // upper: Workspace workspace
    private MutableLiveData< WorkspaceEntity > workspace = new MutableLiveData<>();
    public MutableLiveData< WorkspaceEntity > getWorkspace() {
        return workspace;
    }
    public void setWorkspace( WorkspaceEntity  v) {
        workspace.setValue(v);
    }
    // upper: ZoomClientId zoomClientId
    private MutableLiveData< String > zoomClientId = new MutableLiveData<>();
    public MutableLiveData< String > getZoomClientId() {
        return zoomClientId;
    }
    public void setZoomClientId( String  v) {
        zoomClientId.setValue(v);
    }
    // upper: ZoomClientSecret zoomClientSecret
    private MutableLiveData< String > zoomClientSecret = new MutableLiveData<>();
    public MutableLiveData< String > getZoomClientSecret() {
        return zoomClientSecret;
    }
    public void setZoomClientSecret( String  v) {
        zoomClientSecret.setValue(v);
    }
    // upper: AllowPublicToJoinTheWorkspace allowPublicToJoinTheWorkspace
    private MutableLiveData< Boolean > allowPublicToJoinTheWorkspace = new MutableLiveData<>();
    public MutableLiveData< Boolean > getAllowPublicToJoinTheWorkspace() {
        return allowPublicToJoinTheWorkspace;
    }
    public void setAllowPublicToJoinTheWorkspace( Boolean  v) {
        allowPublicToJoinTheWorkspace.setValue(v);
    }
    }
}