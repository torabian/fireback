package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class NodeDatumEntity extends JsonSerializable {
    public DataNodeEntity node;
    public float valueFloat64;
    public int valueInt64;
    public String valueString;
    public Boolean valueBoolean;
    public String ingestedAt;
    public static class VM extends ViewModel {
    // upper: Node node
    private MutableLiveData< DataNodeEntity > node = new MutableLiveData<>();
    public MutableLiveData< DataNodeEntity > getNode() {
        return node;
    }
    public void setNode( DataNodeEntity  v) {
        node.setValue(v);
    }
    // upper: ValueFloat64 valueFloat64
    private MutableLiveData< Float > valueFloat64 = new MutableLiveData<>();
    public MutableLiveData< Float > getValueFloat64() {
        return valueFloat64;
    }
    public void setValueFloat64( Float  v) {
        valueFloat64.setValue(v);
    }
    // upper: ValueInt64 valueInt64
    private MutableLiveData< Integer > valueInt64 = new MutableLiveData<>();
    public MutableLiveData< Integer > getValueInt64() {
        return valueInt64;
    }
    public void setValueInt64( Integer  v) {
        valueInt64.setValue(v);
    }
    // upper: ValueString valueString
    private MutableLiveData< String > valueString = new MutableLiveData<>();
    public MutableLiveData< String > getValueString() {
        return valueString;
    }
    public void setValueString( String  v) {
        valueString.setValue(v);
    }
    // upper: ValueBoolean valueBoolean
    private MutableLiveData< Boolean > valueBoolean = new MutableLiveData<>();
    public MutableLiveData< Boolean > getValueBoolean() {
        return valueBoolean;
    }
    public void setValueBoolean( Boolean  v) {
        valueBoolean.setValue(v);
    }
    // upper: IngestedAt ingestedAt
    private MutableLiveData< String > ingestedAt = new MutableLiveData<>();
    public MutableLiveData< String > getIngestedAt() {
        return ingestedAt;
    }
    public void setIngestedAt( String  v) {
        ingestedAt.setValue(v);
    }
    }
}