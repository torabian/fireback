package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class AuthResultDto extends JsonSerializable {
    public String workspaceId;
//    public UserRoleWorkspacePermission[] userRoleWorkspacePermissions;
    public String internalSql;
    public String userId;
    public String[] userHas;
    public String[] workspaceHas;
    public UserEntity user;
    public UserAccessLevelDto accessLevel;
    public static class VM extends ViewModel {
    // upper: WorkspaceId workspaceId
    private MutableLiveData< String > workspaceId = new MutableLiveData<>();
    public MutableLiveData< String > getWorkspaceId() {
        return workspaceId;
    }
    public void setWorkspaceId( String  v) {
        workspaceId.setValue(v);
    }
    // upper: UserRoleWorkspacePermissions userRoleWorkspacePermissions
//    private MutableLiveData< UserRoleWorkspacePermission[] > userRoleWorkspacePermissions = new MutableLiveData<>();
//    public MutableLiveData< UserRoleWorkspacePermission[] > getUserRoleWorkspacePermissions() {
//        return userRoleWorkspacePermissions;
//    }
//    public void setUserRoleWorkspacePermissions( UserRoleWorkspacePermission[]  v) {
//        userRoleWorkspacePermissions.setValue(v);
//    }
    // upper: InternalSql internalSql
    private MutableLiveData< String > internalSql = new MutableLiveData<>();
    public MutableLiveData< String > getInternalSql() {
        return internalSql;
    }
    public void setInternalSql( String  v) {
        internalSql.setValue(v);
    }
    // upper: UserId userId
    private MutableLiveData< String > userId = new MutableLiveData<>();
    public MutableLiveData< String > getUserId() {
        return userId;
    }
    public void setUserId( String  v) {
        userId.setValue(v);
    }
    // upper: UserHas userHas
    private MutableLiveData< String[] > userHas = new MutableLiveData<>();
    public MutableLiveData< String[] > getUserHas() {
        return userHas;
    }
    public void setUserHas( String[]  v) {
        userHas.setValue(v);
    }
    // upper: WorkspaceHas workspaceHas
    private MutableLiveData< String[] > workspaceHas = new MutableLiveData<>();
    public MutableLiveData< String[] > getWorkspaceHas() {
        return workspaceHas;
    }
    public void setWorkspaceHas( String[]  v) {
        workspaceHas.setValue(v);
    }
    // upper: User user
    private MutableLiveData< UserEntity > user = new MutableLiveData<>();
    public MutableLiveData< UserEntity > getUser() {
        return user;
    }
    public void setUser( UserEntity  v) {
        user.setValue(v);
    }
    // upper: AccessLevel accessLevel
    private MutableLiveData< UserAccessLevelDto > accessLevel = new MutableLiveData<>();
    public MutableLiveData< UserAccessLevelDto > getAccessLevel() {
        return accessLevel;
    }
    public void setAccessLevel( UserAccessLevelDto  v) {
        accessLevel.setValue(v);
    }
    }
}