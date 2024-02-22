package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class WriteDatumDto extends JsonSerializable {
    public String uniqueId;
    public String key;
    public String value;
    public static class VM extends ViewModel {
    // upper: UniqueId uniqueId
    private MutableLiveData< String > uniqueId = new MutableLiveData<>();
    public MutableLiveData< String > getUniqueId() {
        return uniqueId;
    }
    public void setUniqueId( String  v) {
        uniqueId.setValue(v);
    }
    // upper: Key key
    private MutableLiveData< String > key = new MutableLiveData<>();
    public MutableLiveData< String > getKey() {
        return key;
    }
    public void setKey( String  v) {
        key.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    }
}