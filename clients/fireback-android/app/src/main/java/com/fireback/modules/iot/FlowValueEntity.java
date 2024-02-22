package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class FlowValueEntity extends JsonSerializable {
    public String connectionId;
    public int valueInt;
    public String valueString;
    public float valueFloat;
    public int valueType;
    public static class VM extends ViewModel {
    // upper: ConnectionId connectionId
    private MutableLiveData< String > connectionId = new MutableLiveData<>();
    public MutableLiveData< String > getConnectionId() {
        return connectionId;
    }
    public void setConnectionId( String  v) {
        connectionId.setValue(v);
    }
    // upper: ValueInt valueInt
    private MutableLiveData< Integer > valueInt = new MutableLiveData<>();
    public MutableLiveData< Integer > getValueInt() {
        return valueInt;
    }
    public void setValueInt( Integer  v) {
        valueInt.setValue(v);
    }
    // upper: ValueString valueString
    private MutableLiveData< String > valueString = new MutableLiveData<>();
    public MutableLiveData< String > getValueString() {
        return valueString;
    }
    public void setValueString( String  v) {
        valueString.setValue(v);
    }
    // upper: ValueFloat valueFloat
    private MutableLiveData< Float > valueFloat = new MutableLiveData<>();
    public MutableLiveData< Float > getValueFloat() {
        return valueFloat;
    }
    public void setValueFloat( Float  v) {
        valueFloat.setValue(v);
    }
    // upper: ValueType valueType
    private MutableLiveData< Integer > valueType = new MutableLiveData<>();
    public MutableLiveData< Integer > getValueType() {
        return valueType;
    }
    public void setValueType( Integer  v) {
        valueType.setValue(v);
    }
    }
}