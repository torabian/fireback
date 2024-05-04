package com.fireback.modules.shop;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.ResponseErrorException;
public class ConfirmPurchaseAction {
    public static class Req extends JsonSerializable {
    public String basketId;
    public String currencyId;
    // upper: BasketId basketId
    private MutableLiveData<String> basketIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBasketIdMsg() {
        return basketIdMsg;
    }
    public void setBasketIdMsg(String v) {
        basketIdMsg.setValue(v);
    }
    // upper: CurrencyId currencyId
    private MutableLiveData<String> currencyIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCurrencyIdMsg() {
        return currencyIdMsg;
    }
    public void setCurrencyIdMsg(String v) {
        currencyIdMsg.setValue(v);
    }
    }
    public static class ReqViewModel extends ViewModel {
    // upper: BasketId basketId
    private MutableLiveData< String > basketId = new MutableLiveData<>();
    public MutableLiveData< String > getBasketId() {
        return basketId;
    }
    public void setBasketId( String  v) {
        basketId.setValue(v);
    }
    // upper: CurrencyId currencyId
    private MutableLiveData< String > currencyId = new MutableLiveData<>();
    public MutableLiveData< String > getCurrencyId() {
        return currencyId;
    }
    public void setCurrencyId( String  v) {
        currencyId.setValue(v);
    }
    // upper: BasketId basketId
    private MutableLiveData<String> basketIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBasketIdMsg() {
        return basketIdMsg;
    }
    public void setBasketIdMsg(String v) {
        basketIdMsg.setValue(v);
    }
    // upper: CurrencyId currencyId
    private MutableLiveData<String> currencyIdMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCurrencyIdMsg() {
        return currencyIdMsg;
    }
    public void setCurrencyIdMsg(String v) {
        currencyIdMsg.setValue(v);
    }
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;
    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        if (item.location != null && item.location.equals("basketId")) {
            this.setBasketIdMsg(item.messageTranslated);
        }
        if (item.location != null && item.location.equals("currencyId")) {
            this.setCurrencyIdMsg(item.messageTranslated);
        }
    });
}
    }
}