package com.fireback.modules.licenses;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class LicenseFromPlanIdDto extends JsonSerializable {
    public String machineId;
    public String email;
    public String owner;
    public static class VM extends ViewModel {
    // upper: MachineId machineId
    private MutableLiveData< String > machineId = new MutableLiveData<>();
    public MutableLiveData< String > getMachineId() {
        return machineId;
    }
    public void setMachineId( String  v) {
        machineId.setValue(v);
    }
    // upper: Email email
    private MutableLiveData< String > email = new MutableLiveData<>();
    public MutableLiveData< String > getEmail() {
        return email;
    }
    public void setEmail( String  v) {
        email.setValue(v);
    }
    // upper: Owner owner
    private MutableLiveData< String > owner = new MutableLiveData<>();
    public MutableLiveData< String > getOwner() {
        return owner;
    }
    public void setOwner( String  v) {
        owner.setValue(v);
    }
    }
}