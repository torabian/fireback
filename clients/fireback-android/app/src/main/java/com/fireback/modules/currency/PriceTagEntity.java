package com.fireback.modules.currency;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class PriceTagVariations extends JsonSerializable {
    public CurrencyEntity currency;
    public float amount;
  public static class VM extends ViewModel {
    // upper: Currency currency
    private MutableLiveData< CurrencyEntity > currency = new MutableLiveData<>();
    public MutableLiveData< CurrencyEntity > getCurrency() {
        return currency;
    }
    public void setCurrency( CurrencyEntity  v) {
        currency.setValue(v);
    }
    // upper: Amount amount
    private MutableLiveData< Float > amount = new MutableLiveData<>();
    public MutableLiveData< Float > getAmount() {
        return amount;
    }
    public void setAmount( Float  v) {
        amount.setValue(v);
    }
  }
}
public class PriceTagEntity extends JsonSerializable {
    public PriceTagVariations[] variations;
    public static class VM extends ViewModel {
    // upper: Variations variations
    private MutableLiveData< PriceTagVariations[] > variations = new MutableLiveData<>();
    public MutableLiveData< PriceTagVariations[] > getVariations() {
        return variations;
    }
    public void setVariations( PriceTagVariations[]  v) {
        variations.setValue(v);
    }
    }
}