package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class AppMenuEntity extends JsonSerializable {
    public String href;
    public String icon;
    public String label;
    public String activeMatcher;
    public String applyType;
    public CapabilityEntity capability;
    public static class VM extends ViewModel {
    // upper: Href href
    private MutableLiveData< String > href = new MutableLiveData<>();
    public MutableLiveData< String > getHref() {
        return href;
    }
    public void setHref( String  v) {
        href.setValue(v);
    }
    // upper: Icon icon
    private MutableLiveData< String > icon = new MutableLiveData<>();
    public MutableLiveData< String > getIcon() {
        return icon;
    }
    public void setIcon( String  v) {
        icon.setValue(v);
    }
    // upper: Label label
    private MutableLiveData< String > label = new MutableLiveData<>();
    public MutableLiveData< String > getLabel() {
        return label;
    }
    public void setLabel( String  v) {
        label.setValue(v);
    }
    // upper: ActiveMatcher activeMatcher
    private MutableLiveData< String > activeMatcher = new MutableLiveData<>();
    public MutableLiveData< String > getActiveMatcher() {
        return activeMatcher;
    }
    public void setActiveMatcher( String  v) {
        activeMatcher.setValue(v);
    }
    // upper: ApplyType applyType
    private MutableLiveData< String > applyType = new MutableLiveData<>();
    public MutableLiveData< String > getApplyType() {
        return applyType;
    }
    public void setApplyType( String  v) {
        applyType.setValue(v);
    }
    // upper: Capability capability
    private MutableLiveData< CapabilityEntity > capability = new MutableLiveData<>();
    public MutableLiveData< CapabilityEntity > getCapability() {
        return capability;
    }
    public void setCapability( CapabilityEntity  v) {
        capability.setValue(v);
    }
    }
}