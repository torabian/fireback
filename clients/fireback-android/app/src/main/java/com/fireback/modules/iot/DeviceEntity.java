package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class DeviceDeviceModbusConfig extends JsonSerializable {
    public ModbusTransmissionModeEntity mode;
    public int baudRate;
    public int dataBits;
    public int parity;
    public int stopBit;
    public int timeout;
  public static class VM extends ViewModel {
    // upper: Mode mode
    private MutableLiveData< ModbusTransmissionModeEntity > mode = new MutableLiveData<>();
    public MutableLiveData< ModbusTransmissionModeEntity > getMode() {
        return mode;
    }
    public void setMode( ModbusTransmissionModeEntity  v) {
        mode.setValue(v);
    }
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
    private MutableLiveData< Integer > parity = new MutableLiveData<>();
    public MutableLiveData< Integer > getParity() {
        return parity;
    }
    public void setParity( Integer  v) {
        parity.setValue(v);
    }
    // upper: StopBit stopBit
    private MutableLiveData< Integer > stopBit = new MutableLiveData<>();
    public MutableLiveData< Integer > getStopBit() {
        return stopBit;
    }
    public void setStopBit( Integer  v) {
        stopBit.setValue(v);
    }
    // upper: Timeout timeout
    private MutableLiveData< Integer > timeout = new MutableLiveData<>();
    public MutableLiveData< Integer > getTimeout() {
        return timeout;
    }
    public void setTimeout( Integer  v) {
        timeout.setValue(v);
    }
  }
}
public class DeviceEntity extends JsonSerializable {
    public String name;
    public String model;
    public String ip;
    public String wifiUser;
    public String wifiPassword;
    public String securityType;
    public DeviceTypeEntity type;
    public DeviceDeviceModbusConfig deviceModbusConfig;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Model model
    private MutableLiveData< String > model = new MutableLiveData<>();
    public MutableLiveData< String > getModel() {
        return model;
    }
    public void setModel( String  v) {
        model.setValue(v);
    }
    // upper: Ip ip
    private MutableLiveData< String > ip = new MutableLiveData<>();
    public MutableLiveData< String > getIp() {
        return ip;
    }
    public void setIp( String  v) {
        ip.setValue(v);
    }
    // upper: WifiUser wifiUser
    private MutableLiveData< String > wifiUser = new MutableLiveData<>();
    public MutableLiveData< String > getWifiUser() {
        return wifiUser;
    }
    public void setWifiUser( String  v) {
        wifiUser.setValue(v);
    }
    // upper: WifiPassword wifiPassword
    private MutableLiveData< String > wifiPassword = new MutableLiveData<>();
    public MutableLiveData< String > getWifiPassword() {
        return wifiPassword;
    }
    public void setWifiPassword( String  v) {
        wifiPassword.setValue(v);
    }
    // upper: SecurityType securityType
    private MutableLiveData< String > securityType = new MutableLiveData<>();
    public MutableLiveData< String > getSecurityType() {
        return securityType;
    }
    public void setSecurityType( String  v) {
        securityType.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< DeviceTypeEntity > type = new MutableLiveData<>();
    public MutableLiveData< DeviceTypeEntity > getType() {
        return type;
    }
    public void setType( DeviceTypeEntity  v) {
        type.setValue(v);
    }
    // upper: DeviceModbusConfig deviceModbusConfig
    private MutableLiveData< DeviceDeviceModbusConfig > deviceModbusConfig = new MutableLiveData<>();
    public MutableLiveData< DeviceDeviceModbusConfig > getDeviceModbusConfig() {
        return deviceModbusConfig;
    }
    public void setDeviceModbusConfig( DeviceDeviceModbusConfig  v) {
        deviceModbusConfig.setValue(v);
    }
    }
}