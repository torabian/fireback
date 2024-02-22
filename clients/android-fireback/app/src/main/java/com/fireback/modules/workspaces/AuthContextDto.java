package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class AuthContextDto extends JsonSerializable {
    public Boolean skipWorkspaceId;
    public String workspaceId;
    public String token;
    public String[] capabilities;
    public static class VM extends ViewModel {
    // upper: SkipWorkspaceId skipWorkspaceId
    private MutableLiveData< Boolean > skipWorkspaceId = new MutableLiveData<>();
    public MutableLiveData< Boolean > getSkipWorkspaceId() {
        return skipWorkspaceId;
    }
    public void setSkipWorkspaceId( Boolean  v) {
        skipWorkspaceId.setValue(v);
    }
    // upper: WorkspaceId workspaceId
    private MutableLiveData< String > workspaceId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceId() {
        return workspaceId;
    }
    public void setWorkspaceId( String  v) {
        workspaceId.setValue(v);
    }
    // upper: Token token
    private MutableLiveData< String > token = new MutableLiveData<>();
    public MutableLiveData< String > getToken() {
        return token;
    }
    public void setToken( String  v) {
        token.setValue(v);
    }
    // upper: Capabilities capabilities
    private MutableLiveData< String[] > capabilities = new MutableLiveData<>();
    public MutableLiveData< String[] > getCapabilities() {
        return capabilities;
    }
    public void setCapabilities( String[]  v) {
        capabilities.setValue(v);
    }
    }
}