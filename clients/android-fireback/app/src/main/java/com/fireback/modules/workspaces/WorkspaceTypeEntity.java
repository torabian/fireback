package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class WorkspaceTypeEntity extends JsonSerializable {
    public String title;
    public String description;
    public String slug;
    public RoleEntity role;
    public static class VM extends ViewModel {
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: Slug slug
    private MutableLiveData< String > slug = new MutableLiveData<>();
    public MutableLiveData< String > getSlug() {
        return slug;
    }
    public void setSlug( String  v) {
        slug.setValue(v);
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