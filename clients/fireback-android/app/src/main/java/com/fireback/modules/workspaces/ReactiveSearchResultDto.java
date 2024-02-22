package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ReactiveSearchResultDto extends JsonSerializable {
    public String uniqueId;
    public String phrase;
    public String icon;
    public String description;
    public String group;
    public String uiLocation;
    public String actionFn;
    public static class VM extends ViewModel {
    // upper: UniqueId uniqueId
    private MutableLiveData< String > uniqueId = new MutableLiveData<>();
    public MutableLiveData< String > getUniqueId() {
        return uniqueId;
    }
    public void setUniqueId( String  v) {
        uniqueId.setValue(v);
    }
    // upper: Phrase phrase
    private MutableLiveData< String > phrase = new MutableLiveData<>();
    public MutableLiveData< String > getPhrase() {
        return phrase;
    }
    public void setPhrase( String  v) {
        phrase.setValue(v);
    }
    // upper: Icon icon
    private MutableLiveData< String > icon = new MutableLiveData<>();
    public MutableLiveData< String > getIcon() {
        return icon;
    }
    public void setIcon( String  v) {
        icon.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: Group group
    private MutableLiveData< String > group = new MutableLiveData<>();
    public MutableLiveData< String > getGroup() {
        return group;
    }
    public void setGroup( String  v) {
        group.setValue(v);
    }
    // upper: UiLocation uiLocation
    private MutableLiveData< String > uiLocation = new MutableLiveData<>();
    public MutableLiveData< String > getUiLocation() {
        return uiLocation;
    }
    public void setUiLocation( String  v) {
        uiLocation.setValue(v);
    }
    // upper: ActionFn actionFn
    private MutableLiveData< String > actionFn = new MutableLiveData<>();
    public MutableLiveData< String > getActionFn() {
        return actionFn;
    }
    public void setActionFn( String  v) {
        actionFn.setValue(v);
    }
    }
}