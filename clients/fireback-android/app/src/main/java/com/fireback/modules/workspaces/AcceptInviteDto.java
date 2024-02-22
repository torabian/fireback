package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class AcceptInviteDto extends JsonSerializable {
    public String inviteUniqueId;
    public String visibility;
    public int updated;
    public int created;
    public static class VM extends ViewModel {
    // upper: InviteUniqueId inviteUniqueId
    private MutableLiveData< String > inviteUniqueId = new MutableLiveData<>();
    public MutableLiveData< String > getInviteUniqueId() {
        return inviteUniqueId;
    }
    public void setInviteUniqueId( String  v) {
        inviteUniqueId.setValue(v);
    }
    // upper: Visibility visibility
    private MutableLiveData< String > visibility = new MutableLiveData<>();
    public MutableLiveData< String > getVisibility() {
        return visibility;
    }
    public void setVisibility( String  v) {
        visibility.setValue(v);
    }
    // upper: Updated updated
    private MutableLiveData< Integer > updated = new MutableLiveData<>();
    public MutableLiveData< Integer > getUpdated() {
        return updated;
    }
    public void setUpdated( Integer  v) {
        updated.setValue(v);
    }
    // upper: Created created
    private MutableLiveData< Integer > created = new MutableLiveData<>();
    public MutableLiveData< Integer > getCreated() {
        return created;
    }
    public void setCreated( Integer  v) {
        created.setValue(v);
    }
    }
}