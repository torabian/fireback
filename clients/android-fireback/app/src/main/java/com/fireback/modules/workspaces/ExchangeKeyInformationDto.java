package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ExchangeKeyInformationDto extends JsonSerializable {
    public String key;
    public String visibility;
    public static class VM extends ViewModel {
    // upper: Key key
    private MutableLiveData< String > key = new MutableLiveData<>();
    public MutableLiveData< String > getKey() {
        return key;
    }
    public void setKey( String  v) {
        key.setValue(v);
    }
    // upper: Visibility visibility
    private MutableLiveData< String > visibility = new MutableLiveData<>();
    public MutableLiveData< String > getVisibility() {
        return visibility;
    }
    public void setVisibility( String  v) {
        visibility.setValue(v);
    }
    }
}