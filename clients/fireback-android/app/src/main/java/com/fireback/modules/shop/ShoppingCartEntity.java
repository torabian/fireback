package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class ShoppingCartItems extends JsonSerializable {
    public float quantity;
    public ProductSubmissionEntity product;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
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
    // Handling error message for each field
    // upper: Quantity quantity
    private MutableLiveData<String> quantityMsg = new MutableLiveData<>();
    public MutableLiveData<String> getQuantityMsg() {
        return quantityMsg;
    }
    public void setQuantityMsg(String v) {
        quantityMsg.setValue(v);
    }
    // upper: Product product
    private MutableLiveData<String> productMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductMsg() {
        return productMsg;
    }
    public void setProductMsg(String v) {
        productMsg.setValue(v);
    }
  }
}
public class ShoppingCartEntity extends JsonSerializable {
    public ShoppingCartItems[] items;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Items items
    private MutableLiveData< ShoppingCartItems[] > items = new MutableLiveData<>();
    public MutableLiveData< ShoppingCartItems[] > getItems() {
        return items;
    }
    public void setItems( ShoppingCartItems[]  v) {
        items.setValue(v);
    }
    // Handling error message for each field
    // upper: Items items
    private MutableLiveData<String> itemsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getItemsMsg() {
        return itemsMsg;
    }
    public void setItemsMsg(String v) {
        itemsMsg.setValue(v);
    }
  }
}