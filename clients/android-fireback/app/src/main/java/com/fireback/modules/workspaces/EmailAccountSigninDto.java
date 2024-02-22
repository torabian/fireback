package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class EmailAccountSigninDto extends JsonSerializable {
    public String email;
    public String password;
    public static class VM extends ViewModel {
    // upper: Email email
    private MutableLiveData< String > email = new MutableLiveData<>();
    public MutableLiveData< String > getEmail() {
        return email;
    }
    public void setEmail( String  v) {
        email.setValue(v);
    }
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