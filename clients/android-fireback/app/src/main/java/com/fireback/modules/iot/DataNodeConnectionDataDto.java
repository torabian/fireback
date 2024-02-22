package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DataNodeConnectionDataDto extends JsonSerializable {
    public String subKey;
    public static class VM extends ViewModel {
    // upper: SubKey subKey
    private MutableLiveData< String > subKey = new MutableLiveData<>();
    public MutableLiveData< String > getSubKey() {
        return subKey;
    }
    public void setSubKey( String  v) {
        subKey.setValue(v);
    }
    }
}