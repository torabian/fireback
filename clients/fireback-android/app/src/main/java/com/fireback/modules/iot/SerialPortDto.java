package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class SerialPortDto extends JsonSerializable {
    public String address;
    public static class VM extends ViewModel {
    // upper: Address address
    private MutableLiveData< String > address = new MutableLiveData<>();
    public MutableLiveData< String > getAddress() {
        return address;
    }
    public void setAddress( String  v) {
        address.setValue(v);
    }
    }
}