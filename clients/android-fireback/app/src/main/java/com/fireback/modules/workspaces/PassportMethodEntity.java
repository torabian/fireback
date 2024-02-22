package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PassportMethodEntity extends JsonSerializable {
    public String name;
    public String type;
    public String region;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Region region
    private MutableLiveData< String > region = new MutableLiveData<>();
    public MutableLiveData< String > getRegion() {
        return region;
    }
    public void setRegion( String  v) {
        region.setValue(v);
    }
    }
}