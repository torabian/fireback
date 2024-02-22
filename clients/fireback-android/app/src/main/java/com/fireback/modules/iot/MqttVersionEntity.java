package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MqttVersionEntity extends JsonSerializable {
    public String version;
    public static class VM extends ViewModel {
    // upper: Version version
    private MutableLiveData< String > version = new MutableLiveData<>();
    public MutableLiveData< String > getVersion() {
        return version;
    }
    public void setVersion( String  v) {
        version.setValue(v);
    }
    }
}