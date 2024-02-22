package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GpioModeEntity extends JsonSerializable {
    public String key;
    public int index;
    public String description;
    public static class VM extends ViewModel {
    // upper: Key key
    private MutableLiveData< String > key = new MutableLiveData<>();
    public MutableLiveData< String > getKey() {
        return key;
    }
    public void setKey( String  v) {
        key.setValue(v);
    }
    // upper: Index index
    private MutableLiveData< Integer > index = new MutableLiveData<>();
    public MutableLiveData< Integer > getIndex() {
        return index;
    }
    public void setIndex( Integer  v) {
        index.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    }
}