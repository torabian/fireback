package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnReadSerialPortConfigDto extends JsonSerializable {
    public String address;
    public int baudRate;
    public static class VM extends ViewModel {
    // upper: Address address
    private MutableLiveData< String > address = new MutableLiveData<>();
    public MutableLiveData< String > getAddress() {
        return address;
    }
    public void setAddress( String  v) {
        address.setValue(v);
    }
    // upper: BaudRate baudRate
    private MutableLiveData< Integer > baudRate = new MutableLiveData<>();
    public MutableLiveData< Integer > getBaudRate() {
        return baudRate;
    }
    public void setBaudRate( Integer  v) {
        baudRate.setValue(v);
    }
    }
}