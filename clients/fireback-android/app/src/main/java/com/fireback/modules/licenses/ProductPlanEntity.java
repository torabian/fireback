package com.fireback.modules.licenses;
import com.fireback.modules.currency.*;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class ProductPlanPermissions extends JsonSerializable {
    public com.fireback.modules.workspaces.CapabilityEntity capability;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Capability capability
    private MutableLiveData< CapabilityEntity > capability = new MutableLiveData<>();
    public MutableLiveData< CapabilityEntity > getCapability() {
        return capability;
    }
    public void setCapability( CapabilityEntity  v) {
        capability.setValue(v);
    }
    // Handling error message for each field
    // upper: Capability capability
    private MutableLiveData<String> capabilityMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCapabilityMsg() {
        return capabilityMsg;
    }
    public void setCapabilityMsg(String v) {
        capabilityMsg.setValue(v);
    }
  }
}
public class ProductPlanEntity extends JsonSerializable {
    public String name;
    public int duration;
    public LicensableProductEntity product;
    public com.fireback.modules.currency.PriceTagEntity priceTag;
    public ProductPlanPermissions[] permissions;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Duration duration
    private MutableLiveData< Integer > duration = new MutableLiveData<>();
    public MutableLiveData< Integer > getDuration() {
        return duration;
    }
    public void setDuration( Integer  v) {
        duration.setValue(v);
    }
    // upper: Product product
    private MutableLiveData< LicensableProductEntity > product = new MutableLiveData<>();
    public MutableLiveData< LicensableProductEntity > getProduct() {
        return product;
    }
    public void setProduct( LicensableProductEntity  v) {
        product.setValue(v);
    }
    // upper: PriceTag priceTag
    private MutableLiveData< PriceTagEntity > priceTag = new MutableLiveData<>();
    public MutableLiveData< PriceTagEntity > getPriceTag() {
        return priceTag;
    }
    public void setPriceTag( PriceTagEntity  v) {
        priceTag.setValue(v);
    }
    // upper: Permissions permissions
    private MutableLiveData< ProductPlanPermissions[] > permissions = new MutableLiveData<>();
    public MutableLiveData< ProductPlanPermissions[] > getPermissions() {
        return permissions;
    }
    public void setPermissions( ProductPlanPermissions[]  v) {
        permissions.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Duration duration
    private MutableLiveData<String> durationMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDurationMsg() {
        return durationMsg;
    }
    public void setDurationMsg(String v) {
        durationMsg.setValue(v);
    }
    // upper: Product product
    private MutableLiveData<String> productMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductMsg() {
        return productMsg;
    }
    public void setProductMsg(String v) {
        productMsg.setValue(v);
    }
    // upper: PriceTag priceTag
    private MutableLiveData<String> priceTagMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPriceTagMsg() {
        return priceTagMsg;
    }
    public void setPriceTagMsg(String v) {
        priceTagMsg.setValue(v);
    }
    // upper: Permissions permissions
    private MutableLiveData<String> permissionsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPermissionsMsg() {
        return permissionsMsg;
    }
    public void setPermissionsMsg(String v) {
        permissionsMsg.setValue(v);
    }
  }
}