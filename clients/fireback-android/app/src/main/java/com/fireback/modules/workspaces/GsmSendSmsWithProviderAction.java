package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class GsmSendSmsWithProviderAction {
    public static class Req extends JsonSerializable {
    public GsmProviderEntity gsmProvider;
    public String toNumber;
    public String body;
    }
    public static class ReqViewModel extends ViewModel {
    // upper: GsmProvider gsmProvider
    private MutableLiveData< GsmProviderEntity > gsmProvider = new MutableLiveData<>();
    public MutableLiveData< GsmProviderEntity > getGsmProvider() {
        return gsmProvider;
    }
    public void setGsmProvider( GsmProviderEntity  v) {
        gsmProvider.setValue(v);
    }
    // upper: ToNumber toNumber
    private MutableLiveData< String > toNumber = new MutableLiveData<>();
    public MutableLiveData< String > getToNumber() {
        return toNumber;
    }
    public void setToNumber( String  v) {
        toNumber.setValue(v);
    }
    // upper: Body body
    private MutableLiveData< String > body = new MutableLiveData<>();
    public MutableLiveData< String > getBody() {
        return body;
    }
    public void setBody( String  v) {
        body.setValue(v);
    }
    }
    public static class Res extends JsonSerializable {
    public String queueId;
    }
}