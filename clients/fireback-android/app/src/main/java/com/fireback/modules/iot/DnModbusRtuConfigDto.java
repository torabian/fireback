package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnModbusRtuConfigDto extends JsonSerializable {
    public int baudRate;
    public int dataBits;
    public String parity;
    public int stopBits;
    public int slaveId;
    public int timeout;
    public String address;
    public static class VM extends ViewModel {
    // upper: BaudRate baudRate
    private MutableLiveData< Integer > baudRate = new MutableLiveData<>();
    public MutableLiveData< Integer > getBaudRate() {
        return baudRate;
    }
    public void setBaudRate( Integer  v) {
        baudRate.setValue(v);
    }
    // upper: DataBits dataBits
    private MutableLiveData< Integer > dataBits = new MutableLiveData<>();
    public MutableLiveData< Integer > getDataBits() {
        return dataBits;
    }
    public void setDataBits( Integer  v) {
        dataBits.setValue(v);
    }
    // upper: Parity parity
    private MutableLiveData< String > parity = new MutableLiveData<>();
    public MutableLiveData< String > getParity() {
        return parity;
    }
    public void setParity( String  v) {
        parity.setValue(v);
    }
    // upper: StopBits stopBits
    private MutableLiveData< Integer > stopBits = new MutableLiveData<>();
    public MutableLiveData< Integer > getStopBits() {
        return stopBits;
    }
    public void setStopBits( Integer  v) {
        stopBits.setValue(v);
    }
    // upper: SlaveId slaveId
    private MutableLiveData< Integer > slaveId = new MutableLiveData<>();
    public MutableLiveData< Integer > getSlaveId() {
        return slaveId;
    }
    public void setSlaveId( Integer  v) {
        slaveId.setValue(v);
    }
    // upper: Timeout timeout
    private MutableLiveData< Integer > timeout = new MutableLiveData<>();
    public MutableLiveData< Integer > getTimeout() {
        return timeout;
    }
    public void setTimeout( Integer  v) {
        timeout.setValue(v);
    }
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