package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnReadFsFromHostConfigDto extends JsonSerializable {
    public String path;
    public static class VM extends ViewModel {
    // upper: Path path
    private MutableLiveData< String > path = new MutableLiveData<>();
    public MutableLiveData< String > getPath() {
        return path;
    }
    public void setPath( String  v) {
        path.setValue(v);
    }
    }
}