package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserRoleWorkspaceDto extends JsonSerializable {
    public String roleId;
    public String[] capabilities;
    public static class VM extends ViewModel {
    // upper: RoleId roleId
    private MutableLiveData< String > roleId = new MutableLiveData<>();
    public MutableLiveData< String > getRoleId() {
        return roleId;
    }
    public void setRoleId( String  v) {
        roleId.setValue(v);
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