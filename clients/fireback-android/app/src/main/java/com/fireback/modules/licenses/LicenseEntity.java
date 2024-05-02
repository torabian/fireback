package com.fireback.modules.licenses;
import com.fireback.modules.workspaces.CapabilityEntity;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class LicensePermissions extends JsonSerializable {
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
public class LicenseEntity extends JsonSerializable {
    public String name;
    public String signedLicense;
    public java.util.Date validityStartDate;
    public java.util.Date validityEndDate;
    public LicensePermissions[] permissions;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: SignedLicense signedLicense
    private MutableLiveData< String > signedLicense = new MutableLiveData<>();
    public MutableLiveData< String > getSignedLicense() {
        return signedLicense;
    }
    public void setSignedLicense( String  v) {
        signedLicense.setValue(v);
    }
    // upper: ValidityStartDate validityStartDate
    private MutableLiveData< java.util.Date > validityStartDate = new MutableLiveData<>();
    public MutableLiveData< java.util.Date > getValidityStartDate() {
        return validityStartDate;
    }
    public void setValidityStartDate( java.util.Date  v) {
        validityStartDate.setValue(v);
    }
    // upper: ValidityEndDate validityEndDate
    private MutableLiveData< java.util.Date > validityEndDate = new MutableLiveData<>();
    public MutableLiveData< java.util.Date > getValidityEndDate() {
        return validityEndDate;
    }
    public void setValidityEndDate( java.util.Date  v) {
        validityEndDate.setValue(v);
    }
    // upper: Permissions permissions
    private MutableLiveData< LicensePermissions[] > permissions = new MutableLiveData<>();
    public MutableLiveData< LicensePermissions[] > getPermissions() {
        return permissions;
    }
    public void setPermissions( LicensePermissions[]  v) {
        permissions.setValue(v);
    }
    }
}