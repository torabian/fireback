package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ImportRequestDto extends JsonSerializable {
    public String file;
    public static class VM extends ViewModel {
    // upper: File file
    private MutableLiveData< String > file = new MutableLiveData<>();
    public MutableLiveData< String > getFile() {
        return file;
    }
    public void setFile( String  v) {
        file.setValue(v);
    }
    }
}