package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ExpanderFunctionEntity extends JsonSerializable {
    public String name;
    public String nativeFn;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: NativeFn nativeFn
    private MutableLiveData< String > nativeFn = new MutableLiveData<>();
    public MutableLiveData< String > getNativeFn() {
        return nativeFn;
    }
    public void setNativeFn( String  v) {
        nativeFn.setValue(v);
    }
    }
}