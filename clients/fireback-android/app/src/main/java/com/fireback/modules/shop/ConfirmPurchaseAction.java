package com.fireback.modules.shop;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class ConfirmPurchaseAction {
    public static class Req extends JsonSerializable {
    public String basketId;
    public String currencyId;
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
    }
}