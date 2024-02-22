package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ResetEmailDto extends JsonSerializable {
    public String password;
    public static class VM extends ViewModel {
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    }
}