package com.fireback.modules.workspaces;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class EmailOtpResponseDto extends JsonSerializable {
    public ForgetPasswordEntity request;
    public UserSessionDto userSession;
    public static class VM extends ViewModel {
    // upper: Request request
    private MutableLiveData< ForgetPasswordEntity > request = new MutableLiveData<>();
    public MutableLiveData< ForgetPasswordEntity > getRequest() {
        return request;
    }
    public void setRequest( ForgetPasswordEntity  v) {
        request.setValue(v);
    }
    // upper: UserSession userSession
    private MutableLiveData< UserSessionDto > userSession = new MutableLiveData<>();
    public MutableLiveData< UserSessionDto > getUserSession() {
        return userSession;
    }
    public void setUserSession( UserSessionDto  v) {
        userSession.setValue(v);
    }
    }
}