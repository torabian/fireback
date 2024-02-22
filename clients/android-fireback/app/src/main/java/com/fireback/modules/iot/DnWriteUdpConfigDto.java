package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnWriteUdpConfigDto extends JsonSerializable {
    public String host;
    public String port;
    public static class VM extends ViewModel {
    // upper: Host host
    private MutableLiveData< String > host = new MutableLiveData<>();
    public MutableLiveData< String > getHost() {
        return host;
    }
    public void setHost( String  v) {
        host.setValue(v);
    }
    // upper: Port port
    private MutableLiveData< String > port = new MutableLiveData<>();
    public MutableLiveData< String > getPort() {
        return port;
    }
    public void setPort( String  v) {
        port.setValue(v);
    }
    }
}