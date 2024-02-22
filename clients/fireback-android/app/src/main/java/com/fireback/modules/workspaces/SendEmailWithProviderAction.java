package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class SendEmailWithProviderAction {
    public static class Req extends JsonSerializable {
    public EmailProviderEntity emailProvider;
    public String toAddress;
    public String body;
    }
    public static class ReqViewModel extends ViewModel {
    // upper: EmailProvider emailProvider
    private MutableLiveData< EmailProviderEntity > emailProvider = new MutableLiveData<>();
    public MutableLiveData< EmailProviderEntity > getEmailProvider() {
        return emailProvider;
    }
    public void setEmailProvider( EmailProviderEntity  v) {
        emailProvider.setValue(v);
    }
    // upper: ToAddress toAddress
    private MutableLiveData< String > toAddress = new MutableLiveData<>();
    public MutableLiveData< String > getToAddress() {
        return toAddress;
    }
    public void setToAddress( String  v) {
        toAddress.setValue(v);
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