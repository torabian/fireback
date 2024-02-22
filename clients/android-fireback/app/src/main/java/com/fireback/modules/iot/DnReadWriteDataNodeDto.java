package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnReadWriteDataNodeDto extends JsonSerializable {
    public String nodeId;
    public static class VM extends ViewModel {
    // upper: NodeId nodeId
    private MutableLiveData< String > nodeId = new MutableLiveData<>();
    public MutableLiveData< String > getNodeId() {
        return nodeId;
    }
    public void setNodeId( String  v) {
        nodeId.setValue(v);
    }
    }
}