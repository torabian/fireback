package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MqttClientConnectionDto extends JsonSerializable {
    public String name;
    public Boolean isConnected;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: IsConnected isConnected
    private MutableLiveData< Boolean > isConnected = new MutableLiveData<>();
    public MutableLiveData< Boolean > getIsConnected() {
        return isConnected;
    }
    public void setIsConnected( Boolean  v) {
        isConnected.setValue(v);
    }
    }
}