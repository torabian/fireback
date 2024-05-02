package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class ShoppingCartItems extends JsonSerializable {
    public float quantity;
    public ProductSubmissionEntity product;
  public static class VM extends ViewModel {
    // upper: Quantity quantity
    private MutableLiveData< Float > quantity = new MutableLiveData<>();
    public MutableLiveData< Float > getQuantity() {
        return quantity;
    }
    public void setQuantity( Float  v) {
        quantity.setValue(v);
    }
    // upper: Product product
    private MutableLiveData< ProductSubmissionEntity > product = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionEntity > getProduct() {
        return product;
    }
    public void setProduct( ProductSubmissionEntity  v) {
        product.setValue(v);
    }
  }
}
public class ShoppingCartEntity extends JsonSerializable {
    public ShoppingCartItems[] items;
    public static class VM extends ViewModel {
    // upper: Items items
    private MutableLiveData< ShoppingCartItems[] > items = new MutableLiveData<>();
    public MutableLiveData< ShoppingCartItems[] > getItems() {
        return items;
    }
    public void setItems( ShoppingCartItems[]  v) {
        items.setValue(v);
    }
    }
}