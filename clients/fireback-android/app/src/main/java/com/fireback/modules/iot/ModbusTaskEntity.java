package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ModbusTaskEntity extends JsonSerializable {
    public String name;
    public int modbusId;
    public DeviceEntity device;
    public ModbusConnectionTypeEntity connectionType;
    public ModbusFunctionCodeEntity functionCode;
    public int register;
    public int writeInterval;
    public int readInterval;
    public int range;
    public int length;
    public ModbusVariableTypeEntity variableType;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: ModbusId modbusId
    private MutableLiveData< Integer > modbusId = new MutableLiveData<>();
    public MutableLiveData< Integer > getModbusId() {
        return modbusId;
    }
    public void setModbusId( Integer  v) {
        modbusId.setValue(v);
    }
    // upper: Device device
    private MutableLiveData< DeviceEntity > device = new MutableLiveData<>();
    public MutableLiveData< DeviceEntity > getDevice() {
        return device;
    }
    public void setDevice( DeviceEntity  v) {
        device.setValue(v);
    }
    // upper: ConnectionType connectionType
    private MutableLiveData< ModbusConnectionTypeEntity > connectionType = new MutableLiveData<>();
    public MutableLiveData< ModbusConnectionTypeEntity > getConnectionType() {
        return connectionType;
    }
    public void setConnectionType( ModbusConnectionTypeEntity  v) {
        connectionType.setValue(v);
    }
    // upper: FunctionCode functionCode
    private MutableLiveData< ModbusFunctionCodeEntity > functionCode = new MutableLiveData<>();
    public MutableLiveData< ModbusFunctionCodeEntity > getFunctionCode() {
        return functionCode;
    }
    public void setFunctionCode( ModbusFunctionCodeEntity  v) {
        functionCode.setValue(v);
    }
    // upper: Register register
    private MutableLiveData< Integer > register = new MutableLiveData<>();
    public MutableLiveData< Integer > getRegister() {
        return register;
    }
    public void setRegister( Integer  v) {
        register.setValue(v);
    }
    // upper: WriteInterval writeInterval
    private MutableLiveData< Integer > writeInterval = new MutableLiveData<>();
    public MutableLiveData< Integer > getWriteInterval() {
        return writeInterval;
    }
    public void setWriteInterval( Integer  v) {
        writeInterval.setValue(v);
    }
    // upper: ReadInterval readInterval
    private MutableLiveData< Integer > readInterval = new MutableLiveData<>();
    public MutableLiveData< Integer > getReadInterval() {
        return readInterval;
    }
    public void setReadInterval( Integer  v) {
        readInterval.setValue(v);
    }
    // upper: Range range
    private MutableLiveData< Integer > range = new MutableLiveData<>();
    public MutableLiveData< Integer > getRange() {
        return range;
    }
    public void setRange( Integer  v) {
        range.setValue(v);
    }
    // upper: Length length
    private MutableLiveData< Integer > length = new MutableLiveData<>();
    public MutableLiveData< Integer > getLength() {
        return length;
    }
    public void setLength( Integer  v) {
        length.setValue(v);
    }
    // upper: VariableType variableType
    private MutableLiveData< ModbusVariableTypeEntity > variableType = new MutableLiveData<>();
    public MutableLiveData< ModbusVariableTypeEntity > getVariableType() {
        return variableType;
    }
    public void setVariableType( ModbusVariableTypeEntity  v) {
        variableType.setValue(v);
    }
    }
}