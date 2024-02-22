package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class RoleEntity extends JsonSerializable {
    public String name;
    public CapabilityEntity[] capabilities;
    public String createdFormatted;
    public static class VM extends ViewModel {
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
    }
}