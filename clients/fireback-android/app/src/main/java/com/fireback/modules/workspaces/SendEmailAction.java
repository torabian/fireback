/*
 *	Generated by fireback 1.1.16
 *	Written by Ali Torabi.
 *	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
 */
package com.fireback.modules.workspaces;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.JsonSerializable;
import com.fireback.ResponseErrorException;

public class SendEmailAction {
  public static class Req extends JsonSerializable {
    public String toAddress;
    public String body;
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
    // upper: ToAddress toAddress
    private MutableLiveData<String> toAddress = new MutableLiveData<>();

    public MutableLiveData<String> getToAddress() {
      return toAddress;
    }

    public void setToAddress(String v) {
      toAddress.setValue(v);
    }

    // upper: Body body
    private MutableLiveData<String> body = new MutableLiveData<>();

    public MutableLiveData<String> getBody() {
      return body;
    }

    public void setBody(String v) {
      body.setValue(v);
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
      responseError.error.errors.forEach(
          item -> {
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
