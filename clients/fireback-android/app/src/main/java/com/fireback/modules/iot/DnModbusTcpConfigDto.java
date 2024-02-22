package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnModbusTcpConfigDto extends JsonSerializable {
    public int timeOut;
    public int slaveId;
    public String host;
    public String port;
    public static class VM extends ViewModel {
    // upper: TimeOut timeOut
    private MutableLiveData< Integer > timeOut = new MutableLiveData<>();
    public MutableLiveData< Integer > getTimeOut() {
        return timeOut;
    }
    public void setTimeOut( Integer  v) {
        timeOut.setValue(v);
    }
    // upper: SlaveId slaveId
    private MutableLiveData< Integer > slaveId = new MutableLiveData<>();
    public MutableLiveData< Integer > getSlaveId() {
        return slaveId;
    }
    public void setSlaveId( Integer  v) {
        slaveId.setValue(v);
    }
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