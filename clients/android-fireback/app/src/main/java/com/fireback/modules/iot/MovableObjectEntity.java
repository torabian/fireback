package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MovableObjectEntity extends JsonSerializable {
    public String name;
    public InteractiveMapEntity[] interactiveMaps;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: InteractiveMaps interactiveMaps
    private MutableLiveData< InteractiveMapEntity[] > interactiveMaps = new MutableLiveData<>();
    public MutableLiveData< InteractiveMapEntity[] > getInteractiveMaps() {
        return interactiveMaps;
    }
    public void setInteractiveMaps( InteractiveMapEntity[]  v) {
        interactiveMaps.setValue(v);
    }
    }
}