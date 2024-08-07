/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;

public class UserAccessLevelDto extends JsonSerializable {
  public String[] capabilities;
  public UserRoleWorkspacePermissionDto[] userRoleWorkspacePermissions;
  public String[] workspaces;
  public String SQL;

  public static class VM extends ViewModel {
    // upper: Capabilities capabilities
    private MutableLiveData<String[]> capabilities = new MutableLiveData<>();

    public MutableLiveData<String[]> getCapabilities() {
      return capabilities;
    }

    public void setCapabilities(String[] v) {
      capabilities.setValue(v);
    }

    // upper: UserRoleWorkspacePermissions userRoleWorkspacePermissions
    private MutableLiveData<UserRoleWorkspacePermissionDto[]> userRoleWorkspacePermissions =
        new MutableLiveData<>();

    public MutableLiveData<UserRoleWorkspacePermissionDto[]> getUserRoleWorkspacePermissions() {
      return userRoleWorkspacePermissions;
    }

    public void setUserRoleWorkspacePermissions(UserRoleWorkspacePermissionDto[] v) {
      userRoleWorkspacePermissions.setValue(v);
    }

    // upper: Workspaces workspaces
    private MutableLiveData<String[]> workspaces = new MutableLiveData<>();

    public MutableLiveData<String[]> getWorkspaces() {
      return workspaces;
    }

    public void setWorkspaces(String[] v) {
      workspaces.setValue(v);
    }

    // upper: SQL SQL
    private MutableLiveData<String> SQL = new MutableLiveData<>();

    public MutableLiveData<String> getSQL() {
      return SQL;
    }

    public void setSQL(String v) {
      SQL.setValue(v);
    }
  }
}
