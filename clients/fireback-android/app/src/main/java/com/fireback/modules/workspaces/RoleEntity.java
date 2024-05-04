package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class RoleEntity extends JsonSerializable {
    public String name;
    public CapabilityEntity[] capabilities;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Capabilities capabilities
    private MutableLiveData< CapabilityEntity[] > capabilities = new MutableLiveData<>();
    public MutableLiveData< CapabilityEntity[] > getCapabilities() {
        return capabilities;
    }
    public void setCapabilities( CapabilityEntity[]  v) {
        capabilities.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Capabilities capabilities
    private MutableLiveData<String> capabilitiesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCapabilitiesMsg() {
        return capabilitiesMsg;
    }
    public void setCapabilitiesMsg(String v) {
        capabilitiesMsg.setValue(v);
    }
  }
}