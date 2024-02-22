package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MqttClientConnectDto extends JsonSerializable {
    public Boolean connect;
    public static class VM extends ViewModel {
    // upper: Connect connect
    private MutableLiveData< Boolean > connect = new MutableLiveData<>();
    public MutableLiveData< Boolean > getConnect() {
        return connect;
    }
    public void setConnect( Boolean  v) {
        connect.setValue(v);
    }
    }
}