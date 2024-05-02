package com.fireback.modules.licenses;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class LicenseFromActivationKeyDto extends JsonSerializable {
    public String activationKeyId;
    public String machineId;
    public static class VM extends ViewModel {
    // upper: ActivationKeyId activationKeyId
    private MutableLiveData< String > activationKeyId = new MutableLiveData<>();
    public MutableLiveData< String > getActivationKeyId() {
        return activationKeyId;
    }
    public void setActivationKeyId( String  v) {
        activationKeyId.setValue(v);
    }
    // upper: MachineId machineId
    private MutableLiveData< String > machineId = new MutableLiveData<>();
    public MutableLiveData< String > getMachineId() {
        return machineId;
    }
    public void setMachineId( String  v) {
        machineId.setValue(v);
    }
    }
}