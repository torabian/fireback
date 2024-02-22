package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PhoneNumberAccountCreationDto extends JsonSerializable {
    public String phoneNumber;
    public static class VM extends ViewModel {
    // upper: PhoneNumber phoneNumber
    private MutableLiveData< String > phoneNumber = new MutableLiveData<>();
    public MutableLiveData< String > getPhoneNumber() {
        return phoneNumber;
    }
    public void setPhoneNumber( String  v) {
        phoneNumber.setValue(v);
    }
    }
}