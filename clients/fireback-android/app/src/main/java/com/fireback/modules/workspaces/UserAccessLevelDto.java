package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class UserAccessLevelDto extends JsonSerializable {
    public String[] capabilities;
    public String[] workspaces;
    public String SQL;
    public static class VM extends ViewModel {
    // upper: Capabilities capabilities
    private MutableLiveData< String[] > capabilities = new MutableLiveData<>();
    public MutableLiveData< String[] > getCapabilities() {
        return capabilities;
    }
    public void setCapabilities( String[]  v) {
        capabilities.setValue(v);
    }
    // upper: Workspaces workspaces
    private MutableLiveData< String[] > workspaces = new MutableLiveData<>();
    public MutableLiveData< String[] > getWorkspaces() {
        return workspaces;
    }
    public void setWorkspaces( String[]  v) {
        workspaces.setValue(v);
    }
    // upper: SQL SQL
    private MutableLiveData< String > SQL = new MutableLiveData<>();
    public MutableLiveData< String > getSQL() {
        return SQL;
    }
    public void setSQL( String  v) {
        SQL.setValue(v);
    }
    }
}