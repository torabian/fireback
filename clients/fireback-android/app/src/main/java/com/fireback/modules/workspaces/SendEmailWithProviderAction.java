package com.fireback.modules.workspaces;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class SendEmailWithProviderAction {
    public static class Req extends JsonSerializable {
    public EmailProviderEntity emailProvider;
    public String toAddress;
    public String body;
    // upper: EmailProvider emailProvider
    private MutableLiveData<String> emailProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getEmailProviderMsg() {
        return emailProviderMsg;
    }
    public void setEmailProviderMsg(String v) {
        emailProviderMsg.setValue(v);
    }
    // upper: ToAddress toAddress
    private MutableLiveData<String> toAddressMsg = new MutableLiveData<>();
    public MutableLiveData<String> getToAddressMsg() {
        return toAddressMsg;
    }
    public void setToAddressMsg(String v) {
        toAddressMsg.setValue(v);
    }
    // upper: Body body
    private MutableLiveData<String> bodyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBodyMsg() {
        return bodyMsg;
    }
    public void setBodyMsg(String v) {
        bodyMsg.setValue(v);
    }
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
    // upper: EmailProvider emailProvider
    private MutableLiveData<String> emailProviderMsg = new MutableLiveData<>();
    public MutableLiveData<String> getEmailProviderMsg() {
        return emailProviderMsg;
    }
    public void setEmailProviderMsg(String v) {
        emailProviderMsg.setValue(v);
    }
    // upper: ToAddress toAddress
    private MutableLiveData<String> toAddressMsg = new MutableLiveData<>();
    public MutableLiveData<String> getToAddressMsg() {
        return toAddressMsg;
    }
    public void setToAddressMsg(String v) {
        toAddressMsg.setValue(v);
    }
    // upper: Body body
    private MutableLiveData<String> bodyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBodyMsg() {
        return bodyMsg;
    }
    public void setBodyMsg(String v) {
        bodyMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("emailProvider")) {
            this.setEmailProviderMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("toAddress")) {
            this.setToAddressMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("body")) {
            this.setBodyMsg(item.messageTranslated);
        }
    });
}
    }
    public static class Res extends JsonSerializable {
    public String queueId;
    // upper: QueueId queueId
    private MutableLiveData<String> queueIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getQueueIdMsg() {
        return queueIdMsg;
    }
    public void setQueueIdMsg(String v) {
        queueIdMsg.setValue(v);
    }
    }
}