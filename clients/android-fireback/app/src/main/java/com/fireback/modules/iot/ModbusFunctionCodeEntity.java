package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ModbusFunctionCodeEntity extends JsonSerializable {
    public String name;
    public int code;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Code code
    private MutableLiveData< Integer > code = new MutableLiveData<>();
    public MutableLiveData< Integer > getCode() {
        return code;
    }
    public void setCode( Integer  v) {
        code.setValue(v);
    }
    }
}