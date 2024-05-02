package com.fireback.modules.licenses;
import com.fireback.modules.currency.PriceTagEntity;
import com.fireback.modules.workspaces.CapabilityEntity;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class ProductPlanPermissions extends JsonSerializable {
    public com.fireback.modules.workspaces.CapabilityEntity capability;
  public static class VM extends ViewModel {
    // upper: Capability capability
    private MutableLiveData<CapabilityEntity> capability = new MutableLiveData<>();
    public MutableLiveData< CapabilityEntity > getCapability() {
        return capability;
    }
    public void setCapability( CapabilityEntity  v) {
        capability.setValue(v);
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
    private MutableLiveData<PriceTagEntity> priceTag = new MutableLiveData<>();
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
    }
}