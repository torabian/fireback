package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class ProductFields extends JsonSerializable {
    public ProductEntity product;
    public String name;
    public String type;
  public static class VM extends ViewModel {
    // upper: Product product
    private MutableLiveData< ProductEntity > product = new MutableLiveData<>();
    public MutableLiveData< ProductEntity > getProduct() {
        return product;
    }
    public void setProduct( ProductEntity  v) {
        product.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
  }
}
public class ProductEntity extends JsonSerializable {
    public String name;
    public String description;
    public String uiSchema;
    public String jsonSchema;
    public ProductFields[] fields;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: UiSchema uiSchema
    private MutableLiveData< String > uiSchema = new MutableLiveData<>();
    public MutableLiveData< String > getUiSchema() {
        return uiSchema;
    }
    public void setUiSchema( String  v) {
        uiSchema.setValue(v);
    }
    // upper: JsonSchema jsonSchema
    private MutableLiveData< String > jsonSchema = new MutableLiveData<>();
    public MutableLiveData< String > getJsonSchema() {
        return jsonSchema;
    }
    public void setJsonSchema( String  v) {
        jsonSchema.setValue(v);
    }
    // upper: Fields fields
    private MutableLiveData< ProductFields[] > fields = new MutableLiveData<>();
    public MutableLiveData< ProductFields[] > getFields() {
        return fields;
    }
    public void setFields( ProductFields[]  v) {
        fields.setValue(v);
    }
    }
}